/**************************************************/
/*                                                */
/* Program : Variable Parameters                  */
/*                                                */
/**************************************************/

   /* Scale some measurements on a drawing, say */

#include <stdio.h>

/**************************************************/

   main ()                   /* Scale measurements*/

   { int height,width;

   height = 4;
   width = 5;

   ScaleDimensions (&height,&width);

   printf ("Scaled height = %d\n",height);
   printf ("Scaled width = %d\n",width);
   }

/****************************************************/

   ScaleDimensions (h,w)    /* return scaled values */

   int *h, *w;

   { int hscale = 3;        /* scale factors */
int wscale = 1;

   *h = *h * hscale;
   *w = *w * wscale;
   }

