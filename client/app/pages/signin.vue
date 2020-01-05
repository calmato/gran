<template>
  <gran-login :is-error="isError" @login="login" @close="close" />
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'
import GranLogin from '~/components/templates/GranLogin.vue'
import firebase from '~/plugins/firebase'
import { ILoginForm } from '~/types/form'
import { IUser } from '~/types/store'

export default Vue.extend({
  components: {
    GranLogin
  },
  layout: 'auth',
  data: () => ({
    isError: false
  }),
  methods: {
    ...mapActions({
      setUser: 'auth/setUser'
    }),
    async login(loginForm: ILoginForm) {
      await firebase
        .auth()
        .signInWithEmailAndPassword(loginForm.email, loginForm.password)
        .then((auth: any) => {
          const user: IUser = {
            uid: auth.user.uid,
            email: auth.user.email,
            creationTime: auth.user.metadata.creationTime,
            lastSignInTime: auth.user.metadata.lastSignInTime
          }
          this.setUser(user)
        })
        .catch(() => {
          // firebaseとの接続エラーだとまずいな
          this.isError = true
        })
    },
    close() {
      this.isError = false
    }
  }
})
</script>
