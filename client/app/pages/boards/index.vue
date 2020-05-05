<template>
  <gran-board />
</template>

<script lang="ts">
import Vue from 'vue'
import GranBoard from '~/components/templates/GranBoard.vue'

export default Vue.extend({
  layout: 'home',

  components: {
    GranBoard,
  },

  async fetch({ store }) {
    await store.dispatch('group/groupAll').catch((_error) => {})
  },
})
</script>
