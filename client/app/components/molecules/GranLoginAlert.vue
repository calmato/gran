<template>
  <gran-alert v-if="isError" type="error" dense prominent>
    <v-row align="center">
      <v-col class="grow">認証エラー</v-col>
      <v-col class="shrink">
        <gran-button icon color="" @click="close">
          <gran-icon name="close-circle-outline" />
        </gran-button>
      </v-col>
    </v-row>
  </gran-alert>
</template>

<script lang="ts">
import Vue from 'vue'
import GranAlert from '~/components/atoms/GranAlert.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranAlert,
    GranButton,
    GranIcon
  },
  props: {
    isError: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    close() {
      this.$emit('close')
    }
  }
})
</script>
