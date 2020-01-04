<template>
  <gran-login @login="login" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranLogin from '~/components/templates/GranLogin.vue'
import { ILoginForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranLogin
  },
  layout: 'auth',
  methods: {
    login(loginForm: ILoginForm) {
      // ここでvuexと通信する
      console.log(loginForm)
    }
  }
})
</script>
