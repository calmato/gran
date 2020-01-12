import { action } from '@storybook/addon-actions'
import { text, boolean } from '@storybook/addon-knobs'

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
    color: {
      default: text('Color', 'primary')
    },
    block: {
      default: boolean('Block', false)
    },
    dark: {
      default: boolean('Dark', false)
    },
    outlined: {
      default: boolean('Outlined', false)
    }
  },

  template: `
<div>
  <gran-button
    :color="color"
    :block="block"
    :dark="dark"
    :outlined="outlined"
    @click="action"
  >
    デフォルト
  </gran-button>
  <div>
    <gran-button color="primary">Primary</gran-button>
    <gran-button color="secondary">Secondary</gran-button>
    <gran-button color="accent">Accent</gran-button>
    <gran-button color="error">Error</gran-button>
    <gran-button color="info">Info</gran-button>
    <gran-button color="success">Success</gran-button>
    <gran-button color="warning">Warning</gran-button>
  </div>
</div>
  `,

  methods: {
    action: action('clicked')
  }
})
