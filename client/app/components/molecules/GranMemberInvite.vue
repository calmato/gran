<template>
  <div>
    <v-row>
      <gran-icon :x-large="xLarge" name="account-multiple-plus" />
      <p class="ma-3 pt-1">メンバーを招待</p>
    </v-row>
    <v-layout>
      <v-flex xs12 sm8 md6>
        <v-card>
          <v-card-text>
            <div v-for="(email, index) in emails" :key="index">
              <gran-text-field v-model="email.value" :label="`email-${index + 1}`" />
              <v-btn class="mx-2" fab dark small color="pink" @click="addForm">
                <gran-icon name="plus" />
              </v-btn>
              <v-btn
                v-if="index"
                class="mx-2"
                fab
                dark
                small
                color="cyan"
                @click="deleteForm(index)"
              >
                <gran-icon name="minus" />
              </v-btn>
            </div>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import Vue from 'vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranTextField from '~/components/atoms/GranTextField.vue'

export default Vue.extend({
  components: {
    GranIcon,
    GranTextField
  },
  data: () => ({
    xLarge: true,
    emails: [{ value: '' }]
  }),
  methods: {
    addForm() {
      this.emails.push({ value: '' })
    },
    deleteForm(id) {
      this.emails.splice(id, 1)
    }
  }
})
</script>
