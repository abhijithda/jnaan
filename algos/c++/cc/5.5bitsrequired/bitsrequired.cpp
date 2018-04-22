// Write a function to determine the number of bits required to convert integer A to integer B.
// Input: 31, 14
// Output: 2

#include <iostream>

using namespace std;

int main()
{
    int num1, num2;
    cout << "Enter first number: ";
    cin >> num1;
    cout << "Enter second number: ";
    cin >> num2;

    int num = num1 ^ num2;
    cout << "XOR of " << num1 << " and " << num2 << " is " << num << endl;

    int idx = 1;
    int cnt = 0;
    while(num){
        unsigned bit = num & 1;
        num >>= 1;
        cout << "Bit " << idx++ << ": " << bit << endl;
        if( bit ){
            cnt++;
        }
    }

    cout << "Number of bits that are needed convert " << num1 << " to " << num2 << " is " << cnt << endl;
}