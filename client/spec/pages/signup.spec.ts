import Vuex from 'vuex'
import cloneDeep from 'lodash.clonedeep'
import { mount, RouterLinkStub } from '@vue/test-utils'
import '~~/spec/helpers/component-helper'
import axios from '~~/spec/helpers/axios'
import storeModule from '~~/spec/helpers/store'
import SignUp from '~/pages/SignUp.vue'
import { ISignUpForm, SignUpForm } from '~/types/form'

describe('pages/signup', () => {
  let store: any
  let router: any
  let wrapper: any

  beforeEach(() => {
    store = new Vuex.Store(cloneDeep(storeModule))
    router = { push: jest.fn() }

    wrapper = mount(SignUp, {
      store,
      mocks: {
        $router: router
      },
      stubs: {
        NuxtLink: RouterLinkStub
      }
    })
  })

  afterEach(() => {
    store = null
  })

  describe('script', () => {
    describe('data', () => {
      describe('isError', () => {
        test('isError', () => {
          expect(wrapper.vm.isError).toBeFalsy()
        })

        test('message', () => {
          expect(wrapper.vm.message).toBe('')
        })
      })
    })

    describe('methods', () => {
      let signUpForm: ISignUpForm

      beforeEach(() => {
        store.$axios = axios
        signUpForm = SignUpForm
      })

      describe('doSignUp', () => {
        describe('success', () => {
          beforeEach(() => {
            store.$axios.setSafetyMode(true)
          })

          test('router.pushが実行されること', async () => {
            await wrapper.vm.doSignUp(signUpForm)

            expect(router.push).toBeCalledWith({
              name: 'email-check',
              params: { email: '' }
            })
          })
        })

        describe('failure', () => {
          beforeEach(() => {
            store.$axios.setSafetyMode(false)
          })

          test('isErrorがtrueになること', async () => {
            await wrapper.vm.doSignUp(signUpForm)

            expect(wrapper.vm.isError).toBeTruthy()
          })

          test('messageにエラーメッセージが追加されること', async () => {
            await wrapper.vm.doSignUp(signUpForm)

            expect(wrapper.vm.message).toBe('Error: some error')
          })
        })
      })

      describe('close', () => {
        test('isErrorがfalseになること', async () => {
          await wrapper.vm.close()

          expect(wrapper.vm.isError).toBeFalsy()
        })
      })
    })
  })
})
