/*******************************************************/
/*                                                     */
/* printf Conversion Characters and Types              */
/*                                                     */
/*******************************************************/

#include <stdio.h>

main ()

{ int i = -10;
  unsigned int ui = 10;
  float x = 3.56;
  double y = 3.52;
  char ch = 'z';
  char *string_ptr = "any old string";

printf ("signed integer %d\n", i);
printf ("unsigned integer %u\n",ui);

printf ("This is wrong! %u",i);
printf ("See what happens when you get the ");
printf ("character wrong!");

printf ("Hexadecimal %x %x\n",i,ui);
printf ("Octal %o %o\n",i,ui);

printf ("Float and double %f %f\n",x,y);
printf ("      ditto      %e %e\n",x,y);
printf ("      ditto      %g %g\n",x,y);

printf ("single character %c\n",ch);
printf ("whole string -> %s",string_ptr);
}
