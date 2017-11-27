#include <stdio.h>

int main (void)
{
  int n, triangularNumber;

  printf ("TABLE OF TRIANGULAR NUMBERS\n\n");
  printf (" n   Sum from 1 to n\n");
  printf ("---  ---------------\n");

  n = 1;
  triangularNumber = 0;

  while (n <= 10) {
    triangularNumber += n;
    printf (" %2i          %i\n", n, triangularNumber);
    n++;
  }

  return 0;
}
