const noRequiredAuthorizationPath: Array<string> = ['/signin', '/signup', '/forget', '/email-check']

export default ({ route, store, redirect }) => {
  // 認証が不要の path の場合
  const path: string = route.path
  if (noRequiredAuthorizationPath.includes(path)) {
    return
  }

  // 認証が必要な path の場合
  store.dispatch('auth/authentication').catch(() => {
    redirect('/signin')
  })
}
