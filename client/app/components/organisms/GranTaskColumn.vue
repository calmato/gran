<template>
  <div class="ma-1">
    <gran-kanban-card :name="column.name" :length="tasksLength" :color="column.color">
      <draggable :value="value" group="task" @input="emitter">
        <div v-for="task in column.tasks" :key="task.id" class="pa-1 mb-1">
          <gran-task-card :task="task" />
        </div>
      </draggable>
    </gran-kanban-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import draggable from 'vuedraggable'
import GranTaskCard from '~/components/molecules/GranTaskCard.vue'
import GranKanbanCard from '~/components/molecules/GranKanbanCard.vue'

export default Vue.extend({
  components: {
    draggable,
    GranTaskCard,
    GranKanbanCard,
  },
  props: {
    listIndex: {
      type: Number,
      default: 0,
    },
    value: {
      required: false,
      type: Array,
      default: null,
    },
    column: {
      type: Object,
      default: () => {
        return {
          name: 'ToDo',
          color: '#FF80AB',
          tasks: [
            {
              id: 1,
              name: 'UI設計',
              labels: [],
            },
            {
              id: 2,
              name: 'APIの実装',
              labels: [{ name: 'client', color: 'primary' }],
            },
          ],
        }
      },
    },
  },
  computed: {
    tasksLength(): number {
      return this.column.tasks.length
    },
  },
  methods: {
    emitter(value) {
      this.$emit('input', this.listIndex, value)
    },
  },
})
</script>
