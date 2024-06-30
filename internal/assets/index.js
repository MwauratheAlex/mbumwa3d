import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';
import { STLLoader } from 'three/examples/jsm/loaders/STLLoader.js';


document.addEventListener("login_success", () => {
  console.log("Login success, hello world")
})

const height = 610
const width = 890

// scene
const scene = new THREE.Scene();

// light
const light = new THREE.DirectionalLight(0xfeccb5, 0.9);
light.position.set(0, 2, 20);
scene.add(light);

const directionalLight2 = new THREE.DirectionalLight(0xfeccb5, 0.8);
directionalLight2.position.set(0, -2, -20);
scene.add(directionalLight2);


// Optional: Add an ambient light for softer overall illumination
const ambientLight = new THREE.AmbientLight(0x404040, 0.7); // soft white light
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

// controls
const controls = new OrbitControls(camera, canvas);
controls.enableDamping = true;
controls.enablePan = false;
controls.enableZoom = false;
controls.autoRotate = true;
controls.autoRotateSpeed = 4;


const loader = new STLLoader();
loader.load('public/models/pen_holder.stl',
  function(geometry) {
    geometry.computeBoundingBox();
    const bbox = geometry.boundingBox;
    const center = new THREE.Vector3();
    bbox.getCenter(center)
    geometry.translate(-center.x, -center.y, -center.z);

    const material = new THREE.MeshStandardMaterial({ color: 0x00ff83, roughness: 0.0001 });
    const mesh = new THREE.Mesh(geometry, material);

    const pivot = new THREE.Object3D();
    pivot.add(mesh)

    scene.add(pivot);
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

function updateFileLabel(input) {
  const label = document.getElementById('dropzone-label');
  const fileName = input.files[0].name;

  // Update label text to show file name
  label.innerHTML = `
                <div class="flex flex-col items-center justify-center pt-5 pb-6">
                    <svg class="w-8 h-8 mb-4 text-gray-500"
                         aria-hidden="true"
                         xmlns="http://www.w3.org/2000/svg"
                         fill="none"
                         viewBox="0 0 20 16">
                        <path stroke="currentColor"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                              stroke-width="2"
                              d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"></path>
                    </svg>
                    <p class="mb-2 text-sm text-gray-400">
                        File selected: <span class="font-semibold">${fileName}</span>
                    </p>
                    <p class="text-xs text-gray-500">STL Files</p>
                </div>`;
}


