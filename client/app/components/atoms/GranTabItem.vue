<template>
  <v-tab>
    <slot />
    {{ name }}
  </v-tab>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    name: {
      type: String,
      default: 'default',
    },
  },
})
</script>
