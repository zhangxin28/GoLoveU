package tests

import (
	"goloveu/goinactioncode"
	"testing"
)

func TestGoInActionExecuters(t *testing.T) {
	tests := []struct {
		name        string
		excutername string
	}{
		// TODO: Add test cases.
		{
			name:        "test go in actin gorountines",
			excutername: "grShow", //"MatchersSearch","grShow"
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goinactioncode.RunSample(tt.excutername)
		})
	}
}
