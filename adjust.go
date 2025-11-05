package tushare

import "time"

// Adjust 复权数据
type Adjust struct {
	Code   string    // 股票代码
	Date   time.Time // 日期
	Factor float64   // 复权因子
}

type adjustOpt func(*Args)

// AdjFactor 获取复权数据
func (cli *Client) AdjFactor(opts ...adjustOpt) ([]Adjust, error) {
	var args Args
	for _, o := range opts {
		o(&args)
	}
	fields, data, err := cli.Call("adj_factor", args, []string{"ts_code", "trade_date", "adj_factor"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxDate, idxFactor int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "trade_date":
			idxDate = i
		case "adj_factor":
			idxFactor = i
		}
	}
	items := make([]Adjust, len(data))
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = Adjust{
			Code:   item[idxCode].(string),
			Date:   date,
			Factor: item[idxFactor].(float64),
		}
	}
	return items, nil
}

func WithAdjustCode(code string) adjustOpt {
	return func(args *Args) {
		(*args)["ts_code"] = code
	}
}

func WithAdjustDate(date time.Time) adjustOpt {
	return func(args *Args) {
		(*args)["trade_date"] = date.Format("20060102")
	}
}

func WithAdjustDateRange(start, end time.Time) adjustOpt {
	return func(args *Args) {
		(*args)["start_date"] = start.Format("20060102")
		(*args)["end_date"] = end.Format("20060102")
	}
}
