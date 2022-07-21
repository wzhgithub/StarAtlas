/* eslint-disable */
<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <!-- 123 -->

    <div v-if="loading" class="loadingbox">
      <div data-loader="jumping"></div>
    </div>
    <div v-else class="home">
      <!-- <h3>首页</h3> -->
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import "@/lib/loaders.min.css";

export default {
  name: "DisasterRecovery",
  components: {
    // HelloWorld,
  },
  data() {
    return {
      loading: true,
    };
  },
  methods: {},
  mounted() {
    setTimeout(() => {
      this.loading = false;
    }, 3000);
  },

  created() {},
};
</script>
<style lang="less" scoped>
</style>