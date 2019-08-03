package hexsource

import (
	"starbucks/tools/utils/common"
)

var vendorsMap = map[string]string{
	"929341": "00000000000040001201", //	北京美大咖啡有限公司
	"928400": "00000000000040002001", //	上海普进贸易有限公司
	"963212": "00000000000040002401", //	上海伊诺尔印务有限公司
	"927903": "00000000000040003301", //	德百（上海）包装贸易有限公司
	"928339": "00000000000040004001", //	上海民芳装潢橱柜有限公司
	"928340": "00000000000040004901", //	上海台新食品机械有限公司
	"928347": "00000000000040005001", //	仲野国际贸易上海有限公司
	"963211": "00000000000040007801", //	上海多明工贸有限公司
	"927917": "00000000000020000202", //	美夏国际贸易（上海）有限公司
	"928344": "00000000000020001801", //	昆山统一企业食品有限公司
	"939985": "00000000000020003001", //	上海恒敬贸易有限公司
	"936915": "00000000000020003201", //	西诺迪斯食品（上海）有限公司cd
	"929298": "00000000000020004701", //	维益食品（苏州）有限公司
	"943773": "00000000000020006601", //	捷成饮料（中国）有限公司上海分公司
	"927921": "00000000000040000101", //	上海开展贸易有限公司
	"963411": "00000000000042000004", //	上海颐和装饰设计工程有限公司
	"963268": "00000000000042000006", //	上海泽弘酒店设备用品有限公司
	"963289": "00000000000050000004", //	苏州新浪劳务发展公司
	"928440": "00000000000050000031", //	前锦网络信息技术上海有限公司
	"963349": "00000000000050000059", //	上海日报社（工会）
	"963406": "00000000000050000063", //	上海鑫景建设工程有限公司
	"963354": "00000000000050000080", //	上海复兴明方会计师事务所有限公司
	"963367": "00000000000040011801", //	上海捷威印务技术有限公司
	"928336": "00000000000040011901", //	上海顶印图文设计有限公司
	"963266": "00000000000040012101", //	上海优悦服饰有限公司
	"963173": "00000000000040012501", //	济南海乐.西亚泽食品有限公司
	"963249": "00000000000040012601", //	上海尚标实业有限公司
	"927912": "00000000000040013201", //	上海娜鲁娃实业有限公司
	"927854": "00000000000040014101", //	永能拓食品机械贸易（上海）有限公司
	"963444": "00000000000040014201", //	王力咖啡贸易（上海）有限公司
	"948641": "00000000000040014701", //	上海百吉食品有限公司
	"963272": "00000000000040015101", //	上海资顺化工科技有限公司
	"963153": "00000000000040015401", //	广东爱丽斯包装有限公司
	"957425": "00000000000040016101", //	南京卫岗乳业有限公司
	"932764": "00000000000040018101", //	膳魔师（中国）家庭制品有限公司
	"963230": "00000000000040018401", //	上海康玺实业有限公司
	"943377": "00000000000040018501", //	苏州工业园区尚融科技有限公司
	"928334": "00000000000040018901", //	上海永熹富纸制品有限公司
	"930047": "00000000000040019001", //	深圳市福美信贸易有限公司
	"933113": "00000000000040019201", //	东莞市沃美氏贸易有限公司
	"963269": "00000000000040020101", //	上海泽欣伪钞鉴别仪制造有限公司
	"925023": "00000000000040020701", //	格来纳塑料科技(苏州)有限公司
	"940098": "00000000000040020901", //	苏州市品诺食品贸易有限公司
	"963171": "00000000000040021001", //	江苏海惠沃特包装有限公司
	"927860": "00000000000040022101", //	远东绿色包装（上海）有限公司
	"927880": "00000000000040022201", //	海宁市瑞奇包装系统有限公司
	"963227": "00000000000040022301", //	上海金鸿泰商业道具设计制作有限公司
	"963246": "00000000000040022501", //	上海綦众电子科技有限公司
	"928388": "00000000000040022901", //	东莞黑玫瑰食品有限公司
	"963178": "00000000000040023101", //	明治乳业（苏州）有限公司上海分公司
	"963263": "00000000000040023201", //	上海盈拓包装材料有限公司
	"963250": "00000000000040023301", //	上海山一企业有限公司
	"956130": "00000000000040023601", //	亚弥(深圳)国际贸易有限公司
	"927879": "00000000000040023701", //	北京安德鲁水果食品有限公司
	"963141": "00000000000040024001", //	北京艾莱发喜食品有限公司上海分公司
	"928362": "00000000000060000909", //	上海淳艺家饰有限公司
	"928352": "00000000000060000312", //	大昌华嘉中国有限公司
	"928473": "00000000000060000332", //	上海海彦照明电器工程有限公司
	"963350": "00000000000060000551", //	上海迪晖制冷工程有限公司
	"963408": "00000000000060000676", //	800033上海新欣建设发展有限公司
	"963255": "00000000000060000696", //	上海兴怡清洁服务有限公司
	"963205": "00000000000060000709", //	上海采林物业管理有限公司
	"963586": "00000000000060000728", //	上海市卢湾区企业合同信用促进会
	"963152": "00000000000060000790", //	富士通（中国）信息系统有限公司
	"963233": "00000000000040024401", //	上海乐亿塑料制品有限公司
	"957977": "00000000000040024501", //	上海集邦贸易有限公司
	"963418": "00000000000040023303", //	上海优蕾贸易有限公司
	"963221": "00000000000042000014", //	上海华氏大药房有限公司
	"963206": "00000000000050000584", //	上海晨光科力普办公用品有限公司
	"963222": "00000000000042000015", //	上海惠拓信息技术有限公司
	"963294": "00000000000042000016", //	上海东港安全印刷有限公司
	"963334": "00000000000042000017", //	耀丰（上海）道具制作有限公司
	"133452": "00000000000042000018", //	烟台海湾塑料餐具有限公司
	"963633": "00000000000040024101", //	盛诠纸业（苏州）有限公司
	"963155": "00000000000040024201", //	国誉商业（上海）有限公司
	"957194": "00000000000042000021", //	东莞铭山贸易有限公司
	"957137": "00000000000050002275", //	诺尔供应链管理(上海)有限公司
	"962268": "00000000000050001225", //	上海鑫国食品有限公司
	"963313": "00000000000042000026", //	古德智造文化创意(上海)有限公司
	"955636": "00000000000050000872", //	上海君翌机电设备有限公司
	"963373": "00000000000050001381", //	上海珑灵洁净科技有限公司
	"128571": "00000000000042000028", //	潮州市罗赞娜陶瓷有限公司
	"950680": "00000000000042000029", //	浙江新迪嘉禾食品有限公司
	"955922": "00000000000042000030", //	深圳市骏飞实业有限公司
	"955774": "00000000000042000037", //	通用磨坊贸易（上海）有限公司
	"938147": "00000000000042000038", //	百灵恬（上海）国际贸易有限公司
	"963377": "00000000000042000039", //	上海凌越实业有限公司
	"963177": "00000000000040018001", //	乐清(上海)清洁用具租赁有限公司
	"961736": "00000000000040021501", //	上海南朝印刷有限公司
	"963310": "00000000000042000001", //	大毅装饰工程(上海)有限公司
	"928355": "00000000000060000704", //	富库赛食品技术顾问（北京）有限公司
	"963170": "00000000000060000893", //	上海佰达超市用品有限公司
	"964585": "00000000000042000044", //	泰马克精密铸造（苏州）有限公司
	"966491": "00000000000042000045", //	致品食品（上海）有限公司
	"963404": "00000000000060001067", //	上海暖益机电工程有限公司
	"934292": "00000000000060001075", //	上海共拓机电工程有限公司
	"963336": "00000000000060001098", //	南京慧乾装饰有限公司
	"922473": "00000000000099999501", //	SBI-N
	"410485": "00000000000099999901", //	SBI-SPCC
	"927859": "00000000000040024701", //	潮州市庆发陶瓷有限公司
	"957618": "00000000000040023302", //	南通瑞隆农产品开发有限公司
	"963151": "00000000000050001157", //	上海馥松食品有限公司
	"963229": "00000000000042000020", //	上海君翌气体有限公司
	"941416": "00000000000042000034", //	福建百卡弗食品有限公司
	"943322": "00000000000042000033", //	青岛汉莎天厨食品有限公司
	"928571": "00000000000042000036", //	美心西饼（广州）有限公司
	"963265": "00000000000040006401", //	上海永纯环保科技有限公司
	"928338": "00000000000040023401", //	成象（上海）贸易有限公司
	"963223": "00000000000050000391", //	上海元禺艺术设计有限公司
	"963634": "00000000000050001212", //	上海鑫浩实业有限公司
	"929637": "00000000000040020002", //	金红叶纸业集团有限公司上海分公司
	"931734": "00000000000050000092", //	易百技术（上海）股份有限公司
	"955384": "00000000000042000024", //	北京万邦吉祥贸易有限公司
	"928366": "00000000000060000890", //	上海江丰家具有限公司
	"963201": "00000000000042000040", //	上海葆树实业发展有限公司
	"957589": "00000000000042000041", //	上海亚太国际蔬菜有限公司
	"957621": "00000000000042000043", //	上海台虎食品贸易有限公司
	"928061": "00000000000000928061", //	厦门佰翔空厨食品有限公司
	"959797": "00000000000000959797", //	上海橙意数码科技有限公司
	"963357": "00000000000060000980", //	上海和磬机电设备有限公司
	"955387": "00000000000000955387", //	四川佳美食品工业有限公司
	"928346": "00000000000000928346", //	中粮丰通（北京）食品有限公司
	"952883": "00000000000000952883", //	星巴克（上海）咖啡有限公司
	"970260": "00000000000000970260", //	福建同发食品集团有限公司
	"927894": "00000000000000927894", //	上海祺阜商贸有限公司
	"971938": "00000000000000971938", //	好时（中国）投资管理有限公司
	"955520": "00000000000040024301", //	上海和沁经贸有限公司
	"928345": "00000000000050001024", //	艺康（中国）投资有限公司
	"13551":  "00000000000040002501", //	奔迈(上海)国际贸易有限公司
	"928423": "00000000000040004601", //	上海扬雅国际贸易有限公司
	"929270": "00000000000040008501", //	SSSCC
	"927908": "00000000000040009501", //	林百朋包装（深圳）有限公司
	"928027": "00000000000020004801", //	江苏海创农产品开发有限公司
	"963282": "00000000000050000007", //	绍兴市人才市场
	"963630": "00000000000050000023", //	无锡住房公积金管理中心
	"963351": "00000000000040019601", //	上海地芯信息科技有限公司
	"928549": "00000000000040021101", //	上海一赋食品有限公司
	"963252": "00000000000040023001", //	上海盛优纺织服装有限公司
	"963441": "00000000000040023901", //	苏州益德鸿商贸有限公司
	"963437": "00000000000060000035", //	600012苏州华福物业管理有限公司
	"963217": "00000000000060000550", //	上海豪宇仓储设备有限公司
	"963405": "00000000000060000713", //	传典机电
}

func GetVendor(vendorId string) string {
	vendor := common.GetSafeValue(func() interface{} {
		return vendorsMap[vendorId]
	})
	return vendor.(string)
}
