<template>
  <div>
    <gran-group-name />
    <gran-group-detail />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupName from '~/components/templates/GranGroupName.vue'
import GranGroupDetail from '~/components/templates/GranGroupDetail.vue'

export default Vue.extend({
  components: {
    GranGroupName,
    GranGroupDetail
  }
})
</script>
