'use strict';

angular.module('bbqberry.glance', ['d3', 'ngRadialGauge', 'ngRoute', 'emguo.poller'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/glance', {
            templateUrl: 'views/glance/glance.html',
            controller: 'GlanceController'
        })
    }])

    .controller('GlanceController', ['$scope', '$http', 'poller', 'd3Service',
        function ($scope, $http, poller, d3Service) {
            $scope.interval = 0;
            $scope.noWrapSlides = false;
            $scope.active = false;
            $scope.probes = [];

            d3Service.d3().then(function (d3) {
                var min = 0;
                var max = 800;
                var colorStep = 1;
                var gradStep = 100;
                var steps = (max - min) / colorStep;

                var grads = d3.scale.linear()
                    .range([min, max])
                    .clamp(true)
                    .interpolate(d3.interpolateRound);

                var color = d3.scale.linear()
                    .range(["#00ff00", "#ff0000"])
                    .interpolate(d3.interpolateHcl);

                $scope.ranges = [];
                var pos = 0;
                for (var i = 0; i <= 1; i += 1.0 / steps) {
                    var mn = grads(i);
                    var mx = colorStep + mn;
                    var c = color(i);
                    $scope.ranges[$scope.ranges.length] = {
                        min: mn,
                        max: mx,
                        color: c
                    };
                }
                $scope.lowerLimit = min;
                $scope.upperLimit = $scope.ranges[$scope.ranges.length - 1].max;
                $scope.majorGraduations = ((max - min) / gradStep) + 1;

                var myPoller = poller.get('/api/v1/temperatures/probes', {
                    action: 'get',
                    delay: 1000
                });

                var lastActive = 0;
                myPoller.promise.then(null, null, function (response) {
                    $scope.probes = response['data'];
                    for (var i=0; i<$scope.probes.length; i++) {
                        $scope.probes[i].index = i;
                        if ( i == lastActive ) {
                            $scope.probes[i].active = true;
                        }
                    }

                    $scope.active = true;
                });
            });
        }]);