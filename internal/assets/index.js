import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';
import { STLLoader } from 'three/examples/jsm/loaders/STLLoader.js';

function initThreeJS() {
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


  // ambient light for softer overall illumination
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
  let currentMesh = null;

  function addToScene(geometry) {
    geometry.computeBoundingBox();
    const bbox = geometry.boundingBox;
    const center = new THREE.Vector3();
    bbox.getCenter(center)
    geometry.translate(-center.x, -center.y, -center.z);

    const material = new THREE.MeshStandardMaterial({ color: 0x00ff83, roughness: 0.0001 });
    const mesh = new THREE.Mesh(geometry, material);

    const pivot = new THREE.Object3D();
    pivot.add(mesh)

    if (currentMesh) {
      scene.remove(currentMesh);
    }
    currentMesh = pivot;
    scene.add(pivot);
  }

  function loadInitialModel() {
    loader.load('public/models/pen_holder.stl',
      (geometry) => addToScene(geometry),
      (xhr) => {
        console.log(xhr.loaded / xhr.total * 100 + '% loaded');
      },
      (error) => {
        console.log(error);
      }
    );
  }
  // load model
  function loadModel(file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      const content = e.target.result;
      const geometry = loader.parse(content);
      addToScene(geometry);
    };
    reader.readAsArrayBuffer(file);
  }

  // event listener
  const fileInput = document.querySelector("#dropzone-file")
  fileInput.addEventListener("change", (e) => {
    const file = e.target.files[0]
    if (file) {
      loadModel(file)
    }
  });

  loadInitialModel();
  function loop() {
    controls.update();
    renderer.render(scene, camera);
    window.requestAnimationFrame(loop);
  }

  loop();
}

initThreeJS()

