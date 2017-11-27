#include<stdio.h>

int main()
{
  double result = 0, n, sign = 1;
  for(n = 1; 1.0/n >= 1e-6; n += 2) {
    result += sign*1.0/n;
    sign = -sign;
  }
  printf("%.6f\n", result);
  return 0;
}
