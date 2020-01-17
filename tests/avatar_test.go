package tests

import (
	"goloveu/utils/avatar"
	"image"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{uid: 1234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := avatar.Generate(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateAvatar(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avatar.GenerateAvatar(tt.args.uid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateAvatar() = %v, want %v", got, tt.want)
			}
		})
	}
}
