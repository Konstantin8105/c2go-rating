#define L
#include<stdio.h>

int main()
{
#ifdef L
  freopen("input.txt", "r", stdin);
  /* freopen("output.txt", "w", stdout); */
#endif

  int n, x, sum, max, min, c = 0;
  while (~scanf("%d", &n) && n) {
    sum = 0;
    max = 0;
    min = 1000;
    for(int i=0; i<n; i++) {
      scanf("%d", &x);
      sum += x;
      max = x > max ? x : max;
      min = x < min ? x : min;
    }
    if (c) printf("\n");
    printf("Case %d: %d %d %.3f\n", ++c, min, max, (double)sum/n);
  }
  return 0;
}
