package temp_file

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteTemporaryFile(doc []byte) (string, error) {
	inFile, err := ioutil.TempFile(os.TempDir(), "test_")
	if err != nil {
		return "", err
	}
	_, err = inFile.Write(doc)
	if err != nil {
		return "", err
	}
	err = inFile.Close()
	if err != nil {
		return "", err
	}
	return inFile.Name(), nil
}

func RemoveTemporaryFiles(files []string) {
	for _, fileName := range files {
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}
