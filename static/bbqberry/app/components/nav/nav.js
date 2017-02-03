/**
 * Created by dshanaghy on 2/3/17.
 */
'use strict';

angular.module('bbqberry.nav', [])

.component('bbqberryNav', {
    templateUrl: 'components/nav/nav.html',
    controller: 'BBQBerryNav'
})

.controller('BBQBerryNav', [ '$scope', '$location', function($scope, $location) {
    $scope.isActive = function(viewLocation) {
        return $location.path() == viewLocation;
    }
}]);
