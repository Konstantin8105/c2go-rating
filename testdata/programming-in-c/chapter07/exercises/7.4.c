// Calculates the average of an array of 10 floating-point values.

#include <stdio.h>

int main (void)
{
    int i;
    float average = 0.0, arr_of_floats[10] = { 4.33, 8.43, 593.313, 59.31, 04.13,
        94.13, 493.12, 48.94, 9.12, 91.21 };

    for ( i = 0; i < 10; i++ )
        average += arr_of_floats[i];

    average /= 10;
    printf ("The average of the 10 floating-point numbers is %f.\n", average);

    return 0;
}
