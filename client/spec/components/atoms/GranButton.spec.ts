import { mount } from '@vue/test-utils'
import '~~/spec/helpers/component-helper'
import GranButton from '~/components/atoms/GranButton.vue'

describe('components/atoms/GranButton', () => {
  let wrapper

  beforeEach(() => {
    wrapper = mount(GranButton)
  })

  describe('script', () => {
    describe('props', () => {
      describe('color', () => {
        test('color: 初期値', () => {
          expect(wrapper.props().color).toBe('primary')
        })

        test('color: 代入', () => {
          wrapper.setProps({ color: 'secondary' })
          expect(wrapper.props().color).toBe('secondary')
        })
      })

      describe('block', () => {
        test('block: 初期値', () => {
          expect(wrapper.props().block).toBeFalsy()
        })

        test('block: 代入', () => {
          wrapper.setProps({ block: true })
          expect(wrapper.props().block).toBeTruthy()
        })
      })

      describe('dark', () => {
        test('dark: 初期値', () => {
          expect(wrapper.props().dark).toBeFalsy()
        })

        test('dark: 代入', () => {
          wrapper.setProps({ dark: true })
          expect(wrapper.props().dark).toBeTruthy()
        })
      })

      describe('outlined', () => {
        test('outlined: 初期値', () => {
          expect(wrapper.props().outlined).toBeFalsy()
        })

        test('outlined: 代入', () => {
          wrapper.setProps({ outlined: true })
          expect(wrapper.props().outlined).toBeTruthy()
        })
      })

      describe('icon', () => {
        test('icon: 初期値', () => {
          expect(wrapper.props().icon).toBeFalsy()
        })

        test('icon: 代入', () => {
          wrapper.setProps({ icon: true })
          expect(wrapper.props().icon).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('click', () => {
        test('emitが正常に動作すること', async () => {
          await wrapper.vm.click()
          expect(wrapper.emitted().click).toBeTruthy()
        })
      })
    })
  })
})
