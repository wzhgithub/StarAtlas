<template>
  <div style="cursor: default; margin: 20px 10px 18px">
    <div class="table-header table-row">
      <div class="table-cell" style="width: 25%">任务名称</div>
      <div class="table-cell" style="width: 30%">任务状态</div>
      <div class="table-cell" style="width: 15%">执行人</div>
      <div class="table-cell" style="width: 30%; text-align: right">
        所占资源(GB)
      </div>
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
          <div class="table-cell" style="width: 25%" :title="item.productName">
            {{ item.productName }}
          </div>
          <div class="table-cell" style="width: 30%" :title="item.coreName">
            {{ item.coreName }}
          </div>
          <div class="table-cell" style="width: 15%" :title="item.publish">
            {{ item.publish }}
          </div>
          <div
            class="table-cell"
            style="width: 30%; text-align: right"
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
  line-height: 45px;
  height: 45px;
  transition: all 0.3s;
  border-bottom: 1px solid #21ebff;
}
.table-header {
  color: #21ebff;
}
.table-cell {
  text-align: center;
  font-size: 20px;
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
  height: 240px;
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
    top: -8 * 35px;
  }
}
</style>
