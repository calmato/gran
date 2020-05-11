<template>
  <div>
    <gran-board-search />
    <div v-for="group in groups" :key="group.id">
      <gran-boards-list
        :name="group.name"
        :group-id="group.id"
        :boards="group.boards"
        @handleListItemClick="transitionBoardPage"
        @handleAddButton="transitionNewBoardPage"
      />
    </div>
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
    transitionBoardPage(groupId: string, boardId: string): void {
      this.$router.push({ path: `/boards/${boardId}`, query: { groupId } })
    },
    transitionNewBoardPage(_groupId: string): void {
      // todo: groupIDは新規作成の時に使えそうなのでコード上は書いておく
      this.$router.push('/boards/new')
    },
  },
})
</script>
