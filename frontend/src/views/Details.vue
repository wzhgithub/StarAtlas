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
            <div class="mian_bottom_main_box"></div>
          </el-row>
        </div>
      </el-col>
      <el-col :span="8" class="colBox">
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
          <div v-show="asidvisiber[1]" class="aside_box"></div>
          <div v-show="asidvisiber[2]" class="aside_box"></div>
          <div v-show="asidvisiber[3]" class="aside_box"></div>
          <div class="aside_box_line_bar">
            <p class="title">
              <span>机器历史状况</span>
            </p>
            <div class="canvasbox" id="linebox_">
              <selflineNew inref="linebox_" />
            </div>
          </div>
          <div class="aside_box_task">
            <p class="title">
              <span>分时分区任务</span>
            </p>
            <scrolltable :data="tableData" />
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import CpuInfo from "@/components/CpuInfo.vue";
import selflineNew from "@/components/selflineNew.vue";
import scrolltable from "@/components/scrollTable.vue";
import topNumber from "@/components/topNumber.vue";
export default {
  name: "details",
  components: {
    CpuInfo,
    selflineNew,
    scrolltable,
    topNumber,
  },
  data() {
    return {
      data_: "",
      asidvisiber: [true, false, false, false],
      topdata: [20, 86, 73, 67, 78],
      cpu: {
        name: "",
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
  },
  mounted() {
    let that = this;
    setInterval(() => {
      that.topdata = [
        that.randomRange(0, 100),
        that.randomRange(0, 100),
        that.randomRange(0, 100),
        that.randomRange(0, 100),
        that.randomRange(0, 100),
      ];
    }, 1000);
  },
};
</script>
<style lang="less">
.details {
  width: 100%;
  height: 91.5vh;
  .colBox {
    height: 90vh;
    .grid-content {
      width: 100%;
      height: 100%;
      // background-color: antiquewhite;
      padding: 20px;
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
        background-color: rgb(0, 255, 136);
        margin: 0 !important;
        padding: 2%;
        box-sizing: border-box;
        .mian_bottom_main_box {
          width: 100%;
          height: 100%;
          background-color: pink;
        }
      }
    }
    .grid-content_ {
      width: 100%;
      height: 100%;
      // background-color: greenyellow;
      padding: 1%;
      box-sizing: border-box;
      overflow-y: scroll;
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
        padding: 15px;
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
          line-height: 32px;
        }
      }
    }
  }
}
</style>
