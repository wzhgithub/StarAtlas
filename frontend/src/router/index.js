import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Topo from '../views/Topo.vue'
import DisasterRecovery from '../views/DisasterRecovery.vue'
import RemoteControl from '../views/RemoteControl.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/topo',
    name: 'Topo',
    component: Topo
  },
  {
    path: '/disasterrecovery',
    name: 'DisasterRecovery',
    component: DisasterRecovery
  },
  {
    path: '/remoteControl',
    name: 'RemoteControl',
    component: RemoteControl
  },
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // },
  {
    path: '/details/:id',
    name: 'Details',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Details.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
