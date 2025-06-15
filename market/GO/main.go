package main

import (
	"flag"
	"encoding/csv" 
	"fmt"
	"myapp/controllers"
	"os"
	"time"
)

func main() {
	fmt.Println("Start")
	// コマンドライン引数を定義
	coinID := flag.String("coin", "ripple", "取得する仮想通貨のID（例：bitcoin, ripple, ethereum）")
	startStr := flag.String("start", "2025-05-25", "開始日（例：2025-05-25）")
	endStr := flag.String("end", "2025-06-05", "終了日（例：2025-06-05）")

	// 引数をパース
	flag.Parse()

	// 日付文字列を time.Time に変換
	const layout = "2006-01-02"
	start, err := time.Parse(layout, *startStr)
	if err != nil {
		panic(fmt.Sprintf("開始日が無効です: %v", err))
	}
	end, err := time.Parse(layout, *endStr)
	if err != nil {
		panic(fmt.Sprintf("終了日が無効です: %v", err))
	}
	client := &controllers.CoinGeckoClient{
		BaseURL: "https://api.coingecko.com/api/v3",
	}
	err = os.MkdirAll("/output", os.ModePerm)
	if err != nil {
		fmt.Errorf("出力ディレクトリの作成に失敗しました: %v", err)
	}
	// CSVファイル作成
	file, err := os.Create(fmt.Sprintf("output/prices_%s_F%s_T%s.csv", *coinID, *startStr, *endStr))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ヘッダー
	writer.Write([]string{"date", "usd_price", "jpy_price"})

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		usd, err := client.GetHistoricalUSDPrice(*coinID, d)
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