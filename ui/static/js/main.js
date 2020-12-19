window.addEventListener('DOMContentLoaded', (event) => {

  // Add event listeners to the selection bar on the sightings page
  if (document.getElementById("sightingState")) {
    let stateSelector = document.getElementById("sightingState");

    getAllSightings();

    stateSelector.addEventListener("change", function(event) {
      let state = stateSelector.value;
      displayListOfSightings(state);
    });
  }

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

const sightings = []

// Get JSON from database and return the Promise of the items
async function getJson() {
  const data = await fetch("/sightings/all");
  const json = await data.json();
  return json;
}

async function getAllSightings() {
  getJson().then(d => {
    for (let x of d) {
      sightings.push(x);
    }
  });
}

function displayListOfSightings(state) {
  let displayData;

  if (state !== "") {
    const abbrev = getKeyByValue(states, state);
    displayData = sightings.filter(o => o.state === abbrev);
  } else {
    displayData = sightings.slice();
  }

  const table = document.getElementById("sightingTable");
  table.removeChild(table.getElementsByTagName("tbody")[0]);
  const tbody = document.createElement("tbody");

  
  for (let s of displayData) {
    const row = tbody.insertRow();

    const index = row.insertCell();
    const link = document.createElement("a");
    link.setAttribute("href", "/sighting/" + s.index);
    link.innerHTML = s.index;
    index.appendChild(link);

    const date = row.insertCell();
    date.innerHTML = s.datetime;

    const city = row.insertCell();
    city.innerHTML = s.city;

    const state = row.insertCell();
    state.innerHTML = states[s.state];

    const shape = row.insertCell();
    shape.innerHTML = s.shape.Valid ? s.shape.String : "";

    const duration = row.insertCell();
    duration.innerHTML = s.duration.Valid ? s.duration.Int64 : "";
  }

  table.appendChild(tbody);
}

// Takes in an object and value, then returns the key for that value
// in the object, if it exists
function getKeyByValue(obj, value) {
  return Object.keys(obj).find(key => obj[key] === value);
}

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
