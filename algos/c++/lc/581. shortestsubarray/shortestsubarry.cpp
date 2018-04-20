#include <iostream>


class Solution {
public:
    int findUnsortedSubarray(vector<int>& nums) {
        // If list is empty or of size 1, then function returns 0.
        int saStartIndex = -1;
        int saLength = 0;
        int bigNumIndex = 0;
        for(int i=1; i < nums.size(); i++){
            if(nums[i] < nums[bigNumIndex]){
                if( saStartIndex == -1 || nums[saStartIndex] > nums[i] ){
                    for(int saIndex=0; saIndex<i; saIndex++){
                        if(nums[saIndex] > nums[i]){
                            saStartIndex = saIndex;
                            break;
                        }
                    }
                }
                saLength = i - saStartIndex + 1;
            }
            else if(nums[i] > nums[bigNumIndex]){ // If greater, then update the new bigger number.
                bigNumIndex = i;
            }
        }
        return saLength;
    }
};