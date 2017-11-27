// Program to reverse the digits of a number

#include <stdio.h>
#include <stdbool.h>

int main ()
{
    int number, right_digit;
    bool isNegative = false;

    // Get user input
    printf ("Enter your number to reverse: ");
    scanf  ("%i", &number);

    // Make number positive if it was negative.
    if (number < 0) {
        number = -number;
        isNegative = true;
    }

    // Reverse number
    do {
        right_digit = number % 10;
        printf ("%i", right_digit);
        number = number / 10;
    }
    while (number != 0);

    // Print a negative symbol at the end if
    // number was negative.
    if (isNegative == true)
        printf ("-\n");
    else
        printf("\n");

    return 0;
}
