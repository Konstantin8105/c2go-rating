/************************************************/
/*                                              */
/* static string array                          */
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
  static char *code[] =
     {
     "dummy",             /* index starts at 0 */
     "-----",
     ".----",
     "..---",
     "...--",
     "....-",
     ".....",
     "-....",
     "--...",
     "---..",
     "----.",
     };

printf ("%s\n",code[digit]);
}

