# Backend

## 採用技術

* 言語等
  * 言語: Go
  * フレームワーク: gin
  * ライブラリ: firebase, colog, validator, xerrors

* アーキテクチャ
  * オニオンアーキテクチャ
    * [Qiita: ドメイン駆動+オニオンアーキテクチャ概略](https://qiita.com/little_hand_s/items/2040fba15d90b93fc124)
    * [Qiita: goプロジェクトにオニオンアーキテクチャを導入した](https://qiita.com/nanamen/items/f37d1047368929e377fd)

## 開発ルール

### 開発手順

1. Swaggerの作成
2. Domain層の作成
3. Infrastructure, Interface, Application層の作成
4. Domain, Application層のテストを作成

### 実装詳細

* [ディレクトリ関連図](https://github.com/16francs/gran/tree/master/doc/02_backend/directory.md)

## その他

### マイクロサービス化の粒度

* [User Service](https://github.com/16francs/gran/tree/master/doc/02_backend/03_user_api)

### ディレクトリ構成

```sh
sample_service
├── cmd
├── config
├── internal
│   ├── application
│   │   ├── request
│   │   ├── response
│   │   └── validation
│   ├── domain
│   │   ├── repository
│   │   ├── service
│   │   └── validation
│   ├── infrastructure
│   │   ├── persistence
│   │   └── validation
│   └── interface
│       └── handler
│           └── v1
├── lib
│   └── firebase
│       ├── authentication
│       └── firestore
└── registry
```
