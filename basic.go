package tushare

// StockBasic 股票基本信息
type StockBasic struct {
	Symbol   string // 股票代码
	Name     string // 股票名称
	Area     string // 地域
	Industry string // 行业
}

func (cli *Client) basic(status string) ([]StockBasic, error) {
	fields, data, err := cli.Call("stock_basic", Args{
		"list_status": status,
	}, []string{"symbol", "name", "area", "industry"})
	if err != nil {
		return nil, err
	}
	var idxSymbol, idxName, idxArea, idxIndustry int
	for i, field := range fields {
		switch field {
		case "symbol":
			idxSymbol = i
		case "name":
			idxName = i
		case "area":
			idxArea = i
		case "industry":
			idxIndustry = i
		}
	}
	items := make([]StockBasic, len(data))
	toString := func(v any) string {
		if v == nil {
			return ""
		}
		return v.(string)
	}
	for i, item := range data {
		items[i] = StockBasic{
			Symbol:   toString(item[idxSymbol]),
			Name:     toString(item[idxName]),
			Area:     toString(item[idxArea]),
			Industry: toString(item[idxIndustry]),
		}
	}
	return items, nil
}

// StockBasicL 获取上市的股票列表
func (cli *Client) StockBasicL() ([]StockBasic, error) {
	return cli.basic("L") // 上市
}

// StockBasicD 获取已退市的股票列表
func (cli *Client) StockBasicD() ([]StockBasic, error) {
	return cli.basic("D") // 已退市
}
