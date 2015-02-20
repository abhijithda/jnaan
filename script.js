var app = angular.module('tableExpandCollapseApp', ['ngGrid']);

app.service('shoppingList', function($http, $q) {
  var items;

  return {
    getShoppingList: function() {
      if (angular.isUndefined(items)) {
        items = {};
        return $http.get('items.json').success(function(jsonData) {
          items["Items"] = jsonData["Items"];
        });
      } else {
        return $q.when(items);
      }
    }
  };
});

app.controller('tableExpandCollapseController',
  function($scope, shoppingList) {

    shoppingList.getShoppingList().then(function(result) {
      console.log("Result: ");
      console.log(result);
      $scope.allItems = result.data.Items;
      $scope.selItems = result.data.Items;
      console.log("In then::selItems: ");
      console.log($scope.selItems);

    });


    $scope.gridOptions = {
     data: 'selItems',
      columnDefs: [{
        field: 'Summary',
        displayName: '',
        cellTemplate: '<div class="ngCellText"> <div ng-if="row.getProperty(col.field)"> <button ng-click="toggleDisplay(row.entity.Type)"> {{row.getProperty(col.field)}} </button> </div> </div>'
      }, {
        field: 'Name',
        displayName: 'Name',
      }, {
        field: 'Cost',
        displayName: 'Cost',
      }, {
        field: 'Quantity',
        displayName: 'Quantity',
      }],
      enableCellSelection: false,
      enableColumnResize: true
    };
 // });
  });
