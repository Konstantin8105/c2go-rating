// Program to generate a table of prime numbers

#include <stdio.h>
#include <stdbool.h>

int main(void)
{
    int number, divisor;
    bool isPrime;

    // Loop through each number up to 50.
    for (number = 2; number <= 50; number += (number == 2) ? 1 : 2) {
        isPrime = true;

        // Check if the number is prime by using modulo
        // on all numbers that come before it.
        for (divisor = 3; divisor < number; divisor += 2)
            // Find out if number is prime.
            if (number % divisor == 0)
                isPrime = false;

        // If it's a prime, print it.
        if (isPrime == true)
            printf("%i  ", number);
    }

    printf("\n");
    // All is well.
    return 0;
}
