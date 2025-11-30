// https://tushare.pro/document/2?doc_id=127

package tushare

// FundDaily 获取ETF每日行情
func (cli *Client) FundDaily(opts ...dailyOpt) ([]DailyTick, error) {
	return cli.daily("fund_daily", opts...)
}
