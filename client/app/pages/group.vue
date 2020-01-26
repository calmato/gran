<template>
  <gran-group-new @apply="apply" />
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'
import GranGroupNew from '~/components/templates/GranGroupNew.vue'

export default Vue.extend({
  layout: 'header',

  components: {
    GranGroupNew
  },

  methods: {
    apply(groupNewForm) {
      console.log(groupNewForm)
      this.create(groupNewForm)
        .then(() => {
          console.log('ok')
        })
        .catch(() => {
          console.log('ng')
        })
    },
    ...mapActions('group', ['create'])
  }
})
</script>
