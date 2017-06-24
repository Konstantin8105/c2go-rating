/************************************************/
/*                                              */
/* Formatted strings                            */
/*                                              */
/************************************************/

   /* program rewrites s1 in reverse into s2 */


#include <stdio.h>

#define SIZE   20
#define CODE    0

/************************************************/

main ()

{ static char *s1 = "string 2.3 55x";
  static char *s2 = "....................";
  char ch, *string[SIZE];
  int i,n;
  float x;

sscanf (s1,"%s %f %d %c", string, &x, &i, &ch);

n = sprintf (s2,"%c %d %f %s", ch, i, x, string);

if (n > SIZE)
   {
   printf ("Error: string overflowed!\n");
   exit (CODE);
   }

   puts (s2);
   }
