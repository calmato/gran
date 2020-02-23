# Gran

![Client(Nuxt.js) Build and Test](https://github.com/16francs/gran/workflows/Client(Nuxt.js)%20Build%20and%20Test/badge.svg)
![Client(Nuxt.js) Lighthouse](https://github.com/16francs/gran/workflows/Client(Nuxt.js)%20Lighthouse/badge.svg)  
![Storybook(TypeScript) Build and Test](https://github.com/16francs/gran/workflows/Storybook(TypeScript)%20Build%20and%20Test/badge.svg)  
![User API(Golang) Build and Test](https://github.com/16francs/gran/workflows/User%20API(Golang)%20Build%20and%20Test/badge.svg)
![Group API(Golang) Build and Test](https://github.com/16francs/gran/workflows/Group%20API(Golang)%20Build%20and%20Test/badge.svg)
![ToDo API(Golang) Build and Test](https://github.com/16francs/gran/workflows/ToDo%20API(Golang)%20Build%20and%20Test/badge.svg)

![Client(Nuxt.js) Deploy to Firebase Hosting](https://github.com/16francs/gran/workflows/Client(Nuxt.js)%20Deploy%20to%20Firebase%20Hosting/badge.svg)  
![User API(Golang) Deploy to Staging](https://github.com/16francs/gran/workflows/User%20API(Golang)%20Deploy%20to%20Staging/badge.svg)
![Group API(Golang) Deploy to Staging](https://github.com/16francs/gran/workflows/Group%20API(Golang)%20Deploy%20to%20Staging/badge.svg)
![ToDo API(Golang) Deploy to Staging](https://github.com/16francs/gran/workflows/ToDo%20API(Golang)%20Deploy%20to%20Staging/badge.svg)

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
