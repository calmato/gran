<template>
  <v-app id="main">
    <v-content>
      <v-container class="fill-height" fluid>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <h1 class="font-weight-black white--text text-center">Gran</h1>
            <nuxt />
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<style scoped>
#main {
  background: #039be5;
}
</style>
