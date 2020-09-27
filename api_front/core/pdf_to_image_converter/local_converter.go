package pdf_to_image_converter

type localConverter struct{}

func NewLocalConverter() Converter {
	return localConverter{}
}

func (localConverter) Convert(file []byte, format imageFormat) ([][]byte, error) {
	converter := newMuPdfConverter()
	images, err := converter.Convert(file, ConversionOptions{OutputFormat: format})
	if err != nil {
		return nil, err
	}
	return images, nil
}
