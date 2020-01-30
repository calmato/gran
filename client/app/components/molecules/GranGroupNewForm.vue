<template>
  <form>
    <gran-text-field v-model="groupNewForm.name" label="GroupName" />
    <gran-text-area v-model="groupNewForm.description" label="GroupDescription" />
    <gran-button @click="doSubmit">適用</gran-button>
  </form>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranTextArea from '~/components/atoms/GranTextArea.vue'
import GranButton from '~/components/atoms/GranButton.vue'

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
    }
  }),

  methods: {
    doSubmit() {
      this.$emit('apply', this.groupNewForm)
    }
  }
})
</script>
