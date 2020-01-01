import Vue from 'vue'
import { ValidationProvider, ValidationObserver, localize, extend } from 'vee-validate'
import ja from 'vee-validate/dist/locale/ja.json'
import { required, confirmed, max, min, email } from 'vee-validate/dist/rules'

Vue.component('ValidationProvider', ValidationProvider)
Vue.component('ValidationObserver', ValidationObserver)

extend('required', { ...required })
extend('confirmed', { ...confirmed })
extend('max', { ...max })
extend('min', { ...min })
extend('email', { ...email })

localize('ja', ja)
