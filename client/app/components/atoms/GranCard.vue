<template>
  <v-card :color="color" :min-width="minWidth" :max-width="maxWidth" :flat="flat" @click="click">
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
    },
    minWidth: {
      type: Number,
      default: 100
    },
    maxWidth: {
      type: Number,
      default: 350
    }
  },
  methods: {
    click() {
      this.$emit('click')
    }
  }
})
</script>
