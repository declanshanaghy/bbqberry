'use strict';

angular.module('bbqberry.view2', ['ngRoute'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/view2', {
            templateUrl: 'views/view2/view2.html',
            controller: 'View2Ctrl'
        });
    }])

    .controller('View2Ctrl', ['$scope', function ($scope) {
        $scope.name = 'ARSE';
        $scope.names = [{name: "Chris"}, {name: "Calvin"}];
        $scope.addName = function () {
            $scope.names.push({'name': $scope.name});
            $scope.name = '';
        };
    }]);