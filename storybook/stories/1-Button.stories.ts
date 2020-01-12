import { action } from '@storybook/addon-actions'
import { boolean } from '@storybook/addon-knobs'

import { VBtn } from 'vuetify/lib'

import GranButton from '~/components/atoms/GranButton.vue'

export default {
  title: 'Atoms/Button'
}

export const granButton = () => ({
  components: {
    VBtn,
    GranButton
  },

  props: {
    block: {
      default: boolean('Block', false)
    },
  },

  template: `
    <div>
      <gran-button color="primary" :block="block" @click="action">Primary</gran-button>
      <gran-button color="secondary" :block="block" @@click="action">Secondary</gran-button>
      <gran-button color="accent" :block="block" @click="action">Accent</gran-button>
      <gran-button color="error" :block="block" @click="action">Error</gran-button>
      <gran-button color="info" :block="block" @click="action">Info</gran-button>
      <gran-button color="success" :block="block" @click="action">Success</gran-button>
      <gran-button color="warning" :block="block" @click="action">Warning</gran-button>
    </div>
  `,

  methods: {
    action: action('clicked')
  }
})
