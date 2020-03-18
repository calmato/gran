<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-text-field
      v-model="formData"
      :label="label"
      :name="name"
      :type="type"
      :prepend-icon="prependIconName"
      :append-icon="appendIconName"
      :error-messages="errors"
      :success="valid"
      :counter="count"
      :maxlength="maxlength"
    />
  </validation-provider>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    appendIcon: {
      type: String,
      default: ''
    },
    maxlength: {
      type: String,
      default: ''
    },
    count: {
      type: String,
      default: ''
    },
    label: {
      type: String,
      default: ''
    },
    name: {
      type: String,
      default: ''
    },
    prependIcon: {
      type: String,
      default: ''
    },
    rules: {
      type: Object,
      default: () => {}
    },
    type: {
      type: String,
      default: 'text'
    },
    value: {
      type: String,
      default: ''
    }
  },

  computed: {
    formData: {
      get(): String {
        return this.value
      },
      set(value: String) {
        this.$emit('input', value)
      }
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
    }
  }
})
</script>
