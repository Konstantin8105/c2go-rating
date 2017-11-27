#include<stdio.h>

int main()
{
  FILE *fin, *fout;
  fin = fopen("input.txt", "rb");
  fout = fopen("output.txt", "wb");

  int x, sum = 0, count = 0, max = 0, min = 1000;
  while (fscanf(fin, "%d", &x) > 0) {
    count++;
    sum += x;
    max = x > max ? x : max;
    min = x < min ? x : min;
  }
  fprintf(fout, "%d %d %.3f\n", min, max, (double)sum/count);
  fclose(fin);
  fclose(fout);
  return 0;
}
