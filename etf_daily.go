// https://tushare.pro/document/2?doc_id=127

package tushare

// ETFDaily 获取ETF每日行情
func (cli *Client) ETFDaily(opts ...dailyOpt) ([]DailyTick, error) {
	return cli.daily("etf_daily", opts...)
}
