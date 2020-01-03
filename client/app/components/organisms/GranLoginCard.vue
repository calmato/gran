<template>
  <gran-card>
    <div slot="contents">
      <gran-login-form @login="login" />
      <div class="ma-2 text-center">
        <p>
          <nuxt-link to="singup">新規登録する</nuxt-link>
        </p>
        <p>
          <nuxt-link to="forget">パスワードを忘れた</nuxt-link>
        </p>
      </div>
      <v-divider class="mx-12 ma-4" />
      <gran-social-button name="google" @click="loginWithGoogle"
        >Googleでサインイン
      </gran-social-button>
    </div>
  </gran-card>
</template>

<script>
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranLoginForm from '~/components/molecules/GranLoginForm.vue'
import GranSocialButton from '~/components/molecules/GranSocialButton.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranLoginForm,
    GranSocialButton
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
