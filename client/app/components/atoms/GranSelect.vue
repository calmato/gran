<template>
  <v-select :label="label" :items="items" />
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    label: {
      type: String,
      default: ''
    },
    items: {
      type: Array,
      default: () => ['16frances', 'gran']
    }
  }
})
</script>
