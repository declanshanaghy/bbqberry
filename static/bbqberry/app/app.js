'use strict';

// Declare app level module which depends on views, and components
angular.module('bbqberry', [
    // 3rd Party modules
    'd3',
    'emguo.poller',
    'ngAnimate',
    'ngRoute',
    'ngTouch',
    'ui.bootstrap',
    'mi.AlertService',
    'ngRadialGauge',
    'rzModule',        // Slider
    'mi.AlertService',

    // BBQBerry Components
    'bbqberry.nav',

    //BBQBErry Views
    'bbqberry.overview',
    'bbqberry.version'
])
.constant('ALERT_LEVELS', {
    danger: {timeout: 10000},
    warning: {timeout: 5000},
    success: {timeout: 3000},
    info: {timeout: 3000}
})

.config(['$locationProvider', '$routeProvider', function ($locationProvider, $routeProvider) {
    $locationProvider.hashPrefix('!');

    $routeProvider.otherwise({redirectTo: '/overview'});
}])

;
