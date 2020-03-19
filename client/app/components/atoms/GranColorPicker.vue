<template>
  <v-text-field v-model="color" :label="label" readonly hide-details>
    <template v-slot:append-outer>
      <v-menu v-model="menu" top :close-on-content-click="false">
        <template v-slot:activator="{ on }">
          <v-btn text icon v-on="on">
            <v-icon :color="color">mdi-checkbox-blank</v-icon>
          </v-btn>
        </template>
        <v-card>
          <v-card-text>
            <v-color-picker
              v-model="color"
              hide-canvas
              hide-inputs
              show-swatches
              swatches-max-height="200px"
              flat
            />
          </v-card-text>
        </v-card>
      </v-menu>
    </template>
  </v-text-field>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  props: {
    label: {
      type: String,
      default: '背景'
    }
  },
  data: () => {
    return {
      color: '',
      menu: false
    }
  }
})
</script>
