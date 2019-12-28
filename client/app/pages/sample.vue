<template>
  <v-layout>
    <v-flex class="text-center">
      <v-data-table
        :headers="headers"
        :items="users"
        :items-per-page="5"
        class="elevation-1"
      ></v-data-table>
    </v-flex>

    <v-flex class="text-center">
      <v-card class="mx-auto">
        <v-card-text>
          <v-text-field v-model="sampleUserNewForm.name" label="Name" />
        </v-card-text>

        <v-card-actions>
          <v-btn @click="handleClick">追加</v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters, mapActions } from 'vuex'
import { ISampleNewForm } from '~/types/form'

export default Vue.extend({
  data: () => ({
    headers: [{ text: 'Name', value: 'name' }],
    sampleUserNewForm: {
      name: ''
    } as ISampleNewForm
  }),

  computed: {
    ...mapGetters('sample', ['users'])
  },

  methods: {
    handleClick() {
      this.addUser(this.sampleUserNewForm)
    },
    ...mapActions('sample', ['addUser'])
  }
})
</script>
