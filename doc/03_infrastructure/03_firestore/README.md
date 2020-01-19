# Firestore

## DB設計

### エンティティ一覧

* boards
  * ボード
  * parent - group
  * collection - lists
* check_lists
  * チェックリスト
  * parent - task
  * reference - user_ref
* groups
  * グループ
  * root
  * collection - boards
  * reference - user_refs
* lists
  * ボードリスト
  * parent - boards
* tasks
  * タスク
  * root
  * collection - check_lists
  * reference - list_ref, user_ref
* users
  * ユーザ
  * root
  * reference - group_refs

### 設計での考慮点

1. コスト面を意識した設計にする
    * モバイルアプリでよくみるNoSQL使用時のベストプラクティスはN+1問題がおこる
      * FirestoreはR/Wの回数で料金がかかるので、今回は極力N+1問題がおこらない設計にする

2. 親のドキュメントが変わることが無いものはサブコレクションへ
    * ボード: 別のグループに移動することはない -> サブコレクション
    * リスト: 変更に別のボードに移動することはない -> サブコレクション
    * タスク: 別のリストに移動することがある -> ルート

3. 親コレクションを意識せず検索する必要があるものはルートへ
    * リスト: あるボードのリスト一覧..みたいな検索しかしない -> 特に考慮しなくて良い
    * タスク: どのボード、リストに所属しているか関係なく一覧検索したい -> ルート

4. コレクション間の関連はReference型で表現する

## 参考URL

* [NoSQLデータモデリング技法](https://gist.github.com/matope/2396234)
* [Cloud Firestoreを実践投入するにあたって考えたこと](https://qiita.com/1amageek/items/d606dcee9fbcf21eeec6)
* [Cloud Firestoreのデータ構造の決め方をFirebaseの動画から学ぶ](https://qiita.com/shiz/items/5f4c8ae19083ccdd46b2)
* [Cloud Firestoreの勘所 パート2 — データ設計](https://medium.com/google-cloud-jp/firestore2-920ac799345c)
* [Cloud FirestoreのSubCollectionとQueryっていつ使うの問題](https://qiita.com/1amageek/items/d2ef7a49bccf5b4ea78e)
* [Cloud Firestoreのコレクショングループクエリについて](https://firebase.googleblog.com/2019/06/understanding-collection-group-queries.html)
* [リアルタイムデータベースCloud Firestore入門](https://speakerdeck.com/tetsuyanegishi/riarutaimudetabesu-cloud-firestoreru-men)
