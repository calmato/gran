<template>
  <gran-sign-up @signUp="signUp" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranSignUp from '~/components/templates/GranSignUp.vue'
import { ISignUpForm } from '~/types/form'

export default Vue.extend({
  layout: 'auth',
  components: {
    GranSignUp
  },
  methods: {
    signUp(signUpForm: ISignUpForm) {
      console.log(signUpForm)
    }
  }
})
</script>
