# Description

Allow expand/collapse of rows of a table in order to view/hide details. When a row is expanded, new rows will show up the details of the expanded row.


# Example: 
plnkr [here](http://plnkr.co/edit/RgfI1cinuHPOtpTbMT7u?p=preview).


Below table that shows the overview:

   |Name   |  Quantity |  Type  | Cost
---|-------|-----------|--------|------                                                               
[+](# expand)|Fruits |  5        | Fruit  |    10                                                               
[+](# expand)|Vegetables |  2    |  Vegetable | 5                                                               
[+](# expand)|Pulses |  2        |  Pulse |    40                                                               
[+](# expand)|Drinks |  4        |  Drink |    20


Below table that shows the details of Fruits item when '+' of the Fruits row is clicked.


   |Name   |  Quantity |  Type  | Cost
---|-------|-----------|--------|------                                                               
[-](# collapse)|Fruits |  5        | Fruit  |    10                                                               
   |Banana |	2      | Fruit 	|     1
   |Apple  |	2      | Fruit 	|     4
   |Pomogrande |	1    |	Fruit 	|   5
[+](# expand)|Vegetables |  2    |  Vegetable | 5                                                               
[+](# expand)|Pulses |  2        |  Pulse |    40                                                               
[+](# expand)|Drinks |  4        |  Drink |    20


