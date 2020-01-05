import { IUserStore, IUser } from '~/types/store/auth'

export const state = (): IUserStore => ({
  user: {} as IUser
})

export const mutations = {
  setUser(state: IUserStore, user: IUser) {
    state.user = user
  }
}

export const actions = {
  setUser({ commit }, user: IUser) {
    commit('setUser', user)
  }
}
