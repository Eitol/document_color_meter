package core

import (
	"bytes"
	"fmt"
	"github.com/Eitol/document_color_meter/api_front/core/content_type"
	"github.com/Eitol/document_color_meter/api_front/core/image_converter"
	"image"
	_ "image/jpeg"
	"math"
)

type Color string

const (
	White      = "white"
	Black      = "black"
	Red        = "red"
	Green      = "green"
	Blue       = "blue"
	Gray       = "gray"
	OtherColor = "color"
)

var DefaultWhiteThreshold = uint32(240 + 240 + 240)
var DefaultBlackThreshold = uint32(7 + 7 + 7)
var DefaultGrayThreshold = uint32(15)

const PdfContentType = "application/pdf"

type PageColorMeasureResult struct {
	TotalColorPercentage float32           `json:"totalColorPercentage" bson:"totalColorPercentage"`
	PercentageByColor    map[Color]float32 `json:"percentageByColor" bson:"percentageByColor"`
}

type DocColorMeasureResult struct {
	Pages        map[int]PageColorMeasureResult `json:"pages" bson:"pages"`
	PagesCount   int32                          `json:"pagesCount" bson:"pagesCount"`
	DocumentType string                         `json:"documentType" bson:"documentType"`
}

type Options struct {
	WhiteThreshold uint32 `json:"whiteThreshold" bson:"whiteThreshold"`
	BlackThreshold uint32 `json:"blackThreshold" bson:"blackThreshold"`
	GrayThreshold  uint32 `json:"grayThreshold" bson:"grayThreshold"`
}

type ColorMeter interface {
	MeasureColor(doc []byte, options Options) (*DocColorMeasureResult, [][]byte, error)
}

func NewColorMeter() ColorMeter {
	return &concreteColorMeter{}
}

type concreteColorMeter struct{}

func (c *concreteColorMeter) MeasureColor(doc []byte, options Options) (*DocColorMeasureResult, [][]byte, error) {
	options = mergeOptions(options)
	contentType := content_type.Detect(doc)

	// Phase 1: convert document pages to images
	imgPages, err := getImagesFromFile(doc, contentType)
	if err != nil {
		return nil, nil, err
	}

	// Phase 2: measure the color of every page
	pagesResult, err := measureColorOfPages(imgPages, options)
	if err != nil {
		return nil, nil, err
	}
	return &DocColorMeasureResult{
		Pages:        pagesResult,
		PagesCount:   int32(len(pagesResult)),
		DocumentType: contentType,
	}, imgPages, nil
}

// Convert a document to an array of imagess
func getImagesFromFile(doc []byte, contentType string) ([][]byte, error) {
	var imgPages [][]byte
	var err error
	if content_type.IsImage(contentType) {
		jpegFile := doc
		if !content_type.IsImageJpeg(contentType) {
			jpegFile, err = image_converter.ConvertToJpeg(doc)
			if err != nil {
				return nil, err
			}
		}
		return [][]byte{jpegFile}, nil
	}
	pdfDoc := doc
	if !content_type.IsPDF(contentType) {
		pdfDoc, err = convertDocToPdf(doc)
		if err != nil {
			return nil, err
		}
	}
	imgPages, err = convertPdfToImages(pdfDoc)
	if err != nil {
		return nil, err
	}
	return imgPages, nil
}

// convert 255 rgb to alpha-premultiplied rgb (max 65535)
// see: https://blog.golang.org/image
func normalizeOptions(options Options) Options {
	options.GrayThreshold = (options.GrayThreshold * (math.MaxUint16)) / uint32(255)
	options.BlackThreshold = (options.BlackThreshold * (math.MaxUint16)) / uint32(255)
	options.WhiteThreshold = (options.WhiteThreshold * (math.MaxUint16)) / uint32(255)
	return options
}

func mergeOptions(options Options) Options {
	if options.WhiteThreshold == 0 {
		options.WhiteThreshold = DefaultWhiteThreshold
	}
	if options.GrayThreshold == 0 {
		options.GrayThreshold = DefaultGrayThreshold
	}
	if options.BlackThreshold == 0 {
		options.BlackThreshold = DefaultBlackThreshold
	}
	return options
}

func measureColorOfPages(pages [][]byte, options Options) (map[int]PageColorMeasureResult, error) {
	outPages := map[int]PageColorMeasureResult{}
	count := 1
	for _, page := range pages {
		pageResult, err := measureColorOfPage(page, options)
		if err != nil {
			return nil, err
		}
		if pageResult == nil {
			return nil, fmt.Errorf("nil result")
		}
		outPages[count] = *pageResult
		count += 1
	}
	return outPages, nil
}

func getMin(r, g, b uint32) uint32 {
	if r < g && r < b {
		return r
	}
	if b < g {
		return b
	}
	return g
}

func getMax(r, g, b uint32) uint32 {
	if r > g && r > b {
		return r
	}
	if b > g {
		return b
	}
	return g
}

func isWhite(s, wt uint32) bool {
	return s >= wt
}

func isBlack(s, bt uint32) bool {
	return s <= bt
}

func isGray(r, g, b, grayThreshold uint32) bool {
	min := getMin(r, g, b)
	max := getMax(r, g, b)
	diff := max - min
	return diff <= grayThreshold
}

func getColorOfPixel(r, g, b uint32, options Options) Color {
	s := r + g + b
	if isWhite(s, options.WhiteThreshold) {
		return White
	} else if isBlack(s, options.BlackThreshold) {
		return Black
	} else if isGray(r, g, b, options.GrayThreshold) {
		return Gray
	}
	return OtherColor
}

func computePercentageByColor(totalPixelCount int, colorCount map[Color]int) map[Color]float32 {
	out := make(map[Color]float32)
	for color, val := range colorCount {
		out[color] = float32(val) / float32(totalPixelCount) * float32(100.0)
	}
	return out
}

func measureColorOfPage(page []byte, options Options) (*PageColorMeasureResult, error) {
	options = normalizeOptions(options)
	// Step 1: Count get the color of each pixel and count by color
	img, _, err := image.Decode(bytes.NewReader(page))
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	colorCount := make(map[Color]int)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			color := getColorOfPixel(r, g, b, options)
			_, ok := colorCount[color]
			if !ok {
				colorCount[color] = 1
			} else {
				colorCount[color]++
			}
		}
	}
	// Step 2: Compute the result
	totalPixelCount := bounds.Max.Y * bounds.Max.X
	colourPixelCount, ok := colorCount[OtherColor]
	colorPercentage := float32(0)
	if ok && colourPixelCount > 0 {
		colorPercentage = float32(colourPixelCount) / float32(totalPixelCount)
	}
	return &PageColorMeasureResult{
		TotalColorPercentage: colorPercentage * 100,
		PercentageByColor:    computePercentageByColor(int(totalPixelCount), colorCount),
	}, nil
}
