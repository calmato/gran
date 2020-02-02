import Vue from 'vue'
import Vuetify from 'vuetify'
import { ValidationProvider } from 'vee-validate'

Vue.use(Vuetify)
Vue.component('ValidationProvider', ValidationProvider)
