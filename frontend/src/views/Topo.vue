/* eslint-disable */
<template>
  <div class="topobox">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <!-- 123 -->
    <div v-if="loading" class="loadingbox">
      <div data-loader="jumping"></div>
    </div>
    <div
      v-else
      ref="canvas"
      style="background-color: rgba(0, 0, 0, 0); width: 100%; height: 85vh"
    ></div>
    <el-drawer
      class="drawer_info"
      size="40%"
      :title="`设备：${filterName(activeNodeInfo.name)}的详情`"
      :visible.sync="drawer"
      :modal="false"
    >
      <div v-if="activeNodeInfo.device_type === 'vmc'" class="mainbox">
        <div class="aside_box">
          <div class="grid-content_btn">
            <button @click="mockError(1)">模拟整机故障</button>
          </div>
          <div class="grid-content_btn">
            <button @click="mockError(0)">模拟分区故障</button>
          </div>
          <img
            v-if="activeNodeInfo.device_type === 'vmc'"
            class="miansvg"
            src="../assets/newpng/VMC.svg"
            alt=""
          />
        </div>
        <div class="aside_box_line_bar">
          <p class="title">
            <span>机器性能历史状况</span>
          </p>
          <div class="canvasbox" id="linebox_">
            <selflineNew inref="linebox_" :vmcid="activeNodeInfo.id" />
          </div>
        </div>
        <div class="aside_box_task">
          <p class="title">
            <span>分时分区任务</span>
          </p>
          <scrolltable :data="tableData" />
        </div>
      </div>
      <div v-show="activeNodeInfo.device_type !== 'vmc'" class="mainbox">
        <div class="aside_box">
          <img
            v-if="activeNodeInfo.device_type === 'sw'"
            class="miansvg"
            src="../assets/newpng/sw.svg"
            alt=""
          />
          <img
            v-if="activeNodeInfo.device_type === 'cpu'"
            class="miansvg"
            src="../assets/newpng/CPU.svg"
            alt=""
          />
          <img
            v-if="activeNodeInfo.device_type === 'gpu'"
            class="miansvg"
            src="../assets/newpng/GPU.svg"
            alt=""
          />
          <img
            v-if="activeNodeInfo.device_type === 'dsp'"
            class="miansvg"
            src="../assets/newpng/DSP.svg"
            alt=""
          />
          <img
            v-if="activeNodeInfo.device_type === 'fpga'"
            class="miansvg"
            src="../assets/newpng/FPGA.svg"
            alt=""
          />
          <img
            v-if="activeNodeInfo.device_type === 'vmc'"
            class="miansvg"
            src="../assets/newpng/VMC.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'rtu_0'"
            class="miansvg"
            src="../assets/newpng/op_1.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'rtu_1'"
            class="miansvg"
            src="../assets/newpng/op_2.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'rtu_2'"
            class="miansvg"
            src="../assets/newpng/op_3.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'rtu_3'"
            class="miansvg"
            src="../assets/newpng/op_4.svg"
            alt=""
          />
        </div>
        <div class="aside_box_line_bar">
          <p class="title">
            <span>机器信息</span>
          </p>
          <div class="infoBox">
            <p>设备名：{{ filterName(activeNodeInfo.name) }}</p>
            <p>设备ID：{{ activeNodeInfo.id }}</p>
            <p>设备状态：{{ activeNodeInfo.device_status }}</p>
            <p>设备数量：{{ activeNodeInfo.device_num }}</p>
          </div>
        </div>
      </div>
    </el-drawer>
    <el-drawer
      class="drawer_info_left"
      size="40%"
      :title="`设备：${filterName(activeNodeInfo.name)}的详情`"
      :visible.sync="drawerText"
      :modal="false"
      direction="ltr"
    >
    </el-drawer>
    <el-dialog
      title="容灾迁移计算中..."
      :visible.sync="dialogVisible"
      width="50%"
      :before-close="handleClose"
    >
      <div style="height: 300px; padding: 24px">
        <el-steps direction="vertical" :active="1">
          <el-step title="步骤 1"></el-step>
          <el-step title="步骤 2"></el-step>
          <el-step
            title="步骤 3"
            description="这是一段很长很长很长的描述性文字"
          ></el-step>
        </el-steps>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import "@/lib/loaders.min.css";
