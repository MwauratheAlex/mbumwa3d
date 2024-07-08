/**
  * Dashboard navigation Btns
  *  - script changes opacity
  */
function toggleSelected(button, navBtns) {
  for (let i = 0; i < navBtns.length; i++) {
    const btn = navBtns[i];
    btn.classList.remove("bg-opacity-100");
    btn.classList.add("bg-opacity-30");
  }

  button.classList.remove("bg-opacity-30");
  button.classList.add("bg-opacity-100")
}

(function() {
  const dashPopup = document.querySelector("#dash-popup");
  document.body.addEventListener("dash-pop", (e) => {
    dashPopup.innerHTML = e.detail.value
    dashPopup.classList.remove("hidden")
    dashPopup.classList.add("flex")

    setTimeout(() => {
      dashPopup.innerHTML = "Order selected successfully"
      dashPopup.classList.remove("flex")
      dashPopup.classList.add("hidden")
    }, 3000)
  });

  const dashNavBtns = document.getElementsByClassName("dash-nav-btn");
  for (let i = 0; i < dashNavBtns.length; i++) {
    const btn = dashNavBtns[i];
    btn.addEventListener("click", (e) => toggleSelected(
      e.currentTarget, dashNavBtns)
    );
  }
})();
