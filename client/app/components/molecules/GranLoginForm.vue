<template>
  <form>
    <gran-text-field
      v-model="loginForm.email"
      prepend-icon="email"
      label="Email"
      :rules="loginFormValidate.email"
      @keydown="onKeydown"
    />
    <gran-text-field
      v-model="loginForm.password"
      prepend-icon="lock"
      label="Password"
      type="password"
      :rules="loginFormValidate.password"
      @keydown="onKeydown"
    />
    <gran-button color="light-blue darken-1" block dark @click="doSubmit">ログイン</gran-button>
  </form>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import { ILoginForm, LoginFormValidate } from '~/types/form'

export default Vue.extend({
  components: {
    GranTextField,
    GranButton,
  },
  data: () => ({
    loginForm: {
      email: '',
      password: '',
    } as ILoginForm,
    loginFormValidate: LoginFormValidate,
  }),
  computed: {
    submitDisabled(): Boolean {
      return this.loginForm.email.length === 0 || this.loginForm.password.length === 0
    },
  },
  methods: {
    doSubmit(): void {
      if (!this.submitDisabled) {
        this.$emit('login', this.loginForm)
      }
    },
    onKeydown(keyEvent: KeyboardEvent): void {
      if (keyEvent.keyCode === 13) this.doSubmit() // KeyCode: 13 => enter
    },
  },
})
</script>
