#include <iostream>

#include "myhashset.h"

int main()
{
    MyHashSet *h = new MyHashSet();
    h->add(1);
    h->add(2);

    int ele = 2;
    cout << "Checking whether ele " << ele << " is present: ";
    if (h->contains(ele))
    {
        cout << "Found" << endl;
    }

    ele = 5;
    cout << "Checking whether ele " << ele << " is present: ";
    if (h->contains(ele))
    {
        cout << "Found" << endl;
    }
    else{
        cout << "Not Found" << endl;        
    }

    h->add(ele);
    cout << "Checking whether ele " << ele << " is present: ";
    if (h->contains(ele))
    {
        cout << "Found" << endl;
    }
    else{
        cout << "Not Found" << endl;        
    }

    h->remove(ele);
    cout << "Checking whether ele " << ele << " is present: ";
    if (h->contains(ele))
    {
        cout << "Found" << endl;
    }
    else{
        cout << "Not Found" << endl;        
    }

    h->remove(9);
}