import Vue from 'vue'
import VueRouter from 'vue-router'
import AboutMe from '@/views/AboutMe.vue'
import ContactMe from '@/views/ContactMe.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    component: AboutMe
  },
  {
    path: '/about',
    component: AboutMe
  },
  {
    path: '/contact',
    component: ContactMe
  }
]

const router = new VueRouter({
  routes,
  mode: 'history'
})

export default router
