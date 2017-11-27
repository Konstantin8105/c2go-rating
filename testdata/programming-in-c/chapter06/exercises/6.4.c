// A simple "printing" calculator

#include <stdio.h>

int main(void)
{
    float value, accumulator;
    char operator;

    printf("Enter E as the operator to quit.\n");
    printf("Set an initial accumulator value (ex: 0 S): ");
    do {
        scanf("%f %c", &value, &operator);

        switch (operator) {
            case 'S':
                accumulator = value;
                printf("= %f\n", accumulator);
                break;
            case '+':
                accumulator += value;
                printf("= %f\n", accumulator);
                break;
            case '-':
                accumulator -= value;
                printf("= %f\n", accumulator);
                break;
            case '*':
                accumulator *= value;
                printf("= %f\n", accumulator);
                break;
            case '/':
                accumulator /= value;
                printf("= %f\n", accumulator);
                break;
            default:
                printf("Unrecognized operator.\n");
                break;
        }

    } while(operator != 'E');

    return 0;
}
