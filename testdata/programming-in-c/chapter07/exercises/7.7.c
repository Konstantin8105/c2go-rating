// Sieve of Erastothenes Algorithm

#include <stdio.h>
#include <math.h>

#define MAX 100

int main (void)
{
    int i, j;
    int primes[MAX];

    // Initialize array.
    for ( i = 0; i < MAX; i++ )
        primes[i] = i;

    // Find all primes.
    for ( i = 1; i < sqrt(MAX); i++ )
        // If number is not marked.
        if (primes[i] != 1)
            for ( j = i + 1; j < MAX; j++ )
                // See if any of the other numbers are divisible
                // by the unmarked number, primes[i].
                if ( j % primes[i] == 0)
                    // And mark it as not prime.
                    primes[j] = 1;

    // Print out primes.
    for ( i = 1; i < MAX; i++ )
        if (primes[i] != 1)
            printf("%i\n", primes[i]);

    return 0;
}
