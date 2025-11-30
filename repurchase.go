// https://tushare.pro/document/2?doc_id=124

package tushare

import (
	"time"
)

// Repurchase 股票回购数据
type Repurchase struct {
	Code    string         // 股票代码
	AnnDate time.Time      // 公告日期
	EndDate time.Time      // 截止日期
	ExpDate time.Time      // 过期日期
	Proc    repurchaseProc // 进度
	Volume  float64        // 回购数量
	Amount  float64        // 回购金额
	High    float64        // 最高价
	Low     float64        // 最低价
}

type repurchaseOpt func(Args)

// Repurchase 获取股票回购数据
func (cli *Client) Repurchase(opts ...repurchaseOpt) ([]Repurchase, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("repurchase", args, []string{"ts_code", "ann_date", "end_date", "exp_date", "proc", "vol", "amount", "high_limit", "low_limit"})
	if err != nil {
		return nil, err
	}
	var idxCode, idxAnnDate, idxEndDate, idxExpDate, idxProc, idxVolume, idxAmount, idxHigh, idxLow int
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "ann_date":
			idxAnnDate = i
		case "end_date":
			idxEndDate = i
		case "exp_date":
			idxExpDate = i
		case "proc":
			idxProc = i
		case "vol":
			idxVolume = i
		case "amount":
			idxAmount = i
		case "high_limit":
			idxHigh = i
		case "low_limit":
			idxLow = i
		}
	}
	toDate := func(v any) time.Time {
		if v == nil {
			return time.Time{}
		}
		t, _ := time.ParseInLocation("20060102", v.(string), time.Local)
		return t
	}
	toFloat64 := func(v any) float64 {
		if v == nil {
			return 0
		}
		return v.(float64)
	}
	items := make([]Repurchase, len(data))
	for i, item := range data {
		items[i] = Repurchase{
			Code:    item[idxCode].(string),
			AnnDate: toDate(item[idxAnnDate]),
			EndDate: toDate(item[idxEndDate]),
			ExpDate: toDate(item[idxExpDate]),
			Proc:    repurchaseProc(item[idxProc].(string)),
			Volume:  toFloat64(item[idxVolume]),
			Amount:  toFloat64(item[idxAmount]),
			High:    toFloat64(item[idxHigh]),
			Low:     toFloat64(item[idxLow]),
		}
	}
	return items, nil
}

// WithRepurchaseAnnDate 设置公告日期参数
func WithRepurchaseAnnDate(date time.Time) repurchaseOpt {
	return func(args Args) {
		args["ann_date"] = date.Format("20060102")
	}
}

// WithRepurchaseDateRange 设置公告日期范围参数
func WithRepurchaseDateRange(start, end time.Time) repurchaseOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}

type repurchaseProc string

const (
	RepurchaseProcPrepare   repurchaseProc = "预案"
	RepurchaseProcConfirm   repurchaseProc = "股东大会通过"
	RepurchaseProcImplement repurchaseProc = "实施"
	RepurchaseProcComplete  repurchaseProc = "完成"
)
