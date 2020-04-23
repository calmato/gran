<template>
  <gran-account-new @apply="apply" />
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'
import GranAccountNew from '~/components/templates/GranAccountEdit.vue'

export default Vue.extend({
  layout: 'home',

  components: {
    GranAccountNew,
  },

  methods: {
    apply(acconntForm) {
      this.$nuxt.$loading.start()

      this.edit(acconntForm)
        .then(() => {
          // TODO:成功した時の処理
        })
        .catch(() => {
          // TODO:失敗した時の処理
        })
        .finally(() => this.$nuxt.$loading.finish())
    },
    ...mapActions('account', ['edit']),
  },
})
</script>
