'use strict';

var temperature = require('./temperature.js');

temperature.getTemperature().then(
    function(data) {
        console.log("Success", data);
    },
    function(err) {
        console.log("Error", err);
    }
);
