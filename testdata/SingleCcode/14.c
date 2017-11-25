/***********************************************/
/*                                             */
/* If demo #3                                  */
/*                                             */
/***********************************************/

#include <stdio.h>

/***********************************************/

main ()

{ int persnum,usernum,balance;

persnum = 7462;
balance = -12;

printf ("The Plastic Bank Corporation\n");
printf ("Please enter your personal number :");

usernum = getnumber();

if (usernum == 7462)
   {
   printf ("\nThe current state of your account\n");
   printf ("is %d\n",balance);

   if (balance < 0)
      {
      printf ("The account is overdrawn!\n");
      }
   }
else
   {
   printf ("This is not your account\n");
   }

printf ("Have a splendid day! Thank you.\n");
}

/**************************************************/

getnumber ()     /* get a number from the user */

{ int num = 0;

scanf ("%d",&num);

if ((num > 9999) || (num <= 0))
   {
   printf ("That is not a valid number\n");
   }

return (num);
}

