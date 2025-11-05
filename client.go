package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	token string
	cli   *http.Client
}

type Args map[string]any

func New(token string) *Client {
	return &Client{
		token: token,
		cli: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (cli *Client) call(api string, args Args, fields []string) ([]string, [][]any, error) {
	if args == nil {
		args = make(Args)
	}
	data, err := json.Marshal(map[string]any{
		"api_name": api,
		"token":    cli.token,
		"params":   args,
		"fields":   strings.Join(fields, ","),
	})
	if err != nil {
		return nil, nil, err
	}
	const url = "http://api.tushare.pro"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("http: %d", resp.StatusCode)
	}
	var ret struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Fields []string `json:"fields"`
			Items  [][]any  `json:"items"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, nil, err
	}
	if ret.Code != 0 {
		return nil, nil, fmt.Errorf("code: %d, %s", ret.Code, ret.Msg)
	}
	return ret.Data.Fields, ret.Data.Items, nil
}

func (cli *Client) Call(api string, args Args, fields []string) ([]string, [][]any, error) {
	var columns []string
	var items [][]any
	var err error
	for range 10 {
		columns, items, err = cli.call(api, args, fields)
		if err == nil {
			return columns, items, nil
		}
		time.Sleep(time.Minute)
	}
	return nil, nil, err
}
