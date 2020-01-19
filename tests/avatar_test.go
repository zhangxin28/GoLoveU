package tests

import (
	"fmt"
	"goloveu/utils/avatar"
	"goloveu/utils/common"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test generate file with uid = 1234",
			args: args{uid: 1234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pngBytes, err := avatar.Generate(tt.args.uid)
			if err != nil {
				t.Errorf("Generate() error = %v", err)
			}
			filePath := fmt.Sprintf("%v.png", tt.args.uid)
			fi, err := os.Create(filePath)
			if err != nil {
				t.Errorf("create file %v failed\n", filePath)
			}
			fi.Write(pngBytes)
			fi.Close()
			fmt.Printf("create file %v success\n", filePath)

			fileExistFlag := common.CheckPathExists(filePath)
			if !fileExistFlag {
				t.Errorf("Generated file not existsed\n")
			}
			err = os.Remove(filePath)
			if err != nil {
				t.Errorf("delete file %v failed\n", filePath)
			}
			fmt.Printf("delete file %v success\n", filePath)
		})
	}
}
