<template>
  <gran-group-tab class="gran-group-tab" />
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupTab from '~/components/organisms/GranGroupTab.vue'

export default Vue.extend({
  components: {
    GranGroupTab
  }
})
</script>

<style scoped>
.gran-group-tab {
  margin-bottom: 32px;
}
</style>
