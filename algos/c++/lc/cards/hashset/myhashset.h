#include <iostream>
#include <vector>

using namespace std;

class MyHashSet
{
    vector<int> v; // (1000000, 0);
  public:
    MyHashSet()
    {
    }
    int add(int val)
    {
        cout << "Adding value " << val << endl;
        if (not contains(val))
        {
            v.push_back(val);
        }
        return 0;
    }
    bool contains(int val);
    int remove(int val);
};

bool MyHashSet::contains(int val)
{
    // clog << "Checking for existence of value " << val << endl;
    if (std::find(v.begin(), v.end(), val) != v.end())
    {
        // Element in vector.
        return true;
    }
    return false;
}

int MyHashSet::remove(int val)
{
    cout << "Removing value " << val << endl;
    for (int i = 0; i < v.size(); i++)
    {
        if (v.at(i) == val)
        {
            cout << "Found " << val << " at position " << endl;
            v.erase(v.begin() + i);
            return 0;
        }
    }
    return 0;
}
