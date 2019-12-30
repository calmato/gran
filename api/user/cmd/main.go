package main

import (
	"log"
	"net/http"

	"github.com/16francs/gran/api/user/config"
)

func main() {
	// TODO: 設定の読み込み

	// サーバ起動
	// TODO: ポート番号は環境変数から読み込みする
	r := config.Router()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("alert: %v", err)
	}
}
