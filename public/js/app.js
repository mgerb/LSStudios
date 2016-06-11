var app = angular.module('app', ['ngRoute']);

app.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
    
    $locationProvider.html5Mode({
        enabled: true,
        requireBase: false
    });

    $routeProvider.
    when('/', {
        templateUrl: '/public/view/default.html',
        controller: 'IndexController'
    }).
    otherwise({
        templateUrl: '/'
    });

}]);
