package pdf_to_image_converter

type Converter interface {
	Convert(file []byte, format imageFormat) ([][]byte, error)
}
