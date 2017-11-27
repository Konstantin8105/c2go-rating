// More arithmetic expressions

#include <stdio.h>

int main(void)
{
    int a = 25, b = 2;
    float c = 25.0, d = 2.0;

    printf("6 + a / 5 * b = %i\n", 6 + a / 5 * b);
    printf("a / b * b = %i\n", a / b * b);
    printf("c / d * d = %f\n", c / d * d);
    printf("-a = %i\n", -a);

    return 0;
}

/* Output

6 + a / 5 * b = 16
a / b * b = 24
c / d * d = 25.000000
-a = -25

*/
