<template>
  <v-tabs centered fixed-tabs background-color="#039be5" dark>
    <v-tab v-for="(tab, index) in tabs" :key="index">
      {{ tab }}
    </v-tab>
  </v-tabs>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    tabs: {
      type: Array,
      default: (): Array<string> => {
        return ['グループ', 'ボード']
      }
    }
  }
})
</script>
