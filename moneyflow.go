// https://tushare.pro/document/2?doc_id=170

package tushare

import "time"

// MoneyFlow 资金流数据
type MoneyFlow struct {
	Code       string    // 股票代码
	Date       time.Time // 交易日期
	BuySmVol   float64   // 小单买入成交量
	BuySmAmt   float64   // 小单买入成交金额
	SellSmVol  float64   // 小单卖出成交量
	SellSmAmt  float64   // 小单卖出成交金额
	BuyMdVol   float64   // 中单买入成交量
	BuyMdAmt   float64   // 中单买入成交金额
	SellMdVol  float64   // 中单卖出成交量
	SellMdAmt  float64   // 中单卖出成交金额
	BuyLgVol   float64   // 大单买入成交量
	BuyLgAmt   float64   // 大单买入成交金额
	SellLgVol  float64   // 大单卖出成交量
	SellLgAmt  float64   // 大单卖出成交金额
	BuyElgVol  float64   // 特大单买入成交量
	BuyElgAmt  float64   // 特大单买入成交金额
	SellElgVol float64   // 特大单卖出成交量
	SellElgAmt float64   // 特大单卖出成交金额
	NetMfVol   float64   // 净流入成交量
	NetMfAmt   float64   // 净流入成交金额
}

type moneyflowOpt func(params Args)

// MoneyFlow 获取资金流数据
func (cli *Client) MoneyFlow(opts ...moneyflowOpt) ([]MoneyFlow, error) {
	args := make(Args)
	for _, o := range opts {
		o(args)
	}
	fields, data, err := cli.Call("moneyflow", args, []string{
		"ts_code", "trade_date",
		"buy_sm_vol", "buy_sm_amount", "sell_sm_vol", "sell_sm_amount",
		"buy_md_vol", "buy_md_amount", "sell_md_vol", "sell_md_amount",
		"buy_lg_vol", "buy_lg_amount", "sell_lg_vol", "sell_lg_amount",
		"buy_elg_vol", "buy_elg_amount", "sell_elg_vol", "sell_elg_amount",
		"net_mf_vol", "net_mf_amount",
	})
	if err != nil {
		return nil, err
	}
	var (
		idxCode       int
		idxDate       int
		idxBuySmVol   int
		idxBuySmAmt   int
		idxSellSmVol  int
		idxSellSmAmt  int
		idxBuyMdVol   int
		idxBuyMdAmt   int
		idxSellMdVol  int
		idxSellMdAmt  int
		idxBuyLgVol   int
		idxBuyLgAmt   int
		idxSellLgVol  int
		idxSellLgAmt  int
		idxBuyElgVol  int
		idxBuyElgAmt  int
		idxSellElgVol int
		idxSellElgAmt int
		idxNetMfVol   int
		idxNetMfAmt   int
	)
	for i, field := range fields {
		switch field {
		case "ts_code":
			idxCode = i
		case "trade_date":
			idxDate = i
		case "buy_sm_vol":
			idxBuySmVol = i
		case "buy_sm_amount":
			idxBuySmAmt = i
		case "sell_sm_vol":
			idxSellSmVol = i
		case "sell_sm_amount":
			idxSellSmAmt = i
		case "buy_md_vol":
			idxBuyMdVol = i
		case "buy_md_amount":
			idxBuyMdAmt = i
		case "sell_md_vol":
			idxSellMdVol = i
		case "sell_md_amount":
			idxSellMdAmt = i
		case "buy_lg_vol":
			idxBuyLgVol = i
		case "buy_lg_amount":
			idxBuyLgAmt = i
		case "sell_lg_vol":
			idxSellLgVol = i
		case "sell_lg_amount":
			idxSellLgAmt = i
		case "buy_elg_vol":
			idxBuyElgVol = i
		case "buy_elg_amount":
			idxBuyElgAmt = i
		case "sell_elg_vol":
			idxSellElgVol = i
		case "sell_elg_amount":
			idxSellElgAmt = i
		case "net_mf_vol":
			idxNetMfVol = i
		case "net_mf_amount":
			idxNetMfAmt = i
		}
	}
	items := make([]MoneyFlow, len(data))
	for i, item := range data {
		date, _ := time.ParseInLocation("20060102", item[idxDate].(string), time.Local)
		items[i] = MoneyFlow{
			Code:       item[idxCode].(string),
			Date:       date,
			BuySmVol:   item[idxBuySmVol].(float64),
			BuySmAmt:   item[idxBuySmAmt].(float64),
			SellSmVol:  item[idxSellSmVol].(float64),
			SellSmAmt:  item[idxSellSmAmt].(float64),
			BuyMdVol:   item[idxBuyMdVol].(float64),
			BuyMdAmt:   item[idxBuyMdAmt].(float64),
			SellMdVol:  item[idxSellMdVol].(float64),
			SellMdAmt:  item[idxSellMdAmt].(float64),
			BuyLgVol:   item[idxBuyLgVol].(float64),
			BuyLgAmt:   item[idxBuyLgAmt].(float64),
			SellLgVol:  item[idxSellLgVol].(float64),
			SellLgAmt:  item[idxSellLgAmt].(float64),
			BuyElgVol:  item[idxBuyElgVol].(float64),
			BuyElgAmt:  item[idxBuyElgAmt].(float64),
			SellElgVol: item[idxSellElgVol].(float64),
			SellElgAmt: item[idxSellElgAmt].(float64),
			NetMfVol:   item[idxNetMfVol].(float64),
			NetMfAmt:   item[idxNetMfAmt].(float64),
		}
	}
	return items, nil
}

// WithMoneyFlowCode 设置股票代码参数
func WithMoneyFlowCode(code string) moneyflowOpt {
	return func(args Args) {
		args["ts_code"] = code
	}
}

// WithMoneyFlowDate 设置交易日期参数
func WithMoneyFlowDate(date time.Time) moneyflowOpt {
	return func(args Args) {
		args["trade_date"] = date.Format("20060102")
	}
}

// WithMoneyFlowDateRange 设置日期范围参数
func WithMoneyFlowDateRange(start, end time.Time) moneyflowOpt {
	return func(args Args) {
		args["start_date"] = start.Format("20060102")
		args["end_date"] = end.Format("20060102")
	}
}
