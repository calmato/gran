import { action } from '@storybook/addon-actions'

import MyButton from './MyButton.vue'

export default {
  title: 'Button'
}

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
