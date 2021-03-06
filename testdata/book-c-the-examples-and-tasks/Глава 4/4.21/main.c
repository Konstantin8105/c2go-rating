#include <stdio.h>
#include <stdlib.h>

/* У главной функции есть аргументы */
int main(int argc, char *argv[])
{
    printf("Выполнение программы %s\n", argv[0]);
    printf("Вычисление процентного дохода.\n");

    /* Индексная переменная */
    int k;

    /* Переменная для записи итоговой суммы */
    double s;

    /* Начальная сумма */
    double m;

    /* Процентная ставка */
    double r;

    /* Время размещения депозита */
    int t;

    /* Проверка количества аргументов */
    if (argc != 4)
        printf("Параметры указаны неверно!\n");
    else {
        m = atof(argv[1]); // начальная сумма
        r = atof(argv[2]); // процентная ставка
        t = atof(argv[3]); // время размещения

        /* Вычисление итоговой суммы */
        s = m;

        for (k = 1; k <= t; k++)
            s *= (1 + r / 100);

        printf("Начальная сумма: %.2f\n", m);
        printf("Процентная ставка: %.2f\n", r);
        printf("Время размещения: %d\n", t);
        printf("Итоговая сумма: %.2f\n", s);
    }

    printf("Программа завершила приложение.\n");

    return 0;
}