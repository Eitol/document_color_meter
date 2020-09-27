package office_converter

import (
	"github.com/Eitol/document_color_meter/api_front/core/test_util"
	"io/ioutil"
	"testing"
)

func Test_converterWithLOKit_Convert(t *testing.T) {
	type fields struct {
		InstallationPath string
	}
	type args struct {
		doc    []byte
		format DocumentFormat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "convert doc 1997",
			fields: fields{
				InstallationPath: "",
			},
			args: args{
				doc:    test_util.GetTestFile(test_util.ValidWord1997File),
				format: PDF,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &converterWithLOKit{
				InstallationPath: tt.fields.InstallationPath,
			}
			got, err := c.Convert(tt.args.doc, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = ioutil.WriteFile("test."+string(tt.args.format), got, 0777)
			if err != nil {
				t.Errorf("Convert() writing the file")
			}
		})
	}
}
