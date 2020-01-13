<template>
  <div class="ma-2 text-center">
    <p v-for="(link, index) in linkList" :key="index">
      <nuxt-link :to="link.href">{{ link.title }}</nuxt-link>
    </p>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    linkList: {
      type: Array,
      default: () => {
        return [
          {
            href: '/signup',
            title: '新規登録する'
          },
          {
            href: '/forget',
            title: 'パスワードを忘れた'
          }
        ]
      }
    }
  }
})
</script>
