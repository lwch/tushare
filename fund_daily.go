// https://tushare.pro/document/2?doc_id=127

package tushare

// FundDaily 基金日线行情
func (cli *Client) FundDaily(opts ...dailyOpt) ([]DailyTick, error) {
	return cli.daily("fund_daily", opts...)
}
