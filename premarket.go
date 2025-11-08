// https://tushare.pro/document/2?doc_id=329

package tushare

import "time"

// PreMarket 盘前数据
type PreMarket struct {
	Code       string    // 股票代码
	Date       time.Time // 交易日期
	TotalShare float64   // 总股本(万股)
	FloatShare float64   // 流通股本(万股)
	PreClose   float64   // 昨收盘价
	UpLimit    float64   // 涨停价
	DownLimit  float64   // 跌停价
}

type preMarketOpt func(Args)

// PreMarket 获取盘前数据
func (cli *Client) PreMarket(opts ...preMarketOpt) ([]PreMarket, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("stk_premarket", args,
		[]string{"ts_code", "trade_date", "total_share", "float_share", "pre_close", "up_limit", "down_limit"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxDate, idxTotalShare, idxFloatShare, idxPreClose, idxUpLimit, idxDownLimit int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "trade_date":
			idxDate = i
		case "total_share":
			idxTotalShare = i
		case "float_share":
			idxFloatShare = i
		case "pre_close":
			idxPreClose = i
		case "up_limit":
			idxUpLimit = i
		case "down_limit":
			idxDownLimit = i
		}
	}
	items := make([]PreMarket, len(data))
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = PreMarket{
			Code:       item[idxCode].(string),
			Date:       date,
			TotalShare: item[idxTotalShare].(float64),
			FloatShare: item[idxFloatShare].(float64),
			PreClose:   item[idxPreClose].(float64),
			UpLimit:    item[idxUpLimit].(float64),
			DownLimit:  item[idxDownLimit].(float64),
		}
	}
	return items, nil
}

// WithPreMarketCode 按股票代码查询
func WithPreMarketCode(code string) preMarketOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithPreMarketDate 按交易日期查询
func WithPreMarketDate(date time.Time) preMarketOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithPreMarketDateRange 按开始交易日期查询
func WithPreMarketDateRange(begin, end time.Time) preMarketOpt {
	return func(args Args) {
		args["start_date"] = begin.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
