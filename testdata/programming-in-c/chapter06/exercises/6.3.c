// Check if first number is evenly divisible by second.

#include <stdio.h>

int main(void)
{
    int num1, num2;
    float result;

    // Get user input
    printf("Enter in the first number: ");
    scanf("%i", &num1);
    printf("Enter in the second number: ");
    scanf("%i", &num2);

    // Check if we're dividing by 0
    if (num2 == 0) {
        printf("Division by zero.\n");
        return 1;
    }

    // Typecast nums to floats for result
    result = (float) num1 / (float) num2;
    printf("%i divided by %i equals %.3f.\n", num1, num2, result);

    return 0;
}
