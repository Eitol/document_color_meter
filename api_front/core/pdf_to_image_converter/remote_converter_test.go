package pdf_to_image_converter

import (
	"fmt"
	"github.com/Eitol/document_color_meter/api_front/core/test_util"
	"github.com/gojetpack/pyos"
	"log"
	"testing"
	"time"
)

func Test_remoteConverter_Convert(t *testing.T) {
	type fields struct {
		ServerAddress string
		Timeout       time.Duration
	}
	type args struct {
		file   []byte
		format imageFormat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid pdf",
			fields: fields{
				ServerAddress: "localhost:50052",
				Timeout:       time.Second * 5,
			},
			args: args{
				file:   test_util.GetTestFile(test_util.ValidPDFFile),
				format: ImageFormatJPEG,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := remoteConverter{
				ServerAddress: tt.fields.ServerAddress,
				Timeout:       tt.fields.Timeout,
			}
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
