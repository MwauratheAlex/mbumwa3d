import { initThreeJS } from '/public/assets/index.js';
document.addEventListener('htmx:afterSwap', (event) => {
  if (event.detail.target.id === 'content-container') {
    initThreeJS();
    loadInitialModel();
    setupFileInput();
  }
});
initThreeJS(); // Initial load
loadInitialModel(); // Initial load
setupFileInput(); // Initial load
