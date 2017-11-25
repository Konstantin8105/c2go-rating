/***********************************************************/
/* To compile or not to compile                            */
/***********************************************************/

#define SOMEDEFINITION 6546
#define CHOICE  1     /* Choose this before compiling */

/***********************************************************/

#if (CHOICE == 1)

#define OPTIONSTRING "The programmer selected this"
#define DITTO        "instead of ....             "

#else

#define OPTIONSTRING "The alternative"
#define DITTO        "i.e. This! "


#endif

/***********************************************************/

#ifdef SOMEDEFINITION

#define WHATEVER "Something was defined!"

#else

#define WHATEVER "Nothing was defined"

#endif

/************************************************************/

main ()

{
printf (OPTIONSTRING);
printf (DITTO);
}


