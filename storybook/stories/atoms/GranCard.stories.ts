import { VCard } from 'vuetify/lib'

import GranCard from '~/components/atoms/GranCard.vue'

export default {
  title: 'Atoms/Card'
}

export const component = () => ({
  components: {
    VCard,
    GranCard
  },

  template: `<gran-card>デフォルトテキスト</gran-card>`
})
