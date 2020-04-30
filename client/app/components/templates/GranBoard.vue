<template>
  <div>
    <gran-board-search />
    <gran-boards-list @click="transitionBoardPage" @handleAddButton="transitionNewBoardPage" />
    <gran-boards-list name="グループ" @click="transitionBoardPage" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import GranBoardsList from '~/components/molecules/GranBoardsList.vue'
import GranBoardSearch from '~/components/molecules/GranBoardSearch.vue'

export default Vue.extend({
  components: {
    GranBoardSearch,
    GranBoardsList,
  },
  computed: {
    ...mapGetters('group', ['groups']),
  },
  methods: {
    transitionBoardPage(boardId: string): void {
      this.$router.push(`/boards/${boardId}`)
    },
    transitionNewBoardPage(_groupId: string): void {
      // todo: groupIDは新規作成の時に使えそうなのでコード上は書いておく
      this.$router.push('/boards/new')
    },
  },
})
</script>
