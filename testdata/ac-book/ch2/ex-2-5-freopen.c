#define LOCAL
#include<stdio.h>

int main()
{
#ifdef LOCAL
  freopen("input.txt", "r", stdin);
  freopen("output.txt", "w", stdout);
#endif

  int x, sum = 0, count = 0, max = 0, min = 1000;
  while (~scanf("%d", &x)) {
    count++;
    sum += x;
    max = x > max ? x : max;
    min = x < min ? x : min;
  }
  printf("%d %d %.3f\n", min, max, (double)sum/count);
  return 0;
}
