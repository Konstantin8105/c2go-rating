/* Write a program to generate and display a table of n and n^2,
   for integer values of n ranging from 1 to 10. Be certain to
   print the appropriate column headings.                        */

#include <stdio.h>

int main (void) {

  int n, n2;

  printf (" n        n^2  \n");
  printf ("---    ---------\n");

  for ( n = 1; n <= 10; n++) {
    n2 = n * n;
    printf ("%2i         %i\n", n, n2);
  }

  return 0;
}

/* Output

 n        n^2
---    ---------
 1         1
 2         4
 3         9
 4         16
 5         25
 6         36
 7         49
 8         64
 9         81
 10        100

*/
