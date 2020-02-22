<template>
  <div>
    <gran-form-group>
      <gran-text-field />
      <gran-text-field />
      <gran-text-field />
      <gran-text-field />
    </gran-form-group>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranFormGroup from '~/components/atoms/GranFormGroup.vue'

export default Vue.extend({
  components: {
    GranFormGroup,
    GranTextField
  }
})
</script>
