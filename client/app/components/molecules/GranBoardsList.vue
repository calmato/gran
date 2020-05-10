<template>
  <gran-list>
    <gran-subheader>
      <gran-icon name="account" x-large />
      {{ name }}
    </gran-subheader>
    <div v-for="(board, index) in boards" :key="board.id">
      <gran-list-item color="#ffffff" exact @click="handleListItemClick(board.id)">
        <gran-list-item-content>{{ board.name }}</gran-list-item-content>
      </gran-list-item>
      <v-divider v-if="!isLast(index)" />
      <gran-button
        v-if="isLast(index)"
        color="#757575"
        text
        block
        @click="handleAddButton(groupId)"
      >
        <gran-icon name="plus" />
        Add New Board
      </gran-button>
    </div>
  </gran-list>
</template>

<script lang="ts">
import Vue from 'vue'
import GranList from '~/components/atoms/GranList.vue'
import GranListItem from '~/components/atoms/GranListItem.vue'
import GranListItemContent from '~/components/atoms/GranListItemContent.vue'
import GranSubheader from '~/components/atoms/GranSubheader.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranIcon,
    GranList,
    GranListItem,
    GranListItemContent,
    GranSubheader,
    GranButton,
  },
  props: {
    name: {
      type: String,
      default: 'パーソナルボード',
    },
    groupId: {
      type: String,
      default: 'group-id',
    },
    boards: {
      type: Array,
      default: () => {
        return [
          {
            id: 'abcd',
            name: 'テスト',
          },
          {
            id: 'abc',
            name: 'テスト',
          },
        ]
      },
    },
  },
  methods: {
    handleListItemClick(boardId: string): void {
      this.$emit('handleListItemClick', this.groupId, boardId)
    },
    handleAddButton(groupId: string): void {
      this.$emit('handleAddButton', groupId)
    },
    isLast(index: number): Boolean {
      return index === this.boards.length - 1
    },
  },
})
</script>

<style scoped>
.v-btn {
  text-transform: none;
}
</style>
