<template>
  <v-bottom-navigation color="light-blue darken-1" app>
    <gran-button
      v-for="(item, index) in links"
      :key="index"
      color="white"
      @click="click(item.link)"
    >
      <span>{{ item.name }}</span>
      <gran-icon :name="item.icon" />
    </gran-button>
  </v-bottom-navigation>
</template>

<script lang="ts">
import Vue from 'vue'
import GranButton from '~/components/atoms/GranButton.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranButton,
    GranIcon
  },
  data() {
    return {
      links: [
        {
          name: 'Home',
          link: '/',
          icon: 'home'
        },
        {
          name: 'Board',
          link: '/boards',
          icon: 'view-dashboard'
        },
        {
          name: 'Calendar',
          link: '/',
          icon: 'calendar'
        },
        {
          name: 'Notification',
          link: '/',
          icon: 'information-outline'
        },
        {
          name: 'Account',
          link: '/',
          icon: 'account'
        }
      ]
    }
  },
  methods: {
    click(link: string): void {
      this.$router.push(link)
    }
  }
})
</script>
