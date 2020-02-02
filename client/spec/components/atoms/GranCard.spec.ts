import { mount } from '@vue/test-utils'
import '~~/spec/helpers/component-helper'
import GranCard from '~/components/atoms/GranCard.vue'

describe('components/atoms/GranCard', () => {
  let wrapper

  beforeEach(() => {
    wrapper = mount(GranCard)
  })

  /*
   * テスト箇所がないため、仮のテストを配置
   */
  test('componentsの読み込み', () => {
    expect(wrapper).not.toBeNull()
  })
})
