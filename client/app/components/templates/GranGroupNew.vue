<template>
  <div>
    <gran-group-tab />
    <gran-group-input-name />
    <gran-group-field-name />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupTab from '~/components/templates/GranGroupTab.vue'
import GranGroupInputName from '~/components/templates/GranGroupInputName.vue'
import GranGroupFieldName from '~/components/templates/GranGroupFieldName.vue'

export default Vue.extend({
  components: {
    GranGroupTab,
    GranGroupInputName,
    GranGroupFieldName
  }
})
</script>
