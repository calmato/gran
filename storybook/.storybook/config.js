import { configure, addDecorator } from '@storybook/vue'
import { withKnobs } from '@storybook/addon-knobs'

// automatically import all files ending in *.stories.js, *.stories.ts
configure(require.context('./../stories', true, /\.stories\.(js|ts)$/), module)

import Vue from 'vue'
import Vuetify, { VApp, VContent } from 'vuetify/lib'
import colors from 'vuetify/lib/util/colors'
import 'vuetify/src/styles/main.sass'

Vue.use(Vuetify)

const opts = {
  theme: {
    dark: false,
    themes: {
      light: {
        primary: colors.blue.darken2,
        accent: colors.grey.darken3,
        secondary: colors.amber.darken3,
        info: colors.teal.lighten1,
        warning: colors.amber.base,
        error: colors.deepOrange.accent4,
        success: colors.green.accent3
      }
    }
  }
}

addDecorator(withKnobs)
addDecorator(() => ({
  vuetify: new Vuetify(opts),

  components: {
    VApp,
    VContent
  },

  template: '<v-app><v-content><story/></v-content></v-app>'
}))
