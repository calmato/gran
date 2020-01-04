<template>
  <gran-login @login="login" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranLogin from '~/components/templates/GranLogin.vue'
import { ILoginForm } from '~/types/form'
import firebase from '~/plugins/firebase'

export default Vue.extend({
  components: {
    GranLogin
  },
  layout: 'auth',
  methods: {
    async login(loginForm: ILoginForm) {
      await firebase
        .auth()
        .signInWithEmailAndPassword(loginForm.email, loginForm.password)
        .then((user) => {
          console.log(user)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }
})
</script>
