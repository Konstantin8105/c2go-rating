#include <stdio.h>
#include <stdlib.h>


/*
 * Передача аргументов одномерного массива
 * Функция для отображения содержимого массива
 */
void show(int* nums,int n)
{
    int k;

    for (k = 0; k < n; k++)
        printf("| %d ", nums[k]);

    printf("|\n");
}

/* Функция для заполнения массива случайными числами */
void set(int* nums, int n, int a, int b) {
    /* Индексная переменная */
    int k;

    for (k = 0; k < n; k++)
        nums[k] = a + rand() % (b - a + 1);
}

/* Функция для вычисления суммы элементов массива */
int sum(int* nums,int n)
{
    /* Индексная переменная и переменная для записи суммы */
    int k, s = 0;

    for (k = 0; k < n; k++)
        s += nums[k];

    return s;
}

int main(void)
{
    /* Инициализация генератора случайных чисел */
    srand(5);

    /* Размер массива */
    int n = 12;

    /* Объявление массива */
    int numbers[n];

    /* Заполнение массива случайными числами (от 1 до 10) */
    set(numbers, n, 1, 10);

    /* Отображение содержимого массива */
    show(numbers, n);

    printf("Сумма значений элементов: %d\n", sum(numbers, n));

    /* Отображение части (первой половины) массива */
    show(numbers, n / 2);

    /* Отображение части (второй половины) массива */
    show(numbers + n / 3, n / 3);

    /* Новые значения для массива (случайные числа с -9 до 9) */
    set(numbers, n, -9, 9);

    /* Отображение содержимого массива */
    show(numbers, n);

    /* Оъявление и инициализация массива */
    int m[] = {1, 2, 3, 4, 5};

    /* Отображение содержимого массива */
    show(m, 5);

    printf("Сумма значений элементов: %d\n", sum(m, 5));

    /* Размер для динамического массива */
    int size = 10;

    /* Создание динамического массива */
    int *p = malloc(size * sizeof(int));

    /* Заполнение динамического массива случайными числами от 0 до 9 */
    set(p, size, 0, 9);

    /* Отображение содержимогодинамического массива */
    printf("Сумма значений элементов: %d\n",sum(p, size));

    /* Освобождение памяти */
    free(p);

    return 0;
}