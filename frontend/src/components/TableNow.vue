<template>
  <div style="width: 100%; height: 100%">
    <el-table class="tableList" :data="reallData" style="width: 100%">
      <template v-for="(item, index, nub) in reallData[0]">
        <el-table-column
          width="180"
          :key="index"
          :label="`分区${nub + 1}`"
          align="center"
          type="index"
          :prop="item"
        >
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <p>任务详情</p>
              <p>任务名: {{ scope.row[index].name }}</p>
              <p>所在分区: {{ scope.row[index].zone }}</p>
              <p>任务类型: {{ getType(scope.row[index].type) }}</p>
              <p>任务状态: {{ getStatus(scope.row[index].status) }}</p>
              <div slot="reference" class="name-wrapper">
                {{ scope.row[index].name }}
                <img
                  v-if="
                    scope.row[index].type === 'manage' &&
                    scope.row[index].status === 'running'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/running.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'manage' &&
                    scope.row[index].status === 'block'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/block.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'manage' &&
                    scope.row[index].status === 'ready'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/ready.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'manage' &&
                    scope.row[index].status === 'sleep'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/sleep.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'manage' &&
                    scope.row[index].status === 'uncreated'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/manage/uncreated.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'calculation' &&
                    scope.row[index].status === 'running'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/running.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'calculation' &&
                    scope.row[index].status === 'block'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/block.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'calculation' &&
                    scope.row[index].status === 'ready'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/ready.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'calculation' &&
                    scope.row[index].status === 'sleep'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/sleep.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'calculation' &&
                    scope.row[index].status === 'uncreated'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/calculation/uncreated.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'control' &&
                    scope.row[index].status === 'running'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/running.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'control' &&
                    scope.row[index].status === 'block'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/block.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'control' &&
                    scope.row[index].status === 'ready'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/ready.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'control' &&
                    scope.row[index].status === 'sleep'
                  "
                  style="height: 1rem; vertical-align: middle"
                  src="../assets/othericon/control/sleep.svg"
                  alt=""
                />
                <img
                  v-if="
                    scope.row[index].type === 'control' &&
                    scope.row[index].status === 'uncreated'
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
export default {
  name: "TableNow",
  props: {},
  data() {
    return {
      initMt: 0,
      // getPlayData:[],
      visible: true,
      stop: false,
    };
  },
  methods: {
    getType(typeName) {
      if (typeName === "calculation") {
        return "计算任务";
      }
      if (typeName === "control") {
        return "控制任务";
      }
      if (typeName === "manage") {
        return "管理任务";
      }
      return "未知类型";
    },
    getStatus(typestatus) {
      if (typestatus === "block") {
        return "阻塞";
      }
      if (typestatus === "ready") {
        return "就绪";
      }
      if (typestatus === "running") {
        return "运行中";
      }
      if (typestatus === "sleep") {
        return "睡眠";
      }
      if (typestatus === "uncreated") {
        return "未创建";
      }
      return "未知状态";
    },
  },
  watch: {},
  computed: {
    getPlayData() {},
    reallData() {
      return [
        {
          z_1: {
            name: "任务7684",
            zone: "分区一",
            type: "control",
            status: "block",
          },
          z_2: {
            name: "任务7634",
            zone: "分区二",
            type: "manage",
            status: "ready",
          },
          z_3: {
            name: "任务76842",
            zone: "分区三",
            type: "manage",
            status: "running",
          },
          z_4: {
            name: "任务768455",
            zone: "分区四",
            type: "manage",
            status: "sleep",
          },
          z_5: {
            name: "任务76823",
            zone: "分区五",
            type: "control",
            status: "uncreated",
          },
        },
        {
          z_1: {
            name: "任务7685",
            zone: "分区一",
            type: "manage",
            status: "block",
          },
          z_2: {
            name: "任务7634",
            zone: "分区二",
            type: "control",
            status: "ready",
          },
          z_3: {
            name: "任务76842",
            zone: "分区三",
            type: "manage",
            status: "running",
          },
          z_4: {
            name: "任务768455",
            zone: "分区四",
            type: "manage",
            status: "sleep",
          },
          z_5: {
            name: "任务76823",
            zone: "分区五",
            type: "manage",
            status: "uncreated",
          },
        },
        {
          z_1: {
            name: "任务7687",
            zone: "分区一",
            type: "manage",
            status: "block",
          },
          z_2: {
            name: "任务7634",
            zone: "分区二",
            type: "manage",
            status: "ready",
          },
          z_3: {
            name: "任务76842",
            zone: "分区三",
            type: "manage",
            status: "running",
          },
          z_4: {
            name: "任务768455",
            zone: "分区四",
            type: "manage",
            status: "sleep",
          },
          z_5: {
            name: "任务76823",
            zone: "分区五",
            type: "manage",
            status: "uncreated",
          },
        },
        {
          z_1: {
            name: "任务7684",
            zone: "分区一",
            type: "manage",
            status: "block",
          },
          z_2: {
            name: "任务7634",
            zone: "分区二",
            type: "manage",
            status: "ready",
          },
          z_3: {
            name: "任务76842",
            zone: "分区三",
            type: "manage",
            status: "running",
          },
          z_4: {
            name: "任务768455",
            zone: "分区四",
            type: "manage",
            status: "sleep",
          },
          z_5: {
            name: "任务76823",
            zone: "分区五",
            type: "manage",
            status: "uncreated",
          },
        },
      ];
    },
  },
  mounted() {
    // this.play();
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
