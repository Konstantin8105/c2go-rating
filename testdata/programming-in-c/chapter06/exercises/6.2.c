// Check if first number is evenly divisible by second.

#include <stdio.h>

int main(void)
{
    int num1, num2;

    printf("Enter in the first number: ");
    scanf("%i", &num1);
    printf("Enter in the second number: ");
    scanf("%i", &num2);

    if (num1 % num2 == 0)
        printf("%i is evenly divisible by %i.\n", num1, num2);
    else
        printf("%i is not evenly divisible by %i.\n", num1, num2);
}
