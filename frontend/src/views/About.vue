<template>
  <div class="about">
    <div id="chart"></div>
    <div id="time-log"></div>
  </div>
</template>
<script>
// import * as THREE from "three";
// import satellite from "satellite";
import Globe from "globe.gl";
import { FBXLoader } from "three/examples/jsm/loaders/FBXLoader";
// import satellite_ from "@/assets/lib/satellite.fbx";
export default {
  name: "about",
  data() {
    return {
      data_: "",
    };
  },
  components: {},
  mounted() {
    // const EARTH_RADIUS_KM = 6371; // km
    // const SAT_SIZE = 100; // km
    // const TIME_STEP = 3 * 1000; // per frame

    // const timeLogger = document.getElementById("time-log");

    const world = Globe()(document.getElementById("chart"))
      .globeImageUrl(
        "//unpkg.com/three-globe/example/img/earth-blue-marble.jpg"
      )
      .objectLat("lat")
      .objectLng("lng")
      .objectAltitude("alt")
      .objectLabel("name");

    setTimeout(() => world.pointOfView({ altitude: 3.5 }));

    const loader = new FBXLoader();
    loader.load("/model/eam_026_AR.fbx", function (object) {
      // loader.load("/model/weixing.fbx", function (object) {
      let satGeometry = object;
      console.log(satGeometry);
      satGeometry.position.set(250, 0, 80);
      satGeometry.scale.set(0.1, 0.1, 0.1);
      satGeometry.rotation.set(0, 1, 0);
      // satGeometry.quaternion.set(0, 0, 1, 0);
      // const satMaterial = new THREE.MeshLambertMaterial({
      //   color: "palegreen",
      //   transparent: true,
      //   opacity: 0.7,
      // });
      world.scene().add(satGeometry);
    });

    world.controls().autoRotate = true;
    world.controls().autoRotateSpeed = 1.8;
  },
};
</script>
