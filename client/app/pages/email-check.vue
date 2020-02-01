<template>
  <gran-email-check :email="email" @click="click" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranEmailCheck from '~/components/templates/GranEmailCheck.vue'

export default Vue.extend({
  layout: 'auth',
  components: {
    GranEmailCheck
  },
  asyncData({ route }) {
    return { email: route.params.email }
  },
  validate({ params }) {
    return params.email !== undefined
  },
  methods: {
    click() {
      this.$router.push('/signin')
    }
  }
})
</script>
