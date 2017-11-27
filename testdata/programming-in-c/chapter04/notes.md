# Variables

Pick names that reflect the intended use of the variable.

## Int

Whole numbers.
printf: %i

Extra:

    - Allowed to be expressed in a base other than 10.

    - If first digit is 0, *octal* notation is used (base8).
    printf: %o, or %#o to format and include leading 0.

    - If first symbols are 0x, *hexadecimal* is used (base16).
    printf: %X, %#X

## Float

Numbers containing decimal places.
printf: %f
printf: %2f
printf: %.2f

## Double

Can store roughly twice as many significant digits as float.

%f, %e, or %g can be used to display a double.

## Char

Stores a single character.
printf: %c

Do not confuse a character constant, which is a single character
enclosed in single quotes, with a character string, which is any
number of characters enclosed in double quotes.

'\n' is valid. The backslash is a special character in the C system
and does not actually count as a character.

## \_Bool

Stores 0 or 1.
By convention, 0 is false, 1 is true.

<stdbool.h> allows you to use true or false instead.

# Floating-point Imprecision

Certain floating-point values cannot be exactly represented
inside the computer's memory, there is only so much memory in
a computer. System automatically defaults to round to 6 decimal
places.

# Type Specifiers

## Long

On many systems, an int and a long int both have the same range and
either can be used to store integer values up to 32-bits wide.

A constant value of type long int is formed by optionally appending
the letter L (upper- or lowercase) onto the end of an integer constant.

Ex: long int numberOfPoints = 131071100L;
printf: %1i, %1o, %1x

## Long Long

64-bits wide.

Ex: long long int maxAllowedStorage;
printf: %11i

## Long Double

Ex: long double US_deficit_2004;

A long double constant is written as a floating constant with the letter
l or L immediately following: 1.234e+7L
printf: %Lf, %L

## Short

short int tells the C compiler that the variable being declared is used to
store fairly small integer values. The use of short is primary to conserve
space.

You are guaranteed that the amount of space allocated for a short int will
not be less than 16 bits.

No explicit way to write a constant of type short int in C.

printf: %hi, %ho, %hx

## Unsigned

When a variable will be used to store only positive numbers.

Ex: unsigned int counter;

An unsigned int constant is formed by placed the letter u or U after the
constant: 0x0ffU

You can combine the letters U or L when writing an integer constant to
tell the compiler it is an unsigned long: 20000UL

When declaring variables to be of type long long int, long int, short int,
or unsigned int, you can omit the keyword int.

# Order of Operations

Expressions are evaluated the same way if you were to apply the basic
rules of algebra.

100 + 2 * 25 would equal 150 because to C, it is 100 + (2 * 25)
