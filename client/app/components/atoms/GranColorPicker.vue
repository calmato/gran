<template>
  <v-text-field v-model="value" :label="label" readonly hide-details>
    <template v-slot:append-outer>
      <v-menu v-model="menu" top :close-on-content-click="false">
        <template v-slot:activator="{ on }">
          <v-btn text icon v-on="on">
            <v-icon :color="value">mdi-checkbox-blank</v-icon>
          </v-btn>
        </template>
        <v-card>
          <v-card-text>
            <v-color-picker
              :value="value"
              hide-canvas
              hide-inputs
              show-swatches
              swatches-max-height="200px"
              flat
              @input="emitter"
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
    value: {
      type: String,
      default: '',
    },
    label: {
      type: String,
      default: '背景',
    },
  },
  data: () => {
    return {
      menu: false,
    }
  },
  methods: {
    emitter(value) {
      this.$emit('input', value)
    },
  },
})
</script>
