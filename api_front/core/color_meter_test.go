package core

import (
	"github.com/Eitol/document_color_meter/api_front/core/test_util"
	"reflect"
	"testing"
)

func Test_measureColorOfPage(t *testing.T) {
	type args struct {
		page    []byte
		options Options
	}
	tests := []struct {
		name    string
		args    args
		want    *PageColorMeasureResult
		wantErr bool
	}{
		{
			name: "valid image",
			args: args{
				page: test_util.GetTestFile(test_util.ValidImage),
				options: Options{
					WhiteThreshold: 240 + 240 + 240,
					BlackThreshold: 5 + 5 + 5,
					GrayThreshold:  15,
				},
			},
			want: &PageColorMeasureResult{
				TotalColorPercentage: 0,
				PercentageByColor:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := measureColorOfPage(tt.args.page, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("measureColorOfPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("measureColorOfPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
