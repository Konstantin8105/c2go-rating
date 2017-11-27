#include <stdio.h>

int main(void)
{
    float fahrenheit = 27, celcius;

    celcius = (fahrenheit - 32.0) / 1.8;
    printf("%.0f degrees fahrenheit is %.2f degrees in celcius.\n", fahrenheit, celcius);

    return 0;
}
