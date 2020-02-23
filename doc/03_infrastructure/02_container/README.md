# コンテナに関するメモ

## 本番環境のDockerfileの実行方法

1. プロジェクトのルートディレクトリに移動
2. コンテナのビルド e.g.) user_api
    > $ docker build \  
    >     -f container/api/user/Dockerfile.production \  
    >     -t gran_prod_user_api \  
    >     .
3. コンテナの起動 e.g.) user_api
    > $ docker run \  
    >     --env-file .env \  
    >     -v $PWD/secret:/secret:ro \  
    >     gran_prod_user_api
