package tools

import (
	"bytes"
	"testing"
)

// use test to generate data.txt
func TestGenDistWithPointToFile(t *testing.T) {
	type args struct {
		kmin int
		kmax int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test1", args{1, 13}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w bytes.Buffer
			if err := GenDistWithPointToFile(
				&w,
				tt.args.kmin,
				tt.args.kmax,
			); (err != nil) != tt.wantErr {
				t.Errorf(
					"genDistWithPointToFile() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
