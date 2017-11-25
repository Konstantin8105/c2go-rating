/******************************************************/
/*                                                    */
/* Program : More Value Parameters                    */
/*                                                    */
/******************************************************/

     /* Print out mock exam results etc */

#include <stdio.h>

/******************************************************/

main ()                  /* Print out exam results */

{ int pupil1,pupil2,pupil3;
  int ppr1,ppr2,ppr3;
  float pen1,pen2,pen3;

pupil1 = 87;
pupil2 = 45;
pupil3 = 12;

ppr1 = 200;
ppr2 = 230;
ppr3 = 10;

pen1 = 1;
pen2 = 2;
pen3 = 20;

analyse (pupil1,pupil2,pupil3,ppr1,ppr2,
                    ppr3,pen1,pen2,pen3);

}

/*******************************************************/

analyse (p1,p2,p3,w1,w2,w3,b1,b2,b3)

int p1,p2,p3,w1,w2,w3;
float b1,b2,b3;

{
printf ("Pupil 1 scored %d percent\n",p1);
printf ("Pupil 2 scored %d percent\n",p2);
printf ("Pupil 3 scored %d percent\n",p3);

printf ("However: \n");

printf ("Pupil1 wrote %d sides of paper\n",w1);
printf ("Pupil2 wrote %d sides\n",w2);
printf ("Pupil3 wrote %d sides\n",w3);

if (w2 > w1)
   {
   printf ("Which just shows that quantity");
   printf (" does not imply quality\n");
   }

printf ("Pupil1 used %f biros\n",b1);
printf ("Pupil2 used %f \n",b2);
printf ("Pupil3 used %f \n",b3);

printf ("Total paper used = %d", total(w1,w2,w3));
}

/*****************************************************/

total (a,b,c)                     /* add up total */

int a,b,c;

{
return (a + b + c);
}
