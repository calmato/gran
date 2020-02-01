<template>
  <gran-sign-up :is-error="isError" :message="message" @signUp="doSignUp" @close="close" />
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'

import GranSignUp from '~/components/templates/GranSignUp.vue'
import { ISignUpForm } from '~/types/form'

export default Vue.extend({
  layout: 'auth',
  components: {
    GranSignUp
  },
  data: () => ({
    isError: false,
    message: ''
  }),
  methods: {
    ...mapActions('auth', ['signUp']),
    async doSignUp(signUpForm: ISignUpForm): Promise<void> {
      await this.signUp(signUpForm)
        .then(() => {
          this.$router.push({
            name: 'email-check',
            params: { email: signUpForm.email.value }
          })
        })
        .catch((error) => {
          this.isError = true
          this.message = error.message
        })
    },
    close() {
      this.isError = false
    }
  }
})
</script>
