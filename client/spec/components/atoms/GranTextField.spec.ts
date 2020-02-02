import { mount } from '@vue/test-utils'
import '~~/spec/helpers/component-helper'
import GranTextField from '~/components/atoms/GranTextField.vue'
import { ILoginFormValidate, LoginFormValidate } from '~/types/form'

describe('components/atoms/GranTextField', () => {
  let wrapper: any
  let rules: ILoginFormValidate

  beforeEach(() => {
    wrapper = mount(GranTextField)

    rules = LoginFormValidate
  })

  describe('script', () => {
    describe('appendIcon', () => {
      test('appendIcon: 初期値', () => {
        expect(wrapper.props().appendIcon).toBe('')
      })

      test('appendIcon: 代入', () => {
        wrapper.setProps({ appendIcon: 'home' })
        expect(wrapper.props().appendIcon).toBe('home')
      })
    })

    describe('label', () => {
      test('label: 初期値', () => {
        expect(wrapper.props().label).toBe('')
      })

      test('label: 代入', () => {
        wrapper.setProps({ label: 'test' })
        expect(wrapper.props().label).toBe('test')
      })
    })

    describe('name', () => {
      test('name: 初期値', () => {
        expect(wrapper.props().name).toBe('')
      })

      test('name: 代入', () => {
        wrapper.setProps({ name: 'test' })
        expect(wrapper.props().name).toBe('test')
      })
    })

    describe('prependIcon', () => {
      test('prependIcon: 初期値', () => {
        expect(wrapper.props().prependIcon).toBe('')
      })

      test('prependIcon: 代入', () => {
        wrapper.setProps({ prependIcon: 'home' })
        expect(wrapper.props().prependIcon).toBe('home')
      })
    })

    describe('rules', () => {
      test('rules: 初期値', () => {
        expect(wrapper.props().rules).toEqual(undefined)
      })

      test('rules: 代入', () => {
        wrapper.setProps({ rules })
        expect(wrapper.props().rules).toEqual(rules)
      })
    })

    describe('types', () => {
      test('types: 初期値', () => {
        expect(wrapper.props().type).toBe('text')
      })

      test('types: 代入', () => {
        wrapper.setProps({ type: 'test' })
        expect(wrapper.props().type).toBe('test')
      })
    })

    describe('value', () => {
      test('value: 初期値', () => {
        expect(wrapper.props().value).toBe('')
      })

      test('value: 代入', () => {
        wrapper.setProps({ value: 'input' })
        expect(wrapper.props().value).toBe('input')
      })
    })

    describe('computed', () => {
      describe('formData', () => {
        beforeEach(() => {
          wrapper.vm.$emit('input', 'hoge')
        })

        test('set: emitが正常に動作すること', () => {
          expect(wrapper.emitted().input).toBeTruthy()
        })

        test('set: emitで値が更新されること', () => {
          wrapper.vm.$emit('input', 'test')
          expect(wrapper.emitted().input[1]).toEqual(['test'])
        })

        test('get: 値が取得できること', () => {
          expect(wrapper.vm.formData).toBe('')
        })
      })

      describe('prependIconName', () => {
        test('prependIconName: 初期値', () => {
          expect(wrapper.vm.prependIconName).toBe('')
        })

        test('prependIconName: 代入', () => {
          wrapper.setProps({ prependIcon: 'home' })
          expect(wrapper.vm.prependIconName).toBe('mdi-home')
        })
      })

      describe('appendIconName', () => {
        test('appendIconName: 初期値', () => {
          expect(wrapper.vm.appendIconName).toBe('')
        })

        test('appendIconName: 代入', () => {
          wrapper.setProps({ appendIcon: 'home' })
          expect(wrapper.vm.appendIconName).toBe('mdi-home')
        })
      })
    })
  })
})
