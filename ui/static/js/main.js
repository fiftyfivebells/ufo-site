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


const states = {
  "al": "Alabama",
  "ak": "Alaska",
  "az": "Arizona",
  "ar": "Arkansas",
  "ca": "California",
  "co": "Colorado",
  "ct": "Connecticut",
  "de": "Delaware",
  "dc": "DC",
  "fl": "Florida",
  "ga": "Georgia",
  "hi": "Hawaii",
  "id": "Idaho",
  "il": "Illinois",
  "in": "Indiana",
  "ia": "Iowa",
  "ks": "Kansas",
  "ky": "Kentucky",
  "la": "Louisiana",
  "me": "Maine",
  "md": "Maryland",
  "ma": "Massachusetts",
  "mi": "Michigan",
  "mn": "Minnesota",
  "ms": "Mississippi",
  "mo": "Missouri",
  "mt": "Montana",
  "ne": "Nebraska",
  "nv": "Nevada",
  "nc": "North Carolina",
  "nh": "New Hampshire",
  "nj": "New Jersey",
  "nm": "New Mexico",
  "ny": "New York",
  "nd": "North Dakota",
  "oh": "Ohio",
  "ok": "Oklahoma",
  "or": "Oregon",
  "pa": "Pennsylvania",
  "pr": "Puerto Rico",
  "ri": "Rhode Island",
  "sc": "South Carolina",
  "sd": "South Dakota",
  "tn": "Tennessee",
  "tx": "Texas",
  "ut": "Utah",
  "vt": "Vermont",
  "va": "Virginia",
  "wa": "Washington",
  "wv": "West Virginia",
  "wi": "Wisconsin",
  "wy": "Wyoming",
};
