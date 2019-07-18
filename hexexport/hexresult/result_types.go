package hexresult

// BasicProductInfo 商品基本资料
type BasicProductInfo struct{
	
	// "商品编码"
	Sku string
	// SBI
	SBI string
	// 商品名称
	ProductCNName string
	// 商品英文名称
	ProductENName string
	// 核算单位
	HeSuanUnit string
	// 税别
	Tax string
	// 是否增值税商品
	IsZengZhiTaxProduct string
	// 是否可主配
	IsZhuPei string
	// 商品类别编码
	ProductTypeCode string
	// 存储类型
	StoreType string
	// 行类型
	RowType string
	// 产品属性
	ProductAttr string
	// 用友凭证科目
	YongYouPinZhengClass string
	// BOM类型
	BomType string
	// 适用开始日
	UseStartDate string
	// 条码二
	CodeTwo string
	// 保质期（天）
	KeepDay string
}

// AreaProductAttr 区域商品属性
type AreaProductAttr{
	// -1:   "区域商品属性维护",
	// 1000: "属性区域编码",
	// 0:    "商品编码",
	// 900:  "生效开始日期",
	// 901:  "生效结束日期",
	// 17:   "是否可调拨",
	// 18:   "是否可调整",
	// 19:   "是否可盘点",
}