// Implement an algorithm to determine if a string has all unique characters. What if you can not use additional data structures?

#include <iostream>

using namespace std;

int main(){
    string str = "all unique characters?";

    for(int i=0; i < str.length(); i++ ){
        for(int j=i+1; j < str.length(); j++ ){
            if( str[i] == str[j]){
                cout << "Duplicate character found: " << str[i] << endl;
            }
            
        }
    }
}