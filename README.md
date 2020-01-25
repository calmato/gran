# Gran

![](https://github.com/16francs/gran/workflows/Client(Nuxt.js)%20Build%20and%20Test/badge.svg)
![](https://github.com/16francs/gran/workflows/Client(Nuxt.js)%20Lighthouse/badge.svg)  
![](https://github.com/16francs/gran/workflows/Storybook(TypeScript)%20Build%20and%20Test/badge.svg)  
![](https://github.com/16francs/gran/workflows/User%20API(Golang)%20Build%20and%20Test/badge.svg)

## タスク管理アプリ

## 初期設定

### プロジェクトのダウンロード

> $ git clone https://github.com/16francs/gran.git  
> $ cd gran

### コンテナの初期設定

> $ make setup

### .envファイルの作成

* 以下を参考に `.env` ファイルを編集
  * `FIREBASE_xxxx` はFirebaseプロジェクトにログインし確認
  * `./../../secret/xxx-firebase-adminsdk-xxx.json` は正しいファイル名に置き換え

```env
CLIENT_PATH=./../client
FIREBASE_API_KEY=xxxxxx
FIREBASE_PROJECT_ID=xxxxxx
FIREBASE_MESSAGING_SENDER_ID=xxxxxx
GOOGLE_APPLICATION_CREDENTIALS=./../../secret/xxx-firebase-adminsdk-xxx.json
```

### コンテナの起動

> $ make start
