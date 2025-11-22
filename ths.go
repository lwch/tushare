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
	Type     string    // 指数类型
}

type thsIndexOpt func(Args)

// ThsIndex 获取同花顺行业指数数据
func (cli *Client) ThsIndex(opts ...thsIndexOpt) ([]ThsIndex, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("ths_index", args,
		[]string{"ts_code", "name", "count", "exchange", "list_date", "type"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxName, idxCount, idxExchange, idxDate, idxType int
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
		case "type":
			idxType = i
		}
	}
	toString := func(v any) string {
		if v == nil {
			return ""
		}
		return v.(string)
	}
	toInt := func(v any) int {
		if v == nil {
			return 0
		}
		return int(v.(float64))
	}
	items := make([]ThsIndex, len(data))
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", toString(item[idxDate]), time.Local)
		items[i] = ThsIndex{
			Code:     item[idxCode].(string),
			Name:     item[idxName].(string),
			Count:    toInt(item[idxCount]),
			Exchange: item[idxExchange].(string),
			Date:     date,
			Type:     item[idxType].(string),
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

type thsIndexType string

const ThsIndexTypeN thsIndexType = "N"   // 概念指数
const ThsIndexTypeI thsIndexType = "I"   // 行业指数
const ThsIndexTypeR thsIndexType = "R"   // 地域指数
const ThsIndexTypeS thsIndexType = "S"   // 同花顺特色指数
const ThsIndexTypeST thsIndexType = "ST" // 同花顺风格指数
const ThsIndexTypeTH thsIndexType = "TH" // 同花顺主题指数
const ThsIndexTypeBB thsIndexType = "BB" // 同花顺宽基指数

// WithThsIndexType 指数类型参数
func WithThsIndexType(indexType thsIndexType) thsIndexOpt {
	return func(args Args) {
		args["type"] = indexType
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

type thsDailyOpt func(Args)

// ThsDaily 同花顺行业指数日线行情
func (cli *Client) ThsDaily(opts ...thsDailyOpt) ([]DailyTick, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("ths_daily", args, []string{
		"ts_code", "trade_date",
		"open", "high", "low", "close",
		"pre_close", "change", "pct_change",
		"vol"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxDate int
	var idxOpen, idxHigh, idxLow, idxClose int
	var idxPreClose, idxChange, idxPctChg int
	var idxVolume int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "trade_date":
			idxDate = i
		case "open":
			idxOpen = i
		case "high":
			idxHigh = i
		case "low":
			idxLow = i
		case "close":
			idxClose = i
		case "pre_close":
			idxPreClose = i
		case "change":
			idxChange = i
		case "pct_change":
			idxPctChg = i
		case "vol":
			idxVolume = i
		}
	}
	items := make([]DailyTick, len(data))
	toFloat := func(v any) float64 {
		if v == nil {
			return 0
		}
		return v.(float64)
	}
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = DailyTick{
			Tick: Tick{
				Code:   item[idxCode].(string),
				Time:   date,
				Open:   toFloat(item[idxOpen]),
				High:   toFloat(item[idxHigh]),
				Low:    toFloat(item[idxLow]),
				Close:  toFloat(item[idxClose]),
				Volume: toFloat(item[idxVolume]),
			},
			PreClose: toFloat(item[idxPreClose]),
			Change:   toFloat(item[idxChange]),
			PctChg:   toFloat(item[idxPctChg]),
		}
	}
	return items, nil
}

// WithThsDailyCode 同花顺行业指数代码参数
func WithThsDailyCode(code string) thsDailyOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithThsDailyDate 按日期查询
func WithThsDailyDate(date time.Time) thsDailyOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithThsDailyDateRange 按日期范围查询
func WithThsDailyDateRange(start, end time.Time) thsDailyOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
