package test_util

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

var testDataPath = filepath.Join("test_util", "testdata")

const (
	ValidWord1997File           = "valid_word1997.doc"
	ValidWord2007File           = "valid_word2007.docx"
	ValidWord2007MulticolorFile = "valid_word2007_multicolor.docx"
	ValidPDFFile                = "valid_pdf.pdf"
	ValidPDFFileBig             = "big_file.pdf"
	ValidImage                  = "valid_image.jpg"
)

func getRootPath() string {
	_, path, _, _ := runtime.Caller(1)
	for {
		path = filepath.Dir(path)
		if filepath.Base(path) == "core" {
			return path
		}
	}
}

func GetTestFile(name string) []byte {
	file, err := ioutil.ReadFile(filepath.Join(getRootPath(), testDataPath, name))
	if err != nil {
		log.Fatal("Error invalid test file")
	}
	return file
}
