<template>
  <gran-login-card @login="login" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranLoginCard from '~/components/organisms/GranLoginCard.vue'
import { ILoginForm } from '~/types/form/'

export default Vue.extend({
  components: {
    GranLoginCard
  },
  methods: {
    login(loginForm: ILoginForm) {
      this.$emit('login', loginForm)
    }
  }
})
</script>
