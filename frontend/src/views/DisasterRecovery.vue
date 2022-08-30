/* eslint-disable */
<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <!-- 123 -->
    <div v-if="loading" class="loadingbox">
      <div data-loader="jumping"></div>
    </div>
    <div v-else class="disasterrecovery">
      <!-- <h3>首页</h3> -->
      <el-row
        type="flex"
        class="row-bg"
        justify="space-around"
        style="width: 100%; height: 60%"
      >
        <el-col :span="12" style="width: 100%; height: 100%">
          <div style="width: 100%; height: 100%">
            <div class="topboxforcanvas">
              <p class="title">
                <span>迁移视图</span>
              </p>
              <div class="boxforcanvas" ref="canvas"></div>
            </div>
          </div>
        </el-col>
        <el-col :span="12" style="width: 100%; height: 100%">
          <div style="width: 100%; height: 60%">
            <div class="topboxforcanvas">
              <p class="title">
                <span>相关指标</span>
              </p>
              <div class="content" style="width: 100%; height: 80%">
                <el-carousel :interval="5000" class="carousel" trigger="click">
                  <el-carousel-item v-for="item in vmcNowArr" :key="item.id">
                    <div class="canvasbox" :id="`linebox_${item.id}`">
                      <selflineNewless
                        :inref="`linebox_${item.id}`"
                        :indexNow="item.id"
                      />
                    </div>
                  </el-carousel-item>
                </el-carousel>
              </div>
              <!-- <div class="boxforcanvas" ref="canvas"></div> -->
            </div>
          </div>
          <div style="width: 100%; height: 40%">
            <div class="topboxforcanvas_3">
              <p class="title_3">
                <span>迁移日志</span>
              </p>
              <p class="logp">
                [迁移开始]：容错迁移任务, From: {{this.from.name}}, To: {{this.to.name}}
              </p>
              <p class="logp">[迁移开始时间]: {{this.from.time}}</p>
              <p class="logp">==============================================</p>
              <p class="logp">
                [迁移完成]：容错迁移任务, From: {{this.from.name}}, To: {{this.to.name}}
              </p>
              <p class="logp">[迁移完成时间]: {{this.from.to}}</p>
              <!-- <div class="boxforcanvas" ref="canvas"></div> -->
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row
        type="flex"
        class="row-bg"
        justify="space-around"
        style="width: 100%; height: 40%"
      >
        <el-col :span="24" style="width: 100%; height: 100%">
          <div style="width: 100%; height: 100%">
            <div class="topboxforcanvas_">
              <p class="title_2">
                <!-- <span
                  >{{ `${Nowindex.name || "未知名称"}-${Nowindex.id}` }}相关任务
                </span> -->
                <span>关联任务</span>
              </p>
              <div class="content">
                <el-carousel
                  :interval="5000"
                  class="carousel"
                  trigger="click"
                  @change="changeCarousle"
                >
                  <el-carousel-item
                    v-for="item in vmcNowArr"
                    :key="item.id"
                    style="height: 100%"
                  >
                    <div style="height: 100%">
                      <TableNow
                        style="height: 90%"
                        :key="item.id"
                        :indexNow="item.id"
                      />
                    </div>
                  </el-carousel-item>
                </el-carousel>
              </div>
              <!-- <div class="boxforcanvas" ref="canvas"></div> -->
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
// import "@/lib/loaders.min.css";
import { mapState } from "vuex";
import selflineNewless from "@/components/selflineNewless.vue";
import TableNow from "@/components/TableNow.vue";
import imgvmc from "@/assets/newpng/VMC.svg";
import imgcpu from "@/assets/newpng/CPU.svg";
import imggpu from "@/assets/newpng/GPU.svg";
import imgdsp from "@/assets/newpng/DSP.svg";
import imgfgpa from "@/assets/newpng/FPGA.svg";
import imgsw from "@/assets/newpng/centersw_topo.svg";
import imgop1 from "@/assets/newpng/op_1.svg";
import imgop2 from "@/assets/newpng/op_2.svg";
import imgop3 from "@/assets/newpng/op_3.svg";
import imgop4 from "@/assets/newpng/op_4.svg";
import inswsvg from "@/assets/newpng/sw_in.svg";
import messages from "@/assets/newpng/send_.svg";
import { getTopoShow, filterName } from "@/api";
export default {
  name: "DisasterRecovery",
  components: {
    // HelloWorld,
    selflineNewless,
    TableNow,
  },
  data() {
    return {
      speed: "",
      loading: true,
      flaga: true,
      flagb: true,
      tableData: [],
      vmcedge: {},
      areaedge: {},
      Nowindex: "1",
      vmcs: [],
      vmcNowArr: [],
      vmcNode: {},
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
    };
  },
  computed: {
    ...mapState(["disVmc", "disArea", "from", "to"]),
  },
  methods: {
    filterName,
    async getNameOAll() {
      const { data } = await getTopoShow();
      if (data.code == 0) {
        let tempdata = data.data.node
          .filter((item) => {
            return item.device_type == "vmc";
          })
          .sort((a, b) => {
            return a.id - b.id;
          });
        this.vmcs = tempdata;
        let tempArr = [];
        tempdata.map((item) => {
          if (
            this.from.id === item.id ||
            this.from.parent_id === item.id ||
            this.to.id === item.id ||
            this.to.parent_id === item.id
          ) {
            tempArr.push(item);
          }
        });
        this.vmcNowArr = tempArr;
        this.Nowindex = tempdata[0];
      }
    },
    changeCarousle(before, now) {
      this.Nowindex = this.vmcs[now];
    },
    randomRange(min, max) {
      // min最小值，max最大值
      return Math.floor(Math.random() * (max - min)) + min;
    },
    creatGraph() {
      let graph = new Q.Graph(this.$refs.canvas);
      graph.editable = true;
      graph.enableRectangleSelectionByRightButton = true;
      graph.tooltipDelay = 0;
      graph.tooltipDuration = 10000;
      return graph;
    },
    createFlow(graph) {
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
            var ui = new Q.ImageUI(messages);
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
            scope.timer = setTimeout(A, 100);
          }, 500);
        },
      };
      return FlowingSupport;
    },
    createNode(graph, image, x, y, name, group, randomFlag) {
      var node = graph.createNode(name, x, y);
      if (image) {
        if (Q.isString(image)) {
          image = image;
        }
        node.image = image;
      }
      node.size = { height: 40 };
      if (group) {
        group.addChild(node);
      }
      node.randomAble = randomFlag || false;
      node.setStyle(Q.Styles.LABEL_COLOR, "#ffffff");
      node.setStyle(Q.Styles.LABEL_FONT_SIZE, 16);
      node.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_TOP);
      node.setStyle(Q.Styles.LABEL_ANCHOR_POSITION, Q.Position.CENTER_BOTTOM);
      // node.setStyle(Q.Styles., 25);
      // model.add(node);
      return node;
    },
    createEdge(graph, a, b, color, dashed, name, errType) {
      var edge = graph.createEdge(name, a, b);
      if (dashed) {
        edge.setStyle(Q.Styles.EDGE_LINE_DASH, [8, 5]);
      }
      edge.setStyle(Q.Styles.LABEL_COLOR, "#ffffff");
      edge.setStyle(Q.Styles.EDGE_WIDTH, 2);
      edge.setStyle(Q.Styles.EDGE_COLOR, "#2f6da0");
      edge.setStyle(Q.Styles.ARROW_TO, false);
      edge.tooltip = `当前线路：${name}\n 当前速率：${
        this.speed
      }MB/S\n 当前迁移类型：${
        errType === "vmc" ? "整机迁移" : "分区迁移"
      }\n 迁移开始时间：${
        errType === "vmc"
          ? new Date(parseInt(this.disVmc.time)).toLocaleString()
          : new Date(parseInt(this.disArea.time)).toLocaleString()
      }`;
      return edge;
    },
    createEdegUi(graph) {
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
          var to = toBounds.center;
          path.curveTo(from.x, from.y, 0, 200, to.x, to.y);
        },
      };
      Q.extend(VPNFlexEdgeUI, Q.EdgeUI);
      return VPNFlexEdgeUI;
    },
    drawall() {
      const that = this;
      const graph = this.creatGraph();
      const FlowingSupport = this.createFlow(graph);
      const VPNFlexEdgeUI = this.createEdegUi(graph);
      var flowingSupport = new FlowingSupport(graph);
      if (this.from.type && this.to.type) {
        if (this.from.parent_id === this.to.parent_id && this.to.parent_id) {
          let endarr = that.vmcs.map((items, index) => {
            if (items.id === this.from.parent_id) {
              let tempnode = that.createNode(
                graph,
                imgvmc,
                index * 100,
                (index % 2) * 80,
                that.filterName(items.name),
                null,
                true
              );
              that.vmcNode[items.id] = tempnode;
            }
          });
        } else if (
          this.from.parent_id !== this.to.parent_id &&
          this.from.parent_id &&
          this.to.parent_id
        ) {
          let endarr = that.vmcs.map((items, index) => {
            if (
              items.id === this.from.parent_id ||
              items.id === this.to.parent_id
            ) {
              let tempnode = that.createNode(
                graph,
                imgvmc,
                index * 100,
                (index % 2) * 80,
                that.filterName(items.name),
                null,
                true
              );
              that.vmcNode[items.id] = tempnode;
            }
          });
        }
        if (this.from.type !== "vmc") {
          if (this.from.parent_id === this.to.parent_id) {
            let fromNode = that.createNode(
              graph,
              this.imgObg[this.from.type],
              this.vmcNode[this.from.parent_id].x - 80,
              this.vmcNode[this.from.parent_id].y + 100,
              that.filterName(this.from.name),
              null,
              true
            );
            let toNode = that.createNode(
              graph,
              this.imgObg[this.to.type],
              this.vmcNode[this.to.parent_id].x + 80,
              this.vmcNode[this.to.parent_id].y + 100,
              that.filterName(this.to.name),
              null,
              true
            );
            fromNode.setStyle(
              Q.Styles.LABEL_POSITION,
              Q.Position.CENTER_BOTTOM
            );
            fromNode.setStyle(
              Q.Styles.LABEL_ANCHOR_POSITION,
              Q.Position.CENTTER_TOP
            );
            toNode.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_BOTTOM);
            toNode.setStyle(
              Q.Styles.LABEL_ANCHOR_POSITION,
              Q.Position.CENTTER_TOP
            );
            var edge1_ = this.createEdge(
              graph,
              this.vmcNode[this.from.parent_id],
              fromNode,
              null,
              false,
              "",
              ""
            );
            var edge1_1 = this.createEdge(
              graph,
              this.vmcNode[this.to.parent_id],
              toNode,
              null,
              false,
              "",
              ""
            );
            edge1_.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
            edge1_1.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
            var edge1 = this.createEdge(
              graph,
              fromNode,
              toNode,
              null,
              true,
              "任务迁移流",
              "vmc"
            );
            flowingSupport.addFlowing(edge1, 1, false);
          } else {
            let fromNode = that.createNode(
              graph,
              this.imgObg[this.from.type],
              this.vmcNode[this.from.parent_id].x,
              this.vmcNode[this.from.parent_id].y + 100,
              that.filterName(this.from.name),
              null,
              true
            );
            let toNode = that.createNode(
              graph,
              this.imgObg[this.to.type],
              this.vmcNode[this.to.parent_id].x,
              this.vmcNode[this.to.parent_id].y + 100,
              that.filterName(this.to.name),
              null,
              true
            );
            fromNode.setStyle(
              Q.Styles.LABEL_POSITION,
              Q.Position.CENTER_BOTTOM
            );
            fromNode.setStyle(
              Q.Styles.LABEL_ANCHOR_POSITION,
              Q.Position.CENTTER_TOP
            );
            toNode.setStyle(Q.Styles.LABEL_POSITION, Q.Position.CENTER_BOTTOM);
            toNode.setStyle(
              Q.Styles.LABEL_ANCHOR_POSITION,
              Q.Position.CENTTER_TOP
            );
            var edge1_ = this.createEdge(
              graph,
              this.vmcNode[this.from.parent_id],
              fromNode,
              null,
              false,
              "",
              ""
            );
            var edge1_1 = this.createEdge(
              graph,
              this.vmcNode[this.to.parent_id],
              toNode,
              null,
              false,
              "",
              ""
            );
            edge1_.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
            edge1_1.edgeType = Q.Consts.EDGE_TYPE_ELBOW;
            var edge1 = this.createEdge(
              graph,
              fromNode,
              toNode,
              null,
              true,
              "任务迁移流",
              "vmc"
            );
            flowingSupport.addFlowing(edge1, 1, false);
          }
        } else {
          if (this.from.id !== this.to.id) {
            let fromNode = that.createNode(
              graph,
              this.imgObg[this.from.type],
              -80,
              0,
              that.filterName(this.from.name),
              null,
              true
            );
            let toNode = that.createNode(
              graph,
              this.imgObg[this.to.type],
              80,
              0,
              that.filterName(this.to.name),
              null,
              true
            );
            var edge1 = this.createEdge(
              graph,
              fromNode,
              toNode,
              null,
              true,
              "任务迁移流",
              "vmc"
            );
            flowingSupport.addFlowing(edge1, 1, false);
          } else {
            let fromNode = that.createNode(
              graph,
              this.imgObg[this.from.type],
              -80,
              0,
              that.filterName(this.from.name),
              null,
              true
            );
            var edge1 = this.createEdge(
              graph,
              fromNode,
              fromNode,
              null,
              true,
              "任务迁移流",
              "vmc"
            );
            flowingSupport.addFlowing(edge1, 1, false);
          }
        }
      }
      graph.zoomToOverview({}, 1.4);
      graph.callLater(function () {
        flowingSupport.start();
      });
      var timer = setInterval(() => {
        if (edge1) {
          edge1.tooltip = `当前线路: 整机迁移任务1002\n 当前速率：${that.randomRange(
            128,
            1024
          )}MB/S\n 当前迁移类型：整机迁移\n 迁移开始时间：${new Date(
            parseInt(that.disVmc)
          ).toLocaleString()}`;
        }
        const nowDate = new Date().valueOf();
        if (that.from.time) {
          if (nowDate - that.from.time >= 1000 * 60 * 3 + 1000 * 45) {
            if (this.flaga) {
              this.$message({
                message: "整机迁移已完成，共耗时3分45秒",
                type: "success",
              });
              graph.removeElement(edge1);
              // that.tableData = [[], [...singletable], [...singletable]];
              this.flaga = false;
            }
          }
        }
      }, 1000);
      /// 销毁
      function destroy() {
        flowingSupport.stop();
        clearInterval(timer);
      }
    },
  },
  mounted() {
    setTimeout(() => {
      this.loading = false;
      setTimeout(() => {
        this.getNameOAll();
        this.drawall();
      }, 500);
      this.speed = this.randomRange(1, 10);
      this.getNameOAll();
    }, 3000);
  },
  created() {},
};
</script>
<style lang="less" scoped>
.disasterrecovery {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  overflow: hidden;
  .title {
    display: block;
    height: 15%;
    width: 100%;
    color: azure;
    background: url("../assets/png/part_title.png") no-repeat center;
    background-size: 100% 2rem;
    padding-left: 12%;
    margin: 1%;
    box-sizing: border-box;
    span {
      display: inline-block;
      height: 100%;
      width: 90%;
      font-size: 1.3rem;
      padding-top: -1%;
    }
  }
  .title_1 {
    display: block;
    height: 15%;
    width: 100%;
    color: azure;
    background: url("../assets/png/part_title.png") no-repeat center;
    background-size: 100% 2rem;
    padding-left: 12%;
    margin: 1%;
    box-sizing: border-box;
    span {
      display: inline-block;
      height: 100%;
      width: 90%;
      font-size: 1.3rem;
      padding-top: 0.3%;
    }
  }
  .title_2 {
    display: block;
    height: 15%;
    width: 100%;
    color: azure;
    background: url("../assets/png/part_title.png") no-repeat center;
    background-size: 100% 2rem;
    padding-left: 8%;
    // margin: 1%;
    box-sizing: border-box;
    span {
      display: inline-block;
      height: 100%;
      width: 90%;
      font-size: 1.3rem;
      padding-top: 0.3%;
    }
  }
  .title_3 {
    display: block;
    height: 25%;
    width: 100%;
    color: azure;
    background: url("../assets/png/part_title.png") no-repeat center;
    background-size: 100% 2rem;
    padding-left: 12%;
    // margin: 1%;
    box-sizing: border-box;
    span {
      display: inline-block;
      height: 100%;
      width: 90%;
      font-size: 1.3rem;
      padding-top: 0.3%;
    }
  }
  .content {
    height: 80%;
    width: 92%;
    padding-left: 2.5%;
    .carousel {
      width: 100%;
      height: 100%;
      .el-carousel__arrow {
        background-color: rgba(48, 228, 228, 0.3) !important;
      }
      .canvasbox {
        width: 100%;
        height: 70%;
      }
    }
  }
  .topboxforcanvas_ {
    margin-top: 24px;
    padding-top: 0.5%;
    width: 100%;
    height: 90%;
    background: url("../assets/newpng/tableborder_.png") no-repeat center;
    background-size: 104% 100%;
    box-sizing: border-box;
    background-position-x: -15px;
    overflow: hidden;
    .title {
      height: 8%;
      line-height: 8%;
    }
    .boxforcanvas {
      height: 92%;
      width: 100%;
    }
  }
  .topboxforcanvas {
    padding: 1%;
    width: 100%;
    height: 100%;
    background: url("../assets/newpng/tableborder_.png") no-repeat center;
    background-size: 100% 100%;
    box-sizing: border-box;
    // .title {
    //   height: 8%;
    //   line-height: 8%;
    // }
    .boxforcanvas {
      height: 80%;
      width: 100%;
    }
  }
  .topboxforcanvas_3 {
    padding-top: 0px;
    width: 100%;
    height: 100%;
    background: url("../assets/newpng/tableborder_.png") no-repeat center;
    background-size: 100% 100%;
    box-sizing: border-box;
    // .title {
    //   height: 8%;
    //   line-height: 8%;
    // }
    .boxforcanvas {
      height: 92%;
      width: 100%;
    }
    .logp {
      padding-left: 24px;
      color: aliceblue;
      margin: 0;
    }
  }
  .middleboxfortask {
    padding: 0.5%;
    width: 100%;
    height: 50%;
    background: url("../assets/newpng/tableborder_.png") no-repeat center;
    background-size: 100% 100%;
    box-sizing: border-box;
  }
  .bottomboxforline {
    padding: 0.5%;
    width: 100%;
    height: 50%;
    background: url("../assets/newpng/tableborder_.png") no-repeat center;
    background-size: 100% 100%;
    box-sizing: border-box;
  }
}
</style>