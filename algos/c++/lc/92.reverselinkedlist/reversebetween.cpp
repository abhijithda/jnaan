#include <iostream>

using namespace std;

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        ListNode* trav = head;
        int trav_cnt = 0;
        while(trav_cnt != m - 1){
            if( trav->next == NULL ){
                cout << "Error: List is supposed to have atleast %d nodes", n
                exit(1)
            }
            trav = trav->next;
            trav_cnt++;
        }
        mMinus1Node = trav;

        ListNode* nNode, nPlus1Node = reversePointer(trav_cnt, n, trav->next, trav->next->next);
        mMinus1Node->next->next = nPlus1Node;
        mMinus1Node->next = nNode;
    }

    ListNode* reversePointer(cnt, n, ListNode* first, ListNode* second){
        if (cnt == n){
            return first, second;
        }
        ListNode* nNode, nPlus1Node = reversePointer(++cnt, n, second, second->next)

        second->next = first;

        return nNode, nPlus1Node;
    }
};


