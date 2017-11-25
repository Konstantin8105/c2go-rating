/**********************************************************/
/*                                                        */
/*  Numerical Estimation of Integral                      */
/*                                                        */
/**********************************************************/

#include <stdio.h>
#include <math.h>
#include <limits.h>

#define LIMIT 5

double inc = 0.001;      /* Increment width - arbitrary */
double twopi;

/***********************************************************/
/** LEVEL 0                                                */
/***********************************************************/

main ()

{ double y,integrand();
  double integral = 0;
  twopi = 4 * asin(1.0);

for ( y = inc/2;  y < LIMIT;  y += inc )
   {
   integral += integrand (y) * inc;
   }

printf ("Integral value = %.10f \n",integral);
}

/***************************************************************/
/** LEVEL 1                                                   **/
/***************************************************************/

double integrand (y)

double y;

{ double value;

value = 2*y;

if (value > 1e308)
   {
   printf ("Overflow error\n");
   exit (0);
   }

return (value);
}

