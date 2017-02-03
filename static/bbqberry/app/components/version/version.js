'use strict';

angular.module('bbqberry.version', [
  'bbqberry.version.interpolate-filter',
  'bbqberry.version.version-directive'
])

.value('version', '0.1');
