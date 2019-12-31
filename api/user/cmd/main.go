package main

import (
	"log"
	"net/http"

	"github.com/16francs/gran/api/user/config"
)

func main() {
	// ログ出力設定
	config.Logger()

	// 環境変数
	e, err := config.LoadEnvironment()
	if err != nil {
		log.Fatalf("alert: %+v", err)
	}

	// サーバ起動
	r := config.Router()
	if err = http.ListenAndServe(":"+e.Port, r); err != nil {
		panic(err)
	}
}
