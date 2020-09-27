package image_converter

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
)

// PNG to JPG convertion
func ConvertToJpeg(img []byte) ([]byte, error) {
	imgSrc, err := png.Decode(bytes.NewReader(img))
	if err != nil {
		return nil, err
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
	var opt jpeg.Options
	opt.Quality = 100
	buff := make([]byte, 0)
	w := bytes.NewBuffer(buff)
	err = jpeg.Encode(w, newImg, &opt)
	if err != nil {
		return nil, err
	}
	buff2, err2 := ioutil.ReadAll(w)
	if err2 != nil {
		return nil, err2
	}
	return buff2, nil
}
