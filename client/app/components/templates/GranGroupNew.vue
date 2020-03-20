<template>
  <div>
    <gran-group-new-form @apply="apply" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupNewForm from '~/components/molecules/GranGroupNewForm.vue'

export default Vue.extend({
  components: {
    GranGroupNewForm
  },

  methods: {
    apply(groupNewForm) {
      this.$emit('apply', groupNewForm)
    }
  }
})
</script>
