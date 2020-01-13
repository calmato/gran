export default function({ $axios, store, redirect }) {
  $axios.onRequest((config: any) => {
    // API の ホスト名
    config.baseURL = process.env.apiURL

    // 認証情報の取得
    const token: string = `Bearer ${store.getters['auth/token']}`
    if (token) {
      config.headers.common['Authorization'] = token
    }

    return config
  })

  $axios.onError((err: any) => {
    if (err.response.status === 401) {
      // TODO: ログアウト

      redirect('/signin')
    }
  })
}
