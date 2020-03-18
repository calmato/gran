<template>
  <gran-account-new />
</template>

<script lang="ts">
import Vue from 'vue'
import GranAccountNew from '~/components/templates/GranAccountNew.vue'

export default Vue.extend({
  layout: 'home',

  components: {
    GranAccountNew
  }
})
</script>
