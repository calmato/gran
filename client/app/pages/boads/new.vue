<template>
  <div>
    <gran-board-new />
    <gran-avatar />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranBoardNew from '~/components/templates/GranBoardNew.vue'
import GranAvatar from '~/components/atoms/GranAvatar.vue'

export default Vue.extend({
  layout: 'header',

  components: {
    GranBoardNew,
    GranAvatar
  }
})
</script>
