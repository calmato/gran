import { ILoginForm } from '../types/form/auth/login'
import { IUserStore, IUser } from '~/types/store/auth'

export const state = (): IUserStore => ({
  emailVerified: false,
  token: '',
  user: {} as IUser
})

export const getters = {
  emailVerified: (state: IUserStore) => state.emailVerified,
  token: (state: IUserStore) => state.token,
  user: (state: IUserStore) => state.user
}

export const mutations = {
  setEmailVerified(state: IUserStore, emailVerified: boolean) {
    state.emailVerified = emailVerified
  },

  setToken(state: IUserStore, token: string) {
    state.token = token
  },

  setUser(state: IUserStore, user: IUser) {
    state.user = user
  }
}

export const actions = {
  async setIdToken({ commit }) {
    await this.$firebase
      .auth()
      .currentUser.getIdToken(true)
      .then((token: string) => {
        commit('setToken', token)
      })
      .catch((err: any) => {
        throw err
      })
  },

  loginWithEmailAndPassword({ commit, dispatch }, form: ILoginForm): Promise<void> {
    return this.$firebase
      .auth()
      .signInWithEmailAndPassword(form.email, form.password)
      .then(async (auth: any) => {
        const user: IUser = {
          uid: auth.user.uid,
          email: auth.user.email,
          creationTime: auth.user.metadata.creationTime,
          lastSignInTime: auth.user.metadata.lastSignInTime
        }
        await dispatch('setIdToken')
        commit('setUser', user)
        commit('setEmailVerified', auth.user.emailVerified)
        return Promise.resolve()
      })
      .catch((error: any) => {
        return Promise.reject(error)
      })
  }
}
