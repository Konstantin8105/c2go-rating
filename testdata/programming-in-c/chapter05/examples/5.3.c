// Program to generate a table of triangular numbers

#include <stdio.h>

int main (void)
{
  int n, triangularNumber;

  printf ("TABLE OF TRIANGULAR NUMBERS\n\n");
  printf (" n   Sum from 1 to n\n");
  printf ("---  ---------------\n");

  triangularNumber = 0;

  for (n = 1; n <= 10; n++) {
    triangularNumber += n;
    printf (" %2i          %i\n", n, triangularNumber);
  }

  return 0;
}

/* Output

n   Sum from 1 to n
---  ---------------
 1          1
 2          3
 3          6
 4          10
 5          15
 6          21
 7          28
 8          36
 9          45
10          55

 */
