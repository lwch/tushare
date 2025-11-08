// https://tushare.pro/document/2?doc_id=27

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

type DailyTick struct {
	Tick
	PreClose float64 // 昨收价
	Change   float64 // 涨跌额
	PctChg   float64 // 涨跌幅
}

type dailyOpt func(Args)

func (cli *Client) daily(api string, opts ...dailyOpt) ([]DailyTick, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call(api, args, []string{
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
			PreClose: item[idxPreClose].(float64),
			Change:   item[idxChange].(float64),
			PctChg:   item[idxPctChg].(float64),
		}
	}
	return items, nil
}

// Daily 获取日线数据
func (cli *Client) Daily(opts ...dailyOpt) ([]DailyTick, error) {
	return cli.daily("daily", opts...)
}

// DailyVip 获取VIP日线数据
func (cli *Client) DailyVip(opts ...dailyOpt) ([]DailyTick, error) {
	return cli.daily("daily_vip", opts...)
}

// WithDailyCode 按股票代码查询
func WithDailyCode(code string) dailyOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithDailyDate 按交易日期查询
func WithDailyDate(date time.Time) dailyOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithDailyDateRange 按交易日期范围查询
func WithDailyDateRange(start, end time.Time) dailyOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
