import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  title: 'Logo'
}

export const logo = () => ({
  components: {
    Logo
  },

  template: '<logo />'
})

export const vuetifyLogo = () => ({
  components: {
    VuetifyLogo
  },

  template: '<vuetify-logo />'
})
