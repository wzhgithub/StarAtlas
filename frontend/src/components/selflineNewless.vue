<!--
 * @Author: jacob
 * @Date: 2020-12-10 11:02:20
 * @LastEditTime: 2020-12-17 15:29:51
 * @LastEditors: jacob
 * @Description: echarts通用组件
-->
<template>
  <div ref="echartsa"></div>
</template>

<script>
// import text from "@/assets/data/linejson.json";
import { getVMCDataSeq } from "../api";
export default {
  props: ["value", "color", "inref", "indexNow"],
  name: "selflineNewless",
  data() {
    return {
      // lineData: [],
      category: [],
      lineData: [],
      barData: [],
      allData: [],
      cpu: [],
      gpu: [],
      dps: [],
    };
  },
  methods: {
    async getRawData() {
      let res = await getVMCDataSeq(this.indexNow);
      if (res.status == 200 && res.data.code == 0) {
        this.allData = res.data.data;
        // console.log(this.allData);
      }
    },
    setDatanow() {
      // this.lineData = [];
      // this.category = [];
      // this.barData = [];
      // let dottedBase = +new Date();
      // for (let i = 0; i < 15; i++) {
      //   let date = new Date((dottedBase -= 3600 * 24 * 1000));
      //   this.category.push(
      //     [date.getFullYear(), date.getMonth() + 1, date.getDate()].join("-")
      //   );
      //   let b = Math.random() * 200;
      //   let d = Math.random() * 200;
      //   this.barData.push(b);
      //   this.lineData.push(d + b);
      // }
      this.cpu = [];
      this.gpu = [];
      this.dsp = [];
      this.category = [];
      for (var i = this.allData.length - 1; i >= 0; i--) {
          let date = new Date(this.allData[i].time);
          this.category.push(
            [date.getHours(), date.getMinutes(), date.getSeconds()].join(":")
          );
          this.cpu.push(this.allData[i].totalUsage);
          this.gpu.push(this.allData[i].gpuUsage);
          this.dsp.push(this.allData[i].dspUsage);
      }
    },
    drawEcharts() {
      // var echarts = require("@/lib/echarts.min.js");
      var echarts = require("echarts");
      // this.lineData = JSON.parse(JSON.stringify(text));
      let myChart = echarts.init(document.getElementById(this.inref));
      let option = {
        backgroundColor: "rgba(0,0,0,0)",
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "shadow",
          },
        },
        legend: {
          // data: ["峰值算力", "峰值带宽", "bar", "line_"],
          data: ["cpu利用率", "gpu利用率", "dsp利用率"],
          textStyle: {
            color: "#ccc",
          },
        },
        xAxis: {
          data: this.category,
          axisLine: {
            lineStyle: {
              color: "#ccc",
            },
          },
        },
        yAxis: {
          splitLine: { show: false },
          axisLine: {
            lineStyle: {
              color: "#ccc",
            },
          },
        },
        series: [
          // {
          //   name: "峰值算力",
          //   type: "line",
          //   smooth: true,
          //   showAllSymbol: true,
          //   symbol: "emptyCircle",
          //   symbolSize: 3,
          //   data: this.lineData,
          //   itemStyle: {
          //     color: "#72eedd",
          //   },
          // },
          // {
          //   name: "峰值带宽",
          //   type: "line",
          //   smooth: true,
          //   showAllSymbol: true,
          //   symbol: "emptyCircle",
          //   symbolSize: 3,
          //   data: this.barData,
          // },
          {
            name: "cpu利用率",
            type: "line",
            smooth: true,
            showAllSymbol: true,
            symbol: "emptyCircle",
            symbolSize: 3,
            data: this.cpu,
            itemStyle: {
              color: "#72eedd",
            },
          },
          {
            name: "gpu利用率",
            type: "line",
            smooth: true,
            showAllSymbol: true,
            symbol: "emptyCircle",
            symbolSize: 3,
            data: this.gpu,
          },
          {
            name: "dsp利用率",
            type: "line",
            smooth: true,
            showAllSymbol: true,
            symbol: "emptyCircle",
            symbolSize: 3,
            data: this.dsp,
          },
          // {
          //   name: "bar",
          //   type: "bar",
          //   barWidth: 5,
          //   itemStyle: {
          //     borderRadius: 2,
          //     color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          //       { offset: 0, color: "#14c8d4" },
          //       { offset: 1, color: "#43eec6" },
          //     ]),
          //   },
          //   data: this.barData,
          // },
          // {
          //   name: "line",
          //   type: "bar",
          //   // barGap: '-100%',
          //   barWidth: 5,
          //   itemStyle: {
          //     color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          //       { offset: 0, color: "rgba(20,200,212,0.5)" },
          //       { offset: 0.2, color: "rgba(20,200,212,0.2)" },
          //       { offset: 1, color: "rgba(20,200,212,0)" },
          //     ]),
          //   },
          //   z: -12,
          //   data: this.lineData,
          // },
          // {
          //   name: "line_",
          //   type: "bar",
          //   // barGap: '-100%',
          //   barWidth: 5,
          //   itemStyle: {
          //     borderRadius: 2,
          //     color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          //       { offset: 0, color: "rgba(255,10,22,0.5)" },
          //       { offset: 0.2, color: "rgba(200,20,12,0.2)" },
          //       { offset: 1, color: "rgba(255,10,22,0)" },
          //     ]),
          //   },
          //   data: this.barData,
          // },
          // {
          //   name: "dotted",
          //   type: "pictorialBar",
          //   symbol: "rect",
          //   itemStyle: {
          //     color: "#0f375f",
          //   },
          //   symbolRepeat: true,
          //   symbolSize: [5, 1],
          //   symbolMargin: 1,
          //   z: -10,
          //   data: this.lineData,
          // },
        ],
      };
      myChart.setOption(option);
    },
  },
  mounted() {
    const that = this;
    setTimeout(() => {
      //if (that.indexNow !== 1) {
        that.getRawData();
        that.setDatanow();
      //}
      that.drawEcharts();
    }, 500);
    setInterval(() => {
      that.getRawData();
      that.setDatanow();
      that.drawEcharts();
    }, 3000);
  },
};
</script>

<style lang='less' >
#chart_example {
  height: 100%;
  width: 100%;
}
</style>
