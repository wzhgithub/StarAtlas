<template>
  <div ref="echartsa"></div>
</template>

<script>
import text from "@/assets/data/linejson.json";
export default {
  props: ["value", "color", "inref"],
  name: "selfline",
  data() {
    return {
      lineData: [],
    };
  },
  methods: {
    drawEcharts() {
      // var echarts = require("@/lib/echarts.min.js");
      var echarts = require("echarts");
      this.lineData = JSON.parse(JSON.stringify(text));
      let myChart = echarts.init(document.getElementById(this.inref));
      let option = {
        title: {
          text: "历史峰值变动",
          left: "1%",
          textStyle: {
            color: "#21ebff",
          },
        },
        tooltip: {
          trigger: "axis",
        },
        grid: {
          left: "5%",
          right: "25%",
          bottom: "10%",
        },
        xAxis: {
          data: this.lineData.map(function (item) {
            return item[0];
          }),
        },
        yAxis: {},
        toolbox: {
          right: 10,
          feature: {
            dataZoom: {
              yAxisIndex: "none",
            },
            restore: {},
            saveAsImage: {},
          },
        },
        dataZoom: [
          {
            startValue: "2014-06-01",
          },
          {
            type: "inside",
          },
        ],
        visualMap: {
          top: 50,
          right: 10,
          pieces: [
            {
              gt: 0,
              lte: 50,
              color: "#93CE07",
            },
            {
              gt: 50,
              lte: 100,
              color: "#FBDB0F",
            },
            {
              gt: 100,
              lte: 150,
              color: "#FC7D02",
            },
            {
              gt: 150,
              lte: 200,
              color: "#FD0100",
            },
            {
              gt: 200,
              lte: 300,
              color: "#AA069F",
            },
            {
              gt: 300,
              color: "#AC3B2A",
            },
          ],
          outOfRange: {
            color: "#999",
          },
        },
        series: {
          name: "历史峰值变动",
          type: "line",
          data: this.lineData.map(function (item) {
            return item[1];
          }),
          markLine: {
            silent: true,
            lineStyle: {
              color: "#333",
            },
            data: [
              {
                yAxis: 50,
              },
              {
                yAxis: 100,
              },
              {
                yAxis: 150,
              },
              {
                yAxis: 200,
              },
              {
                yAxis: 300,
              },
            ],
          },
        },
      };
      myChart.setOption(option);
    },
  },
  mounted() {
    const that = this;
    setTimeout(() => {
      that.drawEcharts();
    }, 500);
  },
};
</script>

<style lang='less' >
#chart_example {
  height: 100%;
  width: 100%;
}
</style>
