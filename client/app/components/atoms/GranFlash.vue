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
    :transition="transition"
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
      validator(val) {
        return ['success', 'info', 'warning', 'error'].includes(val)
      },
      default: undefined
    },
    border: {
      type: String,
      validator(val) {
        return ['top', 'right', 'bottom', 'left'].includes(val)
      },
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
      default: ''
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
    },
    transition: {
      type: String,
      default: ''
    }
  },
  computed: {
    iconName(): string | undefined {
      if (this.icon === '') {
        return undefined
      } else {
        return 'mdi-' + this.icon
      }
    }
  }
})
</script>
