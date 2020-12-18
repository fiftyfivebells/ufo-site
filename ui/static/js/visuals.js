window.addEventListener("DOMContentLoaded", (event) => {

  getSightings();

});

const data = [];

// Get JSON from database and return the Promise of the items
async function getJson() {
  const data = await fetch("/sightings/all");
  const json = await data.json();
  return json;
}

// Fill the data array with the items from the database
async function getSightings() {
  getJson().then(d => {
    for (let x of d) {
      data.push(x);
    }
    stateBarChart();
    seasonBarChart();
    latLongChart();
  });
}

// Create scatter plot of lat and long
async function latLongChart() {
  const points = await getLatAndLong(data);
  const margin = {
    top: 10,
    right: 30,
    bottom: 30,
    left: 60,
  };
  const width = 460 - margin.left - margin.right;
  const height = 400 - margin.top - margin.bottom;

  const svg = d3.select("#latLongChart")
        .append("svg")
        .attr("width", width + margin.left + margin.right)
        .attr("height", height + margin.top + margin.bottom)
        .append("g")
        .attr("transform",
              "translate(" + margin.left + ", " + margin.top + ")");


  const x = d3.scaleLinear()
        .domain([-200, 150])
        .range([0, width]);
  svg.append("g")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(x));

  const y = d3.scaleLinear()
        .domain([-60, 90])
        .range([height, 0]);
  svg.append("g")
    .call(d3.axisLeft(y));

  svg.append("g")
    .selectAll("dot")
    .data(points)
    .enter()
    .append("circle")
    .attr("cx", function (d) { return x(d.long);})
    .attr("cy", function (d) { return y(d.lat);})
    .attr("r", 1.5)
    .style("fill", "#69b3a2");

}

// Make bar chart of all the states
async function stateBarChart() {
  const stateData = await countStates(data);
  const ctx = document.getElementById("stateChart").getContext("2d");
  const stateChart = new Chart(ctx, {
    type: "horizontalBar",
    data: {
      labels: Object.keys(stateData),
      datasets: [{
        label: "# of Sightings",
        data: Object.values(stateData),
        borderWidth: 1,
      }]
    },
  });
}

// Make bar chart of sightings by season
async function seasonBarChart() {
  const seasonData = await countSeasons(data);
  const ctx = document.getElementById("seasonChart").getContext("2d");
  const seasonChart = new Chart(ctx, {
    type: "bar",
    data: {
      labels: Object.keys(seasonData),
      datasets: [{
        label: "# of Sightings",
        data: Object.values(seasonData),
        borderWidth: 1,
      }]
    },
    options: {
      scales: {
        yAxes: [{
          ticks: {
            min: 10000
          }
        }]
      }
    }
  });
}

// Get a total count for every season in the dataset
function countSeasons(data) {
  const seasons = {
    "Winter": 0,
    "Spring": 0,
    "Summer": 0,
    "Fall": 0,
  };

  for (let s of data) {
    if (s.season === "fall") {
      seasons.Fall += 1;
    } else if (s.season === "winter") {
      seasons.Winter += 1;
    } else if (s.season === "spring") {
      seasons.Spring += 1;
    } else if (s.season === "summer") {
      seasons.Summer += 1;
    }
  }

  return seasons;
}

// Get a total count for each state in the dataset
function countStates(data) {
  const stateCount = {};

  for (let s of data) {
    if (states[s.state] in stateCount) {
      stateCount[states[s.state]] += 1;
    } else {
      stateCount[states[s.state]] = 1;
    }
  }

  return stateCount;
}

// Get latitude and longitude as a pair of lists
function getLatAndLong(data) {
  const points = [];

  for (let s of data) {
    points.push({long: s.long, lat: s.lat});
  }

  return points;
}
