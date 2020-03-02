<template>
  <div>
    <v-row>
      <gran-icon :x-large="xLarge" name="account" />
      <p class="ma-3 pt-1">参加中のボード一覧</p>
    </v-row>
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
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranIcon
  },
  computed: {
    ...mapGetters('group', ['groups'])
  },
  data: () => ({
    xLarge: true
  })
})
</script>
