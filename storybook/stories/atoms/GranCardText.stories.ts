import { VCardText } from 'vuetify/lib'

import GranCardText from '~/components/atoms/GranCardText.vue'

export default {
  title: 'Atoms/CardText'
}

export const component = () => ({
  components: {
    VCardText,
    GranCardText
  },

  template: `<gran-card-text>デフォルトテキスト</gran-card-text>`
})
