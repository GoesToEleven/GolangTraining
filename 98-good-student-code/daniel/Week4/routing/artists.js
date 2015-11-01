var artistApp = angular.module('artists', ['firebase', 'ngAnimate']);

artistApp.controller('listController', ['$scope', '$firebaseArray', function($scope, $firebaseArray) {
    var ref = new Firebase('https://angular-bootcamp.firebaseio.com/people');

    $scope.people = $firebaseArray(ref);
    $scope.delete = function(itemId) {
        ref.child(itemId).remove();
    };
    $scope.add = function() {
        if($scope.newItem.name && $scope.newItem.reknown && $scope.newItem.bio) {
            var reader = new FileReader();
            reader.addEventListener('load', function(e) {
                $scope.newItem.image = e.target.result;
                ref.push($scope.newItem);
                $scope.newItem = {};
            });
            reader.readAsDataURL(document.forms.addPerson.picture.files[0]);
        }
    }
}]);

artistApp.controller('detailController', ['$scope', '$routeParams', '$firebaseObject', function($scope, $routeParams, $firebaseObject) {
    var ref = new Firebase('https://angular-bootcamp.firebaseio.com/people');
    var person = $firebaseObject(ref.child($routeParams.person));
    person.$bindTo($scope, 'person');
}]);