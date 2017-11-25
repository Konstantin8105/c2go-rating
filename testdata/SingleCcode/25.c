/********************************************************/
/*                                                      */
/* String comparison                                    */
/*                                                      */
/********************************************************/

#include <stdio.h>

#define TRUE   1
#define MAXLEN 30

/********************************************************/

main ()

{ char string1[MAXLEN],string2[MAXLEN];
  int result;

while (TRUE)
   {
   printf ("Type in string 1:\n\n");
   scanf ("%30s",string1);

   printf ("Type in string 2:\n\n");
   scanf ("%30s",string2);

   result = strcmp (string1,string2);

   if (result == 0)
      {
      printf ("Those strings were the same!\n");
      }

   if (result > 0)
      {
      printf ("string1 > string2\n");
      }

   if (result < 0)
      {
      printf ("string1 < string 2\n");
      }
   }
}

