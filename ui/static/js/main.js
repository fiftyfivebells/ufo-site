window.addEventListener('DOMContentLoaded', (event) => {

  // Add event listeners to the radio buttons on the create sighting form. If the
  // radio is No, hide UFO specific details. Otherwise, show them.
  if (document.querySelector("input[name='sighting']")) {
    let ufo = document.getElementById("ufoDetails");

    document.querySelectorAll("input[name='sighting']").forEach((e) => {
      if (e.value === "No" && e.checked) {
        ufo.style.display = "none";
      }

      if (e.value === "Yes" && e.checked) {
        ufo.style.display = "block";
      }
      e.addEventListener("change", function(event) {
        let item = event.target.value;
        if (item === "Yes") {
          ufo.style.display = "block";
        } else if (item === "No") {
          ufo.style.display = "none";
        }
      });
    });
  }
});
