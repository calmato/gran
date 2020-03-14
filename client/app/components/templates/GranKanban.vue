<template>
  <v-layout col>
    <draggable
      v-for="column in board.columns"
      :key="column.id"
      group="column"
      :component-data="getComponentData(column)"
      tag="gran-task-column"
    />
    <gran-add-column-form />
  </v-layout>
</template>

<script lang="ts">
import Vue from 'vue'
import draggable from 'vuedraggable'
import GranAddColumnForm from '~/components/molecules/GranAddColumnForm.vue'

export default Vue.extend({
  components: {
    draggable,
    GranAddColumnForm
  },
  props: {
    board: {
      type: Object,
      default: () => {
        return {
          columns: [
            {
              id: 1,
              name: 'ToDo',
              color: '#FF80AB',
              tasks: [
                {
                  id: 1,
                  title: 'UI設計',
                  labels: []
                },
                {
                  id: 2,
                  title: 'APIの実装',
                  labels: [{ name: 'client', color: 'primary' }]
                },
                {
                  id: 3,
                  title: 'Figma',
                  labels: []
                }
              ]
            },
            {
              id: 2,
              name: 'In Progless',
              color: '#26A69A',
              tasks: [
                {
                  id: 1,
                  title: 'バグ修正',
                  labels: []
                },
                {
                  id: 2,
                  title: '本番環境の構築',
                  labels: [{ name: 'client', color: 'yellow' }]
                }
              ]
            }
          ]
        }
      }
    }
  },
  methods: {
    getComponentData(column) {
      return {
        props: {
          column
        }
      }
    }
  }
})
</script>
