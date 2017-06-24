#include <stdio.h>

#define STRING1        "A macro definition\n"
#define STRING2        "must be all on one line!!\n"
#define EXPRESSION     1 + 2 + 3 + 4
#define EXPR2          EXPRESSION + 10
#define ABS(x)         ((x) < 0) ? -(x) : (x)
#define MAX(a,b)       (a < b) ? (b) : (a)
#define BIGGEST(a,b,c) (MAX(a,b) < c) ? (c) : (MAX(a,b))

/************************************************************/

   main ()             /* No #definitions inside functions! */

   {
   printf (STRING1);
   printf (STRING2);
   printf ("%d\n",EXPRESSION);
   printf ("%d\n",EXPR2);
   printf ("%d\n",ABS(-5));
   printf ("Biggest of 1 2 and 3 is %d",BIGGEST(1,2,3));
   }

