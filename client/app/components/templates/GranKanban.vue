<template>
  <draggable v-model="lists" :component-data="getComponentData()" tag="v-layout">
    <gran-task-column
      v-for="(column, index) in board.lists"
      :key="column.id"
      :value="lists[index].tasks"
      :list-index="index"
      :column="column"
      @input="updateTasks"
    />
    <gran-add-column-form slot="footer" @addColumn="addColumn" />
  </draggable>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapState, mapActions } from 'vuex'
import draggable from 'vuedraggable'
import GranAddColumnForm from '~/components/molecules/GranAddColumnForm.vue'
import GranTaskColumn from '~/components/organisms/GranTaskColumn.vue'

import { IBoardListForm } from '~/types/form'

export default Vue.extend({
  components: {
    draggable,
    GranAddColumnForm,
    GranTaskColumn
  },
  computed: {
    ...mapState('boards', ['board']),
    lists: {
      get() {
        return this.board.lists
      },
      set(value: any) {
        this.$store.commit('boards/updateBoardLists', value)
      }
    }
  },
  methods: {
    ...mapActions('boards', ['addNewColumn']),
    getComponentData(): Object {
      return {
        props: {
          col: true
        }
      }
    },
    updateTasks(listIndex, tasks): void {
      this.$store.commit('boards/updateBoardTasks', { index: listIndex, value: tasks })
    },
    addColumn(formData: IBoardListForm): void {
      this.addNewColumn(formData)
    }
  }
})
</script>
