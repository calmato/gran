import { text, boolean } from '@storybook/addon-knobs'

import { VAlert } from 'vuetify/lib'

import GranFlash from '~/components/atoms/GranFlash.vue'

export default {
  title: 'Atoms/Flash'
}

export const component = () => ({
  components: {
    VAlert,
    GranFlash
  },

  props: {
    type: {
      default: text('Type', '')
    },
    border: {
      default: text('Border', '')
    },
    color: {
      default: text('Color', '')
    },
    coloredBorder: {
      default: boolean('ColoredBorder', false)
    },
    icon: {
      default: text('Icon', '')
    },
    dense: {
      default: boolean('Dense', false)
    },
    text: {
      default: boolean('Text', false)
    },
    dark: {
      default: boolean('Dark', false)
    },
    outlined: {
      default: boolean('Outlined', false)
    },
    prominent: {
      default: boolean('Prominent', false)
    },
    transition: {
      default: text('Transition', '')
    },
  },

  template: `
<gran-flash
  :type="type"
  :border="border"
  :color="color"
  :colored-border="coloredBorder"
  :icon="icon"
  :dense="dense"
  :text="text"
  :dark="dark"
  :outlined="outlined"
  :prominent="prominent"
  :transition="transition"
>
  デフォルトテキスト
</gran-flash>
`
})

export const type = () => ({
  components: {
    VAlert,
    GranFlash
  },

  template: `
<div>
  <gran-flash type="success">Success</gran-flash>
  <gran-flash type="info">Info</gran-flash>
  <gran-flash type="error">Error</gran-flash>
  <gran-flash type="warning">Warning</gran-flash>
</div>
`
})

export const border = () => ({
  components: {
    VAlert,
    GranFlash
  },

  props: {
    type: {
      default: text('Type', 'info')
    }
  },

  template: `
<div>
  <gran-flash :type="type" border="top">Top</gran-flash>
  <gran-flash :type="type" border="right">Right</gran-flash>
  <gran-flash :type="type" border="bottom">Bottom</gran-flash>
  <gran-flash :type="type" border="left">Left</gran-flash>
</div>
`
})
