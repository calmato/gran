<template>
  <gran-card>
    <gran-card-text>
      <gran-login-form @login="login" />
      <gran-link-list />
      <v-divider class="mx-12 ma-4" />
      <gran-social-button name="google" @click="loginWithGoogle"
        >Googleでサインイン
      </gran-social-button>
    </gran-card-text>
  </gran-card>
</template>

<script>
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranLoginForm from '~/components/molecules/GranLoginForm.vue'
import GranSocialButton from '~/components/molecules/GranSocialButton.vue'
import GranLinkList from '~/components/organisms/GranLinkList.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranLoginForm,
    GranSocialButton,
    GranLinkList
  },
  methods: {
    login() {
      this.$emit('login')
    },
    loginWithGoogle() {
      this.$emit('loginWithGoogle')
    }
  }
})
</script>
