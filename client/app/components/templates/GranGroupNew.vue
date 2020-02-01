<template>
  <div>
    <gran-group-tab class="gran-group-tab" />
    <gran-group-new-form @apply="apply" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranGroupTab from '~/components/organisms/GranGroupTab.vue'
import GranGroupNewForm from '~/components/molecules/GranGroupNewForm.vue'

export default Vue.extend({
  components: {
    GranGroupTab,
    GranGroupNewForm
  },

  methods: {
    apply(groupNewForm) {
      this.$emit('apply', groupNewForm)
    }
  }
})
</script>

<style scoped>
.gran-group-tab {
  margin-bottom: 32px;
}
</style>
