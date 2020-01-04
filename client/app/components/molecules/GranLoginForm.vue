<template>
  <form>
    <gran-text-field v-model="loginForm.email" prepend-icon="email" label="Email..." />
    <gran-text-field
      v-model="loginForm.password"
      prepend-icon="lock"
      label="Password..."
      type="password"
    />
    <gran-button color="light-blue darken-1" block dark @click="click">ログイン</gran-button>
  </form>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import { ILoginForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranTextField,
    GranButton
  },
  data: () => ({
    loginForm: {
      email: '',
      password: ''
    } as ILoginForm
  }),
  methods: {
    click() {
      this.$emit('login', this.loginForm)
    }
  }
})
</script>
