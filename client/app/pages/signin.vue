<template>
  <gran-login :is-error="isError" @login="login" @close="close" />
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'
import GranLogin from '~/components/templates/GranLogin.vue'
import { ILoginForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranLogin
  },
  layout: 'header',
  data: () => ({
    isError: false
  }),
  methods: {
    ...mapActions('auth', ['loginWithEmailAndPassword']),
    async login(loginForm: ILoginForm) {
      await this.loginWithEmailAndPassword(loginForm)
        .then(() => {
          this.$router.push('/') // 仮のルーティング
        })
        .catch(() => {
          this.isError = true
        })
    },
    close() {
      this.isError = false
    }
  }
})
</script>
