<template>
  <div class="details">
    <el-row :gutter="24" class="contentBox">
      <el-col :span="16" class="colBox">
        <div class="grid-content">
          <el-row type="flex" justify="space-around" class="row_main_top">
            <el-col :span="4" class="col_main_top">
              <div class="grid-content_top">
                <topNumber
                  :invalue="topdata[0]"
                  suffix_="%"
                  titleTxt="内存使用率"
                />
              </div>
            </el-col>
            <el-col :span="4" class="col_main_top">
              <div class="grid-content_top">
                <topNumber
                  :invalue="topdata[1]"
                  suffix_="%"
                  titleTxt="外存使用率"
                />
              </div>
            </el-col>
            <el-col :span="4" class="col_main_top">
              <div class="grid-content_top">
                <topNumber
                  :invalue="topdata[2]"
                  suffix_="%"
                  titleTxt="CPU使用率"
                />
              </div>
            </el-col>
            <el-col :span="4" class="col_main_top">
              <div class="grid-content_top">
                <topNumber
                  :invalue="topdata[3]"
                  suffix_="%"
                  titleTxt="GPU使用率"
                />
              </div>
            </el-col>
            <el-col :span="4" class="col_main_top">
              <div class="grid-content_top">
                <topNumber
                  :invalue="topdata[4]"
                  suffix_="%"
                  titleTxt="DSP使用率"
                />
              </div>
            </el-col>
          </el-row>
          <el-row type="flex" justify="space-around" class="row_main_bottom">
            <div class="mian_bottom_main_box">
              <div id="wrap">
                <div id="showcase" class="noselect">
                  <img
                    class="cloud9-item"
                    src="../assets/newpng/CPU.svg"
                    alt="Firefox"
                    @click="dealWithTypeClick(0)"
                  />

                  <img
                    class="cloud9-item"
                    src="../assets/newpng/FPGA.svg"
                    alt="Opera"
                    @click="dealWithTypeClick(3)"
                  />
                  <img
                    class="cloud9-item"
                    src="../assets/newpng/DSP.svg"
                    alt="Wyzo"
                    @click="dealWithTypeClick(2)"
                  />
                  <img
                    class="cloud9-item"
                    src="../assets/newpng/GPU.svg"
                    alt="Opera"
                    @click="dealWithTypeClick(1)"
                  />
                </div>
              </div>
              <h3 class="h3_title">设备：vmc1</h3>
            </div>
          </el-row>
        </div>
      </el-col>
      <el-col :span="8" class="colBox">
        <vue-scroll :ops="ops">
          <div class="grid-content_">
            <el-row type="flex" class="row-bg" justify="space-around">
              <el-col :span="6">
                <div
                  :class="
                    asidvisiber[0]
                      ? 'grid-content_btn active_btn'
                      : 'grid-content_btn'
                  "
                >
                  <button @click="dealWithTypeClick(0)">CPU</button>
                </div>
              </el-col>
              <el-col :span="6">
                <div
                  :class="
                    asidvisiber[1]
                      ? 'grid-content_btn active_btn'
                      : 'grid-content_btn'
                  "
                >
                  <button @click="dealWithTypeClick(1)">GPU</button>
                </div>
              </el-col>
              <el-col :span="6">
                <div
                  :class="
                    asidvisiber[2]
                      ? 'grid-content_btn active_btn'
                      : 'grid-content_btn'
                  "
                >
                  <button @click="dealWithTypeClick(2)">DSP</button>
                </div>
              </el-col>
              <el-col :span="6">
                <div
                  :class="
                    asidvisiber[3]
                      ? 'grid-content_btn active_btn'
                      : 'grid-content_btn'
                  "
                >
                  <button @click="dealWithTypeClick(3)">FPGA</button>
                </div>
              </el-col>
            </el-row>
            <div v-show="asidvisiber[0]" class="aside_box">
              <CpuInfo :cpuNow="cpu" />
            </div>
            <div v-show="asidvisiber[1]" class="aside_box">
              <GpuInfo :gpuNow="gpu" />
            </div>
            <div v-show="asidvisiber[2]" class="aside_box">
              <DspInfo :dspNow="dsp" />
            </div>
            <div v-show="asidvisiber[3]" class="aside_box">
              <FpgaInfo :fpgaNow="fpga" />
            </div>
            <div class="aside_box_line_bar">
              <p class="title">
                <span>机器历史状况</span>
              </p>
              <div class="canvasbox" id="linebox_">
                <selflineNew inref="linebox_" :vmcid="vmcid" />
              </div>
            </div>
            <div class="aside_box_task">
              <p class="title">
                <span>分时分区任务</span>
              </p>
              <scrolltable :data="tableData" />
            </div>
          </div>
        </vue-scroll>
      </el-col>
    </el-row>
  </div>
