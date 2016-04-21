var app = angular.module('app', []);

app.controller('controller', function($scope, $http) {
	$http.get('data.json').success(function(data) {
		$scope.people = data;
	})
});