<template>
  <gran-tab centered fixed-tabs dark>
    <gran-tab-item v-for="(item, index) in groups" :key="index" :name="item" />
  </gran-tab>
</template>

<script lang="ts">
import Vue from 'vue'
import GranTab from '~/components/atoms/GranTab.vue'
import GranTabItem from '~/components/atoms/GranTabItem.vue'

export default Vue.extend({
  components: {
    GranTab,
    GranTabItem
  },
  props: {
    groups: {
      type: Array,
      default: (): Array<string> => {
        return ['グループ', 'ボード']
      }
    }
  }
})
</script>
