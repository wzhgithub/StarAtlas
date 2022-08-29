import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    disVmc: {},
    disArea: {}
  },
  mutations: {
    setDisVmc: (state, data) => {
      state.disVmc = data;
    },
    setDisArea: (state, data) => {
      state.disArea = data;
    }
  },
  actions: {
  },
  modules: {
  }
})
