<template>
  <gran-sign-up @signUp="signUpA" />
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
  methods: {
    ...mapActions('auth', ['signUp']),
    async doSignUp(signUpForm: ISignUpForm): Promise<void> {
      await this.signUp(signUpForm)
        .then(() => {})
        .catch(() => {})
    }
  }
})
</script>
