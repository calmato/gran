<template>
  <v-icon :left="left" :right="right" :dark="dark" :x-large="xLarge">{{ iconName }}</v-icon>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    name: {
      type: String,
      default: 'email',
    },
    left: {
      type: Boolean,
      default: false,
    },
    right: {
      type: Boolean,
      default: false,
    },
    dark: {
      type: Boolean,
      default: false,
    },
    xLarge: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    iconName() {
      return 'mdi-' + this.name
    },
  },
})
</script>
