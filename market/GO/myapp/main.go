package main

import (
	"encoding/csv" 
	"fmt"
	"myapp/controllers"
	"os"
	"time"
)

func main() {
	client := &controllers.CoinGeckoClient{
		BaseURL: "https://api.coingecko.com/api/v3",
	}
	// 仮想通貨名：
	// btc:
	// xrp:ripple
	coinID := "ripple"
	// CSVファイル作成
	start := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)
	today := time.Date(2025, 6, 5, 0, 0, 0, 0, time.UTC)
	startFormatted := start.Format("20060102")
	todayFormatted := today.Format("20060102")
	file, err := os.Create(fmt.Sprintf("prices_%s_F%s_T%s.csv", coinID, startFormatted, todayFormatted))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ヘッダー
	writer.Write([]string{"date", "usd_price", "jpy_price"})
	// usdToJpy, err := client.GetUsdToJpyRate()
	for d := start; d.Before(today); d = d.AddDate(0, 0, 1) {
		usd, err := client.GetHistoricalUSDPrice(coinID, d)
		if err != nil {
			fmt.Printf("%s: エラー - %v\n", d.Format("2006-01-02"), err)
			continue
		}
		usdToJpy, err2 := client.GetUsdToJpyRate("usd", d)
		if err2 != nil {
			fmt.Printf("%s: 為替レート取得失敗 - %v\n", d.Format("2006-01-02"), err2)
			continue
		}
		jpy := usd * usdToJpy
		// CSV出力
		record := []string{
			d.Format("2006-01-02"),
			fmt.Sprintf("%.2f", usd),
			fmt.Sprintf("%.2f", jpy),
		}
		writer.Write(record)
		writer.Flush() // 途中停止対策
		// CoinGeckoのAPI制限にかからないよう少し待つ
		time.Sleep(30 * time.Second) // 1.2秒
	}
}