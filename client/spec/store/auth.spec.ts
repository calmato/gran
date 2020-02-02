import Vuex from 'vuex'
import cloneDeep from 'lodash.clonedeep'
import '~~/spec/helpers/store-helper'
import axios from '~~/spec/helpers/axios'
import * as AuthStore from '~/store/auth'
import { IUserStore } from '~/types/store'

describe('store/auth', () => {
  let store

  beforeEach(() => {
    store = new Vuex.Store(cloneDeep(AuthStore))
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
        user: {
          uid: 'JUA1ouY12ickxIupMVdVl3ieM7s2',
          email: 'hoge@hoge.com',
          creationTime: '2020-01-01T00:00:00.000000Z',
          lastSignInTime: '2020-01-01T00:00:00.000000Z'
        }
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
      expect(store.getters['user']).toEqual({
        uid: 'JUA1ouY12ickxIupMVdVl3ieM7s2',
        email: 'hoge@hoge.com',
        creationTime: '2020-01-01T00:00:00.000000Z',
        lastSignInTime: '2020-01-01T00:00:00.000000Z'
      })
    })
  })

  describe('mutations', () => {
    let commit
    beforeEach(() => {
      commit = store.commit
    })

    test('setEmailVerified', () => {
      commit('setEmailVerified', true)

      expect(store.state.emailVerified).toBeTruthy()
    })
  })

  describe('actions', () => {
    beforeEach(() => {
      store.$axios = axios
    })

    describe('success', () => {
      beforeEach(() => {
        store.$axios.setSafetyMode(true)
      })

      test('signUp', async () => {
        const form = {
          email: {
            value: 'hoge@hoge.com'
          },
          password: {
            value: '12345678'
          },
          passwordConfirmation: {
            value: '12345678'
          }
        }

        await store.dispatch('signUp', form)
      })
    })

    describe('failure', () => {
      beforeEach(() => {
        store.$axios.setSafetyMode(false)
      })

      test('signUp', async () => {
        const form = {
          email: {
            value: 'hoge@hoge.com'
          },
          password: {
            value: '12345678'
          },
          passwordConfirmation: {
            value: '12345678'
          }
        }

        await expect(store.dispatch('signUp', form)).rejects.toEqual(new Error('Error: some error'))
      })
    })
  })
})
