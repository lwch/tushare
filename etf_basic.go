// https://tushare.pro/document/2?doc_id=385

package tushare

import "time"

// ETFBasic 获取ETF列表
type ETFBasic struct {
	Code      string      // ETF代码
	Name      string      // ETF名称
	IndexCode string      // 关联指数代码
	IndexName string      // 关联指数名称
	Date      time.Time   // 上市日期
	Status    etfStatus   // ETF状态
	Exchange  etfExchange // 交易所
}

type etfOpt func(Args)

// ETFBasic 获取ETF列表
func (cli *Client) ETFBasic(opts ...etfOpt) ([]ETFBasic, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("etf_basic", args,
		[]string{"ts_code", "csname", "index_code", "index_name", "list_date", "list_status", "exchange"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxName, idxIndexCode, idxIndexName, idxDate, idxStatus, idxExchange int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "csname":
			idxName = i
		case "index_code":
			idxIndexCode = i
		case "index_name":
			idxIndexName = i
		case "list_date":
			idxDate = i
		case "list_status":
			idxStatus = i
		case "exchange":
			idxExchange = i
		}
	}
	items := make([]ETFBasic, len(data))
	toString := func(v any) string {
		if v == nil {
			return ""
		}
		return v.(string)
	}
	toDate := func(v any) time.Time {
		if v == nil {
			return time.Time{}
		}
		s := v.(string)
		t, _ := time.Parse("20060102", s)
		return t
	}
	for i, item := range data {
		items[i] = ETFBasic{
			Code:      toString(item[idxCode]),
			Name:      toString(item[idxName]),
			IndexCode: toString(item[idxIndexCode]),
			IndexName: toString(item[idxIndexName]),
			Date:      toDate(item[idxDate]),
			Status:    etfStatus(toString(item[idxStatus])),
			Exchange:  etfExchange(toString(item[idxExchange])),
		}
	}
	return items, nil
}

// WithETFCode 按ETF代码查询
func WithETFCode(code string) etfOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithETFIndexCode 按关联指数代码查询
func WithETFIndexCode(indexCode string) etfOpt {
	return func(args Args) {
		args["index_code"] = indexCode
	}
}

// WithETFDate 按上市日期查询
func WithETFDate(date time.Time) etfOpt {
	return func(args Args) {
		args["list_date"] = date.Format("20060102")
	}
}

type etfStatus string

const ETFStatusL etfStatus = "L" // 上市
const ETFStatusD etfStatus = "D" // 退市
const ETFStatusP etfStatus = "P" // 待上市

// WithETFStatus 按基金状态查询(上市/退市/待上市)
func WithETFStatus(status etfStatus) etfOpt {
	return func(args Args) {
		args["list_status"] = status
	}
}

type etfExchange string

const ETFExchangeSSE etfExchange = "SSE" // 上交所
const ETFExchangeSZ etfExchange = "SZ"   // 深交所

// WithETFExchange 按交易所查询
func WithETFExchange(exchange etfExchange) etfOpt {
	return func(args Args) {
		args["exchange"] = exchange
	}
}
