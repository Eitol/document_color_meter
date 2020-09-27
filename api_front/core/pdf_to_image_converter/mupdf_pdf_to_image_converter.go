package pdf_to_image_converter

import (
	"bytes"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"runtime"
)

const defaultQuality = 100

const defaultOutputFormat = ImageFormatJPEG

type ConversionOptions struct {
	QualityPercentage int
	OutputFormat      imageFormat
	Password          string
}

type mupdfPdfToImageConverter struct {
}

func newMuPdfConverter() mupdfPdfToImageConverter {
	return mupdfPdfToImageConverter{}
}

func prepareOpts(opts ConversionOptions) ConversionOptions {
	if opts.QualityPercentage > 100 || opts.QualityPercentage <= 0 {
		opts.QualityPercentage = defaultQuality
	}
	if opts.OutputFormat == imageFormat(0) {
		opts.OutputFormat = defaultOutputFormat
	}
	return opts
}

func (mupdfPdfToImageConverter) Convert(pdfFile []byte, opts ConversionOptions) ([][]byte, error) {
	opts = prepareOpts(opts)
	doc, err := fitz.NewFromMemory(pdfFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := doc.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	out := make([][]byte, 0, doc.NumPage())
	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			return nil, fmt.Errorf("conversion error in page %d: %v", n, err)
		}
		imageBytes, err := imageToBytes(img, opts)
		if err != nil {
			return nil, fmt.Errorf("encoding error in page %d: %v", n, err)
		}
		out = append(out, imageBytes)
		runtime.GC()
	}
	return out, nil
}

// convert image.Image to byte array
func imageToBytes(img image.Image, opts ConversionOptions) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	if opts.OutputFormat == ImageFormatJPEG {
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: opts.QualityPercentage})
	} else if opts.OutputFormat == ImageFormatPNG {
		err = png.Encode(buf, img)
	} else {
		err = fmt.Errorf("unsuported format: %v", opts.OutputFormat)
	}
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
