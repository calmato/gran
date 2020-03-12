<template>
  <gran-board />
</template>

<script lang="ts">
import Vue from 'vue'
import GranBoard from '~/components/templates/GranBoard.vue'

export default Vue.extend({
  layout: 'header',

  components: {
    GranBoard
  }
})
</script>
