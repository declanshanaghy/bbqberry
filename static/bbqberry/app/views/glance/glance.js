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
                color: '#DEDEDE'
            },
            {
                min: 100,
                max: 200,
                color: '#8DCA2F'
            },
            {
                min: 200,
                max: 300,
                color: '#FDC702'
            },
            {
                min: 300,
                max: 400,
                color: '#FF7700'
            },
            {
                min: 400,
                max: 500,
                color: '#e82d2b'
            },
            {
                min: 500,
                max: 600,
                color: '#C50200'
            }
        ];

        var myPoller = poller.get('/api/v1/temperatures/probes', {
            action: 'get',
            delay: 5000
        });

        myPoller.promise.then(null, null, function(response) {
            $scope.probes = response['data'];
        });

    }]);