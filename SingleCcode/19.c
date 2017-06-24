/******************************************************/
/*                                                    */
/* Prime Number Sieve                                 */
/*                                                    */
/******************************************************/

#include <stdio.h>

#define SIZE     5000
#define DELETED     0

/*******************************************************/
/* Level 0                                             */
/*******************************************************/

main ()

{ short sieve[SIZE];

printf ("Eratosthenes Sieve \n\n");

FillSeive(sieve);
SortPrimes(sieve);
PrintPrimes(sieve);
}

/*********************************************************/
/* Level 1                                               */
/*********************************************************/

FillSeive (sieve)               /* Fill with integers */

short sieve[SIZE];

{ short i;

for (i = 2; i < SIZE; i++)
   {
   sieve[i] = i;
   }
}

/**********************************************************/

SortPrimes (sieve)                /* Delete non primes */

short sieve[SIZE];

{ short i;

for (i = 2; i < SIZE; i++)
   {
   if (sieve[i] == DELETED)
      {
      continue;
      }
   DeleteMultiplesOf(i,sieve);
   }
}

/***********************************************************/

PrintPrimes (sieve)                  /* Print out array */

short sieve[SIZE];

{ short i;

for (i = 2; i < SIZE; i++)
   {
   if (sieve[i] == DELETED)
      {
      continue;
      }
   else
      {
      printf ("%5d",sieve[i]);
      }
   }
}

/***********************************************************/
/* Level 2                                                 */
/***********************************************************/

DeleteMultiplesOf (i,sieve)   /* Delete.. of an integer */

short i,sieve[SIZE];

{ short j, mult = 2;

for (j = i*2; j < SIZE; j = i * (mult++))
   {
   sieve[j] = DELETED;
   }
}

                      /* end */
