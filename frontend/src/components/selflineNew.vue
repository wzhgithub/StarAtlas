<template>
  <div ref="echartsa"></div>
</template>

<script>
// import text from "@/assets/data/linejson.json";
import { getVMCDataSeq } from "../api";
export default {
  props: ["value", "color", "inref", "vmcid"],
  name: "selflineNew",
  data() {
    return {
      category: [],
      lineData: [],
      barData: [],
      allData: [],
      memory: [],
      disk: [],
      cpu: [],
      gpu: [],
      dsp: [],
    };
  },
  methods: {
    async getRawData() {
      // console.log(this.vmcid);
      let res = await getVMCDataSeq(this.vmcid);
      if (res.status == 200 && res.data.code == 0) {
        this.allData = res.data.data;
        // console.log(this.allData);
      }
    },
    setDatanow() {
      this.lineData = [];
      this.category = [];
      this.barData = [];
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
      this.memory = [];
      this.disk = [];
      this.gpu = [];
      this.dsp = [];
      for (var i = this.allData.length - 1; i >= 0; i--) {
          //console.log(this.allData[i].time);
          let date = new Date(this.allData[i].time);
          this.category.push(
            // [date.getFullYear(), date.getMonth() + 1, date.getDate(), date.getHours()].join("-")
            [date.getHours(), date.getMinutes(), date.getSeconds()].join(":")
          );
          // let b = Math.random() * 200;
          // let d = Math.random() * 200;
          // this.barData.push(b);
          // this.lineData.push(d + b);
          this.cpu.push(this.allData[i].totalUsage);
          this.memory.push(this.allData[i].memoryUsage);
          this.disk.push(this.allData[i].diskUsage);
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
          data: ["cpu", "memory", "disk", "gpu", "dsp"],
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
          {
            name: "cpu",
            type: "line",
            smooth: true,
            showAllSymbol: true,
            symbol: "emptyCircle",
            symbolSize: 3,
            // data: this.lineData,
            data: this.cpu,
            itemStyle: {
              color: "#72eedd",
            },
          },
          {
            name: "memory",
            type: "line",
            smooth: true,
            showAllSymbol: true,
            symbol: "emptyCircle",
            symbolSize: 3,
            // data: this.barData,
            data: this.memory,
          },
          {
            name: "disk",
            type: "line",
            smooth: true,
            showAllSymbol: true,
            symbol: "emptyCircle",
            symbolSize: 3,
            // data: this.barData,
            data: this.disk,
          },
          {
            name: "gpu",
            type: "bar",
            barWidth: 5,
            itemStyle: {
              borderRadius: 2,
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                { offset: 0, color: "#14c8d4" },
                { offset: 1, color: "#43eec6" },
              ]),
            },
            //data: this.barData,
            data: this.gpu,
          },
          {
            name: "dsp",
            type: "bar",
            // barGap: '-100%',
            barWidth: 5,
            itemStyle: {
              // color: "rgba(255,0,0,1)",
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                { offset: 0, color: "rgba(20,200,212,0.5)" },
                { offset: 0.2, color: "rgba(20,200,212,0.2)" },
                { offset: 1, color: "rgba(20,200,212,0)" },
                // { offset: 0, color: "rgba(225, 113, 97, 1)" },
                // { offset: 1, color: "rgba(255,0,0, 1)" },
                // { offset: 0, color: "#b13722" },
                // { offset: 1, color: "#e17261" },
              ]),
            },
            z: -12,
            data: this.dsp,
          },
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
          //   //name: "dotted",
          //   name: "line",
          //   type: "pictorialBar",
          //   symbol: "rect",
          //   itemStyle: {
          //     //color: "#0f375f",
          //     color: "rgba(255,0,0,1)",
          //   },
          //   symbolRepeat: true,
          //   symbolSize: [5, 5],
          //   symbolOffset: [0,0],
          //   symbolMargin: 1,
          //   z: -10,
          //   //data: this.lineData,
          //   data: this.cpuLine,
          // },
        ],
      };
      myChart.setOption(option);
    },
  },
  mounted() {
    const that = this;
    setTimeout(() => {
      that.getRawData();
      that.setDatanow();
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
