// Add inputted number together. Ex: 888 = 8 + 8 + 8 = 24

#include <stdio.h>

int main (void)
{
  int number, right_digit, sum;

  printf ("Enter a number: ");
  scanf  ("%i", &number);

  sum = 0;

  while (number != 0) {
    right_digit = number % 10;
    sum += right_digit;
    number = number / 10;
  }

  printf("%i\n", sum);

  return 0;
}
