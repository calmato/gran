import { Configuration } from '@nuxt/types'
import colors from 'vuetify/es5/util/colors'

const configuration: Configuration = {
  mode: 'spa',
  srcDir: 'app',
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || ''
      }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
  },
  loading: { color: '#fff' },
  plugins: [{ src: '~/plugins/vee-validate.ts' }, '~/plugins/firebase'],
  buildModules: [
    '@nuxt/typescript-build',
    '@nuxtjs/eslint-module',
    '@nuxtjs/stylelint-module',
    '@nuxtjs/vuetify'
  ],
  modules: ['@nuxtjs/axios'],
  typescript: {
    typeCheck: {
      eslint: true
    }
  },
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: false,
      themes: {
        dark: {
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
  },
  env: {
    API_KEY: process.env.API_KEY!,
    AUTH_DOMAIN: process.env.AUTH_DOMAIN!,
    DATABASE_URL: process.env.DATABASE_URL!,
    PROJECT_ID: process.env.PROJECT_ID!,
    STORAGE_BUCKET: process.env.STORAGE_BUCKET!,
    MESSAGING_SENDER_ID: process.env.MESSAGING_SENDER_ID!
  }
}

export default configuration
