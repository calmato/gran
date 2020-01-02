package main

import (
	"context"
	"net/http"

	"github.com/16francs/gran/api/user/config"
	"github.com/16francs/gran/api/user/lib/firebase"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
	"google.golang.org/api/option"
)

func main() {
	// ログ出力設定
	config.Logger()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 環境変数
	e, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	// Firebaseの初期化
	opt := option.WithCredentialsFile(e.GoogleApplicationCredentialsPath)

	fb, err := firebase.InitializeApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	// Firestore
	fs, err := firestore.NewClient(ctx, fb)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	// サーバ起動
	r := config.Router()
	if err = http.ListenAndServe(":"+e.Port, r); err != nil {
		panic(err)
	}
}
