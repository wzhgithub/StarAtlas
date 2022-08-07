<template>
  <div id="app">
    <div class="nav_box">
      <ul class="nav">
        <li class="home_title">高性能星算平台演示系统</li>
        <li
          :class="activenow === 'home' ? 'drop-down active' : 'drop-down'"
          @click="changeRoute('home')"
        >
          <router-link class="nav_btn" to="/home"> 首页 </router-link>
        </li>
        <li
          :class="activenow === 'topo' ? 'drop-down active' : 'drop-down'"
          @click="changeRoute('topo')"
        >
          <router-link class="nav_btn" to="/topo"> 拓扑页 </router-link>
        </li>

        <li :class="activenow === 'details' ? 'drop-down active' : 'drop-down'">
          <a class="nav_btn">详情页</a>
          <ul class="drop-down-content">
            <li @click="changeRoute('details')">
              <router-link class="nav_btn_" to="/details"> vmc1 </router-link>
            </li>
            <li @click="changeRoute('details')">
              <router-link class="nav_btn_" to="/details"> vmc2 </router-link>
            </li>
            <!-- <li>
              <router-link class="nav_btn_" to="/about">vmc3</router-link>
            </li>
            <li>
              <router-link class="nav_btn_" to="/about">vmc4</router-link>
            </li> -->
          </ul>
        </li>
        <li
          :class="
            activenow === 'disasterrecovery' ? 'drop-down active' : 'drop-down'
          "
          @click="changeRoute('disasterrecovery')"
        >
          <router-link class="nav_btn" to="/disasterrecovery">
            容灾演示
          </router-link>
        </li>
      </ul>
      <!-- 
      <router-link class="nav_btn" to="/about">About</router-link>
      <ul>
        <li class="drop-down">
          <span class="nav_btn">详情页</span>
          <ul class="nav_btn drop-down-content">
            <li>
              <router-link class="nav_btn" to="/about">About</router-link>
            </li>
            <li>
              <router-link class="nav_btn" to="/about">About</router-link>
            </li>
            <li>
              <router-link class="nav_btn" to="/about">About</router-link>
            </li>
            <li>
              <router-link class="nav_btn" to="/about">About</router-link>
            </li>
          </ul>
        </li>
      </ul> -->
    </div>
    <router-view class="view_contenter" />
  </div>
</template>
<script>
import "@/lib/qunee-min.js";
export default {
  data() {
    return {
      activenow: "home",
    };
  },
  methods: {
    changeRoute(key) {
      this.activenow = key;
      console.log(this.activenow);
    },
  },
  mounted() {
    if (this.$route.path === "/" || this.$route.path === "/home") {
      this.activenow = "home";
    }
    if (this.$route.path === "/topo") {
      this.activenow = "topo";
    }
    if (this.$route.path === "/disasterrecovery") {
      this.activenow = "disasterrecovery";
    }
    if (this.$route.path === "/details") {
      this.activenow = "details";
    }
  },
};
</script>
<style lang="less">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  // text-align: center;
  color: #2c3e50;
}
.nav_box {
  width: 100%;
  height: 8.5vh;
  // background-color: #2c3e50;
  background: url("./assets/png/header.png") no-repeat;
  background-size: cover;
  vertical-align: middle;
  font-weight: bold;
  ul {
    height: 100%;
    width: 100%;
  }
}
.view_contenter {
  width: 100%;
  height: 91.5vh;
  padding: 2rem;
  box-sizing: border-box;
  // background-color: #42b983;
}
.nav_btn {
  // float: right;
  border: none;
  text-align: center;
  display: inline-block;
  width: 100%;
  height: 100%;
  background: url("./assets/png/normalbtn.png") no-repeat center;
  color: aliceblue;
  text-decoration: none;
  margin-top: 1rem;
  list-style: none;
}

.nav_btn_ {
  // float: right;
  border: none;
  text-align: center;
  display: inline-block;
  width: 100%;
  height: 100%;
  background: url("./assets/png/normalbtn.png") no-repeat center;
  color: aliceblue;
  text-decoration: none;
  margin-top: 1rem;
  list-style: none;
}

ul {
  list-style: none !important;
  text-align: center;
}
.nav {
  width: 100%;
  margin: 0 auto;
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  justify-content: flex-end;
  // float: right;
  position: relative;
  .home_title {
    position: absolute;
    top: 0.5rem;
    left: 3.5rem;
    color: #fff;
    font-size: 2rem;
    font-weight: 700;
  }
}

.drop-down {
  width: 12%;
  height: 5vh;
  z-index: 999;
}

.drop-down-content {
  opacity: 0;
  height: 0;
  overflow: hidden;
  transition: all 1s ease;
}

.nav .drop-down:hover .drop-down-content {
  opacity: 1;
  height: 100%;
  z-index: 999;
  padding: 0;
  // background-color: #fff;
}
.nav_btn:hover {
  background: url("./assets/png/activebtn.png") no-repeat center;
}
.nav_btn_:hover {
  background: url("./assets/png/activebtn.png") no-repeat center;
}
.active {
  .nav_btn {
    background: url("./assets/png/activebtn.png") no-repeat center !important;
  }
}
</style>
