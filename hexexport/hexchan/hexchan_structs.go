package hexchan

// DlChanStruct 表示根据EXCEL数据源生成的需要计算的数据量
type DlChanStruct struct {
	File string
	DataSourceLength int
	HexData map[string][]string
}