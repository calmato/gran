<template>
  <div>
    <gran-error-alert :is-error="isError" :message="message" @close="close" />
    <gran-sign-up-card @signUp="signUp" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranSignUpCard from '~/components/organisms/GranSignUpCard.vue'
import GranErrorAlert from '~/components/molecules/GranErrorAlert.vue'
import { ISignUpForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranSignUpCard,
    GranErrorAlert
  },
  props: {
    isError: {
      type: Boolean,
      default: false
    },
    message: {
      type: String,
      default: 'APIエラー'
    }
  },
  methods: {
    signUp(signUpForm: ISignUpForm) {
      this.$emit('signUp', signUpForm)
    },
    close() {
      this.$emit('close')
    }
  }
})
</script>
