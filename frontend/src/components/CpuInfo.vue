<template>
  <div class="cpu_info_box">
    <p class="title">
      <span>{{ this.cpuNow.name || "CPU信息详情" }}</span>
    </p>
    <el-row class="content">
      <el-col :span="14" class="content_col">
        <div class="left_part">
          <p class="subtitle">
            <span>CPU利用率</span>&nbsp;&nbsp;&nbsp;
            <el-progress
              :text-inside="true"
              :stroke-width="26"
              :percentage="nowdata.cpuuseage"
              stroke-linecap="square"
              class="progressbar"
            ></el-progress>
            <el-row type="flex" justify="space-between">
              <el-col :span="10">
                <div class="">
                  已使用:
                  <countTo
                    class="nub"
                    :startVal="0"
                    :endVal="nowdata.canuse"
                    :duration="1000"
                    suffix=""
                  />
                </div>
              </el-col>
              <el-col :span="10">
                <div class="">
                  可用:
                  <countTo
                    class="nub"
                    :startVal="0"
                    :endVal="nowdata.used"
                    :duration="1000"
                    suffix=""
                  />
                </div>
              </el-col>
            </el-row>
          </p>
        </div>
        <div class="left_part_">
          <p class="subtitle">
            <span>算力状况</span>&nbsp;&nbsp;&nbsp;<span></span>
            <el-progress
              :text-inside="true"
              :stroke-width="26"
              :percentage="nowdata.computeuseage"
              stroke-linecap="square"
              class="progressbar"
            ></el-progress>
            <el-row type="flex" justify="space-between">
              <el-col :span="10">
                <div class="">
                  总算力:
                  <countTo
                    class="nub"
                    :startVal="0"
                    :endVal="nowdata.alluse_"
                    :duration="1000"
                    suffix=""
                  />
                </div>
              </el-col>
              <el-col :span="10">
                <div class="">
                  可用算力:
                  <countTo
                    class="nub"
                    :startVal="0"
                    :endVal="nowdata.canuse_"
                    :duration="1000"
                    suffix=""
                  />
                </div>
              </el-col>
            </el-row>
          </p>
        </div>
      </el-col>
      <el-col :span="10" class="content_col">
        <div class="right_part">
          <p class="core_title">
            <span class="title_span">
              <countTo
                class="nub"
                :startVal="0"
                :endVal="nowdata.canusecore"
                :duration="1000"
                suffix="核"
              />
              /
              <countTo
                class="nub"
                :startVal="0"
                :endVal="nowdata.allcore"
                :duration="1000"
                suffix="核"
              />
            </span>
            <span class="title_span">
              <countTo
                class="nub"
                :startVal="0"
                :endVal="nowdata.canusemb"
                :duration="1000"
                suffix="MB" />/
              <countTo
                class="nub"
                :startVal="0"
                :endVal="nowdata.allmb"
                :duration="1000"
                suffix="MB"
            /></span>
          </p>
          <div class="type_box">
            <el-row type="flex" justify="center" class="type_row">
              <el-col :span="16" class="type_col">
                <span>CPU设备类型</span>
              </el-col>
              <el-col :span="8" class="type_col_">
                <p class="type_text_"></p>
                <p class="type_text">ADM</p>
                <p class="type_text active_text">龙芯</p>
                <p class="type_text">Intel</p>
                <p class="type_text">麒麟</p>
              </el-col>
            </el-row>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import countTo from "vue-count-to";
export default {
  name: "CpuInfo",
  components: {
    countTo,
  },
  props: {
    cpuNow: {
      type: Object,
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
    nowdata: function() {
      return {...this.cpuNow}
    }
  },
  mounted() {
    let that = this;
    setInterval(() => {
      // that.nowdata = {
      //   cpuuseage: that.randomRange(0, 100),
      //   canuse: that.randomRange(0, 100),
      //   used: that.randomRange(0, 100),
      //   computeuseage: that.randomRange(0, 100),
      //   canuse_: that.randomRange(1000, 1600),
      //   alluse_: that.randomRange(2000, 3600),
      //   canusecore: that.randomRange(0, 16),
      //   allcore: 32,
      //   canusemb: that.randomRange(100, 1024),
      //   allmb: 1024,
      // };
    }, 3000);
  },
};
</script>
<style lang="less">
.cpu_info_box {
  width: 100%;
  height: 98%;
  // background-color: #fff;
  .title {
    height: 15%;
    width: 100%;
    color: azure;
    background: url("../assets/png/part_title.png") no-repeat center;
    background-size: 100% 100%;
    padding-left: 12%;
    margin: 1%;
    box-sizing: border-box;
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
        height: 50%;
        // background-color: chartreuse;
        .el-progress-bar__inner {
          background: linear-gradient(
            to right,
            rgb(193, 40, 18),
            rgb(243, 107, 92)
          ) !important;
        }
      }
      .left_part_ {
        height: 50%;
        // background-color: rgb(0, 238, 255);
        .el-progress-bar__inner {
          background: linear-gradient(
            to right,
            rgb(26, 183, 127),
            rgb(29, 255, 242)
          ) !important;
        }
      }
      .right_part {
        height: 100%;
        // background-color: rgb(238, 255, 0);
        .core_title {
          padding: 0;
          margin: 0;
          height: 15%;
          white-space: nowrap;
          overflow: hidden;
          // background-color: chocolate;
          .title_span {
            margin-right: 10%;
            min-width: 50%;
            height: 100%;
            background: url("../assets/png/underline.png") no-repeat center;
            background-size: 100% 12%;
            background-position: 0 0.9rem;
            font-size: 0.7rem;
            color: rgba(15, 248, 248, 0.8);
            text-shadow: rgb(50 220 250 / 44%) 0px 0px 5px,
              rgb(50 220 250 / 44%) 0px 0px 10px,
              rgb(50 220 250 / 44%) 0px 0px 15px;
          }
        }
        .type_box {
          height: 85%;
          .type_row {
            width: 100%;
            height: 100%;
            .type_col {
              height: 100%;
              background: url("../assets/png/typebg.png") no-repeat center;
              background-size: 100% 70%;
              background-position: 0 60%;
              text-align: center;
              span {
                font-size: 1rem;
                margin-top: 0.5rem;
                color: #fff;
              }
              // background-color: #fff;
            }
            .type_col_ {
              height: 100%;
              .type_text_ {
                padding: 0;
                margin: 0;
                display: block;
                height: 25%;
              }
              .type_text {
                padding: 0;
                margin: 0;
                display: block;
                height: 15%;
                font-size: 0.8rem;
                color: #fff;
              }
              .active_text {
                color: rgba(15, 248, 248, 0.8);
                text-shadow: rgb(50 220 250 / 44%) 0px 0px 5px,
                  rgb(50 220 250 / 44%) 0px 0px 10px,
                  rgb(50 220 250 / 44%) 0px 0px 15px;
              }
              // background-color: rgb(100, 22, 22);
            }
          }
          // background-color: #fff;
        }
      }
    }
    .progressbar {
      margin-top: 5%;
      border-radius: 0 !important;
      width: 80%;
      .el-progress-bar__outer {
        border-radius: 0 !important;
        background-color: #24344a !important;
      }
      .el-progress-bar__inner {
        border-radius: 0 !important;
        height: 70%;
        margin-top: 1.85%;
        color: #3267ae;
      }
    }
  }
}
</style>
