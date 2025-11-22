// https://tushare.pro/document/2?doc_id=95

package tushare

import "time"

type indexDailyOpt func(Args)

// IndexDaily 指数日线行情
func (cli *Client) IndexDaily(opts ...indexDailyOpt) ([]DailyTick, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("index_daily", args, []string{
		"ts_code", "trade_date",
		"open", "high", "low", "close",
		"pre_close", "change", "pct_chg",
		"vol", "amount"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxDate int
	var idxOpen, idxHigh, idxLow, idxClose int
	var idxPreClose, idxChange, idxPctChg int
	var idxVolume, idxAmount int
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
		case "pct_chg":
			idxPctChg = i
		case "vol":
			idxVolume = i
		case "amount":
			idxAmount = i
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
				Code:     item[idxCode].(string),
				Time:     date,
				Open:     item[idxOpen].(float64),
				High:     item[idxHigh].(float64),
				Low:      item[idxLow].(float64),
				Close:    item[idxClose].(float64),
				Volume:   item[idxVolume].(float64),
				Turnover: toFloat(item[idxAmount]),
			},
			PreClose: toFloat(item[idxPreClose]),
			Change:   toFloat(item[idxChange]),
			PctChg:   toFloat(item[idxPctChg]),
		}
	}
	return items, nil
}

// WithIndexDailyCode 按指数代码查询
func WithIndexDailyCode(code string) indexDailyOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithIndexDailyDate 按交易日期查询
func WithIndexDailyDate(date time.Time) indexDailyOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithIndexDailyDateRange 按交易日期范围查询
func WithIndexDailyDateRange(start, end time.Time) indexDailyOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
