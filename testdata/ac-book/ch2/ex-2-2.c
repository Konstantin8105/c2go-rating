#include<stdio.h>

int main()
{
  long long n;
  int c;
  scanf("%lld", &n);
  for (c=0 ; n > 1; c++)
  {
    if (n%2 == 0) n = n/2;
    else n = 3*n+1;
  }
  printf("%d\n", c);
  return 0;
}
