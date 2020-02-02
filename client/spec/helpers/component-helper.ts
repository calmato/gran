import Vue from 'vue'
import Vuex from 'vuex'
import Vuetify from 'vuetify'
import { ValidationProvider } from 'vee-validate'

Vue.use(Vuex)
Vue.use(Vuetify)

Vue.component('ValidationProvider', ValidationProvider)
