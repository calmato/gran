<template>
  <gran-flash v-if="isError" type="error" dense prominent>
    <v-row align="center">
      <v-col class="grow">{{ message }}</v-col>
      <v-col class="shrink">
        <gran-button icon color="" @click="close">
          <gran-icon name="close-circle-outline" />
        </gran-button>
      </v-col>
    </v-row>
  </gran-flash>
</template>

<script lang="ts">
import Vue from 'vue'
import GranFlash from '~/components/atoms/GranFlash.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranFlash,
    GranButton,
    GranIcon
  },
  props: {
    isError: {
      type: Boolean,
      default: false
    },
    message: {
      type: String,
      default: '認証エラー'
    }
  },
  methods: {
    close() {
      this.$emit('close')
    }
  }
})
</script>
