<template>
  <div class="ma-1">
    <gran-card :width="width" color="grey lighten-3">
      <gran-card-text v-if="isOpen">
        <form>
          <gran-text-field label="Name" />
          <gran-color-picker label="Color" />
          <gran-button color="light-blue darken-1" dark @click="doSubmit">Add</gran-button>
          <gran-button color="grey darken-1" icon @click="close">
            <gran-icon name="close" />
          </gran-button>
        </form>
      </gran-card-text>
      <v-card-actions v-if="!isOpen">
        <gran-button color="grey darken-1" text block @click="addColumn">
          <gran-icon name="plus" />
          Add Column
        </gran-button>
      </v-card-actions>
    </gran-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranColorPicker from '~/components/atoms/GranColorPicker.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranTextField,
    GranColorPicker,
    GranIcon,
    GranButton
  },
  data: () => {
    return {
      width: 310,
      isOpen: false
    }
  },
  methods: {
    addColumn(): void {
      this.isOpen = true
      this.$emit('addColumn')
    },
    doSubmit(): void {
      this.isOpen = false
      this.$emit('doSubmit')
    },
    close(): void {
      this.isOpen = false
    }
  }
})
</script>

<style scoped>
.v-btn {
  text-transform: none;
}
</style>
