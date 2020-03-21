<template>
  <gran-card :width="width" color="grey lighten-2">
    <gran-card-title :style="{ '--color': color }">
      {{ name }}
      <span class="light-blue--text text--darken-1 mx-2">{{ length }}</span>
    </gran-card-title>
    <slot />
    <v-card-actions v-if="!isOpen">
      <gran-button color="grey darken-1" text block @click="open">
        <gran-icon name="plus" />
        Add Task
      </gran-button>
    </v-card-actions>
    <gran-card-text v-if="isOpen">
      <form @submit.prevent>
        <gran-text-field
          v-model="value"
          :label="newTaskForm.name.label"
          :rules="newTaskFormValidate.name"
          autofocus
          @keydown="onKeydown"
        />
        <gran-button
          color="light-blue darken-1"
          :dark="!submitDisabled"
          :disabled="submitDisabled"
          @click="doSubmit"
        >
          Add
        </gran-button>
      </form>
    </gran-card-text>
  </gran-card>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardTitle from '~/components/atoms/GranCardTitle.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import GranTextField from '~/components/atoms/GranTextField.vue'

import { TaskForm, TaskFormValidate } from '~/types/form'

export default Vue.extend({
  components: {
    GranCard,
    GranCardTitle,
    GranCardText,
    GranIcon,
    GranButton,
    GranTextField
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
  data: () => {
    return {
      width: 310,
      isOpen: false,
      newTaskForm: TaskForm,
      value: '',
      newTaskFormValidate: TaskFormValidate
    }
  },
  computed: {
    submitDisabled(): Boolean {
      return this.value.length === 0
    }
  },
  methods: {
    open(): void {
      this.isOpen = true
    },
    close(): void {
      this.isOpen = false
      this.isOpen = false
    },
    doSubmit(): void {
      if (!this.submitDisabled) {
        this.isOpen = false
        this.$emit('addTask', this.value)
        this.value = ''
      }
    },
    onKeydown(keyEvent: KeyboardEvent): void {
      if (keyEvent.keyCode === 13) this.doSubmit() // KeyCode: 13 => enter
      if (keyEvent.keyCode === 27) this.close() // KeyCode: 27 => esc
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
