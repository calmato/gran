<template>
  <v-list-item :exact="exact" class="tile" :style="{ '--bg-color': color }" @click="click">
    <slot />
  </v-list-item>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    exact: {
      type: Boolean,
      default: false,
    },
    color: {
      type: String,
      default: 'green',
    },
  },
  methods: {
    click(): void {
      this.$emit('click')
    },
  },
})
</script>

<style scoped>
.tile {
  --bg-color: #f5f5f5;

  background-color: var(--bg-color);
}
</style>
