# Gran

## タスク管理アプリ

## 初期設定

### プロジェクトのダウンロード

> $ git clone https://github.com/16francs/gran.git  
> $ cd gran

### コンテナの初期設定

> $ make setup

### .envファイルの作成

* 以下を参考に `.env` ファイルを編集
  * `./../../secret/xxx-firebase-adminsdk-xxx.json` は正しいファイル名に置き換え

```env
CLIENT_PATH=./../client
GOOGLE_APPLICATION_CREDENTIALS=./../../secret/xxx-firebase-adminsdk-xxx.json
```

### コンテナの起動

> $ make start
