<template>
  <gran-group-new @apply="apply" />
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'
import GranGroupNew from '~/components/templates/GranGroupNew.vue'

export default Vue.extend({
  layout: 'header',

  components: {
    GranGroupNew
  },

  methods: {
    apply(groupNewForm) {
      this.create(groupNewForm)
        .then(() => {
          // TODO: 成功した時の処理
        })
        .catch(() => {
          // TODO: 失敗した時の処理
        })
    },
    ...mapActions('group', ['create'])
  }
})
</script>
