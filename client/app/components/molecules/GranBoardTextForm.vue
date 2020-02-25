<template>
  <div>
    <gran-form-group>
      <gran-text-field label="ボード名" />
      <gran-select label="グループ" />
      <gran-select label="公開範囲" />
      <gran-select label="背景" />
      <v-color-picker v-model="color" />
    </gran-form-group>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranSelect from '~/components/atoms/GranSelect.vue'
import GranFormGroup from '~/components/atoms/GranFormGroup.vue'

export default Vue.extend({
  components: {
    GranFormGroup,
    GranSelect,
    GranTextField
  },
  props: {},
  data: () => {
    return {
      color: ''
    }
  }
})
</script>