import exchange from "@/assets/network/exchange.png";
// import server from "@/assets/network/server.png";
import server from "@/assets/network/server_a.png";
import exchange2 from "@/assets/network/exchange2.png";
// import firewall from "@/assets/network/firewall.png";
import firewall from "@/assets/network/firewall_a.png";
// import router from "@/assets/network/router.png";
import router from "@/assets/network/router_a.png";
// import pc from "@/assets/network/pc.png";
import pc from "@/assets/network/pc_a.png";
import flow from "@/assets/flow.png";
import cloud from "@/assets/network/cloud.png";
import { getTopoShow, filterName } from "@/api";
// import text from "../assets/data/topo.json";
import imgcpu from "@/assets/newpng/CPU.svg";
import imggpu from "@/assets/newpng/GPU.svg";
import imgdsp from "@/assets/newpng/DSP.svg";
import imgfgpa from "@/assets/newpng/FPGA.svg";
import imgsw from "@/assets/newpng/centersw_topo.svg";
import imgop1 from "@/assets/newpng/op_1.svg";
import imgop2 from "@/assets/newpng/op_2.svg";
import imgop3 from "@/assets/newpng/op_3.svg";
import imgop4 from "@/assets/newpng/op_4.svg";
import pointsvg from "@/assets/newpng/point_new.svg";
import inswsvg from "@/assets/newpng/sw_in.svg";
// import obc from "@/assets/newpng/cloud.png";
import imgvmc from "@/assets/newpng/VMC.svg";
// import sw from "@/assets/newpng/cloud.png";
import CpuInfo from "@/components/CpuInfo.vue";
import selflineNew from "@/components/selflineNew.vue";
import scrolltable from "@/components/scrollTable.vue";
import selfline from "@/components/selfline.vue";
export default {
  name: "Topo",
  components: {
    // HelloWorld,
    CpuInfo,
    selflineNew,
    scrolltable,
    selfline,
  },
  data() {
    return {
      loading: true,
      oneword: "",
      names: "aa",
      flowColor: "#F00",
      flowColor_: "#ffffff",
      flowColor_vpn: "#00FF00",
      drawer: false,
      drawerText: false,
      dialogVisible: true,
      devicename: "",
      deviceType: "",
      topoData: [],
      tableData: [
        {
          productName: "任务1-1",
          coreName: "任务2-1",
          publish: "任务3-1",
          publishAmount: "任务4-1",
        },
        {
          productName: "任务2",
          coreName: "任务2",
          publish: "任务2",
          publishAmount: "任务2",
        },
        {
          productName: "任务3",
          coreName: "任务3",
          publish: "任务3",
          publishAmount: "任务3",
        },
        {
          productName: "任务4",
          coreName: "任务4",
          publish: "任务4",
          publishAmount: "任务4",
        },
        {
          productName: "任务5",
          coreName: "任务5",
          publish: "任务5",
          publishAmount: "任务5",
        },
        {
          productName: "任务6",
          coreName: "任务6",
          publish: "任务6",
          publishAmount: "任务6",
        },
      ],
      allSwXY: [],
      imgObg: {
        vmc: imgvmc,
        sw: inswsvg,
        sw_c: imgsw,
        cpu: imgcpu,
        gpu: imggpu,
        dsp: imgdsp,
        fpga: imgfgpa,
        rtu_0: imgop1,
        rtu_1: imgop2,
        rtu_2: imgop3,
        rtu_3: imgop4,
      },
      activeNodeInfo: {
        name: "",
      },
    };
  },
  methods: {
    filterName,
    ...mapMutations(["setDisVmc", "setDisArea"]),
    async getTopoData() {
      // const str = {
      //   code: 0,
      //   data: {
      //     created_at: "2022-08-23T10:11:25.066Z",
      //     updated_at: "2022-08-23T10:11:40Z",
      //     id: "topo_table",
      //     node: [
      //       {
      //         id: 1000000,
      //         name: "cpu_all",
      //         device_type: "cpu",
      //         parent_id: 0,
      //         upstream_id: 0,
      //         device_status: "ERROR",
      //         device_num: 2,
      //         other_info: [
      //           { key: "cpu_ids", value: ["0", "1"] },
      //           { key: "cpu_names", value: ["CPU0      ", "CPU1      "] },
      //           { key: "cpu_types", value: ["0", "0"] },
      //           { key: "cpu_cores", value: ["1", "1"] },
      //         ],
      //       },
      //       {
      //         id: 2000000,
      //         name: "gpu_all",
      //         device_type: "gpu",
      //         parent_id: 0,
      //         upstream_id: 0,
      //         device_status: "ERROR",
      //         device_num: 2,
      //         other_info: [
      //           { key: "gpu_ids", value: ["0", "1"] },
      //           { key: "gpu_names", value: ["GPU0      ", "GPU1      "] },
      //           { key: "gpu_types", value: ["0", "0"] },
      //           { key: "gpu_cores", value: ["8", "8"] },
      //         ],
      //       },
      //       {
      //         id: 3000000,
      //         name: "dsp_all",
      //         device_type: "dsp",
      //         parent_id: 0,
      //         upstream_id: 0,
      //         device_status: "ERROR",
      //         device_num: 2,
      //         other_info: [
      //           { key: "dsp_ids", value: ["0", "1"] },
      //           { key: "dsp_names", value: ["DSP0      ", "DSP1      "] },
      //           { key: "dsp_types", value: ["0", "0"] },
      //           { key: "dsp_cores", value: ["8", "8"] },
      //         ],
      //       },
      //       {
      //         id: 4000000,
      //         name: "fpga_all",
      //         device_type: "fpga",
      //         parent_id: 0,
      //         upstream_id: 0,
      //         device_status: "ERROR",
      //         device_num: 2,
      //         other_info: [
      //           { key: "fpga_ids", value: ["0", "1"] },
      //           { key: "fpga_names", value: ["FPGA0     ", "FPGA1     "] },
      //           { key: "fpga_types", value: ["0", "0"] },
      //           { key: "fpga_cores", value: ["0", "0"] },
      //         ],
      //       },
      //       {
      //         id: 0,
      //         name: "VMC0      ",
      //         device_type: "vmc",
      //         parent_id: 0,
      //         upstream_id: 50,
      //         device_status: "ERROR",
      //         device_num: 0,
      //         other_info: [{ key: "proto_type", value: ["85"] }],
      //       },
      //     ],
      //     transfer_info: null,
      //   },
      //   msg: "success",
      // };
      // const { data } = JSON.parse(JSON.stringify(str));
      // console.log(data);
      // this.topoData = data.data.node || [];
      const { data } = await getTopoShow();
      this.topoData = data.data.node || [];
    },
    creatGraph() {
      graph = new Q.Graph(this.$refs.canvas);
      graph.editable = true;
      graph.enableRectangleSelectionByRightButton = true;
    },
    createNode(
      graph,
      image,
      x,
      y,
      name,
      group,
      randomFlag,
      nodesType,
      otherInfo
    ) {
      var node = graph.createNode(name, x, y);
      if (image) {
        if (Q.isString(image)) {
          image = image;
        }
        node.image = image;
      }
      node.size = { height: 70 };
      if (group) {
        group.addChild(node);
      }
      node.randomAble = true;
      node.setStyle(Q.Styles.LABEL_COLOR, "#ffffff");
      node.setStyle(Q.Styles.LABEL_FONT_SIZE, 25);
      node.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_TOP);
      node.setStyle(Q.Styles.LABEL_ANCHOR_POSITION, Q.Position.CENTER_BOTTOM);
      node.nodesType = nodesType;
      node.moreInfo = otherInfo;
      return node;
    },
    createGroup(graph, otherInfo) {
      var group = graph.createGroup();
      group.image = this.imgObg.vmc;
      group.moreInfo = otherInfo;
      group.size = { height: 80 };
      group.setStyle(Q.Styles.GROUP_BACKGROUND_COLOR, "#040f21");
      // group.setStyle(Q.Styles.GROUP_BACKGROUND_A, "#040f21");
      group.setStyle(Q.Styles.ALPHA, 0.8);
      return group;
    },
    createNode_center(graph, image, x, y, name, group, randomFlag, nodesType) {
      var node = graph.createNode(name, x, y);
      if (image) {
        if (Q.isString(image)) {
          image = image;
        }
        node.image = image;
      }
      node.size = { height: 20 };
      if (group) {
        group.addChild(node);
      }
      node.randomAble = false;
      node.setStyle(Q.Styles.SHAPE_FILL_COLOR, "#9bcfee");
      node.setStyle(Q.Styles.SHAPE_STROKE_STYLE, "#9bcfee");
      node.setStyle(Q.Styles.LABEL_COLOR, "#ffffff");
      node.setStyle(Q.Styles.LABEL_FONT_SIZE, 25);
      node.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_TOP);
      node.setStyle(Q.Styles.LABEL_ANCHOR_POSITION, Q.Position.CENTER_BOTTOM);
      node.nodesType = null;
      node.zIndex = 999;
      return node;
    },
    createEdgeForAngle(graph, a, b, angle) {
      var edge = graph.createEdge(null, a, b);
      edge.setStyle(Q.Styles.EDGE_WIDTH, 3);
      edge.setStyle(Q.Styles.EDGE_COLOR, "#8cd1f1");
      edge.setStyle(Q.Styles.ARROW_TO, false);
      edge.zIndex = -2;
      if (angle) {
        edge.angle = angle;
      }
      return edge;
    },
    createEdge(graph, a, b, color, dashed, name, angle) {
      var edge = graph.createEdge(name, a, b);
      if (dashed) {
        edge.setStyle(Q.Styles.EDGE_LINE_DASH, [8, 5]);
      }
      edge.setStyle(Q.Styles.EDGE_WIDTH, 2);
      edge.setStyle(Q.Styles.EDGE_COLOR, "#8cd1f1");
      edge.setStyle(Q.Styles.ARROW_TO, false);
      // edge.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
      // edge.setStyle(Q.Styles.ARROW_FROM, Q.Consts.SHAPE_CIRCLE);
      // edge.setStyle(Q.Styles.ARROW_FROM_STROKE, 7);
      if (angle) {
        edge.angle = angle;
      }
      edge.nodesType = "line";
      return edge;
    },
    createBus(graph) {
      var path = new Q.Path();
      var bus = new Q.Bus(null, path);
      graph.graphModel.add(bus);
      bus.setStyle(Q.Styles.SHAPE_STROKE, 3);
      bus.setStyle(Q.Styles.SHAPE_STROKE_STYLE, "#8cd1f1");
      bus.setStyle(Q.Styles.SHAPE_FILL_COLOR, false);
      return bus;
    },
    initalarm() {
      if (!Q.Element.prototype.initAlarmBalloon) {
        Q.Element.prototype.initAlarmBalloon = function () {
          var alarmUI = new Q.LabelUI();
          alarmUI.position = Q.Position.CENTER_TOP;
          alarmUI.anchorPosition = Q.Position.LEFT_BOTTOM;
          alarmUI.border = 1;
          alarmUI.backgroundGradient = Q.Gradient.LINEAR_GRADIENT_VERTICAL;
          alarmUI.padding = new Q.Insets(2, 5);
          alarmUI.showPointer = true;
          alarmUI.offsetY = -10;
          alarmUI.offsetX = -10;
          alarmUI.rotatable = false;
          alarmUI.showOnTop = true;
          this._alarmBalloon = alarmUI;
        };
        Q.Element.prototype._checkAlarmBalloon = function () {
          if (!this.alarmLabel || !this.alarmColor) {
            if (this._alarmBalloon) {
              this.removeUI(this._alarmBalloon);
            }
            return;
          }
          if (!this._alarmBalloon) {
            this.initAlarmBalloon();
          }
          this._alarmBalloon.data = this.alarmLabel;
          this._alarmBalloon.backgroundColor = this.alarmColor;
          if (this.addUI(this._alarmBalloon) === false) {
            this.invalidate();
          }
        };
        Q.Element.prototype.setAlarm = function (alarmLabel, alarmColor) {
          this.alarmColor = alarmColor;
          this.alarmLabel = alarmLabel;
        };
        Object.defineProperties(Q.Element.prototype, {
          alarmLabel: {
            get: function () {
              return this._alarmLabel;
            },
            set: function (label) {
              if (this._alarmLabel == label) {
                return;
              }
              this._alarmLabel = label;
              this._checkAlarmBalloon();
            },
          },
          alarmColor: {
            get: function () {
              return this._alarmColor;
            },
            set: function (color) {
              if (this._alarmColor == color) {
                return;
              }
              this._alarmColor = color;
              this.setStyle(Q.Styles.RENDER_COLOR, color);
              this._checkAlarmBalloon();
            },
          },
        });
      }
    },
    drawAllCanvas() {
      const that = this;
      const graph = new Q.Graph(this.$refs.canvas);
      var model = graph.graphModel;
      this.initalarm();
      function FlowingSupport(graph) {
        this.flowMap = {};
        this.graph = graph;
      }
      FlowingSupport.prototype = {
        flowMap: null,
        length: 0,
        gap: 40,
        graph: null,
        addFlowing: function (edgeOrLine, count, byPercent, flowColors) {
          var flowList = this.flowMap[edgeOrLine.id];
          if (!flowList) {
            flowList = this.flowMap[edgeOrLine.id] = [];
            this.length++;
          }
          count = count || 1;
          while (--count >= 0) {
            var ui = new Q.ImageUI(pointsvg);
            ui.layoutByPath = true;
            ui.position = { x: 0, y: 0 };
            ui.size = { width: 40 };
            ui.renderColor = flowColors;
            flowList.push(ui);
            flowList.byPercent = byPercent;
            edgeOrLine.addUI(ui);
          }
        },
        removeFlowing: function (id) {
          var flowList = this.flowMap[id];
          if (!flowList) {
            return;
          }
          var edgeOrLine = this.graph.getElement(id);
          if (edgeOrLine) {
            flowList.forEach(function (ui) {
              edgeOrLine.removeUI(ui);
            });
          }
          this._doRemove(id);
        },
        _doRemove: function (id) {
          delete this.flowMap[id];
          this.length--;
        },
        timer: null,
        perStep: 10,
        stop: function () {
          clearTimeout(this.timer);
        },
        start: function () {
          if (this.timer) {
            clearTimeout(this.timer);
          }
          var offset = 0;
          var scope = this;
          scope.timer = setTimeout(function A() {
            if (!scope.length) {
              scope.timer = setTimeout(A, 2000);
              offset = 0;
              return;
            }
            offset += 1;
            for (var id in scope.flowMap) {
              var ui = scope.graph.getUI(id);
              if (!ui) {
                scope._doRemove(id);
                continue;
              }
              var lineLength = ui.length;
              if (!lineLength) {
                continue;
              }
              var flowList = scope.flowMap[id];
              if (flowList.byPercent) {
                var x = offset * 2;
                var gap = 15;
                scope.flowMap[id].forEach(function (ui) {
                  ui.position = { x: (x % 100) / 100, y: 0 };
                  x += gap;
                });
              } else {
                var x = offset * scope.perStep;
                scope.flowMap[id].forEach(function (ui) {
                  ui.position = { x: x % lineLength, y: 0 };
                  x += scope.gap;
                });
              }
              scope.graph.invalidateUI(ui);

              var data = ui.data;
              if (data instanceof Q.Edge) {
                if (data.getStyle(Q.Styles.EDGE_LINE_DASH)) {
                  data.setStyle(Q.Styles.EDGE_LINE_DASH_OFFSET, -offset);
                }
              } else if (data instanceof Q.ShapeNode) {
                if (data.getStyle(Q.Styles.SHAPE_LINE_DASH)) {
                  data.setStyle(Q.Styles.SHAPE_LINE_DASH_OFFSET, -offset);
                }
              }
            }
            scope.timer = setTimeout(A, 50);
          }, 50);
        },
      };
      var VPNFlexEdgeUI = function (edge, graph) {
        Q.doSuperConstructor(this, VPNFlexEdgeUI, arguments);
      };
      VPNFlexEdgeUI.prototype = {
        drawEdge: function (
          path,
          fromUI,
          toUI,
          edgeType,
          fromBounds,
          toBounds
        ) {
          var from = fromBounds.center;
          path.curveTo(from.x, from.y, internet.x, internet.y);
        },
      };

      Q.extend(VPNFlexEdgeUI, Q.EdgeUI);

      let tempobj = this.dealWithDataXY();
      let busnodes = this.drawBusEdgeAndCenterSw(
        graph,
        tempobj.assistantPoints
      );
      let allsw = this.drawSwAndOthers(
        graph,
        tempobj.vmcArrEnd,
        tempobj.rtuArrEnd
      );
      console.log(allsw);
      let ends = this.drawEdgeForSw(
        graph,
        allsw.nodesOfsw,
        allsw.nodesOfOther,
        busnodes
      );
      let ens2 = this.darwTemp(graph, allsw.nodesOfsw, allsw.nodesOfOther);
      var flowingSupport = new FlowingSupport(graph);
      if (busnodes.line1Edge) {
        flowingSupport.addFlowing(
          busnodes.line1Edge,
          1,
          false
          // this.flowColor_vpn
        );
      }
      if (busnodes.line2Edge) {
        flowingSupport.addFlowing(
          busnodes.line2Edge,
          1,
          false
          // this.flowColor_vpn
        );
      }
      if (busnodes.line3Edge) {
        flowingSupport.addFlowing(
          busnodes.line3Edge,
          1,
          false
          // this.flowColor_vpn
        );
      }

      graph.callLater(function () {
        flowingSupport.start();
      });
      /// 报警
      var AlarmSeverity = {
        RUN: { color: Q.toColor(0x7fff00), sortName: "R" },
        WARNING: { color: Q.toColor(0xffa500), sortName: "W" },
        Err: { color: Q.toColor(0xff0000), sortName: "E" },
      };
      var all = [];
      for (var name in AlarmSeverity) {
        all.push(AlarmSeverity[name]);
      }
      AlarmSeverity.all = all;

      Object.defineProperties(AlarmSeverity, {
        random: {
          get: function () {
            return this.all[Q.randomInt(this.all.length)];
          },
        },
      });

      var timer = setTimeout(function A() {
        graph.forEach(function (element) {
          if (Q.randomBool()) {
            element.alarmLabel = null;
            element.alarmColor = null;
            return;
          }
          if (element.randomAble === true && element.moreInfo !== undefined) {
            var alarmSeverity = AlarmSeverity[element.moreInfo.device_status];
            element.alarmColor = alarmSeverity.color;
            // if (element.randomAble === true) {
            //   element.alarmColor = alarmSeverity.color;
            // }
            // element.alarmLabel =
            //   "" +
            //   (1 + Q.randomInt(100)) +
            //   alarmSeverity.sortName +
            //   (Q.randomBool() ? "+" : "");
            // element.alarmColor = alarmSeverity.color;
          }
        });
        timer = setTimeout(A, 1500);
      }, 3000);

      /// 销毁
      function destroy() {
        flowingSupport.stop();
        clearTimeout(timer);
      }
      graph.zoomToOverview(0.8, 0.8);
      // graph.isMovable = false;
      // graph.enableWheelZoom = false;
      graph.onclick = function (evt) {
        if (evt.getData().moreInfo) {
          that.activeNodeInfo = evt.getData().moreInfo;
          that.drawer = true;
          that.drawerText = true;
          if (
            evt.getData().moreInfo.other_info &&
            evt.getData().moreInfo.other_info.length &&
            evt.getData().moreInfo.other_info[0].value[0]
          ) {
            that.deviceType = `${evt.getData().moreInfo.device_type}_${
              evt.getData().moreInfo.other_info[0].value[0]
            }`;
          }
        }
        if (evt.getData().textArea) {
          that.drawerText = true;
        }
      };
    },
    dealWithDataXY() {
      const swArr = [];
      const vmcArr = [];
      const rtuArr = [];
      const otherArr = [];
      let allswData = this.topoData.map((item, index) => {
        if (item.device_type === "sw") {
          swArr.push(item);
        } else if (item.device_type === "vmc") {
          vmcArr.push(item);
        } else if (item.device_type === "rtu") {
          rtuArr.push(item);
        } else {
          otherArr.push(item);
        }
      });
      let swIdForrtu = [];
      rtuArr.map((item) => {
        if (!swIdForrtu.includes(item.upstream_id)) {
          swIdForrtu.push(item.upstream_id);
        }
      });
      let swForRtu = [];
      let swForVmc = [];
      swArr.map((item) => {
        if (!swIdForrtu.includes(item.id)) {
          swForVmc.push(item);
        } else {
          swForRtu.push(item);
        }
      });
      // console.log(swForRtu, swForVmc);
      let incomeSwForOp = [];
      let incomeSwForCal = [];
      let flag = true;
      if (swForVmc.length >= swForRtu.length) {
        incomeSwForCal = swForVmc;
        incomeSwForOp = swForRtu;
        flag = true;
      } else {
        incomeSwForCal = swForRtu;
        incomeSwForOp = swForVmc;
        flag = false;
      }
      let baseLeftX = 0;
      let baseLeftY = 0;
      let baseRightTopX = 0;
      let baseRightBottomX = 0;
      let baseRightTopY = 0;
      let baseRightBottomY = 0;
      let newarr = incomeSwForCal.map((item, index) => {
        let doublecoefficient = Math.floor(incomeSwForCal.length / 2);
        if (index + 1 <= doublecoefficient) {
          return {
            ...item,
            x: -240 * (doublecoefficient - index),
            y: 100 * (doublecoefficient - index - 1),
          };
        } else {
          let baseMultiple = Math.ceil(incomeSwForCal.length / 4);
          if (index % 2 === 0) {
            baseRightTopX = baseRightTopX + 300;
            baseRightTopY = baseRightTopY - 150;
            return {
              ...item,
              x: baseRightTopX,
              y: baseRightTopY,
            };
          } else {
            baseRightBottomX = baseRightBottomX + 300;
            baseRightBottomY = baseRightBottomY + 150;
            return {
              ...item,
              x: baseRightBottomX,
              y: baseRightBottomY,
            };
          }
        }
      });
      let newop = incomeSwForOp.map((items, indexs) => {
        if (newarr[indexs].x <= 0) {
          return {
            ...items,
            x: newarr[indexs].x + 50,
            y: newarr[indexs].y + 100,
          };
        } else {
          if (newarr[indexs].y >= 0) {
            return {
              ...items,
              x: newarr[indexs].x + 100,
              y: newarr[indexs].y + 80,
            };
          } else {
            return {
              ...items,
              x: newarr[indexs].x + 100,
              y: newarr[indexs].y - 80,
            };
          }
        }
      });
      let assistantPoints = [];
      newarr.map((itemNow, indexNow) => {
        if (itemNow.x <= 0) {
          let tempointx = itemNow.x + 40;
          let tempointy = itemNow.y;
          assistantPoints.push(
            {
              no: `${indexNow}_1`,
              type: "busNode",
              x: tempointx - 40,
              y: tempointy + 100,
            },
            {
              no: `${indexNow}_2`,
              type: "busNode",
              x: tempointx + 40,
              y: tempointy,
            }
          );
        } else {
          let tempointx = itemNow.x + 30;
          let tempointy;
          if (itemNow.y >= 0) {
            tempointy = itemNow.y + 40;
          } else {
            tempointy = itemNow.y - 40;
          }
          assistantPoints.push(
            {
              no: `${indexNow}_1`,
              type: "busNode",
              x: tempointx - 100,
              y: tempointy,
            },
            {
              no: `${indexNow}_2`,
              type: "busNode",
              x: tempointx + 100,
              y: tempointy,
            }
          );
        }
      });
      if (flag) {
        return {
          vmcArrEnd: newarr,
          rtuArrEnd: newop,
          assistantPoints,
        };
      }
      return {
        vmcArrEnd: newop,
        rtuArrEnd: newarr,
        assistantPoints,
      };
    },
    dealWithDeviceXY(data) {
      const endObj = {};
      for (const key in data) {
        if (data[key].length) {
          const temp = this.getSwitch(key);
          let ends = data[key].map((item, index) => {
            if (temp.x < 0) {
              return {
                ...item,
                x: temp.x + 120,
                y: temp.y - 20 + index * -40,
              };
            } else {
              if (temp.y < 0) {
                return {
                  ...item,
                  x: temp.x - 150,
                  y: temp.y - 50 + index * -50,
                };
              } else {
                return {
                  ...item,
                  x: temp.x - 150,
                  y: temp.y + 50 + index * 50,
                };
              }
            }
          });
          endObj[key] = ends;
        }
      }
      return endObj;
    },
    dealWithFarendOpXY(data) {
      const endObj = {};
      for (const key in data) {
        if (data[key].length) {
          const temp = this.getSwitch(key);
          let ends = data[key].map((item, index) => {
            if (temp.x < 0) {
              return {
                ...item,
                x: temp.x - 120,
                y: temp.y + 20 + index * 20,
              };
            } else {
              return {
                ...item,
                x: temp.x + 100,
                y: temp.y - 30 + index * 35,
              };
            }
          });
          endObj[key] = ends;
        }
      }
      return endObj;
    },
    getSwitch(id) {
      let result = {};
      this.allSwXY.map((item) => {
        if (item.id === id) {
          result = item;
        }
        return item;
      });
      return result;
    },
    // 传入参数为第一个参数是所有相关主线bus所有的线，参数二为画布实体,目的绘制中心交换机和总线bus
    drawBusEdgeAndCenterSw(graph, pointArr) {
      let tempsarr = pointArr.map((item) => {
        item.y = item.y * -1;
        return item;
      });
      let arrleft = [];
      let arrrightTop = [];
      let arrrightBottom = [];
      tempsarr.map((item) => {
        if (item.x <= 0) {
          arrleft.push(item);
        } else {
          if (item.y <= 0) {
            arrrightBottom.push(item);
          } else {
            arrrightTop.push(item);
          }
        }
        return item;
      });
      let line1 = arrleft.sort(this.compare("x", true));
      let line2 = arrrightTop.sort(this.compare("x", false));
      let line3 = arrrightBottom.sort(this.compare("x", false));
      let mostleftpoint = line1[0];
      let mostrightToppoint = line2[0];
      let mostrightBottompoint = line3[0];
      // let sheapcenter = Q.Shapes.getShape(Q.Consts.SHAPE_CIRCLE, 15, 15, 1, 1);
      var centerNode = graph.createText("TTE", 0, 0);
      centerNode.setStyle(Q.Styles.LABEL_COLOR, "#FFF");
      centerNode.setStyle(Q.Styles.LABEL_BACKGROUND_COLOR, "#9bcfee");
      // centerNode.setStyle(Q.Styles.LABEL_RADIUS, 120);
      centerNode.setStyle(Q.Styles.LABEL_FONT_SIZE, 50);
      centerNode.setStyle(Q.Styles.LABEL_FONT_STYLE, "italic lighter");
      centerNode.zIndex = 9999;
      centerNode.textArea = true;
      // let centerNode = this.createNode_center(graph, textNode, 0, 0);
      let line1Edge = null;
      let line2Edge = null;
      let line3Edge = null;
      if (mostleftpoint) {
        let leftnNode = this.createNode(
          graph,
          pointsvg,
          mostleftpoint.x,
          mostleftpoint.y
        );
        leftnNode.zIndex = 99;
        line1Edge = this.createBus(graph);
        line1.map((item, index) => {
          if (index === 0) {
            line1Edge.moveTo(item.x, item.y);
          } else {
            line1Edge.lineTo(item.x, item.y);
          }
        });
        line1Edge.lineTo(0, 0);
      }
      if (mostrightToppoint) {
        let rightTopNode = this.createNode(
          graph,
          pointsvg,
          mostrightToppoint.x,
          mostrightToppoint.y
        );
        rightTopNode.zIndex = 99;
        line2Edge = this.createBus(graph, rightTopNode, centerNode);
        line2.map((item, index) => {
          if (index === 0) {
            line2Edge.moveTo(item.x, item.y);
          } else {
            line2Edge.lineTo(item.x, item.y);
          }
        });
        line2Edge.lineTo(0, 0);
      }
      if (mostrightBottompoint) {
        let rightBottomNode = this.createNode(
          graph,
          pointsvg,
          mostrightBottompoint.x,
          mostrightBottompoint.y
        );
        rightBottomNode.zIndex = 99;
        line3Edge = this.createBus(graph, rightBottomNode, centerNode);
        line3.map((item, index) => {
          if (index === 0) {
            line3Edge.moveTo(item.x, item.y);
          } else {
            line3Edge.lineTo(item.x, item.y);
          }
        });
        line3Edge.lineTo(0, 0);
      }
      return { line1Edge, line2Edge, line3Edge };
    },
    drawSwAndOthers(graph, vmcSWArr, rtuSWArr, busObj) {
      let tempsarr = vmcSWArr.map((item) => {
        if (item.x <= 0) {
          item.y = item.y * -1;
        } else {
          if (item.y > 0) {
            item.y = (item.y + 80) * -1;
          } else {
            item.y = (item.y - 80) * -1;
          }
        }
        return item;
      });
      let tempsarr_ = rtuSWArr.map((item) => {
        if (item.x <= 0) {
          item.y = item.y * -1;
          item.x = item.x + 40;
        } else {
          item.x = item.x - 50;
          if (item.y > 0) {
            item.y = (item.y - 80) * -1;
          } else {
            item.y = (item.y + 80) * -1;
          }
        }
        return item;
      });
      console.log(tempsarr, tempsarr_);
      let nodesOfsw = tempsarr.map((item) => {
        return this.createNode(
          graph,
          inswsvg,
          item.x,
          item.y,
          null,
          null,
          null,
          "SW",
          item
        );
      });
      let nodesOfOther = tempsarr_.map((item) => {
        return this.createNode(
          graph,
          inswsvg,
          item.x,
          item.y,
          null,
          null,
          null,
          "SW",
          item
        );
      });
      return { nodesOfsw, nodesOfOther };
    },
    drawEdgeForSw(graph, allArrSw, otherArrSw, lineObj) {
      let endline = allArrSw.map((item) => {
        if (item.x <= 0 && lineObj.line1Edge) {
          return this.createEdgeForAngle(
            graph,
            item,
            lineObj.line1Edge,
            -Math.PI / 2.5
          );
        } else {
          if (item.y <= 0) {
            return this.createEdgeForAngle(
              graph,
              item,
              lineObj.line3Edge,
              Math.PI / 2
            );
          } else {
            return this.createEdgeForAngle(
              graph,
              item,
              lineObj.line2Edge,
              Math.PI / 2
            );
          }
        }
      });
      let endOtherArrSw = otherArrSw.map((item) => {
        if (item.x <= 0 && lineObj.line1Edge) {
          return this.createEdgeForAngle(
            graph,
            item,
            lineObj.line1Edge,
            -Math.PI / 2.5
          );
        } else {
          if (item.y <= 0) {
            return this.createEdgeForAngle(
              graph,
              item,
              lineObj.line3Edge,
              Math.PI / 2
            );
          } else {
            return this.createEdgeForAngle(
              graph,
              item,
              lineObj.line2Edge,
              Math.PI / 2
            );
          }
        }
      });
      return endline;
    },
    darwTemp(graph, vmcArrSw, rtuArrSw) {
      let nodesOfvmc = vmcArrSw.map((item, index) => {
        if (item.x <= 0) {
          let tempy = 80;
          let getRelevanceVmc = [];
          this.topoData.map((itemnow) => {
            if (
              itemnow.upstream_id === item.moreInfo.id &&
              itemnow.device_type === "vmc"
            ) {
              getRelevanceVmc.push({
                ...itemnow,
                x: item.x - 100,
                y: item.y + tempy,
              });
              tempy = tempy + 140;
            }
          });
          let vmcArr1 = getRelevanceVmc.map((vmcitem) => {
            let vmcGroup = this.createGroup(graph, { test: "group" });
            let tempnode = this.createNode(
              graph,
              imgvmc,
              vmcitem.x,
              vmcitem.y,
              null,
              vmcGroup,
              null,
              "VMC",
              vmcitem
            );
            this.createEdge(graph, tempnode, item);
            tempnode.GroupSNow = vmcGroup;
            return tempnode;
          });
          let opArr = vmcArr1.map((vmcnode) => {
            let tempx = 100;
            let getRelevanceCal = [];
            this.topoData.map((itemnow) => {
              if (
                itemnow.parent_id === vmcnode.moreInfo.id &&
                itemnow.device_type !== "vmc" &&
                itemnow.device_type !== "sw" &&
                itemnow.device_type !== "rtu"
              ) {
                getRelevanceCal.push({
                  ...itemnow,
                  x: vmcnode.x - tempx,
                  y: vmcnode.y + 50,
                });
                tempx = tempx + 100;
              }
            });
            let calArr1 = getRelevanceCal.map((calitem) => {
              let tempnodes = this.createNode(
                graph,
                this.imgObg[calitem.device_type],
                calitem.x,
                calitem.y,
                null,
                vmcnode.GroupSNow,
                null,
                calitem.device_type,
                calitem
              );
              let tempedge = this.createEdge(graph, tempnodes, vmcnode);
              tempedge.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
              return tempnodes;
            });
          });
        } else {
          if (item.y > 0) {
            let tempy = 80;
            let getRelevanceVmc = [];
            this.topoData.map((itemnow) => {
              if (
                itemnow.upstream_id === item.moreInfo.id &&
                itemnow.device_type === "vmc"
              ) {
                getRelevanceVmc.push({
                  ...itemnow,
                  x: item.x - 100,
                  y: item.y + tempy,
                });
                tempy = tempy + 140;
              }
            });
            let vmcArr1 = getRelevanceVmc.map((vmcitem) => {
              let vmcGroup = this.createGroup(graph, { test: "group" });
              let tempnode = this.createNode(
                graph,
                imgvmc,
                vmcitem.x,
                vmcitem.y,
                null,
                vmcGroup,
                null,
                "VMC",
                vmcitem
              );
              this.createEdge(graph, tempnode, item);
              tempnode.GroupSNow = vmcGroup;
              return tempnode;
            });
            let opArr = vmcArr1.map((vmcnode) => {
              let tempx = 100;
              let getRelevanceCal = [];
              this.topoData.map((itemnow) => {
                if (
                  itemnow.parent_id === vmcnode.moreInfo.id &&
                  itemnow.device_type !== "vmc" &&
                  itemnow.device_type !== "sw" &&
                  itemnow.device_type !== "rtu"
                ) {
                  getRelevanceCal.push({
                    ...itemnow,
                    x: vmcnode.x - tempx,
                    y: vmcnode.y + 50,
                  });
                  tempx = tempx + 100;
                }
              });
              let calArr1 = getRelevanceCal.map((calitem) => {
                let tempnodes = this.createNode(
                  graph,
                  this.imgObg[calitem.device_type],
                  calitem.x,
                  calitem.y,
                  null,
                  vmcnode.GroupSNow,
                  null,
                  calitem.device_type,
                  calitem
                );
                let tempedge = this.createEdge(graph, tempnodes, vmcnode);
                tempedge.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
                return tempnodes;
              });
            });
          } else {
            let tempy = 80;
            let getRelevanceVmc = [];
            this.topoData.map((itemnow) => {
              if (
                itemnow.upstream_id === item.moreInfo.id &&
                itemnow.device_type === "vmc"
              ) {
                getRelevanceVmc.push({
                  ...itemnow,
                  x: item.x + 100,
                  y: item.y - tempy,
                });
                tempy = tempy + 140;
              }
            });
            let vmcArr1 = getRelevanceVmc.map((vmcitem) => {
              let vmcGroup = this.createGroup(graph, { test: "group" });
              let tempnode = this.createNode(
                graph,
                imgvmc,
                vmcitem.x,
                vmcitem.y,
                null,
                vmcGroup,
                null,
                "VMC",
                vmcitem
              );
              this.createEdge(graph, tempnode, item);
              tempnode.GroupSNow = vmcGroup;
              return tempnode;
            });
            let opArr = vmcArr1.map((vmcnode) => {
              let tempx = 100;
              let getRelevanceCal = [];
              this.topoData.map((itemnow) => {
                if (
                  itemnow.parent_id === vmcnode.moreInfo.id &&
                  itemnow.device_type !== "vmc" &&
                  itemnow.device_type !== "sw" &&
                  itemnow.device_type !== "rtu"
                ) {
                  getRelevanceCal.push({
                    ...itemnow,
                    x: vmcnode.x + tempx,
                    y: vmcnode.y - 50,
                  });
                  tempx = tempx + 100;
                }
              });
              let calArr1 = getRelevanceCal.map((calitem) => {
                let tempnodes = this.createNode(
                  graph,
                  this.imgObg[calitem.device_type],
                  calitem.x,
                  calitem.y,
                  null,
                  vmcnode.GroupSNow,
                  null,
                  calitem.device_type,
                  calitem
                );
                let tempedge = this.createEdge(graph, tempnodes, vmcnode);
                tempedge.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
                return tempnodes;
              });
            });
          }
        }
      });
      let nodesOfother = rtuArrSw.map((item, index) => {
        if (item.x <= 0) {
          let tempy = 80;
          let getRelevanceRtu = [];
          this.topoData.map((itemnow) => {
            if (
              itemnow.upstream_id === item.moreInfo.id &&
              itemnow.device_type === "rtu"
            ) {
              getRelevanceRtu.push({
                ...itemnow,
                x: item.x + 100,
                y: item.y - tempy,
              });
              tempy = tempy + 100;
            }
          });
          let rtuArr1 = getRelevanceRtu.map((rtuItem) => {
            let tempnode = this.createNode(
              graph,
              this.imgObg[
                `${rtuItem.device_type}_${rtuItem.other_info[0].value[0]}`
              ],
              rtuItem.x,
              rtuItem.y,
              null,
              null,
              null,
              rtuItem.device_type,
              rtuItem
            );
            this.createEdge(graph, tempnode, item);
            return tempnode;
          });
        } else {
          if (item.y <= 0) {
            let tempy = 80;
            let getRelevanceRtu = [];
            this.topoData.map((itemnow) => {
              if (
                itemnow.upstream_id === item.moreInfo.id &&
                itemnow.device_type === "rtu"
              ) {
                getRelevanceRtu.push({
                  ...itemnow,
                  x: item.x - 100,
                  y: item.y - tempy,
                });
                tempy = tempy + 80;
              }
            });
            let rtuArr1 = getRelevanceRtu.map((rtuItem) => {
              let tempnode = this.createNode(
                graph,
                this.imgObg[
                  `${rtuItem.device_type}_${rtuItem.other_info[0].value[0]}`
                ],
                rtuItem.x,
                rtuItem.y,
                null,
                null,
                null,
                rtuItem.device_type,
                rtuItem
              );
              this.createEdge(graph, tempnode, item);
              return tempnode;
            });
          } else {
            let tempy = 80;
            let getRelevanceRtu = [];
            this.topoData.map((itemnow) => {
              if (
                itemnow.upstream_id === item.moreInfo.id &&
                itemnow.device_type === "rtu"
              ) {
                getRelevanceRtu.push({
                  ...itemnow,
                  x: item.x + 100,
                  y: item.y - tempy,
                });
                tempy = tempy + 80;
              }
            });
            let rtuArr1 = getRelevanceRtu.map((rtuItem) => {
              let tempnode = this.createNode(
                graph,
                this.imgObg[
                  `${rtuItem.device_type}_${rtuItem.other_info[0].value[0]}`
                ],
                rtuItem.x,
                rtuItem.y,
                null,
                null,
                null,
                rtuItem.device_type,
                rtuItem
              );
              this.createEdge(graph, tempnode, item);
              return tempnode;
            });
          }
        }
      });
    },
    /** 两个参数： 参数1 是排序用的字段， 参数2 是：是否升序排序 true 为升序，false为降序*/
    compare(attr, rev) {
      // console.log(attr, rev)
      if (rev == undefined) {
        rev = 1;
      } else {
        rev = rev ? 1 : -1;
      }
      return (a, b) => {
        a = a[attr];
        b = b[attr];
        if (a < b) {
          return rev * -1;
        }
        if (a > b) {
          return rev * 1;
        }
        return 0;
      };
    },
    mockError(type) {
      this.drawer = false;
      if (type) {
        this.$message({
          message:
            "模拟整机故障触发成功，已为您重定向到容灾演示页面查看任务迁移详情",
          type: "success",
        });
        this.setDisVmc({
          ...this.activeNodeInfo,
          time: new Date().valueOf(),
        });
        this.$router.push("/disasterrecovery");
      } else {
        this.$message({
          message:
            "模拟分区故障触发成功，已为您重定向到容灾演示页面查看任务迁移详情",
          type: "success",
        });
        this.setDisArea({
          ...this.activeNodeInfo,
          time: new Date().valueOf(),
        });
        this.$router.push("/disasterrecovery");
      }
    },
  },
  mounted() {
    setTimeout(() => {
      this.loading = false;
      setTimeout(() => {
        this.drawAllCanvas();
      }, 500);
    }, 2000);
  },
  created() {
    this.getTopoData();
  },
};
</script>
<style lang="less">
.topobox {
  background: url("../assets/newpng/BACK2.png") no-repeat center;
  background-size: 100% 100%;
}
.el-drawer {
  background: rgba(0, 0, 0, 0.9) !important;
}
.el-dialog {
  background: #0d2043 !important;
  .el-dialog__title {
    color: #fff !important;
  }
  .el-dialog__body {
    color: #fff !important;
    padding: 24px;
  }
}
.drawer_info_left {
  height: 40% !important;
  color: #fff !important;
}
.mainbox {
  padding: 0;
  margin: 0;
  width: 100%;
  height: 100%;
  background: url("../assets/png/homeasidebg.png") no-repeat center;
  background-size: 100% 98%;
  background-position: 0.5vw 0;
  padding: 1.5rem;
  box-sizing: border-box;
}
.aside_box {
  height: 28vh;
  // background: rgb(207, 160, 160);
  // padding: 1%;
  // box-sizing: border-box;
  text-align: center;
  .grid-content_btn {
    padding: 1rem;
    display: inline-block;
    width: 40%;
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
  .miansvg {
    height: 80%;
    // margin-left: 25%;
  }
}
.aside_box_line_bar {
  height: 28vh;
  // background-color: #fff;
  .canvasbox {
    height: 85%;
    width: 100%;
    p {
      color: #fff;
      margin-left: 10%;
      font-size: 1rem;
    }
  }
  .infoBox {
    height: 85%;
    width: 100%;
    p {
      color: #fff;
      margin-left: 10%;
      font-size: 1rem;
    }
  }
  // background-color: aqua;
  // padding: 1%;
  // box-sizing: border-box;
}
.aside_box_task {
  height: 25vh;
  overflow: hidden;
  // background-color: rgb(204, 0, 255);
  // padding: 1%;
  // box-sizing: border-box;
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
    line-height: 2.29rem;
  }
}
.loadingbox {
  width: 100%;
  height: 85vh;
  background-color: rgba(0, 0, 0, 0);
  display: flex;
  justify-content: center;
  // padding-top: 30%;
  align-items: center;
}
.boxForline {
  // margin-top: 60px;
  width: 100%;
  height: 80%;
}
</style>