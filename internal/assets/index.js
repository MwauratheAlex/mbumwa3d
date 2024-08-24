import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';
import { STLLoader } from 'three/examples/jsm/loaders/STLLoader.js';

let animationFrameId = null;
let scene, camera, renderer, controls, loader, currentMesh;

function initThreeJS() {
	const canvas = document.querySelector("#viewer");
	const canvasDiv = document.querySelector("#canvas-div");
	const height = canvasDiv.offsetHeight;
	const width = canvasDiv.offsetWidth;

	// scene
	scene = new THREE.Scene();

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
	camera = new THREE.PerspectiveCamera(75, width / height, 0.1, 1000);
	camera.position.z = 60;
	scene.add(camera);

	// renderer
	renderer = new THREE.WebGLRenderer({ canvas, alpha: true });
	renderer.setClearColor(0xa2d2ff, 0.03); // the default
	renderer.setSize(width, height);
	renderer.setPixelRatio(2);
	renderer.render(scene, camera);

	// controls
	controls = new OrbitControls(camera, canvas);
	controls.enableDamping = true;
	controls.enablePan = false;
	controls.enableZoom = false;
	controls.autoRotate = true;
	controls.autoRotateSpeed = 4;


	loader = new STLLoader();
	currentMesh = null;

	loop();
}

const fileInput = document.querySelector("#dropzone-file");
const fileUploadContainer = document.getElementById("file-upload-container");

if (fileInput && !fileInput.dataset.listenerAttached) {
	fileInput.addEventListener("change", (e) => {

		const filename = e.target.files[0].name
		const selectedFileLabel = document.getElementById("selected-file");
		selectedFileLabel.innerHTML = `File: ${filename}`;

		fileUploadContainer.classList.add("hidden");

		const file = e.target.files[0]
		if (file) {
			loadModel(file)
		}
	});
	fileInput.dataset.listenerAttached = true;
}

const uploadNewFileBtn = document.getElementById("upload-new-file");
uploadNewFileBtn.addEventListener("click", () => clearScene());

function clearScene() {
	if (currentMesh) {
		scene.remove(currentMesh);
		currentMesh = null;
	}

	if (fileInput) {
		fileInput.value = "";
	}
	fileUploadContainer.classList.remove("hidden");
	document.getElementById("selected-file").innerHTML = "";

	const canvas = document.querySelector("#viewer");
	const uploadNewFileBtn = document.getElementById("upload-new-file");

	canvas.classList.add("hidden");
	uploadNewFileBtn.classList.add("hidden");
}

function loadModel(file) {
	const canvas = document.querySelector("#viewer");
	const uploadNewFileBtn = document.getElementById("upload-new-file");

	canvas.classList.remove("hidden");
	uploadNewFileBtn.classList.remove("hidden");


	const reader = new FileReader()
	reader.onload = (e) => {
		const content = e.target.result;
		const geometry = loader.parse(content);
		addToScene(geometry);
	};
	reader.readAsArrayBuffer(file);
}

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

function loop() {
	controls.update();
	renderer.render(scene, camera);
	animationFrameId = window.requestAnimationFrame(loop);
}

// Function to stop the animation
function stopThreeJS() {
	if (animationFrameId !== null) {
		window.cancelAnimationFrame(animationFrameId);
		animationFrameId = null;
	}
}

document.addEventListener("DOMContentLoaded", () => {
	console.log("anim initialized")
	if (document.querySelector("#viewer")) {
		initThreeJS();
	}
});

// Reinitialize Three.js on HTMX content swap
document.body.addEventListener("htmx:beforeSwap", (event) => {
	if (event.detail.target.id === "content-container"
		&& document.querySelector("#viewer")) {
		stopThreeJS(); // Stop any existing animation loop
		console.log("stopping 2js");
	}
});

// Reinitialize Three.js on HTMX content swap
document.body.addEventListener("htmx:afterSwap", (event) => {
	if (event.detail.target.id === "content-container"
		&& document.querySelector("#viewer")) {
		initThreeJS();
	}
});


// print form stuff
const toastNotification = document.getElementById("toast-notification");
function showToastNotification(message, type) {
	toastNotification.innerHTML = `
	 <div 
       class="daisy-alert daisy-alert-${type} w-min transition-all flex items-center py-1 rounded-md text-center text-sm">
       <span class="text-gray-black tracking-wider">
         ${message}
       </span>
	 </div>
  `;
	setTimeout(
		() => {
			toastNotification.innerHTML = ""
		},
		3000
	);
}

function capitalize(s) {
	return s[0].toUpperCase() + s.slice(1);
}

function beforePostPrint(event) {
	if (!Boolean(fileInput.value)) {
		event.preventDefault();
		showToastNotification("An STL File is required", "error");
		return;
	}

	const form = event.target;
	const formData = new FormData(form);
	for (const [key, value] of formData.entries()) {
		if (value === "Please select") {
			event.preventDefault();
			showToastNotification(`${capitalize(key)} is required`, "error");
			return;
		}
	}
}

window.beforePostPrint = beforePostPrint;



(function() {
	document.body.addEventListener("print-error", (e) => {
		const message = e.detail.message;

		switch (message) {
			case "unauthorized":
				showToastNotification("Please login to continue", "error");
				break;
		}
	});
})();
