'use strict';

// Declare app level module which depends on views, and components
angular.module('bbqberry', [
    // 3rd Party modules
    'd3',
    'emguo.poller',
    'ngAnimate',
    'ngRoute',
    'ngTouch',
    'ngRadialGauge',
    'ui.bootstrap',

    // BBQBerry Components
    'bbqberry.nav',

    //BBQBErry Views
    'bbqberry.glance',
    'bbqberry.view1',
    'bbqberry.view2',
    'bbqberry.version'
])

.config(['$locationProvider', '$routeProvider', function ($locationProvider, $routeProvider) {
    $locationProvider.hashPrefix('!');

    $routeProvider.otherwise({redirectTo: '/glance'});
}])

;
