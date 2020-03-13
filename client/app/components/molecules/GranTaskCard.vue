<template>
  <gran-card :width="width" @click="click">
    <div v-if="hasLables" class="text-left px-1 my-1">
      <gran-chip
        v-for="lable in task.labels"
        :key="lable.name"
        :color="lable.color"
        x-small
        class="px-2 ml-1"
      >
        {{ lable.name }}
      </gran-chip>
    </div>
    <gran-card-title class="headline px-4 pt-2">
      {{ task.title }}
    </gran-card-title>
    <v-card-text>
      <div class="text-left my-2"></div>
    </v-card-text>
  </gran-card>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranChip from '~/components/atoms/GranChip.vue'
import GranCardTitle from '~/components/atoms/GranCardTitle.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranChip,
    GranCardTitle
  },
  props: {
    task: {
      type: Object,
      default: () => {
        return {
          id: 1,
          title: 'タスク',
          labels: [
            { name: 'client', color: 'primary' },
            { name: 'api', color: 'yellow' }
          ]
        }
      }
    }
  },
  data: () => {
    return {
      width: 300
    }
  },
  computed: {
    hasLables(): boolean {
      return this.task.labels.length !== 0
    }
  },
  methods: {
    click(): void {
      this.$emit('click')
    }
  }
})
</script>
