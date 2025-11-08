// https://tushare.pro/document/2?doc_id=25

package tushare

// StockBasic 股票基本信息
type StockBasic struct {
	Code     string // 股票代码
	Symbol   string // 股票代码(无后缀)
	Name     string // 股票名称
	Area     string // 地域
	Industry string // 行业
}

type basicOpt func(Args)

// StockBasic 获取股票列表
func (cli *Client) StockBasic(opts ...basicOpt) ([]StockBasic, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("stock_basic", args,
		[]string{"ts_code", "symbol", "name", "area", "industry"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxSymbol, idxName, idxArea, idxIndustry int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "symbol":
			idxSymbol = i
		case "name":
			idxName = i
		case "area":
			idxArea = i
		case "industry":
			idxIndustry = i
		}
	}
	items := make([]StockBasic, len(data))
	toString := func(v any) string {
		if v == nil {
			return ""
		}
		return v.(string)
	}
	for i, item := range data {
		items[i] = StockBasic{
			Code:     toString(item[idxCode]),
			Symbol:   toString(item[idxSymbol]),
			Name:     toString(item[idxName]),
			Area:     toString(item[idxArea]),
			Industry: toString(item[idxIndustry]),
		}
	}
	return items, nil
}

// WithBasicCode 按股票代码查询
func WithBasicCode(symbol string) basicOpt {
	return func(args Args) {
		args["ts_code"] = symbol
	}
}

// WithBasicName 按股票名称查询
func WithBasicName(name string) basicOpt {
	return func(args Args) {
		args["name"] = name
	}
}

type basicMarket string

const BasicMarket主板 = "主板"
const BasicMarket创业板 = "创业板"
const BasicMarket科创板 = "科创板"
const BasicMarket北交所 = "北交所"

// WithBasicMarket 按市场类型查询
func WithBasicMarket(market basicMarket) basicOpt {
	return func(args Args) {
		args["market"] = market
	}
}

type basicStatus string

const BasicStatusL basicStatus = "L" // 上市
const BasicStatusD basicStatus = "D" // 退市
const BasicStatusP basicStatus = "P" // 暂停上市

// WithBasicStatus 按股票状态查询(上市/退市/暂停上市)
func WithBasicStatus(status basicStatus) basicOpt {
	return func(args Args) {
		args["list_status"] = status
	}
}

type basicExchange string

const BasicExchangeSSE basicExchange = "SSE"   // 上交所
const BasicExchangeSZSE basicExchange = "SZSE" // 深交所
const BasicExchangeBSE basicExchange = "BSE"   // 北交所

// WithBasicExchange 按交易所查询
func WithBasicExchange(exchange basicExchange) basicOpt {
	return func(args Args) {
		args["exchange"] = exchange
	}
}
