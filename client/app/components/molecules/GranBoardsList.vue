<template>
  <gran-list>
    <gran-subheader>
      <gran-icon name="account" x-large />
      {{ name }}
    </gran-subheader>
    <div v-for="(board, index) in boards" :key="board.id">
      <gran-list-item color="#ffffff" exact @click="click(board.id)">
        <gran-list-item-content>
          {{ board.name }}
        </gran-list-item-content>
      </gran-list-item>
      <v-divider v-if="index === boards.length" />
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

export default Vue.extend({
  components: {
    GranIcon,
    GranList,
    GranListItem,
    GranListItemContent,
    GranSubheader,
  },
  props: {
    name: {
      type: String,
      default: 'パーソナルボード',
    },
    boards: {
      type: Array,
      default: () => {
        return [
          {
            id: 'abcd',
            name: 'テスト',
          },
        ]
      },
    },
  },
  methods: {
    click(boardId: string): void {
      this.$emit('click', boardId)
    },
  },
})
</script>
