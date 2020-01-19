package tests

import (
	"goloveu/utils"
	"goloveu/utils/common"
	"reflect"
	"testing"
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

func TestGetFiles(t *testing.T) {
	type args struct {
		filePrefix string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFiles := common.GetFiles(tt.args.filePrefix); !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("GetFiles() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func TestCreateNewFolder(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.CreateNewFolder(tt.args.path)
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

func TestGetArrayIndex(t *testing.T) {
	type args struct {
		value       interface{}
		compareFunc common.CompareFunc
		values      []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.GetArrayIndex(tt.args.value, tt.args.compareFunc, tt.args.values...); got != tt.want {
				t.Errorf("GetArrayIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapKeys(t *testing.T) {
	type args struct {
		m map[interface{}][]interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantKeys []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeys := common.GetMapKeys(tt.args.m); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("GetMapKeys() = %v, want %v", gotKeys, tt.wantKeys)
			}
		})
	}
}

func TestGetSafeValue(t *testing.T) {
	type args struct {
		f func() interface{}
	}
	tests := []struct {
		name          string
		args          args
		wantSafeValue interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSafeValue := common.GetSafeValue(tt.args.f); !reflect.DeepEqual(gotSafeValue, tt.wantSafeValue) {
				t.Errorf("GetSafeValue() = %v, want %v", gotSafeValue, tt.wantSafeValue)
			}
		})
	}
}

func TestDoSafeSave(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.DoSafeSave(tt.args.f)
		})
	}
}

func TestPrintStack(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.PrintStack()
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
