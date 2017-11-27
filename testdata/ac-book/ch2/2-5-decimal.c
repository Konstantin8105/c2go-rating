#include<stdio.h>

int main()
{
  int a, b, c, k = 0;
  while(~scanf("%d %d %d", &a, &b, &c) && a && b && c) {
    double r = (double)a/b;
    printf("Case %d: %.*f\n", ++k, c, r);
  }
  return 0;
}
