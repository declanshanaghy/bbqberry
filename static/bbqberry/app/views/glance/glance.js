'use strict';

angular.module('bbqberry.glance', ['ngRadialGauge', 'ngRoute', 'emguo.poller'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/glance', {
            templateUrl: 'views/glance/glance.html',
            controller: 'GlanceController'
        })
    }])

    .controller('GlanceController', ['$scope', '$http', 'poller', function ($scope, $http, poller) {
        $scope.ranges = [
            {
                min: 0,
                max: 100,
                color: '#22FF44'
            },
            {
                min: 100,
                max: 200,
                color: '#41DF3F'
            },
            {
                min: 200,
                max: 300,
                color: '#61BF3A'
            },
            {
                min: 300,
                max: 400,
                color: '#80A035'
            },
            {
                min: 400,
                max: 500,
                color: '#A08030'
            },
            {
                min: 500,
                max: 600,
                color: '#BF612B'
            },
            {
                min: 600,
                max: 700,
                color: '#DF4126'
            },
            {
                min: 700,
                max: 800,
                color: '#FF2222'
            }
        ];
        $scope.upperLimit = $scope.ranges[$scope.ranges.length - 1].max;
        $scope.majorGraduations = $scope.ranges.length;

        var myPoller = poller.get('/api/v1/temperatures/probes', {
            action: 'get',
            delay: 5000
        });

        myPoller.promise.then(null, null, function(response) {
            $scope.probes = response['data'];
        });

    }]);