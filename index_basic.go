// https://tushare.pro/document/2?doc_id=94

package tushare

// IndexBasic 指数基本信息
type IndexBasic struct {
	Code     string        // 指数代码
	Name     string        // 指数名称
	FullName string        // 指数全称
	Market   indexMarket   // 市场
	Category indexCategory // 分类
}

type indexBasicOpt func(Args)

// IndexBasic 获取指数列表
func (cli *Client) IndexBasic(opts ...indexBasicOpt) ([]IndexBasic, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("index_basic", args,
		[]string{"ts_code", "name", "fullname", "market", "category"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxName, idxFullName, idxMarket, idxCategory int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "name":
			idxName = i
		case "fullname":
			idxFullName = i
		case "market":
			idxMarket = i
		case "category":
			idxCategory = i
		}
	}
	items := make([]IndexBasic, len(data))
	toString := func(v any) string {
		if v == nil {
			return ""
		}
		return v.(string)
	}
	for i, item := range data {
		items[i] = IndexBasic{
			Code:     toString(item[idxCode]),
			Name:     toString(item[idxName]),
			FullName: toString(item[idxFullName]),
			Market:   indexMarket(toString(item[idxMarket])),
			Category: indexCategory(toString(item[idxCategory])),
		}
	}
	return items, nil
}

type indexMarket string

const (
	IndexMarketMSCI  indexMarket = "MSCI" // MSCI
	IndexMarketCSI   indexMarket = "CSI"  // 中证指数
	IndexMarketSSE   indexMarket = "SSE"  // 上交所
	IndexMarketSZSE  indexMarket = "SZSE" // 深交所
	IndexMarketCICC  indexMarket = "CICC" // 中金所
	IndexMarketSW    indexMarket = "SW"   // 申万行业
	IndexMarketOTHER indexMarket = "OTH"  // 其他
)

type indexCategory string

const (
	IndexCategory行业指数   indexCategory = "行业指数"
	IndexCategory一级行业指数 indexCategory = "一级行业指数"
	IndexCategory二级行业指数 indexCategory = "二级行业指数"
	IndexCategory三级行业指数 indexCategory = "三级行业指数"
	IndexCategory四级行业指数 indexCategory = "四级行业指数"
	IndexCategory综合指数   indexCategory = "综合指数"
	IndexCategory主题指数   indexCategory = "主题指数"
	IndexCategory策略指数   indexCategory = "策略指数"
	IndexCategory规模指数   indexCategory = "规模指数"
	IndexCategory风格指数   indexCategory = "风格指数"
	IndexCategory其他指数   indexCategory = "其他"
)

// WithIndexBasicCode 按指数代码查询
func WithIndexBasicCode(symbol string) indexBasicOpt {
	return func(args Args) {
		args["ts_code"] = symbol
	}
}

// WithIndexBasicName 按指数名称查询
func WithIndexBasicName(name string) indexBasicOpt {
	return func(args Args) {
		args["name"] = name
	}
}

// WithIndexBasicMarket 按市场类型查询
func WithIndexBasicMarket(market indexMarket) indexBasicOpt {
	return func(args Args) {
		args["market"] = market
	}
}

// WithIndexBasicCategory 按指数分类查询
func WithIndexBasicCategory(category indexCategory) indexBasicOpt {
	return func(args Args) {
		args["category"] = category
	}
}
