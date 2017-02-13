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
                var myPoller = poller.get('/api/v1/temperatures/probes', {
                    action: 'get',
                    delay: 1000
                });

                myPoller.promise.then(null, null, function (response) {
                    for (var i = 0; i < $scope.probes.length; i++) {
                        $scope.probes[i].reading = response.data[i];
                    }
                    $scope.carouselActive = true;
                });
            };

            var setupProbeSlide = function(d3, probe) {
                var limits = probe.tempLimits;
                var minAbs = limits.minAbsCelsius;
                var maxWarn = limits.maxWarnCelsius;
                var maxAbs = limits.maxAbsCelsius;
                var colorStep = 1;
                var gradStep = 50;
                var steps = (maxAbs - minAbs) / colorStep;

                var grads = d3.scale.linear()
                    .range([minAbs, maxAbs])
                    .clamp(true)
                    .interpolate(d3.interpolateRound);

                var subZero = d3.scale.linear()
                    .range(["#0000ff", "#00ff00"])
                    .interpolate(d3.interpolateHcl);
                var color = d3.scale.linear()
                    .range(["#00ff00", "#FF5F05"])
                    .interpolate(d3.interpolateHcl);

                var guage = {};
                probe.guage = guage;
                guage.ranges = [];

                for (var i = 0; i <= 1; i += 1.0 / steps) {
                    var mn = grads(i);
                    var mx = colorStep + mn;

                    var c = color(i);
                    if (mx < 0) {
                        c = subZero(i * 5);
                    }
                    if (mn > maxWarn) {
                        c = "#FF0000"
                    }

                    guage.ranges[guage.ranges.length] = {
                        min: mn,
                        max: mx,
                        color: c
                    };
                }
                guage.lowerLimit = minAbs;
                guage.upperLimit = maxAbs;
                guage.majorGraduations = ((maxAbs - minAbs) / gradStep) + 1;
            };

            var getHardwareConfig = function() {
                return $q(function(resolve, reject) {
                    $http({
                        method: 'GET',
                        url: '/api/v1/hardware/config'
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