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
    <gran-add-column-form slot="footer" />
  </draggable>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapState } from 'vuex'
import draggable from 'vuedraggable'
import GranTaskColumn from '~/components/organisms/GranTaskColumn.vue'
import GranAddColumnForm from '~/components/molecules/GranAddColumnForm.vue'

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
    getComponentData() {
      return {
        props: {
          col: true
        }
      }
    },
    updateTasks(listIndex, tasks) {
      this.$store.commit('boards/updateBoardTasks', { index: listIndex, value: tasks })
    }
  }
})
</script>
