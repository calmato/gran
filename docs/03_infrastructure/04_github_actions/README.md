# メモ

## GCPのサービスアカウントについて

* jsonで出力される認証情報をbase64に変換して環境変数にコピペする
  > $  cat xxxxxx-xxxxxx.json | base64

## Terraformの認証情報について

* terraformでは `base64` 形式ではなく　`json` 形式で認証情報を渡す必要がある
  * `Github > Settings > Secrets` より認証情報を登録
  * 登録する内容は、SAアカウントのjsonの内容をコピペ
