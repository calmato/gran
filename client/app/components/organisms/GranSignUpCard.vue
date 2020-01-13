<template>
  <gran-card>
    <gran-card-text>
      <gran-sign-up-form @signUp="signUp" />
      <gran-link-list :link-list="linkList" />
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
import { ILink } from '~/types/props'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranLinkList,
    GranSignUpForm
  },
  data: () => ({
    linkList: [
      {
        href: '/signin',
        title: 'ログインへ'
      }
    ] as Array<ILink>
  }),
  methods: {
    signUp(signUpForm: ISignUpForm) {
      this.$emit('signUp', signUpForm)
    }
  }
})
</script>
