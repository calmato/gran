<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app right>
      <v-list dense>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-logout</v-icon>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title>logout</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-contact-mail</v-icon>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title>Contact</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app color="#039BE5" dark>
      <v-spacer />
      <div class="font-weight-black white--text title">
        <v-toolbar-title>Gran</v-toolbar-title>
      </div>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <template v-if="isForm" v-slot:extension>
        <gran-group-tab />
      </template>
    </v-app-bar>

    <v-content>
      <v-container color="grey lighten-5">
        <nuxt />
        <v-footer color="white" padless fixed>
          <v-row justify="center" no-gutters>
            <v-btn v-for="link in links" :key="link" color="light-blue darken-1" text class="my-2">
              {{ link }}
            </v-btn>
          </v-row>
        </v-footer>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import GranGroupTab from '~/components/atoms/GranGroupTab.vue'

export default {
  components: {
    GranGroupTab
  },
  data: () => ({
    drawer: null,
    links: ['Home', 'Board', 'Calender', 'Alert', 'Account']
  }),
  computed: {
    isForm() {
      return this.$route.name === 'boards-new'
    }
  }
}
</script>

<style scoped>
#inspire {
  background-color: #f5f5f5;
}

.title {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
</style>
