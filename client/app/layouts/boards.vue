<template>
  <v-app>
    <v-app-bar color="grey darken-1" dark app>
      <v-app-bar-nav-icon @click="back">
        <gran-icon name="chevron-left" />
      </v-app-bar-nav-icon>
      <v-toolbar-title>ボード名</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container color="grey lighten-5" fluid>
        <nuxt />
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranIcon
  },
  methods: {
    back(): void {
      this.$router.go(-1)
    }
  }
})
</script>

<style scoped>
.v-content {
  overflow-x: auto;
}
</style>
