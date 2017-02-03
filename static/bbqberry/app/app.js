'use strict';

// Declare app level module which depends on views, and components
angular.module('bbqberry', [
    'ngRoute',
    'ui.bootstrap',
    'bbqberry.nav',
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
