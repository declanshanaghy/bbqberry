'use strict';

angular.module('bbqberry.tempguage', [])

    .component('tempGuage', {
        templateUrl: 'components/tempguage/tempguage.html',
        controller: 'TempGuage',
        bindings: {
            probe: '<'
        }
    })

    .controller('TempGuage', ['$scope', function($scope) {

    }])
;
