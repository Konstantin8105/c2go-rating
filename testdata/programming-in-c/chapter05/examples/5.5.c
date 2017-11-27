// Using Nested for Loops

#include <stdio.h>

int main (void)
{
  int n, number, triangularNumber, counter;

  for (counter = 1; counter <= 5; counter++) {
    printf ("What triangular number do you want? ");
    scanf  ("%i", &number);

    triangularNumber = 0;

    for (n = 1; n <= number; n++)
      triangularNumber += n;

    printf ("Triangular number %i is %i\n\n", number, triangularNumber);
  }

  return 0;
}

/* Output

What triangular number do you want? 10
Triangular number 10 is 55

What triangular number do you want? 5
Triangular number 5 is 15

What triangular number do you want? 100
Triangular number 100 is 5050

What triangular number do you want? 1000
Triangular number 1000 is 500500

What triangular number do you want? 5642
Triangular number 5642 is 15918903

*/
