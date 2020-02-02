import Vue from 'vue'
import Vuex from 'vuex'
import Vuetify from 'vuetify'
import { ValidationProvider } from 'vee-validate'
import AxiosHelper from '~~/spec/helpers/axios'
import StoreModule from '~~/spec/helpers/store'

Vue.use(Vuex)
Vue.use(Vuetify)

Vue.component('ValidationProvider', ValidationProvider)

export const Axios = () => AxiosHelper
export const Store = () => new Vuex.Store(StoreModule)
