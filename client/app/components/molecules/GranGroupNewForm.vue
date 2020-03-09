<template>
  <form>
    <gran-text-field
      v-model="groupNewForm.name.value"
      :label="groupNewForm.name.label"
      :rules="groupNewFormValidate.name"
    />
    <gran-text-area
      v-model="groupNewForm.description.value"
      :label="groupNewForm.description.label"
      :rules="groupNewFormValidate.description"
    />
    <gran-member-invite />
    <gran-button class="ma-3 pt-1" @click="doSubmit">適用</gran-button>
  </form>
</template>

<script lang="ts">
import Vue from 'vue'
import { GroupNewForm, IGroupForm, GroupNewFormValidate } from '~/types/form'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranTextArea from '~/components/atoms/GranTextArea.vue'
import GranMemberInvite from '~/components/molecules/GranMemberInvite.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranTextField,
    GranTextArea,
    GranMemberInvite,
    GranButton
  },

  data: () => ({
    groupNewForm: GroupNewForm as IGroupForm,
    groupNewFormValidate: GroupNewFormValidate
  }),
  methods: {
    doSubmit() {
      this.$emit('apply', this.groupNewForm)
    }
  }
})
</script>
