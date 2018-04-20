#include <iostream>
#include <vector>

using namespace std;

class Solution
{
  public:
    bool containsDuplicate(vector<int> &nums)
    {
        unordered_set<int> num_set;

        for (int num : nums)
        {
            if (num_set.count(num) > 0)
            {
                return true;
            }
            else{
                num_set.insert(num);
            }
        }

        return false;
    }
};

int main()
{
    cout << "C++ standard version: " << __cplusplus << endl;
}