package tushare

import (
	"time"
)

// TradeCal 获取指定日期范围内的交易日列表
func (cli *Client) TradeCal(begin, end time.Time) ([]time.Time, error) {
	_, data, err := cli.Call("trade_cal", Args{
		"start_date": begin.Format("20060102"),
		"end_date":   end.Format("20060102"),
		"is_open":    "1",
	}, []string{"cal_date"})
	if err != nil {
		return nil, err
	}
	ret := make([]time.Time, len(data))
	for i := range data {
		t, err := time.ParseInLocation("20060102", data[i][0].(string), time.Local)
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}
