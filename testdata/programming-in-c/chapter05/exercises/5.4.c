/* Program that generates and prints a table of
   the first 10 factorials */

#include <stdio.h>

int main (void)
{
  int i, k, result;

  for (i = 1; i <= 10; i++) { // Loop through 10 factorials
    for (k = i, result = i; k > 0; k--) {
      result *= k;
    }
    printf ("Factorial for %i is %i\n", i, result);
  }

  return 0;
}

/* Output

Factorial for 1 is 1
Factorial for 2 is 4
Factorial for 3 is 18
Factorial for 4 is 96
Factorial for 5 is 600
Factorial for 6 is 4320
Factorial for 7 is 35280
Factorial for 8 is 322560
Factorial for 9 is 3265920
Factorial for 10 is 36288000

*/
