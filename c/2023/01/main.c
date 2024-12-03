//
// Created by pem on 14/12/2023.
//

#include "main.h"

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define INPUTNAME "input.txt"

#define IS_DIGIT(c) ((c) >= '0' && (c) <= '9')
#define BUF_SIZE 128
int findDigit(char* s) {
    int first = -1;
    int last = -1;

    while (*s) {
        if (IS_DIGIT(*s)) {
            if (first < 0) {
                first = *s - '0';
            }
            last = *s - '0';
        }
        s++;
    }
    return 10 * first + last;
}

int is_prefix(char *p, char *s) {
    while (*p) {
        if (*p != *s) {
            return 0;
        }
        p++;
        s++;
    }
    return 1;
}

int get_digit(char *s) {
    char **digits = (char *[]){"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
    if (IS_DIGIT(*s)) {
        return *s - '0';
    }
    for (int i = 0; i < 10; i++) {
        if (is_prefix(digits[i], s)) {
            return i;
        }
    }
    return -1;
}

int findDigitLetter(char* s) {
    int first = -1;
    int last = -1;

    while (*s) {
        int d = get_digit(s);
        if (d >= 0) {
            if (first < 0) {
                first = d;
            }
            last = d;
        }
        s++;
    }
    return 10 * first + last;
}

int part1(char *filename) {
    char buf[BUF_SIZE];

    FILE* f = fopen(filename, "r");
    if (f == NULL) {
        printf("Error opening file!\n");
        exit(1);
    }

    int res = 0;
    while (fgets(buf, sizeof(buf), f) != NULL) {
        res += findDigit(buf);
    }

    fclose(f);
    return res;
}

int part2(char *filename) {
    char buf[BUF_SIZE];

    FILE* f = fopen(filename, "r");
    if (f == NULL) {
        printf("Error opening file!\n");
        exit(1);
    }

    int res = 0;
    while (fgets(buf, sizeof(buf), f) != NULL) {
        res += findDigitLetter(buf);
    }

    fclose(f);
    return res;
}

int main(int argc, char* argv[]) {
    printf("Day 01 of 2023 Advent of Code\n");
    clock_t t = clock();
    printf("Part 1: %d\n", part1(INPUTNAME));
    printf("%f ms\n", 1000 * ((double)(clock() - t))/CLOCKS_PER_SEC);

    t = clock();
    printf("Part 2: %d\n", part2(INPUTNAME));
    printf("%f ms\n", 1000 * ((double)(clock() - t))/CLOCKS_PER_SEC);

    return 0;
}
