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
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({})
</script>
