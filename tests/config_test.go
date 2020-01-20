package tests

import (
	"testing"
	"goloveu/bbs-go-server/config"
	"fmt"
)

func TestYamlFileInitConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name :"test bbs-go.yaml",
			args:args{
				filename: "../config/bbs-go.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitConfig(tt.args.filename)
			fmt.Printf("Config Env = %v\n", config.Conf.Env)
		})
	}
}
