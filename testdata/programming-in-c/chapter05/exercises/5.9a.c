// Replace 5.2 - 5.5 examples for loops with while loops

#include <stdio.h>

int main (void)
{
    int n, triangularNumber;

    n = 1;
    triangularNumber = 0;

    while (n <= 200) {
      triangularNumber += n;
      n++;
    }

    printf ("The 200th triangular number is %i\n", triangularNumber);

    return 0;
}
