<template>
  <v-layout col>
    <gran-kanban-card v-for="(list, index) in lists" :key="index">
      <draggable :list="list.array" class="dragArea list-group">
        <div v-for="(element, i) in list.array" :key="i" class="pa-1 mb-4">
          <gran-task-card />
        </div>
      </draggable>
    </gran-kanban-card>
  </v-layout>
</template>

<script lang="ts">
import Vue from 'vue'
import draggable from 'vuedraggable'
import GranTaskCard from '~/components/molecules/GranTaskCard.vue'
import GranKanbanCard from '~/components/molecules/GranKanbanCard.vue'

export default Vue.extend({
  layout: 'boards',
  components: {
    draggable,
    GranTaskCard,
    GranKanbanCard
  },
  data: () => ({
    lists: [
      {
        array: ['UI設計考える', 'タスク', 'タスク2']
      },
      {
        array: ['UI設計考える2', 'タスク3', 'タスク4']
      }
    ]
  }),
  methods: {
    click() {
      this.$router.push('/boards/1')
    }
  }
})
</script>

<style scoped>
.row {
  overflow-x: auto;
}
</style>
