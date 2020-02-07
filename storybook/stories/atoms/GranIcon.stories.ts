import { text, boolean } from '@storybook/addon-knobs'

import { VIcon } from 'vuetify/lib'

import GranIcon from '~/components/atoms/GranIcon.vue'

export default {
  title: 'Atoms/Icon'
}

export const component = () => ({
  components: {
    VIcon,
    GranIcon
  },

  props: {
    name: {
      default: text('Name', 'email')
    },
    left: {
      default: boolean('Left', false)
    },
    right: {
      default: boolean('Right', false)
    },
    dark: {
      default: boolean('Dark', false)
    },
    xLarge: {
      default: boolean('XLarge', false)
    }
  },

  template: `
<gran-icon
  :name="name"
  :left="left"
  :right="right"
  :dark="dark"
  :x-large="xLarge"
/>
`
})
