<template>
  <div style="width: 100%; height: 100%">
    <el-table class="tableList" :data="reallData" style="width: 100%">
      <template v-for="(item, index, nub) in reallData[0]">
        <el-table-column
          width="245"
          :key="index"
          :label="`分区${item.zone || nub + 1}`"
          align="center"
          type="index"
          :prop="item"
        >
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <p>任务详情</p>
              <p>任务名: {{ scope.row[index].name }}</p>
              <p>所在分区: {{ scope.row[index].zone }}</p>
              <p>任务类型: {{ getType(scope.row[index].task_type) }}</p>
              <p>任务状态: {{ getStatus(scope.row[index].task_status) }}</p>
              <div slot="reference" class="name-wrapper">
                {{ scope.row[index].name }}
                <img
                  v-if="
                    scope.row[index].task_type === 2 &&
                    scope.row[index].task_status === 1
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/running.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 2 &&
                    scope.row[index].task_status === 3
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/block.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 2 &&
                    scope.row[index].task_status === 0
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/ready.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 2 &&
                    scope.row[index].task_status === 2
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/sleep.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 2 &&
                    scope.row[index].task_status === 255
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/uncreated.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 0 &&
                    scope.row[index].task_status === 1
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/running.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 0 &&
                    scope.row[index].task_status === 3
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/block.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 0 &&
                    scope.row[index].task_status === 0
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/ready.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 0 &&
                    scope.row[index].task_status === 2
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/sleep.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 0 &&
                    scope.row[index].task_status === 255
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/uncreated.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 1 &&
                    scope.row[index].task_status === 1
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/running.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 1 &&
                    scope.row[index].task_status === 3
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/block.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 1 &&
                    scope.row[index].task_status === 0
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/ready.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 1 &&
                    scope.row[index].task_status === 2
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/sleep.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].task_type === 1 &&
                    scope.row[index].task_status === 255
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/uncreated.svg"
                  alt=""
                />
              </div>
            </el-popover>
          </template>
        </el-table-column>
      </template>
    </el-table>
  </div>
</template>
<script>
import { getAppInfo } from "../api";
export default {
  name: "TableNow",
  props: ["indexNow"],
  data() {
    return {
      initMt: 0,
      // getPlayData:[],
      visible: true,
      stop: false,
      rawData: [],
    };
  },
  methods: {
    getType(typeName) {
      if (typeName === 0) {
        return "计算任务";
      }
      if (typeName === 1) {
        return "控制任务";
      }
      if (typeName === 2) {
        return "管理任务";
      }
      return "未知类型";
    },
    getStatus(typestatus) {
      if (typestatus === 3) {
        return "阻塞";
      }
      if (typestatus === 0) {
        return "就绪";
      }
      if (typestatus === 1) {
        return "运行中";
      }
      if (typestatus === 2) {
        return "睡眠";
      }
      if (typestatus === 255) {
        return "未创建";
      }
      return "未知状态";
    },
    async getAppInfoData() {
      let res = await getAppInfo(this.indexNow);
      this.rawData = res.data.data.apps || [];
    },
  },
  watch: {},
  computed: {
    getPlayData() {},
    reallData() {
      let endarr = [];
      let maxAppLength = this.rawData.length;
      let maxTaskLength = 0;
      this.rawData.map((item) => {
        if (item.task_set.length > maxTaskLength) {
          maxTaskLength = item.task_set.length;
        }
      });
      for (let i = 0; i < maxTaskLength; i++) {
        let tempDataObj = {};
        for (let j = 0; j < maxAppLength; j++) {
          if (
            this.rawData[j] &&
            this.rawData[j].task_set &&
            this.rawData[j].task_set[i]
          ) {
            tempDataObj[`z_${j + 1}`] = {
              ...this.rawData[j].task_set[i],
              zone: this.rawData[j].app_name,
              app_status: this.rawData[j].app_status,
              belongs_to: this.rawData[j].belongs_to,
              app_id: this.rawData[j].id,
            };
          } else {
            tempDataObj[`z_${j + 1}`] = {
              id: null,
              name: null,
            };
          }
        }
        endarr.push(tempDataObj);
      }
      return endarr;
    },
  },
  mounted() {
    // this.play();
    this.getAppInfoData();
  },
};
</script>
<style lang="less" scoped>
/*最外层透明*/
/deep/ .el-table,
/deep/ .el-table__expanded-cell {
  background-color: transparent;
  color: aqua;
}
/* 表格内背景颜色 */
/deep/ .el-table th,
/deep/ .el-table tr,
/deep/ .el-table td {
  color: aqua;
  background-color: transparent;
}
/deep/ .el-table th.el-table__cell {
  background-color: transparent;
  color: aqua;
}
// .el-table--enable-row-hover .el-table__body tr:hover > td {
//   background-color: rgba(0, 0, 0, 0) !important;
//   color: #000;
// }

/deep/ .el-table--enable-row-hover {
  .el-table__body tr:hover > td {
    background: transparent !important;
  }
}
</style>
