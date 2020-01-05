<template>
  <div>
    <gran-login-alert :is-error="isError" @close="close" />
    <gran-login-card @login="login" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranLoginCard from '~/components/organisms/GranLoginCard.vue'
import GranLoginAlert from '~/components/molecules/GranLoginAlert.vue'
import { ILoginForm } from '~/types/form/'

export default Vue.extend({
  components: {
    GranLoginCard,
    GranLoginAlert
  },
  props: {
    isError: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    login(loginForm: ILoginForm) {
      this.$emit('login', loginForm)
    },
    close() {
      this.$emit('close')
    }
  }
})
</script>
