/* eslint-disable */
<template>
  <div class="disasterrecovery">
    <div class="titleBox">
      <h2>遥控指令触发</h2>
    </div>
    <div class="typeBox">
      指令类型：&nbsp;&nbsp;&nbsp;<el-radio-group v-model="type">
        <el-radio-button label="type1">容错模拟（整机故障）</el-radio-button>
        <el-radio-button label="type2">容错模拟（分区故障）</el-radio-button>
        <el-radio-button label="type3">性能测试</el-radio-button>
      </el-radio-group>
    </div>
    <div class="selectBox">
      <p v-if="type === 'type1'">
        操作目标：&nbsp;&nbsp;&nbsp;
        <el-cascader
          v-model="target"
          :options="options"
          :props="{ checkStrictly: false }"
          clearable
          style="width: 33%; margin-top: 1vh"
        ></el-cascader>
      </p>
      <p v-if="type === 'type2'">
        操作目标：&nbsp;&nbsp;&nbsp;
        <el-cascader
          v-model="target"
          :options="options_task"
          :props="{ checkStrictly: false }"
          clearable
          style="width: 33%; margin-top: 1vh"
        ></el-cascader>
      </p>
      <p v-if="type === 'type3'">
        测试目标：&nbsp;&nbsp;&nbsp;
        <el-cascader
          v-model="cvalue"
          :options="options"
          :props="{ checkStrictly: false }"
          clearable
          style="width: 33%; margin-top: 1vh"
        ></el-cascader>
      </p>
    </div>
    <div class="selectBox">
      <el-button
        v-if="type !== 'type3'"
        style="width: 10%"
        type="success"
        @click="handlerDo(0)"
        round
        >执行</el-button
      >
      <el-button-group style="width: 20%" v-if="type === 'type3'">
        <el-button round style="width: 50%" type="primary" @click="handlerDo(1)"
          >整型算力测试</el-button
        >
        <el-button round style="width: 50%" type="primary" @click="handlerDo(2)"
          >浮点算力测试</el-button
        >
      </el-button-group>
      <!-- <el-button style="width: 10%; margin-left: 5%" type="info" round
        >复位</el-button
      > -->
    </div>
    <div class="endBox">
      <div class="infoBoard">
        <el-steps
          v-if="type !== 'type3'"
          style="width: 50%"
          direction="vertical"
          :active="stepNow"
        >
          <el-step title="建立通信" :description="description0">
            <i v-if="stepNow > 1" slot="icon" class="el-icon-success"></i>
            <img
              v-if="stepNow === 1"
              class="gif"
              slot="icon"
              src="../assets/newpng/loading_new.gif"
              alt=""
            />
          </el-step>
          <el-step title="远程计算" :description="description1">
            <i v-if="stepNow > 2" slot="icon" class="el-icon-success"></i>
            <img
              v-if="stepNow === 2"
              class="gif"
              slot="icon"
              src="../assets/newpng/loading_new.gif"
              alt=""
            />
          </el-step>
          <el-step title="结果处理" :description="description2">
            <i v-if="stepNow > 3" slot="icon" class="el-icon-success"></i>
            <img
              v-if="stepNow === 3"
              class="gif"
              slot="icon"
              src="../assets/newpng/loading_new.gif"
              alt=""
            />
          </el-step>
        </el-steps>
        <el-steps
          v-if="type === 'type3'"
          style="width: 50%"
          direction="vertical"
          :active="stepNow_"
        >
          <el-step title="算力压测准备" :description="description3">
            <i v-if="stepNow_ > 1" slot="icon" class="el-icon-success"></i>
            <img
              v-if="stepNow_ === 1"
              class="gif"
              slot="icon"
              src="../assets/newpng/loading_new.gif"
              alt=""
            />
          </el-step>
          <el-step title="算力测试中" :description="description4">
            <i v-if="stepNow_ > 2" slot="icon" class="el-icon-success"></i>
            <img
              v-if="stepNow_ === 2"
              class="gif"
              slot="icon"
              src="../assets/newpng/loading_new.gif"
              alt=""
            />
          </el-step>
          <el-step title="算力测试结果" :description="description5">
            <i v-if="stepNow_ > 3" slot="icon" class="el-icon-success"></i>
            <img
              v-if="stepNow_ === 3"
              class="gif"
              slot="icon"
              src="../assets/newpng/loading_new.gif"
              alt=""
            />
          </el-step>
        </el-steps>
        <div class="asideBox" v-if="type === 'type3'">
          <div class="top">
            <div class="cell">
              <countTo
                class="nub"
                :startVal="0"
                :endVal="value_a"
                :duration="3000"
                suffix="MIPS"
              ></countTo>
              <p>整型算力</p>
            </div>
          </div>
          <div class="top">
            <div class="cell">
              <countTo
                class="nub"
                :startVal="0"
                :endVal="value_b"
                :duration="3000"
                suffix="TFLOPS"
                :decimals="3"
              ></countTo>
              <p>浮点算力</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
