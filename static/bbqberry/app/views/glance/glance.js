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
            $scope.carouselInterval = 0;
            $scope.carouselNoWrap = false;
            $scope.carouselActive = false;
            $scope.activeSlide = 0;
            $scope.probes = [];

            $scope.swipeLeft = function() {
                // console.log('swipe left: curr=' + $scope.activeSlide);
                var curr = $scope.activeSlide + 1;
                if ( curr > $scope.probes.length - 1 )
                    curr = $scope.probes.length - 1;
                $scope.activeSlide = curr;
                // console.log('swipe left: new=' + $scope.activeSlide);
            };
            $scope.swipeRight = function() {
                // console.log('swipe right: curr=' + $scope.activeSlide);
                var curr = $scope.activeSlide - 1;
                if ( curr < 0 )
                    curr = 0;
                $scope.activeSlide = curr;
                // console.log('swipe right: new=' + $scope.activeSlide);
            };

            d3Service.d3().then(function (d3) {
                var min = 0;
                var warn = 600;
                var max = 800;
                var colorStep = 1;
                var gradStep = 100;
                var steps = (max - min) / colorStep;

                var grads = d3.scale.linear()
                    .range([min, max])
                    .clamp(true)
                    .interpolate(d3.interpolateRound);

                var subZero = d3.scale.linear()
                    .range(["#0000ff", "#00ff00"])
                    .interpolate(d3.interpolateHcl);
                var color = d3.scale.linear()
                    .range(["#00ff00", "#FF5F05"])
                    .interpolate(d3.interpolateHcl);

                $scope.ranges = [];

                for (var i = 0; i <= 1; i += 1.0 / steps) {
                    var mn = grads(i);
                    var mx = colorStep + mn;

                    var c = color(i);
                    if ( mx < 32 ) {
                        c = subZero(i * 25);
                    }
                    if ( mn > warn ) {
                        c = "#FF0000"
                    }

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

                myPoller.promise.then(null, null, function (response) {
                    // console.log(response['data']);
                    $scope.probes = response['data'];
                    // $scope.probes[0].warning = "High temperature limit exceeded: actual=378 °C > threshold=360 °C";
                    $scope.carouselActive = true;
                    // console.log("Probe " + lastActive + " is active");
                });
            });
        }]);