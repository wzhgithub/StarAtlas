/* eslint-disable */
<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <!-- 123 -->

    <div v-if="loading" class="loadingbox">
      <div data-loader="jumping"></div>
    </div>
    <div v-else class="homeInnerBox">
      <el-row style="width: 100%; height: 100%">
        <el-col :span="6" style="height: 100%">
          <div class="leftasid">
            <div class="boxLeft">
              <p class="title_new_left">
                <span>异构计算</span>
              </p>
              <div class="contentBox">
                <ul>
                  <li>支持不同指令集</li>
                  <li>支持不同体系架构</li>
                  <li>CPU、GPU、DSP、FPGA</li>
                </ul>
              </div>
            </div>
            <div class="boxLeft">
              <p class="title_new_left">
                <span>TTE协议</span>
              </p>
              <div class="contentBox">
                <ul>
                  <li>Time-triggered Ethernet</li>
                  <li>时钟同步、时间触发通信</li>
                  <li>速率受约传输、保证传输</li>
                </ul>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="12" style="height: 100%">
          <div class="centermain">
            <div class="topBoxNow">
              <el-row
                type="flex"
                justify="space-around"
                style="width: 100%; height: 100%"
              >
                <el-col :span="5" style="height: 100%">
                  <div class="top_little_box">
                    <div class="toptitle">VMC数量</div>
                    <div class="mainNub"> 
                      <countTo
                        class="nub"
                        :startVal="0"
                        :endVal="vmc_num"
                        :duration="3000"
                      ></countTo>
                  </div>
                  </div>
                </el-col>
                <el-col :span="5" style="height: 100%">
                  <div class="top_little_box">
                    <div class="toptitle">计算单元</div>
                    <div class="mainNub">
                      <countTo
                        class="nub"
                        :startVal="0"
                        :endVal="cpu_num"
                        :duration="3000"
                      ></countTo>
                    </div>
                  </div>
                </el-col>
                <el-col :span="5" style="height: 100%">
                  <div class="top_little_box">
                    <div class="toptitle">远置单元</div>
                    <div class="mainNub">
                      <countTo
                        class="nub"
                        :startVal="0"
                        :endVal="rtu_num"
                        :duration="3000"
                      ></countTo>
                    </div>
                  </div>
                </el-col>
                <el-col :span="5" style="height: 100%">
                  <div class="top_little_box">
                    <div class="toptitle">交换机</div>
                    <div class="mainNub">
                      <countTo
                        class="nub"
                        :startVal="0"
                        :endVal="switch_num"
                        :duration="3000"
                      ></countTo>
                      </div>
                  </div>
                </el-col>
              </el-row>
            </div>
            <div class="boxMainTitleInfo">
              <h3>星辰算力</h3>
            </div>
            <div class="linktopbox">
              <router-link class="link_btn" to="/topo"> 拓扑结构 </router-link>
            </div>
            <div class="linkleftbox">
              <router-link class="link_btn" to="/details">
                设备详情
              </router-link>
            </div>
            <div class="linkrightbox">
              <router-link class="link_btn" to="/disasterrecovery">
                任务容灾
              </router-link>
            </div>
            <div class="linkrightbox_">
              <router-link class="link_btn" to="/disasterrecovery">
                任务容灾1
              </router-link>
            </div>
            <div class="bigbox">
              <div class="middlebox">
                <div class="middlebox2">
                  <div class="middlebox3">
                    <div class="minbox"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6" style="height: 100%">
          <div class="rightasid">
            <div class="boxRight">
              <p class="title_new_right">
                <span>多级容灾</span>
              </p>
              <div class="contentBox">
                <ul>
                  <li>支持三级故障容灾</li>
                  <li>任务、分区、整机级</li>
                  <li>支持主备、多活计算</li>
                </ul>
              </div>
            </div>
            <div class="boxRight">
              <p class="title_new_right">
                <span>以太网络</span>
              </p>
              <div class="contentBox">
                <ul>
                  <li>提高设备接入效率</li>
                  <li>降低接入设备成本</li>
                  <li>丰富接入设备多样性</li>
                </ul>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import "@/lib/loaders.min.css";
import { getTopoShow } from "@/api";
import countTo from "vue-count-to";

