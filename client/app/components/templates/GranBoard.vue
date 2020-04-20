<template>
  <div>
    <gran-board-search />
    <gran-boards-list @click="transition" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranBoardsList from '~/components/molecules/GranBoardsList.vue'
import GranBoardSearch from '~/components/molecules/GranBoardSearch.vue'

export default Vue.extend({
  components: {
    GranBoardSearch,
    GranBoardsList,
  },
  methods: {
    transition(boardId: string): void {
      this.$router.push(`/boards/${boardId}`)
    },
  },
})
</script>
