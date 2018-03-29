'use strict';

angular.module('bbqberry.overview', ['ngRoute', 'ui.bootstrap'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/overview', {
            templateUrl: 'views/overview/overview.html',
            controller: 'OverviewCtrl'
        });
    }])

    .controller('OverviewCtrl', ['$q', '$scope', '$http', 'poller', '$log',
        function ($q, $scope, $http, $poller, $log) {

            $scope.isNavCollapsed = true;
            $scope.isCollapsed = false;
            $scope.isCollapsedHorizontal = false;

            var floor = 0;
            var ceil = 250;
            var interval = 50;
            var nTicks = ((ceil - floor) / interval) + 1;
            var ticksArray = Array.apply(null, {length: nTicks}).map(function(value, index){
                return Math.floor(floor + (index * interval));
            });

            var getHardwareConfig = function() {
                return $q(function(resolve, reject) {
                    $http({
                        method: 'GET',
                        url: '/api/hardware'
                    }).then(function successCallback(response) {
                        resolve(response.data.probes);
                    }, function errorCallback(response) {
                        reject(response.data);
                    });
                });
            };

            getHardwareConfig($scope).then(
                function(probes) {
                    $scope.probes = probes;
                    for ( var i=0; i<$scope.probes.length; i++ ) {
                        $scope.probes[i].index = i;
                    }

                    $scope.probe1Rng = newRangeSlider(probes[0]);
                    $scope.probe1Val = newTempIndicator(probes[0]);

                    $scope.probe2Rng = newRangeSlider(probes[1]);
                    $scope.probe2Val = newTempIndicator(probes[1]);

                    $scope.probe3Rng = newRangeSlider(probes[2]);
                    $scope.probe3Val = newTempIndicator(probes[2]);

                    $scope.probe4Rng = newRangeSlider(probes[3]);
                    $scope.probe4Val = newTempIndicator(probes[3]);

                    pollProbeData();
                },
                function (data) {
                    $log.error("ERROR", data);
                    // $scope.$parent.$glblwarning = data;
                }
            );

            var pollProbeData = function() {
                var myPoller = $poller.get('/api/temperatures', {
                    action: 'get',
                    delay: 1000
                });

                myPoller.promise.then(null, null, function (response) {
                    if ( $scope.probes[0].enabled ) {
                        $scope.probe1Val.value = response.data[0].celsius;
                        $scope.probes[0].data = response.data[0];
                    }
                    if ( $scope.probes[1].enabled ) {
                        $scope.probe2Val.value = response.data[1].celsius;
                        $scope.probes[1].data = response.data[1];
                    }
                    if ( $scope.probes[2].enabled ) {
                        $scope.probe3Val.value = response.data[2].celsius;
                        $scope.probes[2].data = response.data[2];
                    }
                    if ( $scope.probes[3].enabled ) {
                        $scope.probe4Val.value = response.data[3].celsius;
                        $scope.probes[3].data = response.data[3];
                    }
                });
            };

            var newRangeSlider = function(probe) {
                return {
                    minValue: probe.limits.minWarnCelsius,
                    maxValue: probe.limits.maxWarnCelsius,
                    options: {
                        id: probe.index,
                        disabled: !probe.enabled,
                        floor: floor,
                        ceil: ceil,
                        vertical: true,
                        noSwitching: true,
                        showSelectionBar: true,
                        showTicksValues: true,
                        draggableRange: true,
                        ticksArray: ticksArray,
                        onStart: function(id, low, high, type) {
                            probe.stash = {
                                high: high,
                                low: low
                            };
                        },
                        onEnd: function(id, low, high, type) {
                            $http({
                                method: 'PUT',
                                url: '/api/monitors',
                                data: {
                                    "max": high,
                                    "min": low,
                                    "probe": id,
                                    "scale": "celsius"
                                }
                            }).then(function successCallback(response) {
                                $scope.successMessage = "Limits updated successfully";
                            }, function errorCallback(response) {
                                $scope.errorMessage = response.data.message;
                            });
                        }
                    }
                };
            };

            function newTempIndicator(probe) {
                return {
                    value: 0,
                    options: {
                        disabled: true,
                        floor: floor,
                        ceil: ceil,
                        vertical: true,
                        translate: function(value) {
                            return value + "° C";
                        },
                        showTicksValues: false,
                        hideLimitLabels: true,
                        getPointerColor: function(value) {
                            if (value <= 200)
                                return 'blue';
                            if (value <= 400)
                                return 'yellow';
                            if (value <= 600)
                                return 'orange';
                            return 'red';
                        }
                    }
                };
            }
        }]
    );