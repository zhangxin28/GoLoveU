package hexresult

// BasicProductInfo 商品基本资料
type BasicProductInfo struct {

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
type AreaProductAttr struct {

	// 属性区域编码
	PropAreaCode string
	// 商品编码
	Sku string
	// 生效开始日期
	ActiveBeginDate string
	// 生效结束日期
	ActiveEndDate string
	// 是否可调拨
	IsDiaobo string
	// 是否可调整
	IsTiaoZheng string
	// 是否可盘点
	IsPanDian string
}

// AreaProductByPriceInfo 区域商品采购价格
type AreaProductByPriceInfo struct {

	// 采购区域编码
	BuyAreaCode string
	// 供应商编码
	SupplierCode string
	// 商品编码
	Sku string
	// 生效开始日期
	ActiveBeginDate string
	// 生效结束日期
	ActiveEndDate string
	// 采购价格(含税)
	BuyPriceWithTax string
	//应计税率
	TaxRate string
	// 是否主供应商
	IsFirstSupplier string
	// 是否可采购
	IsCanBy string
}

// AreaProductPriceInfo 区域商品售价
type AreaProductPriceInfo struct {

	// 市场区域编码
	MarketAreaCode string
	// 商品编码
	Sku string
	// 生效开始日期
	ActiveBeginDate string
	// 生效结束日期
	ActiveEndDate string
	// 销售价格
	PriceToSale string
	// 是否可销售
	IsCanSale string
}

// AreaProductDeliverRule 区域商品配送规则
type AreaProductDeliverRule struct {

	// 配送区域编码
	DevelieryAreaCode string
	// 商品编码
	Sku string
	// 生效开始日期
	ActiveBeginDate string
	// 生效结束日期
	ActiveEndDate string
	// 物流模式
	DeliveryMode string
	// 配送中心
	DeliveryCenter string
	// 是否可订货
	IsCanDevliery string
	// 周一至周日的送货状况
	DaysDeliveryStatus string
	// 计划到货天数
	PlanArrivedDay string
	// 最小订量
	MinOrder string
	// 最大订量
	MaxOrder string
	// 递增订量
	StepOrder string
}

// AreaProductBom 区域商品配方
type AreaProductBom struct {

	// BOM区域
	BOMArea string
	// BOM商品编码
	BOMProductCode string
	// 生效开始时间
	ActiveBeginDate string
	// 生效结束时间
	ActiveEndDate string
	// 原料商品编码
	MeteriaProductCode string
	// 配方数量
	FormCount string
}



