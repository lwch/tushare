package tushare

import "time"

// https://tushare.pro/document/2?doc_id=259

// ThsIndex 同花顺行业指数数据
type ThsIndex struct {
	Code     string    // 指数代码
	Name     string    // 指数名称
	Count    int       // 成分股数量
	Exchange string    // 交易所代码
	Date     time.Time // 上市日期
}

type thsIndexOpt func(Args)

// ThsIndex 获取同花顺行业指数数据
func (cli *Client) ThsIndex(opts ...thsIndexOpt) ([]ThsIndex, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("ths_index", args,
		[]string{"ts_code", "name", "count", "exchange", "list_date"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxName, idxCount, idxExchange, idxDate int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "name":
			idxName = i
		case "count":
			idxCount = i
		case "exchange":
			idxExchange = i
		case "list_date":
			idxDate = i
		}
	}
	items := make([]ThsIndex, len(data))
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = ThsIndex{
			Code:     item[idxCode].(string),
			Name:     item[idxName].(string),
			Count:    int(item[idxCount].(float64)),
			Exchange: item[idxExchange].(string),
			Date:     date,
		}
	}
	return items, nil
}

// WithThsIndexCode 同花顺行业指数代码参数
func WithThsIndexCode(code string) thsIndexOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

type thsIndexExchange string

const ThsIndexExchangeA thsIndexExchange = "A"   // A股
const ThsIndexExchangeHK thsIndexExchange = "HK" // 港股
const ThsIndexExchangeUS thsIndexExchange = "US" // 美股

// WithThsIndexExchange 交易所参数
func WithThsIndexExchange(exchange thsIndexExchange) thsIndexOpt {
	return func(args Args) {
		args["exchange"] = exchange
	}
}

// https://tushare.pro/document/2?doc_id=261

// ThsMember 同花顺行业成分股
type ThsMember struct {
	IndexCode string // 指数代码
	StockCode string // 成分股代码
	StockName string // 成分股名称
}

type thsMemberOpt func(Args)

// ThsMember 获取同花顺行业成分股
func (cli *Client) ThsMember(opts ...thsMemberOpt) ([]ThsMember, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("ths_member", args,
		[]string{"ts_code", "con_code", "con_name"})
	if err != nil {
		return nil, err
	}
	var idxIndexCode, idxStockCode, idxStockName int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxIndexCode = i
		case "con_code":
			idxStockCode = i
		case "con_name":
			idxStockName = i
		}
	}
	items := make([]ThsMember, len(data))
	for i, item := range data {
		items[i] = ThsMember{
			IndexCode: item[idxIndexCode].(string),
			StockCode: item[idxStockCode].(string),
			StockName: item[idxStockName].(string),
		}
	}
	return items, nil
}

// WithThsMemberIndexCode 同花顺行业指数代码参数
func WithThsMemberIndexCode(code string) thsMemberOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithThsMemberStockCode 成分股代码参数
func WithThsMemberStockCode(code string) thsMemberOpt {
	return func(args Args) {
		args["con_code"] = code
	}
}
