#include<stdio.h>

const int C = 1*2*3*4*5*6*7*8*9;

int main()
{
  for(int i=123; i<334; i++) {
    int abc = i;
    int def = i*2;
    int ghi = i*3;
    int s = 1;
    for(int j=1; j<=100; j*=10) {
      s *= (abc%(j*10)) / j;
      s *= (def%(j*10)) / j;
      s *= (ghi%(j*10)) / j;
    }
    if(s==C) printf("%d %d %d\n", abc, def, ghi);
  }
  return 0;
}
