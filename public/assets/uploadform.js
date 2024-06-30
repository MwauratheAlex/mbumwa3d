/**
  * File Upload Form Stuff
  *   - handlers file upload
  */

document.addEventListener('DOMContentLoaded', () => {
  const paymentForm = document.querySelector("#payment-form")
  const uploadform = document.querySelector("#upload-form")
  const showPaymentFormBtn = document.querySelector("#show-payment-btn")
  const fileInput = document.querySelector("#dropzone-file")
  const fileUploadErrors = document.querySelector("#file-upload-errors")
  const backBtn = document.querySelector("#back-button")
  const fileInputLabel = document.querySelector('#dropzone-label');

  showPaymentFormBtn.addEventListener("click", () => {
    if (!fileInput.files.length) {
      fileUploadErrors.innerHTML = "Please select a file to continue"
      return;
    }
    fileUploadErrors.innerHTML = "";

    paymentForm.style.display = "flex";
    uploadform.style.display = "none";
  });

  backBtn.addEventListener("click", () => {
    paymentForm.style.display = "none";
    uploadform.style.display = "flex";
  });

  fileInput.addEventListener("change", (e) => {
    updateFileLabel(e.target);
  });


  function updateFileLabel(input) {
    const fileName = input.files[0].name;

    // Update label text to show file name
    fileInputLabel.innerHTML = `
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
        </div>
    `;
  }

});
