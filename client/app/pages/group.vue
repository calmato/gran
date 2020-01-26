<template>
  <gran-group-new @apply="apply" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupNew from '~/components/templates/GranGroupNew.vue'

export default Vue.extend({
  layout: 'header',

  components: {
    GranGroupNew
  },

  methods: {
    apply(groupNewForm) {
      console.log(groupNewForm)
    }
  }
})
</script>
