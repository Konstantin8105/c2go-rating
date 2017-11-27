// Turn an integer into words.
// Ex: 505 would be five zero five

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(void)
{
    int number, right_digit, i;

    // Grab a number from the user.
    printf("Enter a number: ");
    scanf("%i", &number);

    // Find the number of digits in the number.
    int num_of_digits = floor(log10(abs(number))) + 1;
    // Declare an array to hold that many digits
    int number_array[num_of_digits];

    // Store each number in an array in reverse order.
    for (i = 0; i < num_of_digits; i++) {
        right_digit = number % 10;
        number /= 10;
        number_array[i] = right_digit;
    }

    // Loop through the array in reverse to get
    // the original order.
    for (i = num_of_digits - 1; i >= 0; i--) {

        // Print out the word for each number.
        switch (number_array[i]) {
            case 0:
                printf("zero ");
                break;
            case 1:
                printf("one ");
                break;
            case 2:
                printf("two ");
                break;
            case 3:
                printf("three ");
                break;
            case 4:
                printf("four ");
                break;
            case 5:
                printf("five ");
                break;
            case 6:
                printf("six ");
                break;
            case 7:
                printf("seven ");
                break;
            case 8:
                printf("eight ");
                break;
            case 9:
                printf("nine ");
                break;
            default:
                printf("%i", number_array[i]);
                break;
        }
    }

    // All is well.
    return 0;
}
