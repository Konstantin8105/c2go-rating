// Program to generate the first 15 Fibonacci numbers.
#include <stdio.h>

int main (void)
{
    int i, current, previous, beforePrevious;

    beforePrevious = 0; // by definition
    previous = 1; // ditto

    printf("%i\n", beforePrevious);
    printf("%i\n", previous);

    for ( i = 2; i < 15;  ++i ) {
        current = beforePrevious + previous;
        printf ("%i\n", current);

        beforePrevious = previous;
        previous = current;
    }

    return 0;
}
