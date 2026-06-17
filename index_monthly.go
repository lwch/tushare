// https://tushare.pro/document/2?doc_id=172

package tushare

import "time"

type indexMonthlyOpt func(Args)

// IndexMonthly 指数月线行情
func (cli *Client) IndexMonthly(code string, opts ...indexMonthlyOpt) ([]DailyTick, error) {
	args := make(Args)
	args["ts_code"] = code
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("index_monthly", args, []string{
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
				Open:     toFloat(item[idxOpen]),
				High:     toFloat(item[idxHigh]),
				Low:      toFloat(item[idxLow]),
				Close:    toFloat(item[idxClose]),
				Volume:   toFloat(item[idxVolume]),
				Turnover: toFloat(item[idxAmount]),
			},
			PreClose: toFloat(item[idxPreClose]),
			Change:   toFloat(item[idxChange]),
			PctChg:   toFloat(item[idxPctChg]),
		}
	}
	return items, nil
}

// WithIndexMonthlyDate 按交易日期查询
func WithIndexMonthlyDate(date time.Time) indexMonthlyOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithIndexMonthlyDateRange 按交易日期范围查询
func WithIndexMonthlyDateRange(start, end time.Time) indexMonthlyOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
