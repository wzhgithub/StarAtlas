<template>
  <div id="app">
    <div class="nav_box">
      <ul class="nav">
        <!-- <li class="drop-down">
          <router-link class="nav_btn" to="/">Home</router-link>
        </li> -->
        <li
          :class="activenow === 'topo' ? 'drop-down active' : 'drop-down'"
          @click="changeRoute('topo')"
        >
          <router-link class="nav_btn" to="/"> 拓扑页 </router-link>
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
export default {
  data() {
    return {
      activenow: "topo",
    };
  },
  methods: {
    changeRoute(key) {
      this.activenow = key;
      console.log(this.activenow);
    },
  },
  mounted() {
    if (this.$route.path === "/") {
      this.activenow = "topo";
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
}
.view_contenter {
  width: 100%;
  height: 91.5vh;
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
  margin-top: 14px;
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
  margin-top: 14px;
  list-style: none;
}

ul {
  list-style: none !important;
  text-align: center;
}
.nav {
  width: 100%;
  margin: 0px auto;
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  justify-content: flex-end;
  // float: right;
}

.drop-down {
  width: 12%;
  height: 5vh;
  // margin-left: -30px;
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
  height: 15vh;
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
