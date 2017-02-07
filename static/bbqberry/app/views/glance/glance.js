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

        // var grads = d3.scale.linear()
        //     .range([0, 800])
        //     .interpolate(d3.interpolateNumber);
        //
        // var min = 0;
        // var max = 800;
        // var step = 100;
        // var steps = (max - min) / step;
        //
        // var color = d3.scale.linear()
        //     .range(["darkblue", "red"])
        //     .interpolate(d3.interpolateHcl);
        //
        // $scope.ranges = [];
        // var pos = 0;
        // for (var i = min; i < max; i += step) {
        //     pos = i / (max - min);
        //     var mn = grads(pos);
        //     var mx = mn + step;
        //     var c = color(pos);
        //     console.log("mn=" + mn + ", mx=" + mx + ", c=" + c);
        //     $scope.ranges[$scope.ranges.length] = {
        //         min: mn,
        //         max: mx,
        //         color: c
        //     };
        // }
        $scope.upperLimit = $scope.ranges[$scope.ranges.length - 1].max;
        $scope.majorGraduations = $scope.ranges.length + 1;

        var myPoller = poller.get('/api/v1/temperatures/probes', {
            action: 'get',
            delay: 5000
        });

        myPoller.promise.then(null, null, function (response) {
            $scope.probes = response['data'];
        });
    }]);