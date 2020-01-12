<template>
  <gran-card>
    <gran-card-text>
      <gran-sign-up-form @login="login" />
      <gran-link-list />
    </gran-card-text>
  </gran-card>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranSignUpForm from '~/components/molecules/GranSignUpForm.vue'
import GranLinkList from '~/components/organisms/GranLinkList.vue'
import { ISignUpForm } from '~/types/form'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranLinkList,
    GranSignUpForm
  },
  methods: {
    login(loginForm: ISignUpForm) {
      this.$emit('login', loginForm)
    },
    loginWithGoogle() {
      this.$emit('loginWithGoogle')
    }
  }
})
</script>
