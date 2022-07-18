<template>
  <div style="cursor: default; margin: 1.5rem 0.8rem 1.3rem">
    <div class="table-header table-row">
      <div class="table-cell" style="width: 25%">分区一</div>
      <div class="table-cell" style="width: 25%">分区二</div>
      <div class="table-cell" style="width: 25%">分区三</div>
      <div class="table-cell" style="width: 25%">分区四</div>
    </div>
    <div class="table-body">
      <div :class="{ 'scroll-wrap': getPlayData.length > 0 }">
        <div
          class="table-row"
          :class="{ hasBgc: index % 2 === 0 }"
          v-for="(item, index) in getPlayData"
          :key="index"
          :ref="'row_' + index"
        >
          <el-tooltip
            class="item"
            effect="dark"
            content=""
            placement="top-start"
          >
            <div slot="content">
              <h3>任务详情</h3>
              <p>任务名：{{ item.productName }}</p>
              <p>所在分区：分区一</p>
            </div>
            <div
              class="table-cell"
              style="width: 25%"
              :title="item.productName"
            >
              {{ item.productName }}
            </div>
          </el-tooltip>
          <div class="table-cell" style="width: 25%" :title="item.coreName">
            {{ item.coreName }}
          </div>
          <div class="table-cell" style="width: 25%" :title="item.publish">
            {{ item.publish }}
          </div>
          <div
            class="table-cell"
            style="width: 25%"
            :title="item.publishAmount"
          >
            {{ item.publishAmount }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "scrolltable",
  props: {
    data: {
      type: Array,
      default: () => {
        return [];
      },
    },
  },
  data() {
    return {
      initMt: 0,
      // getPlayData:[],
      visible: true,
      stop: false,
    };
  },
  methods: {
    play() {
      // const row = this.$refs["row_0"][0];

      setTimeout(() => {
        this.visible = false;

        this.$nextTick(() => {
          this.initMt++;
          if (this.initMt === this.data.length) {
            this.initMt = 0;
          }
          this.visible = true;
        });
        this.play();
      }, 2000);
    },
  },
  watch: {},
  computed: {
    getPlayData() {
      return this.data.concat(this.data.slice(0, 4));
    },
  },
  mounted() {
    // this.play();
  },
};
</script>
<style lang="less" scoped>
.table-row {
  display: flex;
  line-height: 3.5rem;
  height: 3.5rem;
  transition: all 0.3s;
  border-bottom: 0.1rem solid #21ebff;
}
.table-header {
  color: #21ebff;
}
.table-cell {
  text-align: center;
  font-size: 1.4rem;
  text-overflow: ellipsis;
  overflow: hidden;
}
// .hasBgc {
//   background: rgb(0, 59, 81);
// }
.hidden-row {
  height: 0 !important;
  line-height: 0 !important;
  display: none !important;
}
.table-body {
  height: 17.5rem;
  overflow-y: hidden;
  .table-row {
    color: #fff;
  }
}
.scroll-wrap {
  animation: scroll 18s linear infinite;
  position: relative;
}
.scroll-wrap:hover {
  animation-play-state: paused;
}
@keyframes scroll {
  from {
    top: 0;
  }
  to {
    top: -8 * 2.5rem;
  }
}
</style>
