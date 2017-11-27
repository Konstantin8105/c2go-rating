/* Program to find the greatest common divisor of
   two nonnegative integer values                  */

#include <stdio.h>

int main (void)
{
  int u, v, temp;

  printf ("Please type in two nonnnegative integers.\n");
  scanf  ("%i%i", &u, &v);

  while (v != 0) {
    temp = u % v;
    u = v;
    v = temp;
  }

  printf ("Their greatest common divisor is %i\n", u);

  return 0;
}

/* Output

Please type in two nonnnegative integers.
10
10
Their greatest common divisor is 10

*/
