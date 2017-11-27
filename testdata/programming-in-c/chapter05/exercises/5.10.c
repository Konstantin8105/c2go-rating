// See what happens when you type in a negative # in example 5.8

#include <stdio.h>

int main (void)
{
  int number, right_digit;

  printf ("Enter a number to reverse: ");
  scanf  ("%i", &number);

  while (number != 0) {
    right_digit = number % 10;
    printf ("%i", right_digit);
    number = number / 10;
  }

  printf("\n");

  return 0;
}

/*

Enter a number to reverse: -10
0-1

*/
