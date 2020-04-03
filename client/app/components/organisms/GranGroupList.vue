<template>
  <div>
    <v-row>
      <gran-icon :x-large="xLarge" name="account" />
      <p class="ma-3 pt-1">参加中のグループ一覧</p>
    </v-row>
    <v-list three-line>
      <template v-for="(group, index) in groups">
        <v-list-item :key="index" @click="on(group.name, group.description)">
          <v-list-item-avatar>
            <v-img :src="group.avatar"></v-img>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>{{ group.name }}</v-list-item-title>
            <v-list-item-subtitle>{{ group.description }}</v-list-item-subtitle>
            <v-divider :key="index"> </v-divider>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
    <v-dialog v-model="dialog" max-width="600">
      <v-card>
        <v-card-title>
          <span class="headline">グループ情報の編集</span>
        </v-card-title>
        <v-card-title>タイトル</v-card-title>
        <v-row justify="center">
          <v-flex xs12 sm8 md6>
            <gran-text-field v-model="groupName"></gran-text-field>
          </v-flex>
        </v-row>
        <v-card-title>説明</v-card-title>
        <v-row justify="center">
          <v-flex xs12 sm8 md6>
            <gran-text-field v-model="groupDescription"></gran-text-field>
          </v-flex>
        </v-row>
        <v-card-title>メンバー</v-card-title>
        <v-row>
          <gran-icon :x-large="xLarge" class="pl-4 pt-1" name="account-multiple-plus" />
          <v-card-title class="pt-7">メンバーを招待</v-card-title>
        </v-row>
        <v-row justify="center">
          <v-flex xs12 sm8 md6>
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
          </v-flex>
        </v-row>
        <v-row justify="center">
          <v-btn class="ma-3" color="primary" @click="dialog = false">
            適用
          </v-btn>
        </v-row>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranTextField from '~/components/atoms/GranTextField.vue'

export default Vue.extend({
  components: {
    GranIcon,
    GranTextField,
  },
  data: () => ({
    xLarge: true,
    dialog: false,
    emails: [{ value: '' }],
    groupName: '',
    groupDescription: '',
  }),
  computed: {
    ...mapGetters('group', ['groups']),
  },
  methods: {
    on(groupName, groupDescription) {
      this.dialog = true
      this.groupName = groupName
      this.groupDescription = groupDescription
    },
    addForm() {
      this.emails.push({ value: '' })
    },
    deleteForm(id) {
      this.emails.splice(id, 1)
    },
  },
})
</script>
