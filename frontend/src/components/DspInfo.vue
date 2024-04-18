<template>
  <div class="dsp_info_box">
    <p class="title_dsp">
      <!-- <span>{{ this.dspNow.name || "DSP信息详情" }}</span> -->
      <span>{{ "DSP信息详情" }}</span>
    </p>
    <el-row class="content">
      <el-col :span="8" class="content_col">
        <div class="left_part">
          <img class="miansvg" src="../assets/newpng/DSP.svg" alt="" />
        </div>
      </el-col>
      <el-col :span="16" class="content_col">
        <div class="right_part">
          <!-- <p>设备ID：2000044</p>
          <p>设备状态：就绪</p>
          <p>总算力：2376</p>
          <p>设备类型：INTERDSP 1800</p>
          <p>设备温度：24°C</p> -->
          <p>设备类型: {{ nowdata.type }}</p>
          <p>设备数量: {{ nowdata.status }}</p>
          <p>DSP核数: {{ nowdata.cores }}</p>
          <p>浮点算力: {{ Number(nowdata.float_power) }} TFLOPS</p>
          <p>整型算力: {{ Number(nowdata.int_power) }} MIPS</p>
          <p>DSP利用率: {{ nowdata.usage }}%</p>
          <p>内存利用率: {{ nowdata.memory_usage }}%</p>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import imgcpu from "@/assets/newpng/CPU.svg";
import imggpu from "@/assets/newpng/GPU.svg";
import imgdsp from "@/assets/newpng/DSP.svg";
import imgfgpa from "@/assets/newpng/FPGA.svg";
import countTo from "vue-count-to";
export default {
  name: "DspInfo",
  components: {
    countTo,
  },
  props: {
    dspNow: {
      type: Object,
      name: "",
      // default: function () {
      //   return {};
      // },
    },
  },
  data() {
    return {
      // nowdata: {
      //   cpuuseage: 70,
      //   canuse: 12,
      //   used: 80,
      //   computeuseage: 60,
      //   canuse_: 1200,
      //   alluse_: 2400,
      //   canusecore: 8,
      //   allcore: 32,
      //   canusemb: 1024,
      //   allmb: 10240,
      // },
    };
  },
  methods: {
    randomRange(min, max) {
      // min最小值，max最大值
      return Math.floor(Math.random() * (max - min)) + min;
    },
  },
  computed: {
    nowdata: function () {
      return { ...this.dspNow };
    },
  },
  mounted() {
    let that = this;
    // setInterval(() => {
    //   that.nowdata = {
    //     cpuuseage: that.randomRange(0, 100),
    //     canuse: that.randomRange(0, 100),
    //     used: that.randomRange(0, 100),
    //     computeuseage: that.randomRange(0, 100),
    //     canuse_: that.randomRange(1000, 1600),
    //     alluse_: that.randomRange(2000, 3600),
    //     canusecore: that.randomRange(0, 16),
    //     allcore: 32,
    //     canusemb: that.randomRange(100, 1024),
    //     allmb: 1024,
    //   };
    // }, 3000);
  },
};
</script>
<style lang="less">
.dsp_info_box {
  width: 100%;
  height: 98%;
  // background-color: #fff;
  .title_dsp {
    height: 15%;
    width: 100%;
    color: azure;
    background: url("../assets/png/part_title.png") no-repeat center;
    background-size: 100% 100%;
    padding-left: 12%;
    margin: 1%;
    box-sizing: border-box;
    text-align: left;
    span {
      height: 100%;
      line-height: 2rem;
    }
  }
  .content {
    height: 85%;
    padding: 0;
    // background-color: #fff;
    .content_col {
      height: 100%;

      .subtitle {
        padding: 0;
        margin: 0;
        padding-left: 12%;
        height: 30%;
        font-size: 0.8rem;
        color: #fff;
        line-height: 2rem;
        background: url("../assets/png/subtitlebg.png") no-repeat center;
        background-size: 100% 100%;
      }
      .left_part {
        height: 100%;
        text-align: center;
        // background-color: chartreuse;
        .miansvg {
          height: 75%;
          // margin-left: 25%;
        }
      }
      .right_part {
        height: 100%;
        text-align: left;
        // background: #fff;
        // background-color: rgb(238, 255, 0);
        color: aliceblue;
        p {
          margin-left: 25%;
        }
      }
    }
  }
}
</style>
