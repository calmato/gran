<template>
  <v-card max-width="max" class="mx-auto">
    <v-list three-line>
      <template v-for="(group, index) in groups">
        <v-list-item :key="index">
          <v-list-item-avatar>
            <v-img :src="group.avatar"></v-img>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>{{ group.name }}</v-list-item-title>
            <v-list-item-subtitle>{{ group.description }}</v-list-item-subtitle>
            <v-divider :key="index"> </v-divider>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
  </v-card>
</template>

<script lang="ts">
import { mapGetters } from 'vuex'

export default {
  computed: {
    ...mapGetters('group', ['groups'])
  }
}
</script>
