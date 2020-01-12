import { action } from '@storybook/addon-actions'
import { VBtn } from 'vuetify/lib'

import GranButton from '~/components/atoms/GranButton.vue'
import MyButton from './MyButton.vue'

export default {
  title: 'Button'
}

export const granButton = () => ({
  components: {
    VBtn,
    GranButton
  },

  template: '<gran-button @click="action">ボタン</gran-button>',

  methods: {
    action: action('clicked')
  }
})

export const text = () => ({
  components: {
    MyButton
  },

  template: '<my-button @click="action">Hello Button</my-button>',

  methods: {
    action: action('clicked')
  }
})

export const emoji = () => ({
  components: {
    MyButton
  },

  template: '<my-button @click="action">😀 😎 👍 💯</my-button>',

  methods: {
    action: action('clicked')
  }
})
