#include <stdio.h>

int main (void)
{
    int ratingCounters[11], i, response;

    // Add values 1 - 10 to ratingCounters
    for ( i = 1; i <= 10; ++i )
        ratingCounters[i] = 0;

    // Grab responses from user.
    printf ("Enter your responses: \n");
    for ( i = 1; i <= 20; ++i ) {
        scanf ("%i", &response);

        if (response < 1 || response > 10)
            printf ("%i must be between 1 - 10 inclusively.\n", response);
        else
            ++ratingCounters[response];
    }

    // Print table heading.
    printf ("\n\nRating    Number of Responses\n");
    printf ("------ --------------------\n");

    // Print table values.
    for ( i = 1; i <= 10; ++i)
        printf ("%4i%14i\n", i, ratingCounters[i]);

    // All is well.
    return 0;
}
