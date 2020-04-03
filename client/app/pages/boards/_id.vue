<template>
  <gran-kanban />
</template>

<script lang="ts">
import Vue from 'vue'
import GranKanban from '~/components/templates/GranKanban.vue'

export default Vue.extend({
  layout: 'boards',
  components: {
    GranKanban,
  },
  fetch({ store }) {
    store.dispatch('boards/init')
  },
  data: () => ({
    lists: [
      {
        array: ['UI設計考える', 'タスク', 'タスク2'],
      },
      {
        array: ['UI設計考える2', 'タスク3', 'タスク4'],
      },
    ],
  }),
  methods: {
    click() {
      this.$router.push('/boards/1')
    },
  },
})
</script>
