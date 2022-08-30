/* eslint-disable */
<template>
  <div class="disasterrecovery">
    <div class="titleBox">
      <h2>遥控指令触发</h2>
    </div>
    <div class="typeBox">
      指令类型：&nbsp;&nbsp;&nbsp;<el-radio-group v-model="type">
        <el-radio-button label="type1">容错模拟（整机故障）</el-radio-button>
        <el-radio-button disabled label="type2"
          >容错模拟（分区故障）</el-radio-button
        >
        <el-radio-button disabled label="type3">性能测试</el-radio-button>
      </el-radio-group>
    </div>
    <div class="selectBox">
      <p v-if="type !== 'type3'">
        操作目标：&nbsp;&nbsp;&nbsp;
        <el-cascader
          v-model="target"
          :options="options"
          :props="{ checkStrictly: true }"
          clearable
          style="width: 33%; margin-top: 1vh"
        ></el-cascader>
      </p>
      <p v-else>
        选择算法：&nbsp;&nbsp;&nbsp;
        <el-cascader
          v-model="cvalue"
          :options="options_a"
          :props="{ checkStrictly: true }"
          clearable
          style="width: 33%; margin-top: 1vh"
        ></el-cascader>
      </p>
    </div>
    <div class="selectBox">
      <el-button style="width: 10%" type="success" @click="handlerDo" round
        >执行</el-button
      >
      <el-button style="width: 10%" type="info" round>复位</el-button>
    </div>
    <div class="endBox">
      <div class="infoBoard">
        <el-steps direction="vertical" :active="stepNow">
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
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
// import "@/lib/loaders.min.css";
import { mapState, mapMutations } from "vuex";
import { getTopoShow, filterName, getAppInfo } from "@/api";
export default {
  name: "RemoteControl",
  components: {
    // HelloWorld,
  },
  data() {
    return {
      type: "type1",
      target: "123",
      cvalue: "",
      options: [],
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
      description0: "",
      description1: "",
      description2: "",
      allNodes: [],
      allTasks: [],
    };
  },
  computed: {
    ...mapState(["disVmc", "disArea"]),
  },
  methods: {
    ...mapMutations(["setFrom", "setTo"]),
    filterName,
    mcokLoading() {
      let that = this;
      setTimeout(() => {
        that.stepNow = 1;
        that.description0 = "客户端正在发起建立连接.。。。。";
        that.description1 = "未执行到当前步骤";
        that.description2 = "未执行到当前步骤";
      }, 1000);
      setTimeout(() => {
        that.stepNow = 1;
        that.description0 = "连接建立中，等待远端响应。。。";
      }, 3000);
      setTimeout(() => {
        that.stepNow = 1;
        that.description0 = "客户端建立连接成功！";
      }, 5000);
      setTimeout(() => {
        that.stepNow = 1;
        that.description0 = "客户端正在进行数据交互。。。。";
      }, 7000);
      setTimeout(() => {
        that.stepNow = 1;
        that.description0 = "客户端数据传输完成。。。。";
      }, 9000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description0 = "连接建立成功";
        that.description1 = "平台正在进行数据计算。。。。";
      }, 10000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（10%）";
      }, 10000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（15%）";
      }, 12000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（25%）";
      }, 15000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（35%）";
      }, 19000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（55%）";
      }, 20000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（75%），";
      }, 25000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（95%），";
      }, 27000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度（100%），";
      }, 30000);
      setTimeout(() => {
        that.stepNow = 2;
        that.description1 = "平台正在进行数据计算进度完成，数据集开始处理";
      }, 31000);
      setTimeout(() => {
        that.stepNow = 3;
        that.description1 = "平台正在进行数据计算进度完成";
        that.description2 = "数据集处理中";
      }, 32000);
      setTimeout(() => {
        that.stepNow = 3;
        that.description2 = "数据集格式比对中";
      }, 34000);
      setTimeout(() => {
        that.stepNow = 3;
        that.description2 = "数据集校验中";
      }, 36000);
      setTimeout(() => {
        that.stepNow = 3;
        that.description2 = "数据集处理完成，返回数据中";
      }, 38000);
      setTimeout(() => {
        that.stepNow = 3;
        that.description2 = "安全断开链接中";
      }, 40000);
      setTimeout(() => {
        that.stepNow = 4;
        that.description2 = "数据中返回成功，目标迁移设备为“计算节点——测试”";
      }, 42000);
      setTimeout(() => {
        this.$message({
          message:
            "迁移目标计算成功，已为您重定向到容灾演示页面查看任务迁移详情",
          type: "success",
        });
        this.$router.push("/disasterrecovery");
      }, 45000);
    },
    handlerDo() {
      if (this.type === "type1" && this.target) {
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
        });
        // this.setTo({
        //   id: this.target[1],
        //   type: "vmc",
        //   parent_id: null,
        //   name: "测试",
        // });
        this.mcokLoading();
      }
      if (this.type === "type2" && this.target) {
        let ids = this.target[1].split("_");
        let objTep = {};
        let appdata = {};
        this.allNodes.map((item) => {
          if (item.id === this.target[0]) {
            objTep = item;
          }
        });
        objTep.apps.map((item) => {
          if (item.id === ids[1]) {
            appdata = item;
          }
        });
        this.setFrom({
          id: appdata.id,
          type: "cpu",
          parent_id: this.target[0],
          name: this.filterName(appdata.app_name),
        });
        this.mcokLoading();
      }
      if (this.type === "type3" && this.cvalue) {
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
      });
      this.allNodes.map((nodes) => {
        let tee = nodes.apps.map((app) => {
          return {
            label: this.filterName(app.app_name),
            value: `${nodes.id}_${app.id}`,
          };
        });
        nodes.children = tee;
      });
      this.options = this.allNodes.map((nodes) => {
        return {
          label: this.filterName(nodes.name),
          value: nodes.id,
          children: nodes.children,
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
  mounted() {},
  created() {
    this.getTopoData();
  },
};
</script>
<style lang="less" scoped>
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
    }
  }
}
</style>