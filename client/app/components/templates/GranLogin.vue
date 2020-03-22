<template>
  <div>
    <gran-error-alert :is-error="isError" @close="close" />
    <gran-login-card @login="login" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranLoginCard from '~/components/organisms/GranLoginCard.vue'
import GranErrorAlert from '~/components/molecules/GranErrorAlert.vue'
import { ILoginForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranLoginCard,
    GranErrorAlert,
  },
  props: {
    isError: {
      type: Boolean,
      default: false,
    },
  },
  methods: {
    login(loginForm: ILoginForm) {
      this.$emit('login', loginForm)
    },
    close() {
      this.$emit('close')
    },
  },
})
</script>
