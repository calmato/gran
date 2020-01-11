<template>
  <v-app>
    <v-content>
      <v-container>
        <v-layout column align-center>
          <h1>Internal Server Error</h1>
          <v-img src="/error-500.png" />

          <v-btn @click="handleClick">ホームへ戻る</v-btn>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  layout: 'empty',

  methods: {
    handleClick() {
      this.$router.push('/')
    }
  }
})
</script>
