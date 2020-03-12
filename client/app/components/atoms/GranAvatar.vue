<template>
  <figure>
    <img :src="src" width="128" height="128" @click="click" />
  </figure>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    src: {
      type: String,
      default: 'https://cdn.vuetifyjs.com/images/john.jpg'
    }
  },
  methods: {
    click() {
      this.$emit('click')
    }
  }
})
</script>

<style scoped>
figure img {
  border-radius: 15px;
  -webkit-filter: grayscale(0);
  filter: grayscale(0);
  -webkit-transition: 0.3s ease-in-out;
  transition: 0.3s ease-in-out;
}

figure:hover img {
  -webkit-filter: grayscale(100%);
  filter: grayscale(100%);
</style>
