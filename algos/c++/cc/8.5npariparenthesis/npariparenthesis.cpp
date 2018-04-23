// Implement an algorithm to print all valid (e.g., properly opened and closed) combinations of n-pairs of parentheses.
// EXAMPLE:
// input: 3 (e.g., 3 pairs of parentheses)
// output: ()()(), ()(()), (())(), ((()))

#include <iostream>
#include <string>

using namespace std;

void printPar(int l, int r, char str[], int count)
{
    if (l < 0 || r < l)
    {
        cout << "Invalid state!" << endl;
        return; // invalid state
    }
    if (l == 0 && r == 0)
    {
        cout << "String: " << str << endl; // found one, so print it
    }
    else
    {
        if (l > 0)
        {
            // try a left paren, if there are some available
            str[count] = '(';
            printPar(l - 1, r, str, count + 1);
        }
        if (r > l)
        { // try a right paren, if there's a matching left
            str[count] = ')';
            printPar(l, r - 1, str, count + 1);
        }
    }
}

void printPar(int count)
{
    char str[20];
    printPar(count, count, str, 0);
}

int main()
{
    int cnt = 0;
    cout << "Enter parenthesis count: ";
    cin >> cnt ;
    printPar(cnt);
}