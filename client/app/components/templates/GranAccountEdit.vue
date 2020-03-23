<template>
  <div>
    <div class="centering_item">
      <gran-avatar />
      <p>
        Change Profile Photo
      </p>
    </div>
    <v-row justify="center">
      <v-flex xs12 sm8 md6>
        <gran-text-field
          v-model="accountForm.name.value"
          :label="accountForm.name.label"
          :rules="accountFormValidate.name"
        />
        <gran-text-field
          v-model="accountForm.userName.value"
          :label="accountForm.userName.label"
          :rules="accountFormValidate.userName"
        />
        <gran-text-area v-model="accountForm.profile.value" :label="accountForm.profile.label" />
      </v-flex>
    </v-row>
    <h2>Private Information</h2>
    <v-row justify="center">
      <v-flex xs12 sm8 md6>
        <gran-text-field v-model="accountForm.email.value" :label="accountForm.email.label" />
        <gran-text-field v-model="accountForm.phone.value" :label="accountForm.phone.label" />
      </v-flex>
    </v-row>
    <v-row justify="center">
      <gran-button color="" class="ma-5">CANCEL</gran-button>
      <gran-button class="ma-5" @click="done">DONE</gran-button>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { AccountForm, IAccountForm, AccountFormValidate } from '~/types/form'
import GranAvatar from '~/components/atoms/GranAvatar.vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranTextArea from '~/components/atoms/GranTextArea.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranAvatar,
    GranTextField,
    GranTextArea,
    GranButton,
  },

  data: () => ({
    accountForm: AccountForm as IAccountForm,
    accountFormValidate: AccountFormValidate,
  }),
  methods: {
    done() {
      this.$emit('apply', this.accountForm)
    },
  },
})
</script>

<style scoped>
.centering_item {
  text-align: center;
  font-weight: bold;
}

p {
  color: #03a9f4;
}
</style>
