// https://tushare.pro/document/2?doc_id=385

package tushare

import "time"

// FundBasic 基金基本信息
type FundBasic struct {
	Code      string       // 基金代码
	Name      string       // 基金名称
	IndexCode string       // 关联指数代码
	IndexName string       // 关联指数名称
	Date      time.Time    // 上市日期
	Status    fundStatus   // 基金状态
	Exchange  fundExchange // 交易所
}

type fundOpt func(Args)

// FundBasic 获取基金列表
func (cli *Client) FundBasic(opts ...fundOpt) ([]FundBasic, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("fund_basic", args,
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
	items := make([]FundBasic, len(data))
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
		items[i] = FundBasic{
			Code:      toString(item[idxCode]),
			Name:      toString(item[idxName]),
			IndexCode: toString(item[idxIndexCode]),
			IndexName: toString(item[idxIndexName]),
			Date:      toDate(item[idxDate]),
			Status:    fundStatus(toString(item[idxStatus])),
			Exchange:  fundExchange(toString(item[idxExchange])),
		}
	}
	return items, nil
}

// WithFundCode 按基金代码查询
func WithFundCode(code string) fundOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithIndexCode 按关联指数代码查询
func WithIndexCode(indexCode string) fundOpt {
	return func(args Args) {
		args["index_code"] = indexCode
	}
}

// WithDate 按上市日期查询
func WithDate(date time.Time) fundOpt {
	return func(args Args) {
		args["list_date"] = date.Format("20060102")
	}
}

type fundStatus string

const FundStatusL fundStatus = "L" // 上市
const FundStatusD fundStatus = "D" // 退市
const FundStatusP fundStatus = "P" // 待上市

// WithFundStatus 按基金状态查询(上市/退市/待上市)
func WithFundStatus(status fundStatus) fundOpt {
	return func(args Args) {
		args["list_status"] = status
	}
}

type fundExchange string

const FundExchangeSSE fundExchange = "SSE" // 上交所
const FundExchangeSZ fundExchange = "SZ"   // 深交所

// WithFundExchange 按交易所查询
func WithFundExchange(exchange fundExchange) fundOpt {
	return func(args Args) {
		args["exchange"] = exchange
	}
}
