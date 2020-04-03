<template>
  <gran-button color="gray" block outlined @click="click">
    <gran-icon :name="name" left />
    <slot />
  </gran-button>
</template>

<script lang="ts">
import Vue from 'vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranIcon,
    GranButton,
  },
  props: {
    name: {
      type: String,
      default: 'google',
    },
  },
  methods: {
    click() {
      this.$emit('click')
    },
  },
})
</script>

<style scoped>
.v-btn {
  text-transform: none;
}
</style>