export default {
  name: "Home",
  components: {
    // HelloWorld,
    countTo,
  },
  data() {
    return {
      loading: true,
      vmc_num: 0,
      cpu_num: 0,
      rtu_num: 0,
      switch_num: 0,
    };
  },
  methods: {
    async getNameOAll() {
      const { data } = await getTopoShow();
      if (data.code == 0) {
        this.vmc_num = 0;
        this.cpu_num = 0;
        this.rtu_num = 0;
        this.switch_num = 0;
        this.vmcs = data.data.node.map((item) => {
          // return item.device_type == "vmc" || item.device_type == "cpu";
          if (item.device_type == "vmc") {
            this.vmc_num += 1;
          } else if (item.device_type == "sw") {
            this.switch_num += 1;
          } else if (item.device_type == "rtu") {
            this.rtu_num += 1;
          } else {
            this.cpu_num += item.other_info[0].value.length;
          }
          return item;
        });
        // console.log(this.vmcs);
      }
    },
  },
  mounted() {
    let that = this;
    setTimeout(() => {
      that.loading = false;
      that.getNameOAll();
    }, 3000);
  },

  created() {},
};
</script>
<style lang="less" scoped>
.home {
  background: url("../assets/newpng/home_in_bgi.png") no-repeat center;
  background-size: 100% 100%;
  .homeInnerBox {
    width: 100%;
    height: 100%;
    overflow: hidden;
    .littlebox {
      width: 100%;
      height: 50%;
      // margin-bottom: 1%;
    }
    .centermain {
      height: 100%;
      // background-color: #fff;
      position: relative;
      .topBoxNow {
        width: 100%;
        height: 15%;
        // background-color: #fff;
        position: absolute;
        top: 0;
        .top_little_box {
          width: 100%;
          height: 100%;
          text-align: center;
          // background-color: aqua;
          // background: url("../assets/newpng/center_top_main_little_bgi.png")
          //   no-repeat center;
          background: url("../assets/tt.gif")
            no-repeat center;
          background-size: 100% 100%;
          //color: rgba(126, 225, 50, 0.823);
          color: #94f0e7;
          font-size: 1.5rem;
          .toptitle {
            font-size: 1.5rem;
            margin-bottom: 1rem;
            font-weight: bold;
          }
          .mainNub {
            text-shadow: 0 0 10px #3ee4f0;
            font-weight: 700;
          }
        }
      }
      .boxMainTitleInfo {
        width: 100%;
        position: absolute;
        top: 42%;
        text-align: center;
        // left: 45%;
        color: #fff;

        h3 {
          font-size: 2.5rem;
          font-weight: 700;
        }
      }
      .linktopbox {
        width: 50%;
        height: 15%;
        // background-color: #fff;
        position: absolute;
        top: 20%;
        right: 25%;
        background: url("../assets/newpng/center_mian_big_bgi.png") no-repeat
          center;
        background-size: 100% 100%;
        font-size: 1.25rem;
        font-weight: 700;
      }
      .linkleftbox {
        width: 50%;
        height: 15%;
        // background-color: #fff;
        position: absolute;
        top: 45%;
        left: -10%;
        background: url("../assets/newpng/center_mian_big_bgi.png") no-repeat
          center;
        background-size: 100% 100%;
        font-size: 1.25rem;
        font-weight: 700;
      }
      .linkrightbox {
        width: 50%;
        height: 15%;
        // background-color: #fff;
        position: absolute;
        top: 45%;
        right: -10%;
        background: url("../assets/newpng/center_mian_big_bgi.png") no-repeat
          center;
        background-size: 100% 100%;
        font-size: 1.25rem;
        font-weight: 700;
      }
      .linkrightbox_ {
        width: 50%;
        height: 15%;
        position: absolute;
        bottom: 15%;
        right: 25%;
        background: url("../assets/newpng/center_mian_big_bgi.png") no-repeat
          center;
        background-size: 100% 100%;
        font-size: 1.25rem;
        font-weight: 700;
      }
      .bigbox {
        position: absolute;
        width: 120%;
        height: 100%;
        z-index: 10;
        left: -10%;
        background: url("../assets/newpng/center_bgi_1.png") no-repeat center;
        background-size: 100% 70%;
        display: flex;
        align-items: center;
        justify-content: center;
        background-position: 0 15%;
        z-index: -100;
        .middlebox {
          width: 60vh;
          height: 60vh;
          background: url("../assets/newpng/center_bgi_2.png") no-repeat center;
          background-size: 100% 100%;
          transform: rotate(0deg);
          animation-timing-function: linear;
          animation-iteration-count: infinite;
          animation-name: minus;
          animation-duration: 8s;
          .middlebox2 {
            width: 100%;
            height: 100%;
            background: url("../assets/newpng/center_bgi_3.png") no-repeat
              center;
            background-size: 92% 92%;
            background-position: 50% 55%;
            transform: rotate(0deg);
            animation-timing-function: linear;
            animation-iteration-count: infinite;
            animation-name: minus_2;
            animation-duration: 8s;
            .middlebox3 {
              width: 100%;
              height: 100%;
              background: url("../assets/newpng/center_bgi_4.png") no-repeat
                center;
              background-size: 92% 92%;
              background-position: 50% 55%;
              transform: rotate(0deg);
              animation-timing-function: linear;
              animation-iteration-count: infinite;
              animation-name: minus_3;
              animation-duration: 8s;
              .minbox {
                width: 100%;
                height: 100%;
                background: url("../assets/newpng/center_bgi_5.png") no-repeat
                  center;
                background-size: 92% 92%;
                background-position: 50% 55%;
                transform: rotate(0deg);
                animation-timing-function: linear;
                animation-iteration-count: infinite;
                animation-name: minus_4;
                animation-duration: 8s;
              }
            }
          }
        }
      }
    }
    .leftasid {
      height: 100%;
      width: 100%;
      padding: 1%;
      box-sizing: border-box;
      .boxLeft {
        position: relative;
        text-align: right;
        width: 100%;
        height: 50%;
        background: url("../assets/newpng/homeTextbgileft.svg") no-repeat center;
        background-size: 105%;
        // background-size: cover !important;
      }
      // background-color: #fff;
    }
    .rightasid {
      height: 100%;
      width: 100%;
      padding: 1%;
      box-sizing: border-box;
      // background-color: #fff;
      .boxRight {
        position: relative;
        text-align: left;
        width: 100%;
        height: 50%;
        background: url("../assets/newpng/homeTextbgiright.svg") no-repeat
          center;
        background-size: 105%;
        // background-size: cover !important;
      }
    }
    .title_new_left {
      height: 12%;
      width: 100%;
      color: azure;
      position: absolute;
      top: 24%;
      span {
        margin-right: 8%;
        font-size: 1.5rem;
        font-weight: 700;
        background-image: linear-gradient(
          #33bcdf,
          #dcf2ff
        ); //背景色渐变，默认从上到下
        -webkit-background-clip: text; //CSS3属性：设定背景的绘制范围为文字
        color: transparent; //将字的颜色设置透明，只露出背景色
      }
    }
    .title_new_right {
      height: 12%;
      width: 100%;
      color: azure;
      position: absolute;
      top: 24%;
      span {
        margin-left: 8%;
        font-size: 1.5rem;
        font-weight: 700;
        background-image: linear-gradient(
          #33bcdf,
          #dcf2ff
        ); //背景色渐变，默认从上到下
        -webkit-background-clip: text; //CSS3属性：设定背景的绘制范围为文字
        color: transparent; //将字的颜色设置透明，只露出背景色
        // height: 100%;
        // line-height: 2.29rem;
        // font-size: 1rem;
        // font-weight: bold;
      }
    }
    .contentBox {
      width: 100%;
      height: 85%;
      // background-color: aqua;
      // border-bottom: rgba(72, 159, 193, 0.5) 1px solid;
      // border-left: rgba(72, 159, 193, 0.5) 1px solid;
      // border-right: rgba(72, 159, 193, 0.5) 1px solid;
      color: #fff;
      position: absolute;
      top: 35%;
      left: 10%;
      font-size: 1.25rem;
      ul li {
        list-style: disc;
        text-align: left;
      }
    }

    .typing {
      font-size: 1.1rem;
      color: #fff;
      text-shadow: 0 0 10px #0ebeff, 0 0 20px #0ebeff, 0 0 50px #0ebeff,
        0 0 100px #0ebeff, 0 0 200px #0ebeff;
      padding: 0.75rem;
      box-sizing: border-box;
    }

    .link_btn {
      // float: right;
      border: none;
      text-align: center;
      display: inline-block;
      width: 100%;
      height: 100%;
      color: aliceblue;
      text-decoration: none;
      margin-top: 1rem;
      list-style: none;
    }
  }
}
</style>
