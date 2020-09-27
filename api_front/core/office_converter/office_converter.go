package office_converter

import (
	"fmt"
	"github.com/Eitol/document_color_meter/api_front/core/temp_file"
	"github.com/dveselov/go-libreofficekit"
	"github.com/gojetpack/pyos"
	"io/ioutil"
	"os"
	"path/filepath"
)

type OfficeConverter interface {
	Convert(doc []byte, outFormat DocumentFormat) ([]byte, error)
}

func NewOfficeConverter() OfficeConverter {
	return &converterWithLOKit{}
}

type converterWithLOKit struct {
	InstallationPath string
}

var possibleLibreOfficeInstallDir = []string{
	"/usr/lib",
	"/usr/lib32",
	"/usr/lib64",
	"/opt",
	"/usr/local",
	"/usr/bin",
}

var libreOfficeNames = []string{
	"libreoffice",
	"openoffice",
	"ooo",
}

var libreOfficeCommonLibsPath = []string{
	"program",
}

const notFoundInstallationError = `
the libre office installation path was not found.
If you are on debian/ubuntu you can install libre office with:
sudo apt-get update && apt-get install libreoffice libreofficekit-dev
`

// returns the possible installation path of libre office
func findLibreOfficeInstallationPath() (string, error) {
	for _, path := range possibleLibreOfficeInstallDir {
		if pyos.Path.IsDir(path) {
			for _, name := range libreOfficeNames {
				p := filepath.Join(path, name)
				if pyos.Path.Exist(p) {
					commonLibPath := findCommonLibsPath(p)
					if commonLibPath != "" {
						return commonLibPath, nil
					}
					programPath := findLibsPathRecursive(p)
					if programPath != "" {
						return programPath, nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf(notFoundInstallationError)
}

func findCommonLibsPath(path string) string {
	for _, commonLibPath := range libreOfficeCommonLibsPath {
		p := filepath.Join(path, commonLibPath)
		if pyos.Path.IsDir(p) {
			return p
		}
	}
	return ""
}

func findLibsPathRecursive(path string) string {
	out := ""
	err := filepath.Walk(path,
		func(f string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if out != "" {
				return nil
			}
			if info.IsDir() {
				return nil
			}

			if filepath.Base(f) == "pyuno.so" {
				out = filepath.Dir(f)
			}
			return nil
		},
	)
	if err != nil {
		return ""
	}
	return out
}

func (c *converterWithLOKit) init() error {
	if c.InstallationPath != "" {
		return nil
	}
	instPath, err := findLibreOfficeInstallationPath()
	if err != nil {
		return err
	}
	c.InstallationPath = instPath
	return nil
}

func (c converterWithLOKit) libreOfficeDocumentConvert(fileIn, fileOut string, type_ DocumentFormat) error {
	office, _ := libreofficekit.NewOffice(c.InstallationPath)
	document, _ := office.LoadDocument(fileIn)
	defer func() {
		office.Close()
		if document != nil {
			document.Close()
		}
	}()
	err := document.SaveAs(fileOut, string(type_), "")
	if err != nil {
		return err
	}
	return nil
}

func (c converterWithLOKit) convert(doc []byte, format DocumentFormat) ([]byte, error) {
	inFileName, err := temp_file.WriteTemporaryFile(doc)
	if err != nil {
		return nil, err
	}
	outFileName := filepath.Join(os.TempDir(), inFileName+"."+string(format))
	defer func() {
		temp_file.RemoveTemporaryFiles([]string{inFileName, outFileName})
	}()
	err = c.libreOfficeDocumentConvert(inFileName, outFileName, format)
	if err != nil {
		return nil, err
	}
	convertedDocumentBytes, err := ioutil.ReadFile(outFileName)
	if err != nil {
		return nil, err
	}
	return convertedDocumentBytes, nil
}

func (c *converterWithLOKit) Convert(doc []byte, format DocumentFormat) ([]byte, error) {
	err := c.init()
	if err != nil {
		return nil, err
	}
	convertedDoc, err := c.convert(doc, format)
	if err != nil {
		return nil, err
	}
	return convertedDoc, nil
}
