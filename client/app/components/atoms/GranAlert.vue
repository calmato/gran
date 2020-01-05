<template>
  <v-alert
    :type="type"
    :border="border"
    :color="color"
    :colored-border="coloredBorder"
    :icon="iconName"
    :dense="dense"
    :text="text"
    :dark="dark"
    :outlined="outlined"
    :prominent="prominent"
  >
    <slot />
  </v-alert>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    type: {
      type: String,
      in: ['success', 'info', 'warning', 'error'],
      default: undefined
    },
    border: {
      type: String,
      in: ['top', 'right', 'bottom', 'left'],
      default: undefined
    },
    coloredBorder: {
      type: Boolean,
      default: false
    },
    color: {
      type: String,
      default: ''
    },
    icon: {
      type: String,
      default: undefined
    },
    dense: {
      type: Boolean,
      default: false
    },
    text: {
      type: Boolean,
      default: false
    },
    dark: {
      type: Boolean,
      default: false
    },
    outlined: {
      type: Boolean,
      default: false
    },
    prominent: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    iconName(): string | undefined {
      if (!this.icon === undefined) return 'mdi-' + this.icon
      else return undefined
    }
  }
})
</script>
