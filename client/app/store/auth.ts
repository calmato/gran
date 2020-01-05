import { ILoginForm } from '../types/form/auth/login'
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
  loginWithEmailAndPassword({ commit }, form: ILoginForm): Promise<any> {
    return this.$firebase
      .auth()
      .signInWithEmailAndPassword(form.email, form.password)
      .then((auth: any) => {
        const user: IUser = {
          uid: auth.user.uid,
          email: auth.user.email,
          creationTime: auth.user.metadata.creationTime,
          lastSignInTime: auth.user.metadata.lastSignInTime
        }
        commit('setUser', user)
      })
      .catch((error: any) => {
        throw Promise.reject(error)
      })
  }
}
