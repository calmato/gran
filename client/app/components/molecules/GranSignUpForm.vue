<template>
  <gran-form-group>
    <gran-text-field
      v-model="signUpForm.email.value"
      prepend-icon="email"
      label="Email"
      :rules="signUpFormValidate.email"
    />
    <gran-text-field
      v-model="signUpForm.password.value"
      prepend-icon="lock"
      :label="signUpForm.password.label"
      type="password"
      :rules="signUpFormValidate.password"
    />
    <gran-text-field
      v-model="signUpForm.passwordConfirmation.value"
      prepend-icon="lock"
      :label="signUpForm.passwordConfirmation.label"
      type="password"
      :rules="signUpFormValidate.passwordConfirmation"
    />
    <gran-button color="light-blue darken-1" block dark @click="click">登録する</gran-button>
  </gran-form-group>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import GranFormGroup from '~/components/atoms/GranFormGroup.vue'
import { SignUpForm, SignUpFormValidate } from '~/types/form'

export default Vue.extend({
  components: {
    GranTextField,
    GranButton,
    GranFormGroup,
  },
  data: () => ({
    signUpForm: SignUpForm,
    signUpFormValidate: SignUpFormValidate,
  }),
  methods: {
    click() {
      this.$emit('signUp', this.signUpForm)
    },
  },
})
</script>
