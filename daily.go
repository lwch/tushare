package tushare

import (
	"time"
)

// Tick 行情数据
type Tick struct {
	Code     string  // 股票代码
	Open     float64 // 开盘价
	High     float64 // 最高价
	Low      float64 // 最低价
	Close    float64 // 收盘价
	Volume   float64 // 成交量(手)
	Turnover float64 // 成交额(千元)
}

func daily(fields []string, data [][]any) []Tick {
	var idxCode, idxOpen, idxHigh, idxLow, idxClose, idxVolume, idxAmount int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
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
		items[i] = Tick{
			Code:     item[idxCode].(string),
			Open:     item[idxOpen].(float64),
			High:     item[idxHigh].(float64),
			Low:      item[idxLow].(float64),
			Close:    item[idxClose].(float64),
			Volume:   item[idxVolume].(float64),
			Turnover: item[idxAmount].(float64),
		}
	}
	return items
}

// DailyByDate 获取指定日期的股票行情数据
func (cli *Client) DailyByDate(date time.Time) ([]Tick, error) {
	fields, data, err := cli.Call("daily", Args{
		"trade_date": date.Format("20060102"),
	}, []string{"ts_code", "open", "high", "low", "close", "vol", "amount"})
	if err != nil {
		return nil, err
	}
	return daily(fields, data), nil
}

// DailyByCode 获取指定股票的行情数据
func (cli *Client) DailyByCode(code string) ([]Tick, error) {
	fields, data, err := cli.Call("daily", Args{
		"ts_code": code,
	}, []string{"ts_code", "open", "high", "low", "close", "vol", "amount"})
	if err != nil {
		return nil, err
	}
	return daily(fields, data), nil
}

// DailyVIPByDate 获取指定日期的股票行情数据(vip接口)
func (cli *Client) DailyVIPByDate(date time.Time) ([]Tick, error) {
	fields, data, err := cli.Call("daily_vip", Args{
		"trade_date": date.Format("20060102"),
	}, []string{"ts_code", "open", "high", "low", "close", "vol", "amount"})
	if err != nil {
		return nil, err
	}
	return daily(fields, data), nil
}

// DailyVIPByCode 获取指定股票的行情数据(vip接口)
func (cli *Client) DailyVIPByCode(code string) ([]Tick, error) {
	fields, data, err := cli.Call("daily_vip", Args{
		"ts_code": code,
	}, []string{"ts_code", "open", "high", "low", "close", "vol", "amount"})
	if err != nil {
		return nil, err
	}
	return daily(fields, data), nil
}
