<template>
  <v-chip :color="color" :x-small="xSmall">
    <slot />
  </v-chip>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    color: {
      type: String,
      default: 'primary'
    },
    xSmall: {
      type: Boolean,
      default: false
    }
  }
})
</script>
