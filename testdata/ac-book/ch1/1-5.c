// Discount
#include<stdio.h>

int main()
{
  int n, price;
  scanf("%d", &n);
  price = n*95;
  if (price >= 300)
    printf("%.2f\n", price*0.85);
  else
    printf("%.2f\n", price*1.0);
  return 0;
}
