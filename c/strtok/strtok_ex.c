#include <stdio.h>
#include <string.h>

int main()
{
    char values[] = "foo,bar,zoo";

    printf("Tokenizing values=%s, delimiter=%s\n", values, ",");
    char *token = strtok(values, ",");

    /*
     * strtok can modify the original string. if the delimiter is found,
     * then a null terminator is inserted in place of th delimiter.
     * captures from gdb
        (gdb) p token
        $1 = 0x7fffffffe340 "foo"
        (gdb) p &values
        $2 = (char (*)[32]) 0x7fffffffe340
        (gdb)
     */
    printf("first token=%s, original values=%s\n", token, values);

    /*
     (gdb) x/20xb token
     0x7fffffffe340:	0x66	0x6f	0x6f	0x00	0x62	0x61	0x72	0x2c
     0x7fffffffe348:	0x7a	0x6f	0x6f	0x00	0x00	0x00	0x00	0x00
     0x7fffffffe350:	0x00	0x00	0x00	0x00
     */
    printf("addr(first token)=%p, addr(original values)=%p\n", token, values);  
    while (token != NULL) {
        printf("token=%s\n", token);
        token = strtok(NULL, ",");
    }

    // Tokenzing a string with multiple delimiters
    char values_with_multiple_delims[] = "quick,brown;fox,jumps;over,";
    char *delimiters = ";,";
    printf("Tokenizing values=%s, delimiters=%s\n", values_with_multiple_delims, delimiters);
    token = strtok(values_with_multiple_delims, delimiters);
    while (token != NULL) {
        printf("token=%s\n", token);
        /*
         * delimiter specified in the successive call need not be same as the one passed
         * initially. strtok just scans forward for the delim specified in the current call.
         */
        token = strtok(NULL, delimiters);
    }

    return 0;
}