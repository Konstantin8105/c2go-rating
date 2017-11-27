#include <stdio.h>

int main(void)
{
    float x, answer;

    x = 2.55;
    answer = 3 * (x * x * x) - 5 * (x * x) + 6;

    printf("x is %.2f.\n", x);
    printf("3x^3 - 5x^2 + 6 = %.2f", answer);

    return 0;
}
