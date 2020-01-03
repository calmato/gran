<template>
  <form>
    <gran-text-field prepend-icon="email" label="Email..." />
    <gran-text-field prepend-icon="lock" label="Password..." />
    <gran-button color="light-blue darken-1" block dark @click="click">ログイン</gran-button>
  </form>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranTextField,
    GranButton
  },
  methods: {
    click() {
      this.$emit('login')
    }
  }
})
</script>
