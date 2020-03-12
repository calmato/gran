<template>
  <v-tabs
    :centered="centered"
    :fixed-tabs="fixedTabs"
    :dark="dark"
    :background-color="backgroundColor"
  >
    <slot />
  </v-tabs>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    centered: {
      type: Boolean,
      default: false
    },
    fixedTabs: {
      type: Boolean,
      default: false
    },
    dark: {
      type: Boolean,
      default: false
    },
    backgroundColor: {
      type: String,
      default: '#039be5'
    }
  }
})
</script>
