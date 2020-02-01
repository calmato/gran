<template>
  <gran-card>
    <gran-card-text>
      <div class="text-center">
        <gran-icon :x-large="xLarge" class="pa-3" />
        <h1 class="pb-4">Check your email</h1>
        <p class="ma-0 pa-0 subtitle-2 subtitle-2 font-weight-bold">{{ email }}</p>
        <p class="ma-0 pa-0">にメールを送信しました．</p>
        <p class="ma-0 pb-4">
          メール内のリンクをチェックして登録を完了してください．
        </p>
        <gran-button color="light-blue darken-1" block dark @click="click">
          ログイン画面へ
        </gran-button>
      </div>
    </gran-card-text>
  </gran-card>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranButton from '~/components/atoms/GranButton.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranButton,
    GranIcon
  },
  props: {
    email: {
      type: String,
      default: 'hogehoge@email.com'
    }
  },
  data: () => ({
    xLarge: true
  }),
  methods: {
    click() {
      this.$emit('click')
    }
  }
})
</script>
