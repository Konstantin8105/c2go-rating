/**************************************/
/*                                    */
/* Operators Demo # 2                 */
/*                                    */
/**************************************/

#include <stdio.h>

/**************************************/

   main ()

   { int i;

   printf ("Assignment Operators\n\n");

   i = 10;                           /* Assignment */
   printf("i = 10 : %d\n",i);

   i++;                              /* i = i + 1 */
   printf ("i++ : %d\n",i);

   i += 5;                           /* i = i + 5 */
   printf ("i += 5 : %d\n",i);

   i--;                              /* i = i = 1 */
   printf ("i-- : %d\n",i);

   i -= 2;                           /* i = i - 2 */
   printf ("i -= 2 : %d\n",i);

   i *= 5;                           /* i = i * 5 */
   printf ("i *= 5 :%d\n",i);

   i /= 2;                           /* i = i / 2 */
   printf ("i /= 2 : %d\n",i);

   i %= 3;                           /* i = i % 3 */
   printf ("i %%= 3 : %d\n",i);
   }
