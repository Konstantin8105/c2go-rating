/**********************************************/
/*                                            */
/* Swapping Pointers                          */
/*                                            */
/**********************************************/

   /* Program swaps the variables which a,b */
   /* point to. Not pointless really !      */

#include <stdio.h>


main ()

{ int *a,*b,*c;               /* Declr ptrs */
  int  A,B;               /* Declare storage */

A = 12;               /* Initialize storage */
B = 9;

a = &A;              /* Initialize pointers */
b = &B;

printf ("%d %d\n",*a,*b);

c = a;                     /* swap pointers */
a = b;
b = c;

printf ("%d %d\n",*a,*b);

}
