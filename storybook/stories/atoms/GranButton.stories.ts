import { action } from '@storybook/addon-actions'
import { text, boolean } from '@storybook/addon-knobs'

import { VBtn } from 'vuetify/lib'

import GranButton from '~/components/atoms/GranButton.vue'

export default {
  title: 'Atoms/Button'
}

export const component = () => ({
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
    },
    icon: {
      default: boolean('Icon', false)
    }
  },

  template: `
<gran-button
  :color="color"
  :block="block"
  :dark="dark"
  :outlined="outlined"
  :icon="icon"
  @click="action"
>
  デフォルトボタン
</gran-button>
`,

  methods: {
    action: action('clicked')
  }
})

export const color = () => ({
  components: {
    VBtn,
    GranButton
  },

  template: `
<div>
  <gran-button color="primary">Primary</gran-button>
  <gran-button color="secondary">Secondary</gran-button>
  <gran-button color="accent">Accent</gran-button>
  <gran-button color="error">Error</gran-button>
  <gran-button color="info">Info</gran-button>
  <gran-button color="success">Success</gran-button>
  <gran-button color="warning">Warning</gran-button>
</div>
`
})
