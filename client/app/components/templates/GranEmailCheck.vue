<template>
  <gran-email-check-card :email="email" @click="click" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranEmailCheckCard from '~/components/organisms/GranEmailCheckCard.vue'

export default Vue.extend({
  components: {
    GranEmailCheckCard
  },
  props: {
    email: {
      type: String,
      default: 'hogehoge@email.com'
    }
  },
  methods: {
    click() {
      this.$emit('click')
    }
  }
})
</script>
