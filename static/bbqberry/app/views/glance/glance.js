'use strict';

angular.module('bbqberry.glance', ['ngRoute'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/glance', {
            templateUrl: 'views/glance/glance.html',
            controller: 'GlanceController'
        })
    }])

    .controller('GlanceController', ['$scope', '$http', function ($scope, $http) {
        $http.get('/api/v1/temperatures/probes').then(function(response) {
            $scope.probes = response.data;
        });
    }]);