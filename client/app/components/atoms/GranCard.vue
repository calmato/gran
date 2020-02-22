<template>
  <v-card :color="color" flat>
    <slot />
  </v-card>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    color: {
      type: String,
      default: 'white'
    },
    flat: {
      type: Boolean,
      default: false
    }
  }
})
</script>