// import "@/lib/loaders.min.css";
import countTo from "vue-count-to";
import { mapState, mapMutations } from "vuex";
import {
  getTopoShow,
  filterName,
  getAppInfo,
  doFailureOver,
  getFailureResult,
} from "@/api";
export default {
  name: "RemoteControl",
  components: {
    // HelloWorld,
    countTo,
  },
  data() {
    return {
      type: "type1",
      target: "123",
      cvalue: "",
      options: [],
      options_task: [],
      options_a: [
        {
          label: "算法1",
          value: "c1",
        },
        {
          label: "算法2",
          value: "c2",
        },
      ],
      stepNow: 0,
      stepNow_: 0,
      description0: "",
      description1: "",
      description2: "",
      description3: "",
      description4: "",
      description5: "",
      allNodes: [],
      allTasks: [],
      value_a: 0,
      value_b: 0,
    };
  },
  methods: {
    ...mapMutations(["setFrom", "setTo"]),
    filterName,
    async postDoFailureOver(data, flag) {
      let res = await doFailureOver(data);
      if (res.data.code === 200) {
        this.mcokLoading(flag);
      }
    },
    async getDoFailureRuselt(data) {
      let res = await getFailureResult(data);
      return res.data;
    },
    randomNum(minNum, maxNum, decimalNum) {
      var max = 0,
        min = 0;
      minNum <= maxNum
        ? ((min = minNum), (max = maxNum))
        : ((min = maxNum), (max = minNum));
      return (Math.random() * (max - min) + min).toFixed(decimalNum);
    },
    romoteGetResult(data, type) {
      let that = this;
      setTimeout(() => {
        let end = this.getDoFailureRuselt(data);
        let res = {};
        end
          .then((resault) => {
            res = resault;
            if (res.code !== 200) {
              that.romoteGetResult(data);
            } else {
              let objTep = {};
              let divicedata = {};
              that.allNodes.map((item) => {
                if (item.id === res.data.to.vmc_id) {
                  objTep = item;
                }
              });
              objTep.id &&
                objTep.chArr.map((item) => {
                  if (item.value === res.data.to.device_id) {
                    divicedata = item;
                  }
                });
              if (type === "vmc") {
                that.setTo({
                  id: res.data.to.vmc_id,
                  type: type,
                  parent_id: null,
                  name: that.filterName(objTep.name),
                });
              } else {
                that.setTo({
                  id: res.data.to.device_id,
                  type: type,
                  parent_id: res.data.to.vmc_id,
                  name: that.filterName(divicedata.label),
                });
              }
            }
          })
          .catch(() => {
            that.romoteGetResult(data);
          });
      }, 1000);
    },
    mcokLoading(flag) {
      let that = this;
      if (flag === 0) {
        setTimeout(() => {
          that.stepNow = 1;
          that.description0 = "客户端正在发起建立连接.。。。。";
          that.description1 = "未执行到当前步骤";
          that.description2 = "未执行到当前步骤";
        }, 500);
        setTimeout(() => {
          that.stepNow = 1;
          that.description0 = "连接建立中，等待远端响应。。。";
        }, 1000);
        setTimeout(() => {
          that.stepNow = 1;
          that.description0 = "客户端建立连接成功！";
        }, 1500);
        setTimeout(() => {
          that.stepNow = 1;
          that.description0 = "客户端正在进行数据交互。。。。";
        }, 2000);
        setTimeout(() => {
          that.stepNow = 1;
          that.description0 = "客户端数据传输完成。。。。";
        }, 2500);
        setTimeout(() => {
          that.stepNow = 2;
          that.description0 = "连接建立成功";
          that.description1 = "平台正在进行数据计算。。。。";
        }, 3000);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（10%）";
        }, 3500);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（15%）";
        }, 4000);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（25%）";
        }, 4500);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（35%）";
        }, 5000);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（55%）";
        }, 5500);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（75%）";
        }, 6000);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（95%）";
        }, 6500);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度（100%）";
        }, 7000);
        setTimeout(() => {
          that.stepNow = 2;
          that.description1 = "平台正在进行数据计算进度完成，数据集开始处理";
        }, 7500);
        setTimeout(() => {
          that.stepNow = 3;
          that.description1 = "平台正在进行数据计算进度完成";
          that.description2 = "数据集处理中";
        }, 8000);
        setTimeout(() => {
          that.stepNow = 3;
          that.description2 = "数据集格式比对中";
        }, 8500);
        setTimeout(() => {
          that.stepNow = 3;
          that.description2 = "数据集校验中";
        }, 9000);
        setTimeout(() => {
          that.stepNow = 3;
          that.description2 = "数据集处理完成，返回数据中";
        }, 9500);
        setTimeout(() => {
          that.stepNow = 3;
          that.description2 = "安全断开链接中";
        }, 10000);
        setTimeout(() => {
          that.stepNow = 4;
          that.description2 = "数据中返回成功，目标迁移设备已确定";
          // if (!that.to.type && that.target.length === 1) {
          //   that.setTo({
          //     id: that.options[0].value,
          //     type: that.from.type,
          //     parent_id: null,
          //     name: "for_test",
          //     time: "",
          //   });
          // }
          // if (!that.to.type && that.target.length > 1) {
          //   that.setTo({
          //     id: 999,
          //     type: that.from.type,
          //     parent_id: that.target[0],
          //     name: "for_test",
          //     time: "",
          //   });
          // }
        }, 11500);
        setTimeout(() => {
          this.$message({
            message:
              "迁移目标计算成功，已为您重定向到容错演示页面查看任务迁移详情",
            type: "success",
          });
          this.$router.push("/disasterrecovery");
        }, 12000);
      }
      if (flag === 2) {
        setTimeout(() => {
          that.stepNow_ = 1;
          that.description3 = "正在选择相关算法";
          that.description4 = "未执行到当前步骤";
          that.description5 = "未执行到当前步骤";
        }, 500);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。";
        }, 1000);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。。";
        }, 1500);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。。。";
        }, 2500);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。。。。";
        }, 4500);
        setTimeout(() => {
          that.stepNow_ = 2;
          that.description3 = "已成功选择算法【drystone】，算法准备成功！";
          that.description4 = "平台正在进行算力测试。。。。";
        }, 5000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（10%）";
        }, 7500);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（15%）";
        }, 9000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（25%）";
        }, 11000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（35%）";
        }, 12000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（55%）";
        }, 13000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（75%）";
        }, 14000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（95%）";
        }, 15500);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（100%）";
        }, 16000);
        setTimeout(() => {
          that.description4 = "算力测试完成，正在输出相关指标";
        }, 17500);
        setTimeout(() => {
          that.stepNow_ = 3;
          that.description4 = "算力测试完成，相关指标已输出";
          that.description5 = "输出结果校验中";
        }, 18000);
        setTimeout(() => {
          that.description5 = "输出结果校验中。";
        }, 18500);
        setTimeout(() => {
          that.description5 = "输出结果校验中。。";
        }, 19000);
        setTimeout(() => {
          that.description5 = "输出结果校验中。。。";
        }, 19500);
        setTimeout(() => {
          let nub = that.randomNum(1510, 1790, 0);
          that.stepNow_ = 4;
          that.description5 = `当前ID为【${that.cvalue[0]}】的计算设备的整型算力为【${nub}MIPS】`;
          that.value_a = nub;
        }, 20000);
      }
      if (flag === 3) {
        setTimeout(() => {
          that.stepNow_ = 1;
          that.description3 = "正在选择相关算法";
          that.description4 = "未执行到当前步骤";
          that.description5 = "未执行到当前步骤";
        }, 500);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。";
        }, 2000);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。。";
        }, 3500);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。。。";
        }, 4000);
        setTimeout(() => {
          that.description3 = "正在选择相关算法。。。";
        }, 5500);
        setTimeout(() => {
          that.stepNow_ = 2;
          that.description3 = "已成功选择算法【whetstone】，算法准备成功！";
          that.description4 = "平台正在进行算力测试。。。。";
        }, 7000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（10%）";
        }, 8500);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（15%）";
        }, 9000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（25%）";
        }, 10000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（35%）";
        }, 11000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（55%）";
        }, 13500);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（75%）";
        }, 14000);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（95%）";
        }, 15500);
        setTimeout(() => {
          that.description4 = "平台正在进行算力测试（100%）";
        }, 16000);
        setTimeout(() => {
          that.description4 = "算力测试完成，正在输出相关指标";
        }, 17500);
        setTimeout(() => {
          that.stepNow_ = 3;
          that.description4 = "算力测试完成，相关指标已输出";
          that.description5 = "输出结果校验中";
        }, 18000);
        setTimeout(() => {
          that.description5 = "输出结果校验中。";
        }, 18500);
        setTimeout(() => {
          that.description5 = "输出结果校验中。。";
        }, 19000);
        setTimeout(() => {
          that.description5 = "输出结果校验中。。。";
        }, 19500);
        setTimeout(() => {
          let nub = that.randomNum(1.1, 1.38, 3);
          that.stepNow_ = 4;
          that.description5 = `当前ID为【${that.cvalue[0]}】的计算设备的浮点算力为【${nub}TFLOPS】`;
          that.value_b = nub;
        }, 20000);
      }
    },
    handlerDo(types) {
      if (types === 0) {
        if (this.type === "type1") {
          let objTep = {};
          this.allNodes.map((item) => {
            if (item.id === this.target[0]) {
              objTep = item;
            }
          });
          this.setFrom({
            id: this.target[0],
            type: "vmc",
            parent_id: null,
            name: objTep.name,
            time: new Date().valueOf(),
          });
          this.postDoFailureOver(
            {
              transType: 0,
              fromVmcId: this.target[0],
              isFault: 0,
            },
            0
          );
          this.romoteGetResult(
            {
              transType: 0,
              fromVmcId: this.target[0],
              isFault: 0,
            },
            "vmc"
          );
        }
        if (this.type === "type2") {
          let objTep = {};
          let divicedata = {};
          this.allNodes.map((item) => {
            if (item.id === this.target[0]) {
              objTep = item;
            }
          });
          objTep.chArr.map((item) => {
            if (item.value === this.target[1]) {
              divicedata = item;
            }
          });
          this.setFrom({
            id: this.target[1],
            type: divicedata.label.split("_")[0],
            parent_id: this.target[0],
            name: this.filterName(divicedata.label),
            time: new Date().valueOf(),
          });
          this.postDoFailureOver(
            {
              transType: 1,
              fromVmcId: this.target[0],
              isFault: 0,
              deviceId: this.target[1],
              appId: this.target[2],
            },
            0
          );
          this.romoteGetResult(
            {
              transType: 1,
              fromVmcId: this.target[0],
              isFault: 0,
              deviceId: this.target[1],
              appId: this.target[2],
            },
            divicedata.label.split("_")[0]
          );
        }
      }
      if (types === 1) {
        this.mcokLoading(2);
      }
      if (types === 2) {
        this.mcokLoading(3);
      }
    },
    dealWithOPs() {
      let tempObj = {};
      for (let i = 0; i < this.allTasks.length; i++) {
        let temps = this.allTasks[i];
        let arr = temps.config.url.split("=");
        tempObj[arr[1]] = temps.data.data.apps;
      }
      this.allNodes.map((nodes) => {
        nodes.apps = tempObj[nodes.id];
        let arrTempCh = [];
        nodes.child.map((ch) => {
          ch.other_info.map((info) => {
            if (
              info.key === "cpu_ids" ||
              info.key === "gpu_ids" ||
              info.key === "dsp_ids" ||
              info.key === "fpga_ids"
            ) {
              info.value.map((cpinfo) => {
                arrTempCh.push({
                  label: `${info.key.split("_")[0]}_${cpinfo}`,
                  value: Number(cpinfo),
                });
              });
            }
          });
        });
        nodes.chArr = arrTempCh;
      });
      this.allNodes.map((nodes) => {
        nodes.chArr.map((cpuinfo) => {
          let tempCpuinfoCharr = [];
          nodes.apps.map((app) => {
            if (app.device_id == cpuinfo.value) {
              tempCpuinfoCharr.push({
                label: this.filterName(app.app_name),
                value: app.id,
              });
            }
          });
          cpuinfo.children = tempCpuinfoCharr;
        });
      });
      this.options_task = this.allNodes.map((nodes) => {
        return {
          label: this.filterName(nodes.name),
          value: nodes.id,
          children: nodes.chArr,
        };
      });
      this.options = this.allNodes.map((nodes) => {
        return {
          label: this.filterName(nodes.name),
          value: nodes.id,
        };
      });
    },
    async getTopoData() {
      const { data } = await getTopoShow();
      let tempData = data.data.node || [];
      let vmctemp = [];
      tempData.map((item) => {
        if (item.device_type === "vmc") {
          vmctemp.push(item);
        }
      });
      let haschArr = vmctemp.map((vmcitem) => {
        let chArr = [];
        tempData.map((item) => {
          if (item.parent_id && item.parent_id === vmcitem.id) {
            chArr.push(item);
          }
        });
        vmcitem.child = chArr;
        return vmcitem;
      });
      this.allNodes = haschArr;
      this.getallAppInfo();
    },
    getallAppInfo() {
      let that = this;
      let pArr = [];
      this.allNodes.map((item, index) => {
        pArr.push(getAppInfo(item.id));
      });
      Promise.allSettled(pArr).then((results) => {
        results.forEach((result) => {
          that.allTasks.push(result.value);
        });
      });
      setTimeout(() => {
        that.dealWithOPs();
      }, 300);
    },
  },
  computed: {
    ...mapState(["disVmc", "disArea", "from", "to"]),
  },
  mounted() {},
  created() {
    this.getTopoData();
  },
};
</script>
<style lang="less">
.disasterrecovery {
  padding: 24px;
  box-sizing: border-box;
  width: 100%;
  height: 91.5vh;
  background: url("../assets/newpng/BACK2.png") no-repeat center;
  background-size: 100% 100%;
  color: aliceblue !important;
  .titleBox {
    width: 100%;
    text-align: center;
  }
  .typeBox {
    width: 100%;
    text-align: center;
  }
  .selectBox {
    width: 100%;
    text-align: center;
    margin-top: 2vh;
  }
  .endBox {
    width: 100%;
    text-align: center;
    margin-top: 2vh;
    height: 50vh;
    .infoBoard {
      position: relative;
      margin-left: 25%;
      width: 50%;
      height: 100%;
      background: url("../assets/newpng/tableborder_.png") no-repeat center;
      background-size: 100% 100%;
      padding: 7vh 7vw;
      box-sizing: border-box;
      .gif {
        width: 4vw;
      }
      .asideBox {
        position: absolute;
        width: 40%;
        height: 94%;
        right: 5%;
        top: 3%;
        .top {
          width: 100%;
          height: 50%;
          text-align: center;
          vertical-align: middle;
          font-size: 2rem;
          display: table;
          .cell {
            background: url("../assets/newpng/center_top_main_little_bgi.png")
              no-repeat center;
            display: table-cell;
            vertical-align: middle;
            width: 100%;
            height: 100%;
          }
          p {
            margin: 0;
            padding: 0;
          }
        }
      }
    }
  }
  .el-step__title {
    font-size: 1.5rem !important;
  }
  .el-step__description {
    font-size: 1rem !important;
  }
}
</style>