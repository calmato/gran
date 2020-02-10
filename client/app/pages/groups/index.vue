<template>
  <gran-group-list />
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupList from '~/components/organisms/GranGroupList.vue'

export default Vue.extend({
  layout: 'header',

  components: {
    GranGroupList
  },

  async fetch({ store }) {
    await store.dispatch('group/groupAll').catch((err: any) => {
    })
  }
})
</script>
