# ディレクトリ設計

![ディレクトリ関連図](https://github.com/16francs/gran/tree/master/doc/02_backend/01_development_rules/directory_dependencies.jpeg)

## 詳細

### 1. Interface層での処理

1. `Interface/Handler` で、リクエストを受け取る
2. `Interface/Handler` で、 `Application/Request` にリクエストを代入
3. `Interface/Handler` から `Application` を呼び出し

### 2. Application層での処理

1. `Application` から `Application/Validation` を呼び出し
2. `Application/Validation` で、リクエスト値の検証
3. `Application` で、 `Domain` にリクエスト値を代入し、エンティティの生成
4. `Application` から `Domain/Service` を呼び出し

### 3. Domain層での処理

1. `Domain/Service` から `Domain/Validation` を呼び出し
2. `Domain/Service` から `Domain/Repository` を呼び出し
3. `Application` へ値を返す

### 4. Application層での処理

1. `Interface/Handler` へ値を返す

### 5. Interface層での処理

1. `Interface/Handler` で、 `Application/Response` にレスポンスを代入
2. レスポンスを返す

### ※ Infrastructure層での処理

* Persistenceについて
  * `Domain/Repository` のInterfaceを満たした処理を実装
  * `lib` 内の `firebase` ライブラリを使用して実装

* Validationについて
  * `Domain/Validation` のInterfaceを満たした処理を実装
