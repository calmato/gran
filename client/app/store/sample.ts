import { ISampleNewForm } from '~/types/form'
import { ISampleStore, ISampleUser } from '~/types/store'

export const state = (): ISampleStore => ({
  users: [{ name: '山田 侑樹' }, { name: '濵田 広大' }, { name: '西川 直志' }]
})

export const getters = {
  users: (state: ISampleStore) => state.users
}

export const mutations = {
  addUser(state: ISampleStore, user: ISampleUser) {
    state.users.push(user)
  }
}

export const actions = {
  addUser({ commit }, newForm: ISampleNewForm) {
    commit('addUser', newForm)
  }
}
