import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import jquery from "jquery";
import vuescroll from 'vuescroll'
import 'vuescroll/dist/vuescroll.css'
import 'element-ui/lib/theme-chalk/index.css';
import './assets/icons/iconfont.css';
import './assets/icons/iconfont.js'

Vue.config.productionTip = false
Vue.prototype.$ = jquery;
Vue.use(ElementUI);
Vue.use(vuescroll);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
