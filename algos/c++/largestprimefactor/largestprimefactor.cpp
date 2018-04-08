// Problem 3 - Largest prime factor

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

#include <iostream>

using namespace std;

typedef struct primes {
    long prime;
    struct primes* next;
}P;

P* g_head = NULL;
P* g_tail = NULL;
long gl_primecount = 0;

int addPrimeNum(long num){
    P *cur = new(P);
    cout << "Add prime number " << num << " to the list of " << gl_primecount++ <<" elements..." <<endl;
    cur->prime = num;
    cur->next = NULL;
    if( g_head == NULL){
        g_head = cur;
        g_tail = cur;
        return 0;
    }
    g_tail->next = cur;
    g_tail = cur;
    return 0;
}


long isPrime(long num){
    cout << endl;
    cout << " ================ Checking for prime: " << num << "? ";
    short isPrime = 1;
    // for(long i=2; i<=long(num/2); i++){}
    for(P* cur=g_head; cur != NULL; cur=cur->next){
        long p = cur->prime;
        cout << "." << flush;
        // cout << "Divisible by " << p << "? " ;
        if(num%p == 0){
            isPrime = 0;
            // cout << "Yes."  << endl;
            break;
        }
    }
    cout << endl;
    cout << " ================ Is prime: " << num << "? ";
    if(isPrime){
        cout << "Yes." << endl;
        addPrimeNum(num);
    }
    else{
        cout << "No." << endl;
    }
    return isPrime;
}

long largestprimefactor(long num){
    long largestPrime = 1;

    for (long i=2; i<=long(num/2); i++){
        if(not isPrime(i)){
            // cout << i << " is a NOT a prime number" << endl;
            continue;
        }
        // cout << i << " is a prime number" << endl;

        if(num%i == 0){
            largestPrime = i;
        }

    }
    cout << largestPrime << " is the largest prime factor of " << num << endl;        
    return largestPrime;
}




int main(){
    long num = 600851475143;
    cout << "Largest prime factor of " << num << ": ";
    cout << largestprimefactor(num);
    cout << endl;
    cout << "Number of elements in the prime list: " << gl_primecount << endl;
}