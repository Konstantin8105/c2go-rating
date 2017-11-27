#include<stdio.h>
#include<math.h>

int main()
{
  // A1
  printf("%d\n", 11111*11111);
  printf("%d\n", 111111*111111); // overflow
  printf("%d\n", 111111111*111111111); // overflow

  // A2
  printf("%f\n", 11111.0*11111.0);
  printf("%f\n", 111111.0*111111.0);
  printf("%f\n", 111111111.0*111111111.0);

  // A3
  printf("%d\n", sqrt(-10)); // Some strange value
  printf("%f\n", sqrt(-10)); // -nan
  printf("%lf\n", sqrt(-10));// -nan

  // A4
  //printf("%f\n", 1.0/0.0); // floating pt exception
  //printf("%f\n", 0.0/0.0); // floating pt exception

  // A5
  //printf("%d\n", 1/0); // floating pt exception
  //printf("%d\n", 0/0); // floating pt exception

  // %d
  printf("%d\n"); // Some strange value
  printf("%%d\n");

  return 0;
}
