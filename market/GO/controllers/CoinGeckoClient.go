package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CoinGecko API への接続を抽象化した構造体
type CoinGeckoClient struct {
	BaseURL string
}

type ExchangeRateResponse struct {
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
	} `json:"market_data"`
}

// USD→JPY の為替レートを取得
func (c *CoinGeckoClient) GetUsdToJpyRate(coinID string, date time.Time) (float64, error) {
	formatted := date.Format("02-01-2006") // dd-mm-yyyy
	url := fmt.Sprintf("%s/coins/%s/history?date=%s", c.BaseURL, coinID, formatted)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTPエラー: %d - %s", resp.StatusCode, url)
	}
	defer resp.Body.Close()

	var data ExchangeRateResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	rate, ok := data.MarketData.CurrentPrice["jpy"]
	if !ok {
		return 0, fmt.Errorf("JPYレートが取得できませんでした")
	}
	return rate, nil
}

// 仮想通貨の USD 価格を一括取得
func (c *CoinGeckoClient) GetHistoricalUSDPrice(coinID string, date time.Time) (float64, error) {
	formatted := date.Format("02-01-2006") // dd-mm-yyyy
	url := fmt.Sprintf("%s/coins/%s/history?date=%s", c.BaseURL, coinID, formatted)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data ExchangeRateResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}
	rate, ok := data.MarketData.CurrentPrice["usd"]
	if !ok {
		return 0, fmt.Errorf("USDレートが取得できませんでした")
	}
	return rate, nil
}