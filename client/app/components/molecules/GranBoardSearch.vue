<template>
  <div>
    <v-card class="mx-auto" max-width="max">
      <v-sheet class="pa-4 lighten-2">
        <v-text-field
          v-model="search"
          label="Search for a board"
          flat
          background-color="white"
          color="black"
          clearable
          clear-icon="mdi-close-circle-outline"
        ></v-text-field>
      </v-sheet>
    </v-card>
    <v-row>
      <gran-icon :x-large="xLarge" name="account" />
      <p class="ma-3 pt-1">パーソナルボード</p>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranIcon
  },
  data: () => ({
    xLarge: true
  })
})
</script>
