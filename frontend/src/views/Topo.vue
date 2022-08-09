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
      :title="`设备：${devicename}的详情`"
      :visible.sync="drawer"
      :modal="false"
    >
      <div v-if="devicename === 'cpu'" class="mainbox">
        <div class="aside_box">
          <CpuInfo :cpuNow="cpu" />
        </div>
        <div class="aside_box_line_bar">
          <p class="title">
            <span>机器性能历史状况</span>
          </p>
          <div class="canvasbox" id="linebox_">
            <selflineNew inref="linebox_" />
          </div>
        </div>
        <div class="aside_box_task">
          <p class="title">
            <span>分时分区任务</span>
          </p>
          <scrolltable :data="tableData" />
        </div>
      </div>
      <div v-else class="mainbox">
        <div class="aside_box">
          <div class="grid-content_btn">
            <button @click="mockError(1)">模拟整机故障</button>
          </div>
          <div class="grid-content_btn">
            <button @click="mockError(0)">模拟分区故障</button>
          </div>
          <img
            v-if="deviceType === 'SW'"
            class="miansvg"
            src="../assets/newpng/sw.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'CPU'"
            class="miansvg"
            src="../assets/newpng/CPU.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'GPU'"
            class="miansvg"
            src="../assets/newpng/GPU.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'DSP'"
            class="miansvg"
            src="../assets/newpng/DSP.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'FPGA'"
            class="miansvg"
            src="../assets/newpng/FPGA.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'VMC'"
            class="miansvg"
            src="../assets/newpng/VMC.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'OP1'"
            class="miansvg"
            src="../assets/newpng/op_1.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'OP2'"
            class="miansvg"
            src="../assets/newpng/op_2.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'OP3'"
            class="miansvg"
            src="../assets/newpng/op_3.svg"
            alt=""
          />
          <img
            v-if="deviceType === 'OP4'"
            class="miansvg"
            src="../assets/newpng/op_4.svg"
            alt=""
          />
        </div>
        <div class="aside_box_line_bar">
          <p class="title">
            <span>机器信息</span>
          </p>
          <div class="canvasbox">
            <!-- <selflineNew inref="linebox_" /> -->
            <p>
              信息字段1:
              信息信息&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;信息字段5:
              信息信息
            </p>
            <p>
              信息字段2:
              信息信息&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;信息字段6:
              信息信息
            </p>
            <p>
              信息字段3:
              信息信息&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;信息字段7:
              信息信息
            </p>
            <p>
              信息字段4:
              信息信息&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;信息字段8:
              信息信息
            </p>
            <!-- <p>信息字段5: 信息信息</p> -->
          </div>
        </div>
        <!-- <div class="aside_box_task">
          <p class="title">
            <span>交换机历史数据信息</span>
          </p>
          <div class="boxForline" id="linebox">
            <selfline inref="linebox" />
          </div>
        </div> -->
      </div>
    </el-drawer>
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
import { getTopoShow } from "@/api";
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
    };
  },
  methods: {
    ...mapMutations(["setDisVmc", "setDisArea"]),
    async getTopoData() {
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
      node.randomAble = randomFlag || false;
      node.setStyle(Q.Styles.LABEL_COLOR, "#ffffff");
      node.setStyle(Q.Styles.LABEL_FONT_SIZE, 25);
      node.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_TOP);
      node.setStyle(Q.Styles.LABEL_ANCHOR_POSITION, Q.Position.CENTER_BOTTOM);
      node.nodesType = nodesType;
      node.moreInfo = otherInfo;
      return node;
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
        CRITICAL: { color: Q.toColor(0xeeff0000), sortName: "C" },
        MAJOR: { color: Q.toColor(0xeeffaa00), sortName: "M" },
        MINOR: { color: Q.toColor(0xeeffff00), sortName: "m" },
        WARNING: { color: Q.toColor(0xee00ffff), sortName: "W" },
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
          if (element.randomAble) {
            // var alarmSeverity = AlarmSeverity.random;
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
        if (evt.getData()) {
          if (evt.getData().nodesType) {
            that.deviceType = evt.getData().nodesType;
            that.drawer = true;
            if (evt.getData().nodesType === "SW") {
              that.devicename = "交换机_60002";
            }
            if (evt.getData().nodesType === "VMC") {
              that.devicename = "VMC_10000";
            }
            if (evt.getData().nodesType === "CPU") {
              that.devicename = "CPU_2001";
            }
            if (evt.getData().nodesType === "GPU") {
              that.devicename = "GPU_2002";
            }
            if (evt.getData().nodesType === "DSP") {
              that.devicename = "DSP_2003";
            }
            if (evt.getData().nodesType === "FPGA") {
              that.devicename = "FPGA_2004";
            }
            if (evt.getData().nodesType === "OP1") {
              that.devicename = "相机_3001";
            }
            if (evt.getData().nodesType === "OP2") {
              that.devicename = "激光扫描仪_3002";
            }
            if (evt.getData().nodesType === "OP3") {
              that.devicename = "激光雷达_3003";
            }
            if (evt.getData().nodesType === "OP4") {
              that.devicename = "推进器_3004";
            }
          }
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
      let sheapcenter = Q.Shapes.getShape(Q.Consts.SHAPE_CIRCLE, 15, 15, 1, 1);
      let centerNode = this.createNode_center(graph, sheapcenter, 0, 0);
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
        leftnNode.zIndex = 999;
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
        rightTopNode.zIndex = 999;
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
        rightBottomNode.zIndex = 999;
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
          let tempy = 50;
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
              tempy = tempy + 60;
            }
          });
          getRelevanceVmc.map((vmcitem) => {
            return this.createNode(
              graph,
              imgvmc,
              vmcitem.x,
              vmcitem.y,
              null,
              null,
              null,
              "VMC"
            );
          });
          // let tempsnow = this.createNode(
          //   graph,
          //   imgvmc,
          //   item.x - 100,
          //   item.y + 50,
          //   null,
          //   null,
          //   null,
          //   "VMC"
          // );
          // let tempsnow1 = this.createNode(
          //   graph,
          //   imgvmc,
          //   item.x - 100,
          //   item.y + 110,
          //   null,
          //   null,
          //   null,
          //   "VMC"
          // );
          // let tempsnow_1 = this.createNode(
          //   graph,
          //   imgcpu,
          //   item.x - 250,
          //   item.y + 50,
          //   null,
          //   null,
          //   null,
          //   "CPU"
          // );
          // let tempsnow_2 = this.createNode(
          //   graph,
          //   imggpu,
          //   item.x - 250,
          //   item.y + 130,
          //   null,
          //   null,
          //   null,
          //   "GPU"
          // );
          // let tempsnow_3 = this.createNode(
          //   graph,
          //   imgdsp,
          //   item.x - 250,
          //   item.y + 210,
          //   null,
          //   null,
          //   null,
          //   "DSP"
          // );
          // let tempsnow_4 = this.createNode(
          //   graph,
          //   imgfgpa,
          //   item.x - 250,
          //   item.y + 290,
          //   null,
          //   null,
          //   null,
          //   "FPGA"
          // );
          // this.createEdge(graph, tempsnow, item);
          // this.createEdge(graph, tempsnow1, item);
          if (index % 2 === 0) {
            // this.createEdge(graph, tempsnow_1, tempsnow1);
            // this.createEdge(graph, tempsnow_2, tempsnow1);
            // this.createEdge(graph, tempsnow_3, tempsnow1);
            // this.createEdge(graph, tempsnow_4, tempsnow1);
          } else {
            // this.createEdge(graph, tempsnow_1, tempsnow);
            // this.createEdge(graph, tempsnow_2, tempsnow);
            // this.createEdge(graph, tempsnow_3, tempsnow);
            // this.createEdge(graph, tempsnow_4, tempsnow);
          }
          // return tempsnow;
        } else {
          // if (item.y > 0) {
          //   let tempsnow_ = this.createNode(
          //     graph,
          //     imgvmc,
          //     item.x - 150,
          //     item.y,
          //     null,
          //     null,
          //     null,
          //     "VMC"
          //   );
          //   let tempsnow_1 = this.createNode(
          //     graph,
          //     imgcpu,
          //     item.x - 250,
          //     item.y,
          //     null,
          //     null,
          //     null,
          //     "CPU"
          //   );
          //   let tempsnow_2 = this.createNode(
          //     graph,
          //     imggpu,
          //     item.x - 250,
          //     item.y + 80,
          //     null,
          //     null,
          //     null,
          //     "GPU"
          //   );
          //   let tempsnow_3 = this.createNode(
          //     graph,
          //     imgdsp,
          //     item.x - 250,
          //     item.y + 160,
          //     null,
          //     null,
          //     null,
          //     "DSP"
          //   );
          //   let tempsnow_4 = this.createNode(
          //     graph,
          //     imgfgpa,
          //     item.x - 250,
          //     item.y + 240,
          //     null,
          //     null,
          //     null,
          //     "FPGA"
          //   );
          //   this.createEdge(graph, tempsnow_1, tempsnow_);
          //   this.createEdge(graph, tempsnow_2, tempsnow_);
          //   this.createEdge(graph, tempsnow_3, tempsnow_);
          //   this.createEdge(graph, tempsnow_4, tempsnow_);
          //   this.createEdge(graph, tempsnow_, item);
          //   return tempsnow_;
          // } else {
          //   let tempsnow_ = this.createNode(
          //     graph,
          //     imgvmc,
          //     item.x - 150,
          //     item.y,
          //     null,
          //     null,
          //     null,
          //     "VMC"
          //   );
          //   this.createEdge(graph, tempsnow_, item);
          // }
        }
      });
      let nodesOfother = rtuArrSw.map((item, index) => {
        if (item.x <= 0) {
          let tempsnow_1 = this.createNode(
            graph,
            imgop1,
            item.x + 120,
            item.y,
            null,
            null,
            null,
            "OP1"
          );
          let tempsnow_2 = this.createNode(
            graph,
            imgop2,
            item.x + 120,
            item.y - 80,
            null,
            null,
            null,
            "OP2"
          );
          let tempsnow_3 = this.createNode(
            graph,
            imgop3,
            item.x + 120,
            item.y - 160,
            null,
            null,
            null,
            "OP3"
          );
          let tempsnow_4 = this.createNode(
            graph,
            imgop4,
            item.x + 120,
            item.y - 240,
            null,
            null,
            null,
            "OP4"
          );
          this.createEdge(graph, tempsnow_1, item);
          this.createEdge(graph, tempsnow_2, item);
          this.createEdge(graph, tempsnow_3, item);
          this.createEdge(graph, tempsnow_4, item);
          return tempsnow_1;
        } else {
          if (item.y <= 0) {
            let tempsnow_1 = this.createNode(
              graph,
              imgop1,
              item.x + 120,
              item.y,
              null,
              null,
              null,
              "OP1"
            );
            let tempsnow_2 = this.createNode(
              graph,
              imgop2,
              item.x + 120,
              item.y + 80,
              null,
              null,
              null,
              "OP2"
            );
            let tempsnow_3 = this.createNode(
              graph,
              imgop3,
              item.x + 120,
              item.y + 160,
              null,
              null,
              null,
              "OP3"
            );
            let tempsnow_4 = this.createNode(
              graph,
              imgop4,
              item.x + 120,
              item.y + 240,
              null,
              null,
              null,
              "OP4"
            );
            this.createEdge(graph, tempsnow_1, item);
            this.createEdge(graph, tempsnow_2, item);
            this.createEdge(graph, tempsnow_3, item);
            this.createEdge(graph, tempsnow_4, item);
            return tempsnow_1;
          } else {
            let tempsnow_2 = this.createNode(
              graph,
              imgop2,
              item.x - 120,
              item.y,
              null,
              null,
              null,
              "OP2"
            );
            this.createEdge(graph, tempsnow_2, item);
            return tempsnow_2;
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
          message: "模拟整机故障触发成功，可前往容灾演示页面查看任务迁移详情",
          type: "success",
        });
        this.setDisVmc(new Date().valueOf());
      } else {
        this.$message({
          message: "模拟分区故障触发成功，可前往容灾演示页面查看任务迁移详情",
          type: "success",
        });
        this.setDisArea(new Date().valueOf());
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