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

    <v-app-bar app color="#039BE5">
      <v-spacer />
      <div class="title">
        <v-toolbar-title>Gran</v-toolbar-title>
      </div>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    drawer: null
  })
}
</script>

<style>
body {
  background-color: #f7f7f7;
}

.title {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
</style>
