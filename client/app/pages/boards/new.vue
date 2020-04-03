<template>
  <div>
    <gran-board-new />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranBoardNew from '~/components/templates/GranBoardNew.vue'

export default Vue.extend({
  layout: 'home',

  components: {
    GranBoardNew,
  },
})
</script>
