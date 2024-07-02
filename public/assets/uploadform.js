/**
  * File Upload Form Stuff
  *   - handlers file upload
  */

document.addEventListener('DOMContentLoaded', () => {
  const fileInput = document.querySelector("#dropzone-file")
  const fileInputLabel = document.querySelector('#selected-file');

  fileInput.addEventListener("change", (e) => {
    updateFileLabel(e.target);
  });

  function updateFileLabel(input) {
    const fileName = input.files[0].name;
    // Update label text to show file name
    fileInputLabel.innerHTML = `
            <p class="mb-2 text-sm text-gray-400">
                File selected: <span class="font-semibold">${fileName}</span>
            </p>
    `;
  }
});
