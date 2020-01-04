<template>
  <gran-login />
</template>

<script lang="ts">
import Vue from 'vue'
import GranLogin from '~/components/templates/GranLogin.vue'

export default Vue.extend({
  components: {
    GranLogin
  },
  layout: 'auth'
})
</script>
