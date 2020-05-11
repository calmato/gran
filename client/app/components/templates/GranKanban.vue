<template>
  <draggable v-model="columns" :component-data="getComponentData()" tag="v-layout">
    <gran-task-column
      v-for="(column, index) in board.lists"
      :key="column.id"
      :value="columns[index].tasks"
      :list-index="index"
      :column="column"
      @input="updateTasks"
      @addTask="addTask"
    />
    <gran-add-column-form slot="footer" @addColumn="addColumn" />
  </draggable>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters, mapActions } from 'vuex'
import draggable from 'vuedraggable'
import GranAddColumnForm from '~/components/molecules/GranAddColumnForm.vue'
import GranTaskColumn from '~/components/organisms/GranTaskColumn.vue'

import { IBoardListForm } from '~/types/form'

export default Vue.extend({
  components: {
    draggable,
    GranAddColumnForm,
    GranTaskColumn,
  },
  computed: {
    ...mapGetters('boards', ['board', 'lists']),
    columns: {
      get() {
        return this.board.lists
      },
      set(value: any) {
        this.$store.commit('boards/updateBoardLists', value)
      },
    },
  },
  watch: {
    lists: {
      handler(_val, _old) {
        this.updateKanban()
      },
      deep: true,
    },
  },
  methods: {
    ...mapActions('boards', ['addNewColumn', 'addNewTask', 'updateKanban']),
    getComponentData(): Object {
      return {
        props: {
          col: true,
        },
      }
    },
    updateTasks(listIndex, tasks): void {
      this.$store.commit('boards/updateBoardTasks', { index: listIndex, value: tasks })
    },
    addColumn(formData: IBoardListForm): void {
      this.addNewColumn(formData)
    },
    addTask(index, value): void {
      this.addNewTask({ index, value })
    },
  },
})
</script>
