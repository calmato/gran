import * as AuthStore from '~/store/auth'

export default {
  namespaced: true,

  state: AuthStore.state,
  getters: AuthStore.getters,
  mutations: AuthStore.mutations,
  actions: AuthStore.actions,
}
