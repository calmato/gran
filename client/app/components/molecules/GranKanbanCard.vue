<template>
  <gran-card color="grey lighten-2">
    <gran-card-title :style="{ '--color': color }">
      {{ name }}
      <span class="light-blue--text text--darken-1 mx-2">{{ length }}</span>
    </gran-card-title>
    <slot />
    <v-card-actions>
      <gran-button color="grey darken-1" text block @click="add">
        <gran-icon name="plus" />
        Add Task
      </gran-button>
    </v-card-actions>
  </gran-card>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardTitle from '~/components/atoms/GranCardTitle.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranCardTitle,
    GranIcon,
    GranButton
  },
  props: {
    name: {
      type: String,
      default: 'ToDo'
    },
    color: {
      type: String,
      default: 'green'
    },
    length: {
      type: Number,
      default: 0
    }
  },
  methods: {
    add(): void {
      this.$emit('add')
    }
  }
})
</script>

<style scoped>
.v-btn {
  text-transform: none;
}

.v-card__title {
  --color: green;

  border-top: 5px solid var(--color);
}
</style>
