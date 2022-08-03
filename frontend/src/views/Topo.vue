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
            v-if="deviceType === 'sw'"
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
        <div class="aside_box_task">
          <p class="title">
            <span>交换机历史数据信息</span>
          </p>
          <div class="boxForline" id="linebox">
            <selfline inref="linebox" />
          </div>
        </div>
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
import { nameOrAll } from "@/api";
// import text from "../assets/data/topo.json";
import imgcpu from "@/assets/newpng/CPU.svg";
import imggpu from "@/assets/newpng/GPU.svg";
import imgdsp from "@/assets/newpng/DSP.svg";
import imgfgpa from "@/assets/newpng/FPGA.svg";
import imgsw from "@/assets/newpng/sw.svg";
import imgop1 from "@/assets/newpng/op_1.svg";
import imgop2 from "@/assets/newpng/op_2.svg";
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
      topoData: [
        {
          id: 0,
          name: "CPU_00",
          device_type: "cpu",
          parent_id: 0,
          upstream_id: 0,
          device_status: "RUN",
          other_info: [
            {
              key: "cpu_type",
              value: "3",
            },
            {
              key: "cpu_cores",
              value: "154",
            },
          ],
        },
        {
          id: 2,
          name: "DSP_02",
          device_type: "dsp",
          parent_id: 0,
          upstream_id: 0,
          device_status: "RUN",
          other_info: [
            {
              key: "dsp_type",
              value: "0",
            },
            {
              key: "dsp_cores",
              value: "181",
            },
          ],
        },
        {
          id: 0,
          name: "FPGA_00",
          device_type: "fpga",
          parent_id: 0,
          upstream_id: 0,
          device_status: "RUN",
          other_info: [
            {
              key: "fpga_type",
              value: "0",
            },
            {
              key: "fpga_cores",
              value: "0",
            },
          ],
        },
        {
          id: 0,
          name: "vmc_00@",
          device_type: "vmc",
          parent_id: 403,
          upstream_id: 0,
          device_status: "RUN",
          other_info: [
            {
              key: "proto_type",
              value: "85",
            },
          ],
        },
      ],
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
    async getNameOAll() {
      const { data } = await nameOrAll();
      // data = JSON.parse(data)
      console.log(data);
      this.names = data[0]["names"];
    },
    creatGraph() {
      graph = new Q.Graph(this.$refs.canvas);
      graph.editable = true;
      graph.enableRectangleSelectionByRightButton = true;
    },
    createNode(graph, image, x, y, name, group, randomFlag, nodesType) {
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
      // node.setStyle(Q.Styles., 25);
      model.add(node);
      return node;
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
            var ui = new Q.ImageUI(flow);
            ui.layoutByPath = true;
            ui.position = { x: 0, y: 0 };
            ui.size = { width: 20 };
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
      function createNode(image, x, y, name, group, randomFlag, nodesType) {
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
        // node.setStyle(Q.Styles., 25);
        model.add(node);
        return node;
      }
      function createSw(image, x, y, name, group, randomFlag) {
        var node = graph.createNode(name, x, y);
        if (image) {
          if (Q.isString(image)) {
            image = image;
          }
          node.image = image;
        }
        node.size = { height: 120 };
        if (group) {
          group.addChild(node);
        }
        node.randomAble = randomFlag || false;
        node.setStyle(Q.Styles.LABEL_COLOR, "#ffffff");
        node.setStyle(Q.Styles.LABEL_FONT_SIZE, 25);
        node.nodesType = "sw";
        model.add(node);
        return node;
      }
      function createText(name, x, y, fontSize, color, parent) {
        var text = graph.createText(name, x, y);
        text.setStyle(Q.Styles.LABEL_ANCHOR_POSITION, Q.Position.CENTER_MIDDLE);
        text.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_MIDDLE);
        text.setStyle(Q.Styles.LABEL_FONT_SIZE, fontSize);
        text.setStyle(Q.Styles.LABEL_COLOR, color);
        text.setStyle(Q.Styles.LABEL_BACKGROUND_COLOR, null);
        if (parent) {
          parent.addChild(text);
        }
        return text;
      }
      function createGroup(padding, flag, image, host, name) {
        var group = graph.createGroup(name);
        // group.name = name;
        group.image = image;
        group.groupType = Q.Consts.GROUP_TYPE_RECT;
        group.padding = padding || 40;
        // group.minSize = {
        //   width: 300,
        //   height: 440,
        // };
        if (flag) {
          group.setStyle(Q.Styles.GROUP_BACKGROUND_COLOR, "#000000");
          group.setStyle(Q.Styles.ALPHA, "0");
        } else {
          group.setStyle(Q.Styles.GROUP_BACKGROUND_COLOR, "#000000");
          group.setStyle(Q.Styles.ALPHA, "0.3");
        }
        group.setStyle(Q.Styles.GROUP_STROKE_LINE_DASH, [3, 2]);
        group.setStyle(Q.Styles.LABEL_COLOR, "#fff");
        group.setStyle(Q.Styles.LABEL_FONT_SIZE, 30);
        // group.setStyle(Q.Styles.GROUP_BACKGROUND_COLOR, false);
        // group.setStyle(Q.Styles.ALPHA, "0");
        group.nodesType = "VMC";
        return group;
      }
      function createEdge(a, b, color, dashed, name) {
        var edge = graph.createEdge(name, a, b);
        if (dashed) {
          edge.setStyle(Q.Styles.EDGE_LINE_DASH, [8, 5]);
        }
        edge.setStyle(Q.Styles.EDGE_WIDTH, 5);
        edge.setStyle(Q.Styles.EDGE_COLOR, "#8cd1f1");
        edge.setStyle(Q.Styles.ARROW_TO, false);
        edge.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
        edge.setStyle(Q.Styles.ARROW_FROM, Q.Consts.SHAPE_CIRCLE);
        edge.setStyle(Q.Styles.ARROW_TO, Q.Consts.SHAPE_CIRCLE);
        edge.nodesType = "line";
        // edge.setStyle(Q.Styles.EDGE_OUTLINE_STYLE, "#0F0");
        return edge;
      }
      function createBus(path) {
        var bus = new Q.Bus(null, path);
        graph.graphModel.add(bus);
        bus.setStyle(Q.Styles.SHAPE_STROKE, 25);
        bus.setStyle(Q.Styles.SHAPE_FILL_COLOR, "#fff");
        // bus.image = server;
        return bus;
      }
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

      // createText("公共事业服务中心\n网络拓扑", 859, 100, 40, "#F00");

      // var bus = createBus();
      // bus.moveTo(-600, 460);
      // bus.lineTo(-600, 540);
      // bus.moveTo(-600, 500);
      // bus.lineTo(600, 500);
      // var node = graph.createNode(name, x, y);
      // node.image = Q.Shapes.getShape(Q.Consts.SHAPE_RECT, 30, 15);
      // node.setStyle(Q.Styles.SHAPE_FILL_COLOR, "#888");

      var vmc1 = createGroup(50, false, imgvmc, null, "vmc1");
      // var obc1 = createGroup(20, false, exchange2, vmc1, "obc1");
      // var obc2 = createGroup(20, false, exchange2, vmc1, "obc2");
      // var obc3 = createGroup(20, false, exchange2, vmc1, "obc3");
      // var obc4 = createGroup(20, false, exchange2, vmc1, "obc4");
      var cpu = createNode(imgcpu, -120, 100, "cpu", vmc1, true, "CPU");
      var gpu = createNode(imggpu, -270, 100, "gpu", vmc1, true, "GPU");
      var dsp = createNode(imgdsp, -420, 100, "dsp", vmc1, true, "DSP");
      var fpga = createNode(imgfgpa, -570, 100, "fpga", vmc1, true, "FPGA");
      // vmc1.addChild(obc1);
      // vmc1.addChild(obc2);
      // vmc1.addChild(obc3);
      // vmc1.addChild(obc4);

      var vmc2 = createGroup(50, false, imgvmc, null, "vmc2");
      // var obc1_ = createGroup(20, false, exchange2, vmc2, "obc1");
      // var obc2_ = createGroup(20, false, exchange2, vmc2, "obc2");
      // var obc3_ = createGroup(20, false, exchange2, vmc2, "obc3");
      // var obc4_ = createGroup(20, false, exchange2, vmc2, "obc4");
      var cpu_ = createNode(imgcpu, 120, 900, "cpu", vmc2, true, "CPU");
      var gpu_ = createNode(imggpu, 270, 900, "gpu", vmc2, true, "GPU");
      var dsp_ = createNode(imgdsp, 420, 900, "dsp", vmc2, true, "DSP");
      var fpga_ = createNode(imgfgpa, 570, 900, "fpga", vmc2, true, "FPGA");
      // vmc2.addChild(obc1_);
      // vmc2.addChild(obc2_);
      // vmc2.addChild(obc3_);
      // vmc2.addChild(obc4_);

      var exchange1_ = createSw(imgsw, -345, 350, "Switch1", null, true);
      var exchange2_ = createSw(imgsw, 345, 700, "Switch2", null, true);
      var exchange3_ = createSw(imgsw, -300, 700, "Switch3", null, true);
      var exchange4_ = createSw(imgsw, 370, 350, "Switch4", null, true);
      var exchange5_ = createSw(imgsw, 0, 525, "centerSwitch", null, true);
      var edge = createEdge(exchange1_, exchange5_);
      var edge_ = createEdge(exchange2_, exchange5_);
      var edge3_ = createEdge(exchange5_, exchange3_);
      var edge4_ = createEdge(exchange5_, exchange4_);
      // edge.angle = Math.PI / 2;
      // edge_.angle = Math.PI / 2;
      // edge3_.angle = Math.PI / 2;
      // edge4_.angle = Math.PI / 2;

      var edgevc11 = createEdge(cpu, exchange1_);
      var edgevc12 = createEdge(gpu, exchange1_);
      var edgevc13 = createEdge(dsp, exchange1_);
      var edgevc14 = createEdge(fpga, exchange1_);
      // edgevc11.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
      // edgevc12.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
      // edgevc13.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
      // edgevc14.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;

      var edgevc11_ = createEdge(cpu_, exchange2_);
      var edgevc12_ = createEdge(gpu_, exchange2_);
      var edgevc13_ = createEdge(dsp_, exchange2_);
      var edgevc14_ = createEdge(fpga_, exchange2_);
      // edgevc11_.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
      // edgevc12_.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
      // edgevc13_.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;
      // edgevc14_.edgeType = Q.Consts.EDGE_TYPE_VERTICAL_HORIZONTAL;

      var h1 = createSw(imgop1, -150, 900, "远置单元1", null, true);
      var h2 = createSw(imgop2, -300, 900, "远置单元2", null, true);
      var h3 = createSw(imgop1, -450, 900, "远置单元3", null, true);
      var edgevc1_1 = createEdge(exchange3_, h1);
      var edgevc1_2 = createEdge(exchange3_, h2);
      var edgevc1_3 = createEdge(exchange3_, h3);

      var h4 = createSw(imgop1, 150, 100, "远置单元4", null, true);
      var h5 = createSw(imgop2, 300, 100, "远置单元5", null, true);
      var h6 = createSw(imgop1, 450, 100, "远置单元6", null, true);
      var edgevc2_1 = createEdge(exchange4_, h4);
      var edgevc2_2 = createEdge(exchange4_, h5);
      var edgevc2_3 = createEdge(exchange4_, h6);

      var flowingSupport = new FlowingSupport(graph);
      flowingSupport.addFlowing(edge, 1, false, this.flowColor);
      flowingSupport.addFlowing(edge_, 1, false, this.flowColor);
      flowingSupport.addFlowing(edge3_, 1, false, this.flowColor);
      flowingSupport.addFlowing(edge4_, 1, false, this.flowColor);

      flowingSupport.addFlowing(edgevc11, 1, false);
      flowingSupport.addFlowing(edgevc12, 1, false);
      flowingSupport.addFlowing(edgevc13, 1, false);
      flowingSupport.addFlowing(edgevc14, 1, false);

      flowingSupport.addFlowing(edgevc11_, 1, false);
      flowingSupport.addFlowing(edgevc12_, 1, false);
      flowingSupport.addFlowing(edgevc13_, 1, false);
      flowingSupport.addFlowing(edgevc14_, 1, false);

      flowingSupport.addFlowing(edgevc1_1, 1, false, this.flowColor_vpn);
      flowingSupport.addFlowing(edgevc1_2, 1, false, this.flowColor_vpn);
      flowingSupport.addFlowing(edgevc1_3, 1, false, this.flowColor_vpn);

      flowingSupport.addFlowing(edgevc2_1, 1, false, this.flowColor_vpn);
      flowingSupport.addFlowing(edgevc2_2, 1, false, this.flowColor_vpn);
      flowingSupport.addFlowing(edgevc2_3, 1, false, this.flowColor_vpn);

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
            var alarmSeverity = AlarmSeverity.random;
            element.alarmLabel =
              "" +
              (1 + Q.randomInt(100)) +
              alarmSeverity.sortName +
              (Q.randomBool() ? "+" : "");
            element.alarmColor = alarmSeverity.color;
          }
        });
        timer = setTimeout(A, 1500);
      }, 3000);

      /// 销毁
      function destroy() {
        flowingSupport.stop();
        clearTimeout(timer);
      }
      graph.zoomToOverview(0.5);
      // graph.isMovable = false;
      graph.enableWheelZoom = false;
      graph.onclick = function (evt) {
        if (evt.getData()) {
          if (evt.getData().nodesType) {
            that.deviceType = evt.getData().nodesType;
            that.drawer = true;
            that.devicename = evt.getData().name;
          }
        }
      };
    },
    dealWithDataXY() {
      let incomeSwForOp = [
        { name: "a1", id: 10011 },
        { name: "b1", id: 10021 },
        { name: "c1", id: 10031 },
        { name: "d1", id: 10041 },
        { name: "e1", id: 10051 },
      ];
      let incomeSwForCal = [
        { name: "a", id: 1001 },
        { name: "b", id: 1002 },
        { name: "c", id: 1003 },
        { name: "d", id: 1004 },
        { name: "e", id: 1005 },
      ];
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
            baseRightTopY = baseRightTopY + 150;
            return {
              ...item,
              x: baseRightTopX,
              y: baseRightTopY,
            };
          } else {
            baseRightBottomX = baseRightBottomX + 300;
            baseRightBottomY = baseRightBottomY - 150;
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
            x: newarr[indexs].x,
            y: newarr[indexs].y + 100,
          };
        } else {
          if (newarr[indexs].y >= 0) {
            return {
              ...items,
              x: newarr[indexs].x + 100,
              y: newarr[indexs].y - 80,
            };
          } else {
            return {
              ...items,
              x: newarr[indexs].x + 100,
              y: newarr[indexs].y + 80,
            };
          }
        }
      });
      let assistantPoints = [];
      newarr.map((itemNow, indexNow) => {
        if (itemNow.x <= 0) {
          let tempointx = itemNow.x - 25;
          let tempointy = itemNow.y;
          assistantPoints.push(
            {
              no: `${indexNow}_1`,
              type: "busNode",
              x: tempointx - 30,
              y: tempointy + 100,
            },
            {
              no: `${indexNow}_2`,
              type: "busNode",
              x: tempointx + 30,
              y: tempointy,
            }
          );
        } else {
          let tempointx = itemNow.x + 50;
          let tempointy;
          if (itemNow.y >= 0) {
            tempointy = itemNow.y - 40;
          } else {
            tempointy = itemNow.y + 40;
          }
          assistantPoints.push(
            {
              no: `${indexNow}_1`,
              type: "busNode",
              x: tempointx - 160,
              y: tempointy,
            },
            {
              no: `${indexNow}_2`,
              type: "busNode",
              x: tempointx + 160,
              y: tempointy,
            }
          );
        }
      });
      console.log(newarr, newop, assistantPoints);
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
    drawBusEdgeAndCenterSw(pointArr, graph) {
      let arrleft = [];
      let arrrightTop = [];
      let arrrightBottom = [];
      pointArr.map((item) => {
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

      // var edge = graph.createEdge(name, a, b);
      // if (dashed) {
      //   edge.setStyle(Q.Styles.EDGE_LINE_DASH, [8, 5]);
      // }
      // edge.setStyle(Q.Styles.EDGE_WIDTH, 5);
      // edge.setStyle(Q.Styles.EDGE_COLOR, "#8cd1f1");
      // edge.setStyle(Q.Styles.ARROW_TO, false);
      // edge.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
      // edge.setStyle(Q.Styles.ARROW_FROM, Q.Consts.SHAPE_CIRCLE);
      // edge.setStyle(Q.Styles.ARROW_TO, Q.Consts.SHAPE_CIRCLE);
      // edge.nodesType = "line";
      // // edge.setStyle(Q.Styles.EDGE_OUTLINE_STYLE, "#0F0");
      // return edge;
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
        this.dealWithDataXY();
      }, 500);
    }, 2000);
  },

  created() {
    this.getNameOAll();
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
    width: 40%;
    margin-left: 25%;
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