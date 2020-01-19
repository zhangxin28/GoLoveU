package tests

import (
	"fmt"
	"goloveu/utils"
	"testing"
)

func TestTimeParse(t *testing.T) {
	type args struct {
		value  string
		layout string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test Time Parse with format : yyyyMMdd",
			args: args{
				value:  utils.FMT_DATE_TIME,
				layout: utils.FMT_DATE_TIME,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.TimeParse(tt.args.value, tt.args.layout)
			if err != nil {
				t.Errorf("utils.TimeParse() Time = %v, Layout = %v", tt.args.value, tt.args.layout)
			}
			fmt.Printf("utils.TimeParse() Time = %v, Layout = %v, Result = %v\n", tt.args.value, tt.args.layout, result)
			timeUnix := utils.Timestamp(result)
			fmt.Printf("Result Millisecond Time Stamp = %v\n", timeUnix)
			timeFromUnix := utils.TimeFromTimestamp(timeUnix)
			fmt.Printf("TimeFromUnix = %v\n", timeFromUnix)
		})
	}
}
