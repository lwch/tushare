// https://tushare.pro/document/2?doc_id=96

package tushare

import "time"

type indexWeightOpt func(Args)

type IndexWeight struct {
	Code   string    // 成分股代码
	Date   time.Time // 交易日期
	Weight float64   // 权重
}

// IndexWeight 指数成分股权重
func (cli *Client) IndexWeight(code string, opts ...indexWeightOpt) ([]IndexWeight, error) {
	args := make(Args)
	args["index_code"] = code
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("index_weight", args, []string{
		"con_code", "trade_date", "weight"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxDate, idxWeight int
	for i, field := range fields {
		switch field {
		case "con_code":
			idxCode = i
		case "trade_date":
			idxDate = i
		case "weight":
			idxWeight = i
		}
	}
	items := make([]IndexWeight, len(data))
	toFloat := func(v any) float64 {
		if v == nil {
			return 0
		}
		return v.(float64)
	}
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = IndexWeight{
			Code:   item[idxCode].(string),
			Date:   date,
			Weight: toFloat(item[idxWeight]),
		}
	}
	return items, nil
}

// WithIndexWeightDate 按交易日期查询
func WithIndexWeightDate(date time.Time) indexWeightOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithIndexWeightDateRange 按交易日期范围查询
func WithIndexWeightDateRange(start, end time.Time) indexWeightOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
