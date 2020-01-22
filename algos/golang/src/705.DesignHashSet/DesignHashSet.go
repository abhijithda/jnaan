package DesignHashSet

type MyHashSet struct {
	hash []bool
	size int
}

// Constructor /** Initialize your data structure here. */
func Constructor() MyHashSet {
	// Note:

	// All values will be in the range of [0, 1000000].
	// The number of operations will be in the range of [1, 10000].
	// Please do not use the built-in HashSet library.
	hs := MyHashSet{size: 1000000}
	hs.hash = make([]bool, hs.size)
	return hs
}

// Add (value): Insert a value into the HashSet.
func (this *MyHashSet) Add(key int) {
	// if key >= this.size {
	// 	log.Fatal("Not meeting expected constaints")
	// }
	this.hash[key] = true
}

// Remove (value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.
func (this *MyHashSet) Remove(key int) {
	this.hash[key] = false
}

// Contains (value) : Return whether the value exists in the HashSet or not. /** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.hash[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
