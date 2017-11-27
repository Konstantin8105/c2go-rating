// Triangle
#include<stdio.h>
#include<math.h>

int main()
{
  int a, b, c;
  scanf("%d%d%d", &a, &b, &c);

  // sort
  if (a>b) { a = a+b; b = a-b; a = a-b; }
  if (a>c) { a = a+c; c = a-c; a = a-c; }
  if (b>c) { b = b+c; b = b-c; b = b-c; }

  if ( (a+b)<=c )
    printf("not a triangle\n");
  else
  {
    if ( (pow(a,2)+pow(b,2)) == pow(c,2) )
      printf("yes\n");
    else
      printf("no\n");
  }

  return 0;
}
