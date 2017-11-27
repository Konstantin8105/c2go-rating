#include<stdio.h>

int main()
{
  int m, n, k = 0;
  double r;
  while(~scanf("%d %d", &n, &m) && n && m) {
    r = 0;
    for(int i=n; i<=m; i++) {
      double k = (double)i;
      r += 1./(k*k);
    }
    printf("Case %d: %.5f\n", ++k, r);
  }
  return 0;
}
