#include<stdio.h>

int main()
{
  int a, b, c, k = 0;
  while(~scanf("%d %d %d", &a, &b, &c)) {
    int i;
    for(i=10; i<=100; i++)
      if(i%3 == a && i%5 == b && i%7 == c) {
        printf("Case %d: %d\n", ++k, i);
        break;
      }
    if(i == 101)
      printf("Case %d: No answer\n", ++k);
  }
  return 0;
}
