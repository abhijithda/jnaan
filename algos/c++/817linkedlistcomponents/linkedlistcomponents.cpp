#include <iostream>
#include <vector>

/**
 Definition for singly-linked list.
 */
typedef struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(NULL) {}
 }ListNode;
 
 using namespace std;

class Solution
{
  public:
    int numComponents(ListNode *head, vector<int> &G)
    {
        int count = 0;
        int consequtive = 0;
        ListNode *trav = head;
        while (trav != nullptr)
        {
            if (contains(G, trav->val) == true)
            {
                if(consequtive == 0)
                {
                    consequtive = 1;
                    count++;
                }
            }
            else
            {
                consequtive = 0;
            }
            trav = trav->next;
        }
        return count;
    }

    bool contains(vector<int> &v, int val)
    {
        // clog << "Checking for existence of value " << val << endl;
        if (std::find(v.begin(), v.end(), val) != v.end())
        {
            // Element in vector.
            return true;
        }
        return false;
    }
};

int main()
{
    
}