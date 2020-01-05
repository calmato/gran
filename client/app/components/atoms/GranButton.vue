<template>
  <v-btn
    :color="color"
    :block="block"
    :dark="dark"
    :outlined="outlined"
    :icon="icon"
    @click="click"
  >
    <slot />
  </v-btn>
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
  props: {
    color: {
      type: String,
      default: 'primary'
    },
    block: {
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
    icon: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    click() {
      this.$emit('click')
    }
  }
})
</script>
