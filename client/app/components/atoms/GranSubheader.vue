<template>
  <v-subheader>
    <slot />
  </v-subheader>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    ccolor: {
      type: String,
      default: '#f5f5f5',
    },
  },
})
</script>

<style scoped>
.v-subheader {
  --color: #f5f5f5;

  background-color: var(--color);
}
</style>
