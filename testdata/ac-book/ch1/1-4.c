// Sin Cos
#include<stdio.h>
#include<math.h>

int main()
{
  float n;
  scanf("%f", &n);
  printf("%f %f\n", sin(n*M_PI/180.0), cos(n*M_PI/180.0));
  return 0;
}
