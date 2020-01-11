# Backend

## 採用技術

* 言語等
  * 言語: TypeScript
  * フレームワーク: Vue.js, Nuxt.js, Vuetify
  * ライブラリ: firebase, vee-validate, vuedraggable, eslint, jest

* VeeValidate
  * [公式Docs](https://logaretm.github.io/vee-validate/)
  * [GitHub](https://github.com/logaretm/vee-validate)

* Vue.Draggable
  * [公式Docs](https://sortablejs.github.io/Vue.Draggable/)
  * [GitHub](https://github.com/SortableJS/Vue.Draggable)

## 開発ルール

### 開発手順

1. PagesでUI作成
2. Typesに型定義
3. Componentsに分割
4. Storeの実装
5. Components, Storeのテストを作成
6. Storybookの実装

## その他

### Typesディレクトリ

* store
  * stateの型定義
  * modelの型定義
* form
  * 入力フォームの型定義
  * VeeValidate -> バリデーションの型定義

### 命名規則

* コンポーネントの命名規則はスタイルガイドを参考にする．
  * 単一ファイルコンポーネントは，基本的にPrefixに Gran をつけて命名する．

### ディレクトリ構成

```sh
client
├── app
|   ├── assets
|   ├── components
|   │   ├── atoms
|   │   ├── molecules
|   │   ├── organisms
|   │   └── templates
|   ├── layouts
|   ├── middleware
|   ├── pages
|   ├── plugins
|   ├── static
|   ├── store
|   └── types
|       ├── form
|       └── store
└── spec
    ├── components
    │   ├── atoms
    │   ├── molecules
    │   ├── organisms
    │   └── templates
    ├── helpers
    │   ├── response
    │   └── store
    └── store
```
