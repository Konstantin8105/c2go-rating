/************************************************/
/*                                              */
/* switch  .. case                              */
/*                                              */
/************************************************/

   /* Morse code program. Enter a number and */
   /* find out what it is in Morse code      */

#include <stdio.h>

#define CODE 0

/*************************************************/

main ()

{ short digit;

printf ("Enter any digit in the range 0..9");

scanf ("%h",&digit);

if ((digit < 0) || (digit > 9))
   {
   printf ("Number was not in range 0..9");
   return (CODE);
   }

printf ("The Morse code of that digit is ");
Morse (digit);
}

/************************************************/

Morse (digit)        /* print out Morse code */

short digit;

{
switch (digit)
   {
   case 0 : printf ("-----");
            break;
   case 1 : printf (".----");
            break;
   case 2 : printf ("..---");
            break;
   case 3 : printf ("...--");
            break;
   case 4 : printf ("....-");
            break;
   case 5 : printf (".....");
            break;
   case 6 : printf ("-....");
            break;
   case 7 : printf ("--...");
            break;
   case 8 : printf ("---..");
            break;
   case 9 : printf ("----.");
   }
}
