package pdf_to_image_converter

import (
	"fmt"
	"github.com/Eitol/document_color_meter/api_front/core/test_util"
	"github.com/gojetpack/pyos"
	"log"
	"testing"
)

func Test_localConverter_Convert(t *testing.T) {
	type args struct {
		file   []byte
		format imageFormat
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid pdf",
			args: args{
				file:   test_util.GetTestFile(test_util.ValidPDFFile),
				format: ImageFormatJPEG,
			},
			wantErr: false,
		},
		{
			name: "valid pdf big file",
			args: args{
				file:   test_util.GetTestFile(test_util.ValidPDFFileBig),
				format: ImageFormatJPEG,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := localConverter{}
			got, err := client.Convert(tt.args.file, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, page := range got {
				tmpFile, err := pyos.File.WriteTemp(page)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Print(tmpFile)
			}
		})
	}
}
