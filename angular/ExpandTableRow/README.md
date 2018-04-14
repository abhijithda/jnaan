# Description

Allow `+`expand/`-`collapse of rows of a table in order to view/hide details. When a row is expanded, new rows will show up the details of the expanded row.

## Example

* [plnkr](http://plnkr.co/edit/RgfI1cinuHPOtpTbMT7u?p=preview)
* [rawgit](http://rawgit.com/abhijithda/jnaan/master/angular/index.html)

Below table that shows the overview:

|   |Name   |  Quantity |  Type  | Cost
---|-------|-----------|--------|------
|[+](# expand)|Fruits |  5        | Fruit  |    10
|[+](# expand)|Vegetables |  2    |  Vegetable | 5
|[+](# expand)|Pulses |  2        |  Pulse |    40
|[+](# expand)|Drinks |  4        |  Drink |    20

Clicking `+` on the `Fruits` row shows it's details as can be seen in the below table.


|  |  Name   |  Quantity |  Type  | Cost
---|-------|-----------|--------|------
|[-](# collapse)  |  Fruits |  5        | Fruit  |    10
|   |  Banana |	2      | Fruit 	|     1
|   |  Apple  |	2      | Fruit 	|     4
|   |Pomogrande |	1    |	Fruit 	|   5
|[+](# expand)|Vegetables |  2    |  Vegetable | 5
|[+](# expand)|Pulses |  2        |  Pulse |    40
|[+](# expand)|Drinks |  4        |  Drink |    20

---

<p align="center">
<img alt="Expand and Collapse Table Rows" src=img/expandtable.gif>
</p>

---