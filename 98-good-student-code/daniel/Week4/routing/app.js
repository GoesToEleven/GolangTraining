myApp = angular.module('myApp', ['ngRoute', 'artists']);

myApp.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/list', {
        templateUrl: 'list.html',
        controller: 'listController'
    })
    .when('/details/:person', {
        templateUrl: 'detail.html',
        controller: 'detailController'
    })
    .otherwise({
        redirectTo: '/list'
    });
}]);