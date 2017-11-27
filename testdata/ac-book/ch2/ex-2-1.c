#include<stdio.h>
#include<math.h>

int main()
{
  int n;
  for(int a=1; a <= 9; a++)
  {
    for(int b=0; b <= 9; b++)
    {
      int aabb = 1100*a + 11*b;
      int s = floor(sqrt(aabb)+0.5);
      if (s*s==aabb) printf("%d\n", aabb);
    }
  }
  for(int x=1; ; x++)
  {
    int n = x*x;
    if (n<1000) continue;
    if (n>9999) break;
    if ( (n/100)%11 == 0 && (n%100)%11 == 0 ) printf("%d\n", n);
  }
  return 0;
}
