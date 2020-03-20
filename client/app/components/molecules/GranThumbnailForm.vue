<template>
  <gran-card color="#f5f5f5" flat>
    <gran-card-text class="ma-0 pa-0 text-center">
      <gran-avatar @click="click" />
    </gran-card-text>
    <gran-card-text class="ma-0 pa-0 text-center">
      <gran-button color="grey-darken-2" text @click="click">サムネイルを編集</gran-button>
    </gran-card-text>
  </gran-card>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranAvatar from '~/components/atoms/GranAvatar.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranButton from '~/components/atoms/GranButton.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranAvatar,
    GranButton
  },
  methods: {
    click() {
      this.$emit('click')
    }
  }
})
</script>
