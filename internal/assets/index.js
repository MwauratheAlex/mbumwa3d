import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';
import { STLLoader } from 'three/examples/jsm/loaders/STLLoader.js';

const height = 610
const width = 890

// scene
const scene = new THREE.Scene();

// sphere
const geometry = new THREE.SphereGeometry(3, 64, 64);
const material = new THREE.MeshStandardMaterial({ color: 0x00ff83, roughness: 0.01 });
const mesh = new THREE.Mesh(geometry, material);
//scene.add(mesh);
//

// light
const light = new THREE.PointLight(0xfeccb5, 100, 100);
light.position.set(0, 10, 10);
scene.add(light);


// Optional: Add an ambient light for softer overall illumination
const ambientLight = new THREE.AmbientLight(0x404040, 0.8); // soft white light
scene.add(ambientLight);

// camera
const camera = new THREE.PerspectiveCamera(75, width / height, 0.1, 1000);
camera.position.z = 60;
scene.add(camera);


// renderer
const canvas = document.querySelector("#viewer");
const renderer = new THREE.WebGLRenderer({ canvas, alpha: true });
renderer.setClearColor(0xa2d2ff, 0.03); // the default
renderer.setSize(width, height);
renderer.setPixelRatio(2);
renderer.render(scene, camera);
// mesh.position.x = 3;
// mesh.position.y = 1;

// controls
const controls = new OrbitControls(camera, canvas);
controls.enableDamping = true;
controls.enablePan = false;
controls.enableZoom = false;
controls.autoRotate = true;
controls.autoRotateSpeed = 5;

const loader = new STLLoader();
loader.load('public/models/pen_holder.stl',
  function(geometry) {
    geometry.computeBoundingBox();
    const bbox = geometry.boundingBox;
    const center = new THREE.Vector3();
    bbox.getCenter(center)
    geometry.translate(-center.x, -center.y, -center.z);

    const material = new THREE.MeshStandardMaterial({ color: 0x00ff83, roughness: 0.01 });
    const mesh = new THREE.Mesh(geometry, material);


    scene.add(mesh);
  },
  (xhr) => {
    console.log(xhr.loaded / xhr.total * 100 + '% loaded');
  },
  (error) => {
    console.log(error);
  }
);


// resize
function loop() {
  controls.update();
  renderer.render(scene, camera);
  window.requestAnimationFrame(loop);
}

loop();
