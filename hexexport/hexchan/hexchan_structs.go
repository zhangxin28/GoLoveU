package hexchan

// HexGenerateResult 表示根据EXCEL数据源生成的需要计算的数据量
type HexGenerateResult struct {
	File string
	HexData map[string][]string
}