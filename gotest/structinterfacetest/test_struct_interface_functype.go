package structinterfacetest

import (
	"fmt"
)

type requestStruct struct {
}

type IHandlerRequest interface {
	DoHandle(string)
}

type FuncDoHandle func(string)

func (funcDoH FuncDoHandle) DoHandle(message string) {
	funcDoH(message)
}

func (rs *requestStruct) DoHandle(message string) {
	rs.doHandleDefault(message)
}

func (rs *requestStruct) doHandleDefault(message string) {
	fmt.Println("Default Handle Message ", message)
}

func (rs *requestStruct) doHandleFunc(message string, ihr IHandlerRequest) {
	rs.doHandleDefault(message)
	ihr.DoHandle(message)
}

func (rs *requestStruct) HandleFunc(message string, funcDH func(string)) {
	if funcDH != nil {
		rs.doHandleFunc(message, FuncDoHandle(funcDH))
	} else {
		rs.doHandleDefault(message)
	}
}

var DefaultRequestStruct = &requestStruct{}

func HandleFunc(message string, funcDoHandle func(string)) {
	DefaultRequestStruct.HandleFunc(message, funcDoHandle)
}
