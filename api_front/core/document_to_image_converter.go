package core

import (
	"github.com/Eitol/document_color_meter/api_front/config"
	"github.com/Eitol/document_color_meter/api_front/core/office_converter"
	"github.com/Eitol/document_color_meter/api_front/core/pdf_to_image_converter"
	"net/http"
	"os"
)

const imgConversionFormat = pdf_to_image_converter.ImageFormatJPEG

type ConverterSource string

const (
	Remote = ConverterSource("remote")
	Local  = ConverterSource("local")
)

const defaultSource = Local

func convertDocToPdf(doc []byte) ([]byte, error) {
	contentType := http.DetectContentType(doc)
	if contentType == PdfContentType {
		return doc, nil
	}
	converter := office_converter.NewOfficeConverter()
	var err error
	doc, err = converter.Convert(doc, office_converter.PDF)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func getConverterToUse() ConverterSource {
	val := os.Getenv("PDF_TO_IMAGE_CONVERTER_SOURCE")
	if val == "" {
		return defaultSource
	}
	if val == "remote" {
		return Remote
	}
	return defaultSource
}

func convertPdfToImages(doc []byte) ([][]byte, error) {
	source := getConverterToUse()
	var pdfToImageConv pdf_to_image_converter.Converter
	if source == Remote {
		pdfToImageConv = pdf_to_image_converter.NewRemoteConverter(
			config.GetConfig().ImageConversionHost,
			config.GetConfig().ExternalServicesTimeout,
		)
	} else {
		pdfToImageConv = pdf_to_image_converter.NewLocalConverter()
	}
	pagesImgs, err := pdfToImageConv.Convert(doc, imgConversionFormat)
	if err != nil {
		return nil, err
	}
	return pagesImgs, nil
}
