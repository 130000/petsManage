import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import './less/index.less'
import store from './store/index'
import api from './api/api'

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
store.commit('addMenu', router)
app.config.globalProperties.$api = api
function checkRouter(path) {
  let hasCheck = router.getRoutes().filter(route => route.path == path).length
  if (hasCheck) return true
  else return false
}
router.beforeEach((to, _, next) => {
  NProgress.start()
  store.commit('getToken')
  const token = store.state.token
  if (!token && to.name !== 'login') {
    next({
      name: 'login'
    })
  }
  else if (!checkRouter(to.path)) {
    next({
      name: 'home'
    })
  }
  else {
    next()
  }
})
router.afterEach(() => {
  NProgress.done()
})
app.use(router).use(store)
app.mount('#app')