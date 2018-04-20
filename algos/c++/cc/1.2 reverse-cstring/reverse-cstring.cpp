// Write code to reverse a C-Style String. (C-String means that “abcd” is represented as five characters, including the null character.)

#include <iostream>

using namespace std;

int main(){
    char a[] = "abcdef";
    cout << "Enter C-string to reverse: ";
    cin >> a;
    cout << "C-string to reverse is " << a << " of "<< sizeof(a) << "characters." << endl;
    for(int i=0; i<sizeof(a); i++){
        cout << a[i] << endl;
    }

    cout << "Reversing the string..." << endl;
    for(int i=0, j=sizeof(a)-2; i<sizeof(a)/2; i++, j--){
        cout << "Swapping characters: " << a[i] << a[j] << endl;
        a[i] ^= a[j];
        a[j] ^= a[i];
        a[i] ^= a[j];
    }

    cout << "After reverse: " << a << endl;
}