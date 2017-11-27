/* Round off an integer i to the next largest even multiple
   of another integer j.

   Ex: 7 goes into 371 fifty-three times evenly. It is the
   next possible divible number for 7 to go into starting
   from 365.                                                */


#include <stdio.h>

int main (void)
{
    int i = 365, j = 7;
    printf ("%i\n", i + j - i % j); // 371

    i = 12258, j = 23;
    printf ("%i\n", i + j - i % j); // 12259

    i = 996, j = 4;
    printf ("%i\n", i + j - i % j); // 1000

    return 0;

}
