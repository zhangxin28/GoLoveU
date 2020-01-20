package tests

import (
	"goloveu/utils"
	"goloveu/utils/common"
	"goloveu/utils/email"
	"testing"
	"fmt"
)

func TestCheckError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.CheckError(tt.args.err)
		})
	}
}

func TestWaitUserEnterKeyToExit(t *testing.T) {
	type args struct {
		exit bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.WaitUserEnterKeyToExit(tt.args.exit)
		})
	}
}

func TestGetFileName(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name             string
		args             args
		wantFileName     string
		wantFileOnlyName string
		wantFileSuffix   string
	}{
		// TODO: Add test cases.
		{
			name:             "test file go.mod",
			args:             args{file: "go.mod"},
			wantFileName:     "go.mod",
			wantFileOnlyName: "go",
			wantFileSuffix:   ".mod",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileName, gotFileOnlyName, gotFileSuffix := common.GetFileName(tt.args.file)
			if gotFileName != tt.wantFileName {
				t.Errorf("GetFileName() gotFileName = %v, want %v", gotFileName, tt.wantFileName)
			}
			if gotFileOnlyName != tt.wantFileOnlyName {
				t.Errorf("GetFileName() gotFileOnlyName = %v, want %v", gotFileOnlyName, tt.wantFileOnlyName)
			}
			if gotFileSuffix != tt.wantFileSuffix {
				t.Errorf("GetFileName() gotFileSuffix = %v, want %v", gotFileSuffix, tt.wantFileSuffix)
			}
		})
	}
}

func TestCheckPathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test file exists for file go.mod",
			args: args{path: "go.mod"},
			want: true,
		},
		{
			name: "test file not exists for file xxxnotexists.go",
			args: args{path: "xxxnotexists.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.CheckPathExists(tt.args.path); got != tt.want {
				t.Errorf("Case %s , CheckPathExists() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPasswordEncryDecrp(t *testing.T) {
	type args struct {
		originalPwd string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test password logic",
			args: args{originalPwd: "goloveu"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodePwd := utils.EncodePassword(tt.args.originalPwd)
			if result := utils.ValidatePassword(encodePwd, tt.args.originalPwd); result != true {
				t.Errorf("PasswordEncodeValidate Logic OriginalPwd = %v, EncodePwd %v", tt.args.originalPwd, encodePwd)
			}
		})
	}
}


func TestBuildEmailTemplate(t *testing.T) {
	type args struct {
		title        string
		content      string
		quoteContent string
		url          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:"test building email template with no error",
			args: args{
				title: "test title",
				content: "oh, this is a test",
				quoteContent: "oh yes, this is a test",
				url:"https://www.emailtest.com/testurl",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := email.BuildEmailTemplate(tt.args.title, tt.args.content, tt.args.quoteContent, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildEmailTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("BuildEmailTemplate() no any template")
			}
			fmt.Println(got)
		})
	}
}
