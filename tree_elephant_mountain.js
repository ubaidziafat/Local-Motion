// Local Motion.js

// Step 1: Initialize Global Variables

let currentLocation;
let userInput;
let destination;
let route;

// Step 2: Function to get user's current location

function getLocation() {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(function(position) {
        currentLocation = {
          lat: position.coords.latitude,
          lng: position.coords.longitude
       };
     });
   }
}

// Step 3: Function to get user's destination

function getDestination() {
   userInput = prompt("Please enter your destination's address: ");
   destination = userInput;
}

// Step 4: Function to calculate the route

function calculateRoute() {
   let directionsService = new google.maps.DirectionsService();
   let start = currentLocation;
   let end = destination;
   let request = {
      origin: start,
      destination: end,
      travelMode: 'DRIVING'
   };

   directionsService.route(request, function(response, status) {
     if (status == 'OK') {
        route = response.routes[0].overview_path;
     }
  });
}

// Step 5: Function to draw the route

function drawRoute() {
   let directionsDisplay = new google.maps.DirectionsRenderer();
   let mapProp = {
      center: currentLocation,
      zoom: 15
   };
   let map = new google.maps.Map(document.getElementById("map"), mapProp);
   directionsDisplay.setMap(map);
   directionsDisplay.setDirections({routes: [{
      overview_path: route
   }]});
}

// Step 6: Function to initiate navigation

function startNavigate() {
   getLocation();
   getDestination();
   calculateRoute();
   drawRoute();
}

// Step 7: Call startNavigate() function

startNavigate();