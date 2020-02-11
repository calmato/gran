<template>
  <form>
    <gran-text-field
      v-model="groupNewForm.name"
      label="GroupName"
      :rules="groupNewFormValidate.name"
    />
    <gran-text-area
      v-model="groupNewForm.description"
      label="GroupDescription"
      :rules="groupNewFormValidate.description"
    />
    <gran-button @click="doSubmit">適用</gran-button>
  </form>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranTextArea from '~/components/atoms/GranTextArea.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import { IGroupForm, GroupNewFormValidate } from '~/types'

export default Vue.extend({
  components: {
    GranTextField,
    GranTextArea,
    GranButton
  },

  data: () => ({
    groupNewForm: {
      name: '',
      description: ''
    } as IGroupForm,
    groupNewFormValidate: GroupNewFormValidate
  }),
  methods: {
    doSubmit() {
      this.$emit('apply', this.groupNewForm)
    }
  }
})
</script>
