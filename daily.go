package tushare

import (
	"time"
)

// Tick 行情数据
type Tick struct {
	Code     string    // 股票代码
	Time     time.Time // 时间
	Open     float64   // 开盘价
	High     float64   // 最高价
	Low      float64   // 最低价
	Close    float64   // 收盘价
	Volume   float64   // 成交量(手)
	Turnover float64   // 成交额(千元)
}

type dailyOpt func(Args)

// Daily 获取日线数据
func (cli *Client) Daily(opts ...dailyOpt) ([]Tick, error) {
	var args Args
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("daily", args, []string{"ts_code", "trade_date", "open", "high", "low", "close", "vol", "amount"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxDate, idxOpen, idxHigh, idxLow, idxClose, idxVolume, idxAmount int
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
		case "vol":
			idxVolume = i
		case "amount":
			idxAmount = i
		}
	}
	items := make([]Tick, len(data))
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = Tick{
			Code:     item[idxCode].(string),
			Time:     date,
			Open:     item[idxOpen].(float64),
			High:     item[idxHigh].(float64),
			Low:      item[idxLow].(float64),
			Close:    item[idxClose].(float64),
			Volume:   item[idxVolume].(float64),
			Turnover: item[idxAmount].(float64),
		}
	}
	return items, nil
}

func WithDailyCode(code string) dailyOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

func WithDailyDate(date time.Time) dailyOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

func WithDailyDateRange(start, end time.Time) dailyOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
