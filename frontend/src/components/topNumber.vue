<!--
 * @Author: jacob
 * @Date: 2020-12-10 11:02:20
 * @LastEditTime: 2020-12-17 15:29:51
 * @LastEditors: jacob
 * @Description: echarts通用组件
-->
<template>
  <div class="topnumber_box">
    <p class="title_top">{{ this.titleTxt }}</p>
    <div class="top_bottom_content">
      <countTo
        class="nub"
        :startVal="0"
        :endVal="invalue"
        :duration="300"
        :suffix="suffix_"
      ></countTo>
      <el-progress
        :text-inside="true"
        :stroke-width="16"
        :percentage="invalue"
        stroke-linecap="square"
        :color="colorTop"
        class="progressbar"
      ></el-progress>
    </div>
  </div>
</template>

<script>
import countTo from "vue-count-to";
import {
  getVMCData,
  getDeviceData,
  getAppInfo,
  getTopoShow,
  insertNode,
  deleteNode,
  failureOver,
} from "../api";
export default {
  props: ["invalue", "color", "suffix_", "titleTxt"],
  name: "topNumber",
  components: { countTo },
  data() {
    return {
      lineData: [],
      colorTop: [
        { color: "rgb(238, 73, 73)", percentage: 90 },
        { color: "rgb(223, 94, 35)", percentage: 80 },
        { color: "rgb(35, 223, 214)", percentage: 60 },
        { color: "rgb(35, 223, 113)", percentage: 40 },
        { color: "rgb(126, 223, 35)", percentage: 20 },
      ],
      // colorTop: [
      //   {
      //     color: "linear-gradient(to right,rgb(193, 40, 18),rgb(243, 107, 92))",
      //     percentage: 90,
      //   },
      //   {
      //     color:
      //       "linear-gradient(to right,rgb(26, 183, 127),rgb(29, 255, 242))",
      //     percentage: 0,
      //   },
      //   // {color:'linear-gradient(to right,rgb(193, 40, 18),rgb(243, 107, 92))', percentage:90},
      //   // {color:'linear-gradient(to right,rgb(193, 40, 18),rgb(243, 107, 92))', percentage:90},
      //   // {color:'linear-gradient(to right,rgb(193, 40, 18),rgb(243, 107, 92))', percentage:90},
      // ],
    };
  },
  methods: {
    async getinfo() {
      let res = await getTopoShow();
      console.log(res);
    },
  },
  mounted() {
    this.getinfo();
  },
};
</script>

<style lang='less' >
.topnumber_box {
  height: 100%;
  width: 100%;
  background: url("../assets/png/topnubbg.png") no-repeat center;
  background-size: 100% 100%;
  .title_top {
    display: block;
    padding: 0;
    margin: 0;
    height: 30%;
    width: 100%;
    color: #fff;
    font-size: 1rem;
    text-align: center;
    line-height: 100%;
    padding-top: 0.2rem;
    box-sizing: border-box;
    // background-color: #fff;
  }
  .top_bottom_content {
    height: 70%;
    color: turquoise;
    // background-color: rgb(126, 223, 35);
    text-align: center;
    .nub {
      font-size: 2.3rem;
      text-shadow: 0 0 1rem rgb(148, 240, 231);
    }
    .progressbar {
      border-radius: 0 !important;
      width: 80%;
      height: 1rem;
      margin-left: 10%;
    }
    .el-progress-bar__outer {
      // border-radius: 0 !important;
      background-color: #24344a !important;
    }
  }
}
</style>
