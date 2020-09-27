package temp_file

import (
	"github.com/Eitol/document_color_meter/api_front/core/test_util"
	"github.com/gojetpack/pyos"
	"testing"
)

func Test_writeTemporaryFile(t *testing.T) {
	type args struct {
		doc []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid word 1997",
			args: args{
				doc: test_util.GetTestFile(test_util.ValidWord1997File),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, err := WriteTemporaryFile(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteTemporaryFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !pyos.Path.Exist(name) {
				t.Errorf("WriteTemporaryFile() the file was not created")
			}
			RemoveTemporaryFiles([]string{name})
		})
	}
}
