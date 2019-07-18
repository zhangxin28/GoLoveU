package utils

type CompareFunc func(interface{}, interface{}) bool

type DispatchHandled struct{
	AreaSheetName string
	AreaName string
	Data [][]string
}
