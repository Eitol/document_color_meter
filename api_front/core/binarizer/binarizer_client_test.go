package binarizer

import (
	"github.com/Eitol/document_color_meter/api_front/core/test_util"
	"testing"
	"time"
)

func Test_remoteBinarizer_Binarize(t *testing.T) {
	type fields struct {
		ServerAddress string
		Timeout       time.Duration
	}
	type args struct {
		images             [][]byte
		binarizationMethod BinarizationMethod
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid images",
			fields: fields{
				ServerAddress: "localhost:50053",
				Timeout:       time.Second * 5,
			},
			args: args{
				images:             [][]byte{test_util.GetTestFile(test_util.ValidImage)},
				binarizationMethod: SAUVOLA,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := remoteBinarizer{
				ServerAddress: tt.fields.ServerAddress,
				Timeout:       tt.fields.Timeout,
			}
			got, _, err := client.Binarize(tt.args.images, tt.args.binarizationMethod)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binarize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf(" Binarize() empty response")
				return
			}
		})
	}
}
