/*
 * This is a RANDOMLY GENERATED PROGRAM.
 *
 * Generator: csmith 2.2.0
 * Git version: dcef523
 * Options:   (none)
 * Seed:      43448120
 */

#include "csmith.h"


static long __undefined;

/* --- Struct/Union Declarations --- */
struct S0 {
   volatile signed f0 : 10;
   const volatile signed f1 : 23;
   volatile unsigned f2 : 29;
};

/* --- GLOBAL VARIABLES --- */
static struct S0 g_3[1][8] = {{{-20,-464,1490},{-17,-373,4158},{-20,-464,1490},{-20,-464,1490},{-17,-373,4158},{-20,-464,1490},{-20,-464,1490},{-17,-373,4158}}};
static struct S0 g_5[7] = {{-21,-2387,2889},{-21,-2387,2889},{-21,-2387,2889},{-21,-2387,2889},{-21,-2387,2889},{-21,-2387,2889},{-21,-2387,2889}};
static struct S0 *g_4 = &g_5[3];


/* --- FORWARD DECLARATIONS --- */
static const int64_t  func_1(void);


/* --- FUNCTIONS --- */
/* ------------------------------------------ */
/* 
 * reads : g_5.f2
 * writes: g_4
 */
static const int64_t  func_1(void)
{ /* block id: 0 */
    struct S0 *l_2[8];
    int i;
    for (i = 0; i < 8; i++)
        l_2[i] = &g_3[0][1];
    g_4 = l_2[2];
    return g_5[3].f2;
}




/* ---------------------------------------- */
int main (int argc, char* argv[])
{
    int i, j;
    int print_hash_value = 0;
    if (argc == 2 && strcmp(argv[1], "1") == 0) print_hash_value = 1;
    platform_main_begin();
    crc32_gentab();
    func_1();
    for (i = 0; i < 1; i++)
    {
        for (j = 0; j < 8; j++)
        {
            transparent_crc(g_3[i][j].f0, "g_3[i][j].f0", print_hash_value);
            transparent_crc(g_3[i][j].f1, "g_3[i][j].f1", print_hash_value);
            transparent_crc(g_3[i][j].f2, "g_3[i][j].f2", print_hash_value);
            if (print_hash_value) printf("index = [%d][%d]\n", i, j);

        }
    }
    for (i = 0; i < 7; i++)
    {
        transparent_crc(g_5[i].f0, "g_5[i].f0", print_hash_value);
        transparent_crc(g_5[i].f1, "g_5[i].f1", print_hash_value);
        transparent_crc(g_5[i].f2, "g_5[i].f2", print_hash_value);
        if (print_hash_value) printf("index = [%d]\n", i);

    }
    platform_main_end(crc32_context ^ 0xFFFFFFFFUL, print_hash_value);
    return 0;
}

/************************ statistics *************************
XXX max struct depth: 0
breakdown:
   depth: 0, occurrence: 2
XXX total union variables: 0

XXX non-zero bitfields defined in structs: 3
XXX zero bitfields defined in structs: 0
XXX const bitfields defined in structs: 1
XXX volatile bitfields defined in structs: 3
XXX structs with bitfields in the program: 2
breakdown:
   indirect level: 0, occurrence: 0
   indirect level: 1, occurrence: 2
XXX full-bitfields structs in the program: 0
breakdown:
XXX times a bitfields struct's address is taken: 2
XXX times a bitfields struct on LHS: 0
XXX times a bitfields struct on RHS: 0
XXX times a single bitfield on LHS: 0
XXX times a single bitfield on RHS: 1

XXX max expression depth: 1
breakdown:
   depth: 1, occurrence: 3

XXX total number of pointers: 2

XXX times a variable address is taken: 2
XXX times a pointer is dereferenced on RHS: 0
breakdown:
XXX times a pointer is dereferenced on LHS: 0
breakdown:
XXX times a pointer is compared with null: 0
XXX times a pointer is compared with address of another variable: 0
XXX times a pointer is compared with another pointer: 0
XXX times a pointer is qualified to be dereferenced: 0
XXX number of pointers point to pointers: 0
XXX number of pointers point to scalars: 0
XXX number of pointers point to structs: 2
XXX percent of pointers has null in alias set: 0
XXX average alias set size: 1

XXX times a non-volatile is read: 1
XXX times a non-volatile is write: 1
XXX times a volatile is read: 1
XXX    times read thru a pointer: 0
XXX times a volatile is write: 0
XXX    times written thru a pointer: 0
XXX times a volatile is available for access: 2
XXX percentage of non-volatile access: 66.7

XXX forward jumps: 0
XXX backward jumps: 0

XXX stmts: 2
XXX max block depth: 0
breakdown:
   depth: 0, occurrence: 2

XXX percentage a fresh-made variable is used: 66.7
XXX percentage an existing variable is used: 33.3
FYI: the random generator makes assumptions about the integer size. See platform.info for more details.
********************* end of statistics **********************/

