package content_type

import (
	"net/http"
	"strings"
)

const PdfContentType = "application/pdf"

var OfficeMimeTypes = []string{"application/msword",
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.template",
	"application/vnd.ms-word.document.macroEnabled.12",
	"application/vnd.ms-word.template.macroEnabled.12",
	"application/vnd.ms-excel",
	"application/vnd.ms-excel",
	"application/vnd.ms-excel",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.template",
	"application/vnd.ms-excel.sheet.macroEnabled.12",
	"application/vnd.ms-excel.template.macroEnabled.12",
	"application/vnd.ms-excel.addin.macroEnabled.12",
	"application/vnd.ms-excel.sheet.binary.macroEnabled.12",
	"application/vnd.ms-powerpoint",
	"application/vnd.ms-powerpoint",
	"application/vnd.ms-powerpoint",
	"application/vnd.ms-powerpoint",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"application/vnd.openxmlformats-officedocument.presentationml.template",
	"application/vnd.openxmlformats-officedocument.presentationml.slideshow",
	"application/vnd.ms-powerpoint.addin.macroEnabled.12",
	"application/vnd.ms-powerpoint.presentation.macroEnabled.12",
	"application/vnd.ms-powerpoint.template.macroEnabled.12",
	"application/vnd.ms-powerpoint.slideshow.macroEnabled.12",
}

var ImageJpegMimeTypes = []string{
	"image/jpeg",
	"image/pjpeg",
	"image/jpeg",
	"image/jpeg",
	"image/pjpeg",
	"image/jpeg",
	"image/pjpeg",
	"image/jpeg",
	"image/pjpeg",
}

func IsPDF(ct string) bool {
	return ct == PdfContentType
}

func IsImage(ct string) bool {
	return strings.Split(ct, "/")[0] == "image"
}

func IsImageJpeg(ct string) bool {
	for _, mimeType := range ImageJpegMimeTypes {
		if mimeType == ct {
			return true
		}
	}
	return false
}

func IsOffice(ct string) bool {
	for _, mimeType := range OfficeMimeTypes {
		if mimeType == ct {
			return true
		}
	}
	return false
}

func IsUndetermined(ct string) bool {
	return ct == "application/octet-stream"
}

func Detect(f []byte) string {
	return http.DetectContentType(f)
}
