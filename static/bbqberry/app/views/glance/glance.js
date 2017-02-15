'use strict';

angular.module('bbqberry.glance', ['d3', 'ngRadialGauge', 'ngRoute', 'emguo.poller'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/glance', {
            templateUrl: 'views/glance/glance.html',
            controller: 'GlanceController'
        })
    }])

    .controller('GlanceController', ['$q', '$scope', '$http', 'poller', 'd3Service',
        function ($q, $scope, $http, poller, d3Service) {
            var pollProbeData = function() {
                var myPoller = poller.get('/api/temperatures', {
                    action: 'get',
                    delay: 250
                });

                myPoller.promise.then(null, null, function (response) {
                    for (var i = 0; i < $scope.probes.length; i++) {
                        $scope.probes[i].reading = response.data[i];
                    }
                    $scope.carouselActive = true;
                });
            };

            var setupProbeSlide = function(d3, probe) {
                var freeze = celsiusToFahrenheit(0);

                var limits = probe.tempLimits;
                var minAbs = celsiusToFahrenheit(limits.minAbsCelsius);
                if (minAbs < 0)
                    minAbs = 0;
                var maxWarn = celsiusToFahrenheit(limits.maxWarnCelsius);
                var maxAbs = roundUpTo100(celsiusToFahrenheit(limits.maxAbsCelsius));
                var colorStep = 1;
                var gradStep = 100;
                var steps = ((maxWarn - minAbs) / colorStep) - 1;

                var grads = d3.scale.linear()
                    .range([minAbs, maxWarn])
                    .clamp(true)
                    .interpolate(d3.interpolate);

                // d3.scale.linear().range([0, 1]).domain([minLimit, maxLimit]);
                var subZero = d3.scale.linear()
                    .range(["#0000FF", "#00FF00"])
                    .interpolate(d3.interpolateHcl);
                var color = d3.scale.linear()
                    .range(["#00ff00", "#FF0000"])
                    .interpolate(d3.interpolateHcl);

                var guage = {};
                probe.guage = guage;
                probe.reading = {};
                probe.reading.fahrenheit = 0;
                guage.ranges = [];

                for (var i=0, j=0; i <= 1; i+=(1.0 / steps), j++) {
                    var mn = minAbs + (j * colorStep);
                    var mx = mn + colorStep;

                    var c = color(i);
                    if (mn < freeze) {
                        c = subZero(i * (maxWarn - freeze) / (freeze - minAbs));
                    }
                    // else if ( mn == freeze || mn == freeze ) {
                    //     c = "#000000";
                    // }
                    guage.ranges[guage.ranges.length] = {
                        min: mn,
                        max: mx,
                        color: c,
                        stroke: c,
                        needleColor: c,
                        needleStroke: "#FFFFFF"
                    };
                }

                // guage.ranges[guage.ranges.length - 1] = {
                //     min: guage.ranges[guage.ranges.length - 1].max,
                //     max: guage.ranges[guage.ranges.length - 1].max + 1,
                //     color: "#000000",
                //     stroke: "#FF0000",
                //     needleColor: guage.ranges[guage.ranges.length - 1].needleColor,
                //     needleStroke: guage.ranges[guage.ranges.length - 1].needleStroke
                // };
                guage.ranges[guage.ranges.length - 1] = {
                    min: guage.ranges[guage.ranges.length - 1].max,
                    max: maxAbs,
                    color: "url(#CrossHatch)",
                    stroke: color(1),
                    needleColor: "#FF0000",
                    needleStroke: "#FFFFFF"
                };

                guage.lowerLimit = minAbs;
                guage.upperLimit = maxAbs;
                guage.majorGraduations = ((maxAbs - minAbs) / gradStep) + 1;
                guage.minorGraduations = 10;
            };

            var getHardwareConfig = function() {
                return $q(function(resolve, reject) {
                    $http({
                        method: 'GET',
                        url: '/api/hardware'
                    }).then(function successCallback(response) {
                        $scope.probes = response['data'].probes;
                        for (var i = 0; i < $scope.probes.length; i++) {
                            $scope.probes[i].index = i;
                        }
                        resolve();
                    }, function errorCallback(response) {
                        reject(response['data']);
                    });
                });
            };

            $scope.carouselInterval = 0;
            $scope.carouselNoWrap = false;
            $scope.carouselActive = false;
            $scope.activeSlide = 0;
            $scope.probes = [];

            $scope.swipeLeft = function() {
                var curr = $scope.activeSlide + 1;
                if ( curr > $scope.probes.length - 1 ) {
                    if ( $scope.carouselNoWrap )
                        curr = $scope.probes.length - 1;
                    else
                        curr = 0;
                }
                $scope.activeSlide = curr;
            };
            $scope.swipeRight = function() {
                var curr = $scope.activeSlide - 1;
                if ( curr < 0 ) {
                    if ( $scope.carouselNoWrap )
                        curr = 0;
                    else
                        curr = $scope.probes.length - 1;
                }
                $scope.activeSlide = curr;
            };

            d3Service.d3().then(function (d3) {
                getHardwareConfig().then(function() {
                        for (var i = 0; i < $scope.probes.length; i++) {
                            setupProbeSlide(d3, $scope.probes[i]);
                        }
                        pollProbeData();
                    },
                    function (data) {
                        $scope.$parent.$glblwarning = data;
                    }
                )
            });
        }]);