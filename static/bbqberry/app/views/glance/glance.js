'use strict';

angular.module('bbqberry.glance', ['ngRoute', 'emguo.poller'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/glance', {
            templateUrl: 'views/glance/glance.html',
            controller: 'GlanceController'
        })
    }])

    .controller('GlanceController', ['$scope', '$http', 'poller', function ($scope, $http, poller) {
        $http.get('/api/v1/temperatures/probes').then(function(response) {
            $scope.probes = response.data;
        });

        var myPoller = poller.get('/api/v1/temperatures/probes', {
            action: 'get',
            delay: 1000
        });

        myPoller.promise.then(null, null, function(response) {
            $scope.probes = response['data'];
        });

    }]);