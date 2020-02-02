import Vuex from 'vuex'
import cloneDeep from 'lodash.clonedeep'
import '~~/spec/helpers/store-helper'
import axios from '~~/spec/helpers/axios'
import * as AuthStore from '~/store/auth'
import { ISignUpForm, SignUpForm } from '~/types/form'
import { IUser, IUserStore } from '~/types/store'

describe('store/auth', () => {
  let store: any
  let user: IUser

  beforeEach(() => {
    store = new Vuex.Store(cloneDeep(AuthStore))

    user = {
      uid: 'JUA1ouY12ickxIupMVdVl3ieM7s2',
      email: 'hoge@hoge.com',
      creationTime: '2020-01-01T00:00:00.000000Z',
      lastSignInTime: '2020-01-01T00:00:00.000000Z'
    }
  })

  afterEach(() => {
    store = null
  })

  describe('state', () => {
    test('emailVerified: 初期値', () => {
      expect(store.state.emailVerified).toBeFalsy()
    })

    test('token: 初期値', () => {
      expect(store.state.token).toBe('')
    })

    test('user: 初期値', () => {
      expect(store.state.user).toEqual({})
    })
  })

  describe('getters', () => {
    beforeEach(() => {
      const userStore: IUserStore = {
        emailVerified: true,
        token: 'xL7QdFig7HOv7btzH8gAKuK81xI2',
        user
      }

      store.replaceState(userStore)
    })

    test('emailVerified', () => {
      expect(store.getters['emailVerified']).toBeTruthy()
    })

    test('token', () => {
      expect(store.getters['token']).toBe('xL7QdFig7HOv7btzH8gAKuK81xI2')
    })

    test('user', () => {
      expect(store.getters['user']).toEqual(user)
    })
  })

  describe('mutations', () => {
    let commit: any
    beforeEach(() => {
      commit = store.commit
    })

    test('setEmailVerified', () => {
      commit('setEmailVerified', true)

      expect(store.state.emailVerified).toBeTruthy()
    })

    test('setEmailVerified', () => {
      commit('setToken', 'xL7QdFig7HOv7btzH8gAKuK81xI2')

      expect(store.state.token).toBe('xL7QdFig7HOv7btzH8gAKuK81xI2')
    })
  })

  describe('actions', () => {
    beforeEach(() => {
      store.$axios = axios
    })

    describe('success', () => {
      let signUpForm: ISignUpForm

      beforeEach(() => {
        store.$axios.setSafetyMode(true)

        signUpForm = SignUpForm
      })

      test('signUp', async () => {
        await store.dispatch('signUp', signUpForm)
      })
    })

    describe('failure', () => {
      beforeEach(() => {
        store.$axios.setSafetyMode(false)
      })

      test('signUp', async () => {
        await expect(store.dispatch('signUp', SignUpForm)).rejects.toEqual(
          new Error('Error: some error')
        )
      })
    })
  })
})
