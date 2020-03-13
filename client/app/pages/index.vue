<template>
  <v-layout column justify-center align-center>
    <h1>今日の予定</h1>
    <v-divider />
    <gran-task-card @click="click" />
  </v-layout>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTaskCard from '~/components/molecules/GranTaskCard.vue'

export default Vue.extend({
  layout: 'home',
  components: {
    GranTaskCard
  },
  data: () => ({}),
  methods: {
    click() {
      this.$router.push('/boards/1')
    }
  }
})
</script>