</template>
<script>
// import "@/lib/jq_a.js";
import $ from "@/lib/jq_a.js";
// import $ from "jquery";
import "@/lib/jquery.reflection.js";
import "@/lib/jquery.cloud9carousel.js";
import "@/lib/main.css";
import "@/lib/jquerysctipttop.css";

import CpuInfo from "@/components/CpuInfo.vue";
import GpuInfo from "@/components/GpuInfo.vue";
import DspInfo from "@/components/DspInfo.vue";
import FpgaInfo from "@/components/FpgaInfo.vue";
import selflineNew from "@/components/selflineNew.vue";
import scrolltable from "@/components/scrollTable.vue";
import topNumber from "@/components/topNumber.vue";
import { getVMCData } from "@/api";
import { getVMCDataSeq } from "../api";
export default {
  name: "details",
  components: {
    CpuInfo,
    GpuInfo,
    DspInfo,
    FpgaInfo,
    selflineNew,
    scrolltable,
    topNumber,
  },
  data() {
    return {
      ops: {
        bar: {
          hoverStyle: true,
          onlyShowBarOnScroll: true, //是否只有滚动的时候才显示滚动条
          showDelay: 0, // 在鼠标离开容器后多长时间隐藏滚动条。
          keepShow: true, // 滚动条是否保持显示。
          background: "#ccc", // 滚动条背景色。
          "overflow-x": "hidden",
        },
      },
      data_: "",
      asidvisiber: [true, false, false, false],
      topdata: [20, 86, 73, 67, 78],
      vmcid: "",
      cpu: {
        name: "CPU信息详情",
        cpuuseage: 70,
        canuse: 12,
        used: 80,
        computeuseage: 60,
        canuse_: 1200,
        alluse_: 2400,
        canusecore: 8,
        allcore: 32,
        canusemb: 1024,
        allmb: 10240,
      },
      gpu: {
        name: "",
        type: "NIDIA AGX",
        float_power: 0,
        usage: 60,
        cores: 1,
        memory_usage: 50,
        total_memory: 65535,
        status: "健康",
      },
      dsp: {
        name: "",
        type: "",
        float_power: 0,
        int_power: 0,
        usage: 0,
        cores: 1,
        memory_usage: 0,
        total_memory: 65535,
        status: "健康",
      },
      fpga: {
        name: "",
        type: "FPGA",
        float_power: 0,
        int_power: 0,
        usage: 0,
        cores: 1,
        memory_usage: 0,
        total_memory: 65535,
        status: "健康",
      },
      tableData: [
        {
          productName: "核心任务_12654_456",
          coreName: "执行完成",
          publish: "james",
          publishAmount: 23567,
        },
        {
          productName: "核心任务_2654_456",
          coreName: "执行完成",
          publish: "james",
          publishAmount: 356,
        },
        {
          productName: "普通任务_654_456",
          coreName: "执行完成",
          publish: "james",
          publishAmount: 56,
        },
        {
          productName: "非必要任务_456",
          coreName: "执行失败",
          publish: "anna",
          publishAmount: 6,
        },
        {
          productName: "普通任务_12654_456",
          coreName: "执行失败",
          publish: "李明",
          publishAmount: 678954,
        },
        {
          productName: "核心任务_54_456",
          coreName: "执行完成",
          publish: "anna",
          publishAmount: 5765489,
        },
        {
          productName: "普通任务_1884_456",
          coreName: "执行失败",
          publish: "李明",
          publishAmount: 773245,
        },
        {
          productName: "核心任务_4_456",
          coreName: "执行完成",
          publish: "李明",
          publishAmount: 978654,
        },
        {
          productName: "普通任务_2654_456",
          coreName: "执行完成",
          publish: "anna",
          publishAmount: 3856,
        },
      ],
    };
  },
  methods: {
    async getTopData() {
      const { data } = await getVMCData(this.$route.params.id);
      console.log(data);
      this.topdata = [data.data.memory_usage, data.data.total_disk_usage,
              data.data.total_cpu_usage, data.data.total_gpu_usage, data.data.total_dsp_usage];
      this.cpu = {
        name: "CPU信息详情",
        cpuuseage: data.data.total_cpu_usage,
        canuse: 100 - data.data.total_cpu_usage,
        used: data.data.total_cpu_usage,
        // computeuseage: 32768 / 32768.0 * 100,
        // alluse_: 32768,
        // canuse_: 32768 - data.data.cpu_set[0].int_computing_power,
        computeuseage: data.data.memory_usage,
        canuse_: data.data.total_memory * (100 - data.data.memory_usage) / 100,
        alluse_: data.data.total_memory,
        canusecore: data.data.cpu_number,
        allcore: data.data.cpu_number,
        canusemb: data.data.cpu_set[0].float_computing_power ? data.data.cpu_set[0].float_computing_power : 0,
        allmb: data.data.cpu_set[0].int_computing_power ? data.data.cpu_set[0].float_computing_power : 0,
        // canusemb: data.data.total_memory * (100 - data.data.memory_usage) / 100,
        // allmb: data.data.total_memory
      }
      this.gpu = {
        name: data.data.gpu_set[0].name,
        float_power: data.data.gpu_set[0].float_computing_power,
        cores: data.data.gpu_set[0].num,
        usage: data.data.gpu_set[0].usage,
        memory_usage: data.data.gpu_set[0].memory_usage,
        total_memory: data.data.gpu_set[0].total_memory,
        type: "NIDIA AGX",
        status: "健康",
      }
      this.dsp = {
        name: data.data.dsp_set[0].name,
        float_power: data.data.dsp_set[0].float_computing_power,
        int_power: data.data.dsp_set[0].int_computing_power,
        cores: data.data.dsp_set[0].num,
        usage: data.data.dsp_set[0].usage,
        memory_usage: data.data.dsp_set[0].memory_usage,
        total_memory: data.data.dsp_set[0].total_memory,
        type: data.data.dsp_set[0].type,
        status: "健康",
      }
      this.fpga = {
        name: data.data.fpga_set[0].name,
        float_power: data.data.fpga_set[0].float_computing_power,
        int_power: data.data.fpga_set[0].int_computing_power,
        cores: data.data.fpga_set[0].num,
        usage: data.data.fpga_set[0].usage,
        memory_usage: data.data.fpga_set[0].memory_usage,
        total_memory: data.data.fpga_set[0].total_memory,
        type: data.data.fpga_set[0].type,
        status: "健康",
      }
    },
    dealWithTypeClick(num) {
      // this.asidvisiber = [false,false,false,false]
      let arr = [false, false, false, false];
      arr[num] = true;
      this.asidvisiber = arr;
    },
    randomRange(min, max) {
      // min最小值，max最大值
      return Math.floor(Math.random() * (max - min)) + min;
    },
    drawall() {
      console.log($);
      $(function () {
        var showcase = $("#showcase");
        showcase.Cloud9Carousel({
          yPos: 42,
          yRadius: 48,
          mirrorOptions: {
            gap: 12,
            height: 0.2,
          },
          buttonLeft: $(".nav > .left"),
          buttonRight: $(".nav > .right"),
          autoPlay: true,
          bringToFront: true,
          onRendered: showcaseUpdated,
          onLoaded: function () {
            showcase.css("visibility", "visible");
            showcase.css("display", "none");
            showcase.fadeIn(1500);
          },
        });

        function showcaseUpdated(showcase) {
          var title = $("#item-title").html(
            $(showcase.nearestItem()).attr("alt")
          );

          var c = Math.cos((showcase.floatIndex() % 1) * 2 * Math.PI);
          title.css("opacity", 0.5 + 0.5 * c);
        }

        // Simulate physical button click effect
        $(".nav > button").click(function (e) {
          var b = $(e.target).addClass("down");
          setTimeout(function () {
            b.removeClass("down");
          }, 80);
        });

        $(document).keydown(function (e) {
          switch (e.keyCode) {
            /* left arrow */
            case 37:
              $(".nav > .left").click();
              break;

            /* right arrow */
            case 39:
              $(".nav > .right").click();
          }
        });
      });
    },
  },
  mounted() {
    let that = this;
    this.vmcid = this.$route.params.id;
    setInterval(() => {
      that.getTopData();
      // that.topdata = [
      //   that.randomRange(0, 100),
      //   that.randomRange(0, 100),
      //   that.randomRange(0, 100),
      //   that.randomRange(0, 100),
      //   that.randomRange(0, 100),
      // ];
    }, 5000);
    this.drawall();
  },
};
</script>
<style lang="less" scoped>
.details {
  width: 100%;
  height: 91.5vh;
  .colBox {
    height: 90vh;
    .grid-content {
      width: 100%;
      height: 100%;
      // background-color: antiquewhite;
      padding: 0.7rem;
      box-sizing: border-box;
      .row_main_top {
        width: 100%;
        height: 18%;
        // background-color: yellow;
        margin: 0 !important;
        padding: 1%;
        box-sizing: border-box;
        .col_main_top {
          height: 100%;
          .grid-content_top {
            width: 100%;
            height: 100%;
            // background-color: pink;
          }
        }
      }
      .row_main_bottom {
        width: 100%;
        height: 82%;
        // background-color: rgb(0, 255, 136);
        margin: 0 !important;
        padding: 2%;
        box-sizing: border-box;
        .mian_bottom_main_box {
          position: relative;
          width: 100%;
          height: 100%;
          background: url("../assets/png/deallitle.png") no-repeat center;
          background-size: 100% 100%;
          background-position-y: 3.6rem;
          // background-color: pink;src/assets/png/deallitle.png
          img {
            height: 45%;
          }
        }
      }
    }
    .grid-content_ {
      width: 100%;
      height: 100%;
      // background-color: greenyellow;
      padding: 1%;
      box-sizing: border-box;
      // overflow-y: scroll;
      .aside_box {
        height: 30vh;
        // background: rgb(207, 160, 160);
        // padding: 1%;
        // box-sizing: border-box;
      }
      .aside_box_line_bar {
        height: 30vh;
        // background-color: #fff;
        .canvasbox {
          height: 85%;
          width: 100%;
        }
        // background-color: aqua;
        // padding: 1%;
        // box-sizing: border-box;
      }
      .aside_box_task {
        height: 30vh;
        // background-color: rgb(204, 0, 255);
        // padding: 1%;
        // box-sizing: border-box;
      }
      .grid-content_btn {
        padding: 1rem;
        button {
          width: 100%;
          height: 4vh;
          background: url("../assets/png/btnbgc_.png") no-repeat center;
          background-size: 100% 100%;
          border: none;
          color: rgba(255, 255, 255, 0.6);
          &:hover {
            background: url("../assets/png/btnbgc.png") no-repeat center;
            background-size: 100% 100%;
            color: rgb(222, 241, 255);
            text-shadow: rgb(30 222 255 / 37%) 0px 0px 9px;
          }
        }
      }
      .active_btn {
        button {
          background: url("../assets/png/btnbgc.png") no-repeat center;
          background-size: 100% 100%;
          color: rgb(222, 241, 255);
          text-shadow: rgb(30 222 255 / 37%) 0px 0px 9px;
        }
      }
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
          line-height: 2.29rem;
        }
      }
    }
  }
}
.h3_title {
  position: absolute;
  color: aqua;
  font-size: 2rem;
  font-weight: 700;
  bottom: 0.8%;
  left: 40%;
}
</style>
