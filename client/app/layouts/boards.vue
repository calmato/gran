<template>
  <v-app>
    <v-app-bar color="grey darken-1" dark app>
      <v-app-bar-nav-icon @click="back">
        <gran-icon name="chevron-left" />
      </v-app-bar-nav-icon>
      <v-toolbar-title>{{ board.name }}</v-toolbar-title>
    </v-app-bar>

    <v-content :style="{ '--color': board.backgroundColor }">
      <v-container fluid>
        <nuxt />
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapState } from 'vuex'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranIcon
  },
  computed: {
    ...mapState('boards', ['board'])
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
  --color: #fafafa;

  overflow-x: auto;
  background-color: var(--color);
}
</style>
