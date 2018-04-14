var app = angular.module('tableExpandCollapseApp', ['ngGrid']);

app.service('shoppingList', function($http, $q) {
  var items;

  return {
    getShoppingList: function() {
      if (angular.isUndefined(items)) {
        items = {};
        return $http.get('https://cdn.rawgit.com/abhijithda/jnaan/master/angular/ExpandTableRow/items.json').success(function(jsonData) {
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

    $scope.displayItemDetails = {};


    $scope.toggleDisplay = function(iType) {
      console.log("Toggling display of " + iType + " from " + $scope.displayItemDetails[iType]);
      $scope.displayItemDetails[iType] = $scope.displayItemDetails[iType] ? 0 : 1;
      console.log(" to " + $scope.displayItemDetails[iType]);

      console.log(" All display details value:");
      console.log($scope.displayItemDetails);

      $scope.selItems = $scope.updateTable();
      console.log("In toggleDisplay()::selItems: ");
      console.log($scope.selItems);
    };


    $scope.updateTable = function() {
      var selItems = [];
      for (var i in $scope.allItems) {
        var iType = $scope.allItems[i]["Type"];

        if (angular.isUndefined($scope.displayItemDetails[iType])) {
          $scope.displayItemDetails[iType] = 0;
        }

        if (1 == $scope.displayItemDetails[iType]) {
          $scope.allItems[i]["Summary"] = '-';
          console.log("Expanding(" + $scope.allItems[i]["Summary"] + ") details of " + iType);
        } else {
          $scope.allItems[i]["Summary"] = '+';
          console.log("Collapsing(" + $scope.allItems[i]["Summary"] + ") details of " + iType);
        }
        selItems.push($scope.allItems[i]);

        console.log("Display details of " + iType + ": " + $scope.displayItemDetails[iType]);
        if ($scope.displayItemDetails[iType]) {
          console.log("Viewing details of " + iType);
          for (var j in $scope.allItems[i]["Details"]) {
            $scope.allItems[i]["Details"][j]["Summary"] = "";
            selItems.push($scope.allItems[i]["Details"][j]);
          }
        }

      }
      console.log("In updateTable()::selItems: ");
      console.log(selItems);
      return selItems;
    };


    shoppingList.getShoppingList().then(function(result) {
      console.log("Result: ");
      console.log(result);
      $scope.allItems = result.data.Items;
      $scope.selItems = $scope.updateTable();
      console.log("In then::selItems: ");
      console.log($scope.selItems);
    });



    var summaryCellTemplate = "<div ng-if='row.getProperty(col.field)'>" +
      "<button ng-click='toggleDisplay(row.entity.Type)'> {{row.getProperty(col.field)}} </button>" +
      "</div>";

    $scope.gridOptions = {
      data: 'selItems',
      columnDefs: [{
        field: 'Summary',
        displayName: '',
        cellTemplate: summaryCellTemplate,
        width: 30
      }, {
        field: 'Name',
        displayName: 'Name',
      }, {
        field: 'Type',
        displayName: 'Type',
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

  });
