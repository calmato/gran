<template>
  <v-card
    :color="color"
    :width="width"
    :min-width="minWidth"
    :max-width="maxWidth"
    :flat="flat"
    @click.native="click"
  >
    <slot />
  </v-card>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    color: {
      type: String,
      default: 'white',
    },
    flat: {
      type: Boolean,
      default: false,
    },
    width: {
      type: Number,
      default: undefined,
    },
    minWidth: {
      type: Number,
      default: undefined,
    },
    maxWidth: {
      type: Number,
      default: undefined,
    },
  },
  methods: {
    click(): void {
      this.$emit('click')
    },
  },
})
</script>
