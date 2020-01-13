<template>
  <gran-sign-up-card @signUp="signUp" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranSignUpCard from '~/components/organisms/GranSignUpCard.vue'
import { ISignUpForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranSignUpCard
  },
  methods: {
    signUp(signUpForm: ISignUpForm) {
      this.$emit('signUp', signUpForm)
    }
  }
})
</script>
