/* Program that finds every fifth triangular number
   between 5 and 50 */

#include <stdio.h>

int main (void)
{
  int n, triangularNumber;

  for (n = 5; n <= 50; n = n + 5 ) {
    triangularNumber =  n * (n + 1) / 2;
    printf ("Triangular number for %i is %i\n", n, triangularNumber);
  }

  return 0;
}

/* Output

Triangular number for 5 is 15
Triangular number for 10 is 55
Triangular number for 15 is 120
Triangular number for 20 is 210
Triangular number for 25 is 325
Triangular number for 30 is 465
Triangular number for 35 is 630
Triangular number for 40 is 820
Triangular number for 45 is 1035
Triangular number for 50 is 1275

*/
