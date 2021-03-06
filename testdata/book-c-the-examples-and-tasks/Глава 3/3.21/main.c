#include <stdio.h>
#include <stdlib.h>

/* Создание "треугольного" массива */
int main(void)
{
    /* Количество строк в массиве */
    int m;

    /* Индексная переменная и переменная счётчик */
    int i, j, count = 0;

    printf("Укажите количество строк: ");
    scanf("%d", &m);

    /* Создание одномерного динамического массива указателей на целые числа */
    int **nums = malloc(m * sizeof(int*));

    /* Создание строк для двумерного массива */
    for (i = 0; i < m; i++)
        nums[i] = malloc((i + 1) * sizeof(int));

    /* Заполнение массива и отображение значений элементов массива */
    for (i = 0; i < m; i++) {
        for (j = 0; j < i + 1; j++) {
            count++;

            /* Присваивание значения элементу */
            nums[i][j]=count;

            printf("%3d", nums[i][j]);
        }

        printf("\n");
    }

    /* Удаление строк в двумерном массиве */
    for (i = 0; i < m; i++)
        free(nums[i]);

    /* Удаление одномерного массива указателей */
    free(nums);

    return 0;
}