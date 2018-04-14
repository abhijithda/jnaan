# Design HashSet

Design a HashSet without using any built-in hash table libraries. Refer [here](https://leetcode.com/explore/learn/card/hash-table/182/practical-applications/1139/) for online version.

To be specific, your design should include these two functions:

* add(value): Insert a value into the HashSet.
* contains(value) : Return whether the value exists in the HashSet or not.
* remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

Example:

```c++
MyHashSet hashSet = new MyHashSet();
hashSet.add(1);
hashSet.add(2);
hashSet.contains(1);    // returns true
hashSet.contains(3);    // returns false (not found)
hashSet.add(2);
hashSet.contains(2);    // returns true
hashSet.remove(2);
hashSet.contains(2);    // returns false (already removed)
```

Note:

* All values will be in the range of [1, 1000000].
* The number of operations will be in the range of [1, 10000].
* Please do not use the built-in HashSet library.