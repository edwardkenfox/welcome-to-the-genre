import Vue from 'vue'
import VueRouter from 'vue-router'
import Top from './Top'
import SourceSelect from './SourceSelect'
import Result from './Result'

import './assets/main.css'

Vue.use(VueRouter)
Vue.config.productionTip = false

const routes = [
  { path: '/', component: Top },
  { path: '/source', component: SourceSelect },
  { path: '/results/:genre', component: Result, props: true }
]

const router = new VueRouter({
  routes
})

/* eslint-disable no-new */
new Vue({
  router,
  components: {
    Top,
    SourceSelect,
    Result
  },
  computed: {
    resultStyle () {
      let bgColor
      switch (this.$route.params.genre) {
        case 'blues':
          bgColor = 'rgb(57, 26, 6)'
          break
        case 'classical':
          bgColor = 'rgb(209, 170, 69)'
          break
        case 'country':
          bgColor = 'rgb(57, 128, 35)'
          break
        case 'disco':
          bgColor = 'rgb(172, 38, 124)'
          break
        case 'hiphop':
          bgColor = 'rgb(193, 55, 62)'
          break
        case 'jazz':
          bgColor = 'rgb(44, 16, 128)'
          break
        case 'metal':
          bgColor = 'rgb(72, 72, 72)'
          break
        case 'pop':
          bgColor = 'rgb(93, 206, 127)'
          break
        case 'reggae':
          bgColor = 'rgb(202, 102, 39)'
          break
        case 'rock':
          bgColor = 'rgb(0, 0, 0)'
          break
        default:
          bgColor = 'rgb(0, 0, 0)'
          break
      }

      return {
        backgroundColor: bgColor
      }
    }
  }
}).$mount('#app')
