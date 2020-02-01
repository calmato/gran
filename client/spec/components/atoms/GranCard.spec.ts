import Vue from 'vue'
import Vuetify from 'vuetify'
import { mount } from '@vue/test-utils'
import GranCard from '~/components/atoms/GranCard.vue'

Vue.use(Vuetify)

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
