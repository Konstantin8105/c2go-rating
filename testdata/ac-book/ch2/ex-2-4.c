#include<stdio.h>

const int M = 1000000;

int factorial(int n)
{
  if (n == 1)
    return 1;
  else
    return (factorial(n-1) * n) % M;
}

int main()
{
  int n, result = 0;
  scanf("%d", &n);
  for(int i = 1; i <= n && i <= 25; i++) {
    result = (result + factorial(i)) % M;
  }
  printf("%d\n", result);
  return 0;
}
