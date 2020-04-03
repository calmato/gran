<template>
  <div>
    <gran-thumbnail-form />
    <gran-board-text-form />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranThumbnailForm from '~/components/molecules/GranThumbnailForm.vue'
import GranBoardTextForm from '~/components/molecules/GranBoardTextForm.vue'

export default Vue.extend({
  components: {
    GranThumbnailForm,
    GranBoardTextForm,
  },
})
</script>
