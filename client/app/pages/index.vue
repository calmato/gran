<template>
  <v-layout column justify-center align-center>
    <hr />
    <p>タスクカードコンポーネント</p>
    <v-window :show-arrows="showArrows" reverse>
      <v-window-item v-for="(list, i) in lists" :key="i">
        <draggable :list="list.array" class="dragArea list-group">
          <transition-group>
            <div v-for="(element, index) in list.array" :key="index" class="pa-1 mb-4">
              <v-card>
                <div class="text-left pa-1 my-1">
                  <v-chip class="px-2" x-small>Default</v-chip>
                  <v-chip class="px-2" color="primary" x-small>Default</v-chip>
                  <v-chip class="px-2" color="secondary" x-small>Default</v-chip>
                </div>
                <v-card-title class="headline px-4 pt-0">
                  {{ element }}
                </v-card-title>
                <v-card-text>
                  <div class="text-left my-2"></div>
                </v-card-text>
              </v-card>
            </div>
          </transition-group>
        </draggable>
      </v-window-item>
    </v-window>
  </v-layout>
</template>

<script lang="ts">
import draggable from 'vuedraggable'
import Vue from 'vue'

export default Vue.extend({
  layout: 'home',
  components: {
    draggable
  },
  data: () => ({
    lists: [
      {
        array: ['UI設計考える', 'タスク', 'タスク2']
      },
      {
        array: ['UI設計考える', 'タスク3', 'タスク4']
      }
    ],
    showArrows: true
  })
})
</script>
