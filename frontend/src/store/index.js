import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    disVmc: {},
    disArea: {},
    from: {
      // id: 80000,
      // type: 'cpu',
      // parent_id: 120,
      // name: 'cpu_123',
      // time: '',
    },
    to: {
      // id: 180000,
      // type: 'vmc',
      // parent_id: null,
      // name: '测试',
    },
  },
  mutations: {
    setDisVmc: (state, data) => {
      state.disVmc = data;
    },
    setDisArea: (state, data) => {
      state.disArea = data;
    },
    setFrom: (state, data) => {
      state.from = data;
    },
    setTo: (state, data) => {
      state.to = data;
    }
  },
  actions: {
  },
  modules: {
  }
})
