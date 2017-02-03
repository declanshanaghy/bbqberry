'use strict';

// Declare app level module which depends on views, and components
angular.module('bbqberry', [
    // 3rd Party modules
    'emguo.poller',
    'ngRoute',
    'ui.bootstrap',

    // BBQBerry Components
    'bbqberry.nav',
    'bbqberry.tempguage',

    //BBQBErry Views
    'bbqberry.glance',
    'bbqberry.view1',
    'bbqberry.view2',
    'bbqberry.version'
])

.config(['$locationProvider', '$routeProvider', function ($locationProvider, $routeProvider) {
    $locationProvider.hashPrefix('!');

    $routeProvider.otherwise({redirectTo: '/view1'});
}])

;
