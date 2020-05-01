<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-text-field
      v-model="formData"
      :label="label"
      :name="name"
      :type="type"
      :prepend-icon="prependIconName"
      :prepend-inner-icon="prependInnerIconName"
      :append-icon="appendIconName"
      :error-messages="errors"
      :success="valid"
      :autofocus="autofocus"
      :solo="solo"
      :flat="flat"
      :clearable="clearable"
      @keydown="keydown"
    />
  </validation-provider>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    appendIcon: {
      type: String,
      default: '',
    },
    label: {
      type: String,
      default: '',
    },
    name: {
      type: String,
      default: '',
    },
    prependIcon: {
      type: String,
      default: '',
    },
    prependInnerIcon: {
      type: String,
      default: '',
    },
    rules: {
      type: Object,
      default: () => {},
    },
    type: {
      type: String,
      default: 'text',
    },
    value: {
      type: String,
      default: '',
    },
    autofocus: {
      type: Boolean,
      default: false,
    },
    solo: {
      type: Boolean,
      default: false,
    },
    flat: {
      type: Boolean,
      default: false,
    },
    clearable: {
      type: Boolean,
      default: false,
    },
  },

  computed: {
    formData: {
      get(): String {
        return this.value
      },
      set(value: String) {
        this.$emit('input', value)
      },
    },
    prependIconName() {
      if (this.prependIcon === '') {
        return ''
      }

      return 'mdi-' + this.prependIcon
    },
    appendIconName() {
      if (this.appendIcon === '') {
        return ''
      }

      return 'mdi-' + this.appendIcon
    },
    prependInnerIconName() {
      if (this.prependInnerIcon === '') {
        return ''
      }

      return 'mdi-' + this.prependInnerIcon
    },
  },
  methods: {
    keydown(keyEvent: KeyboardEvent) {
      this.$emit('keydown', keyEvent)
    },
  },
})
</script>
