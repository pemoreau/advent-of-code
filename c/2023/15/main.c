//
// Created by pem on 15/12/2023.
//

#include "main.h"

#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "hashtable.h"

#define INPUTNAME "input.txt"
#define BUF_SIZE 128


char *read_file(char *filename) {
    FILE *f = fopen(filename, "r");
    if (f == NULL) {
        printf("Error opening file!\n");
        exit(EXIT_FAILURE);
    }
    fseek(f, 0, SEEK_END);
    size_t length = ftell(f);
    fseek(f, 0, SEEK_SET);

    char *buf = (char *) malloc(length + 1);

    size_t read = length > 0 ? fread(buf, 1, length, f) : 0;
    if (read != length) {
        free(buf);
        printf("[ERROR] Failed to read file\n");
        exit(EXIT_FAILURE);
    }

    fclose(f);
    return buf;
}

int part1(char *filename) {
    char *buf = read_file(filename);

    int res = 0;
    char *s = buf;
    int h = 0;
    while (*s) {
        if (*s == ',') {
            res += h;
            h = 0;
            s++;
            continue;
        }
        h += *s;
        h *= 17;
        h &= 0xff;
        s++;
    }
    res += h;
    return res;
}

int part2(char *filename) {
    char *buf = read_file(filename);

    Hashtable *ht = newHashtable();

    char *s = buf;
    char name[BUF_SIZE];
    char *pos = name;

    while (*s) {
        if (*s == ',') {
            s++;
            pos = name;
            continue;
        }
        if (*s == '-') {
            *pos = '\0';
            s++;
            removeEntry(ht, name);
            continue;
        }
        if (*s == '=') {
            *pos = '\0';
            s++;
            put(ht, name, *s - '0');
            s++;
            continue;
        }
        if (*s >= 'a' && *s <= 'z') {
            *pos = *s;
            s++;
            pos++;
            continue;
        }
        printf("unknown char %c\n", *s);
    }

    int res = 0;
    for (int i = 0; i < ht->size; i++) {
        Entry *e = ht->entries[i];
        for (int j = 0; e; j++) {
            res += (i + 1) * (j + 1) * e->value;
            e = e->next;
        }
    }

    return res;
}

#ifndef TEST

int main(int argc, char *argv[]) {
    printf("Day 15 of 2023 Advent of Code\n");
    clock_t t = clock();
    printf("Part 1: %d\n", part1(INPUTNAME));
    printf("%f ms\n", 1000 * ((double) (clock() - t)) / CLOCKS_PER_SEC);

    t = clock();
    printf("Part 2: %d\n", part2(INPUTNAME));
    printf("%f ms\n", 1000 * ((double) (clock() - t)) / CLOCKS_PER_SEC);

    return 0;
}

#endif

#ifdef TEST
#include <assert.h>
#include <string.h>
int main() {
    printf("Day 15 of 2023 Advent of Code (Test)\n");
    assert(part1("input_test.txt") == 1320);
    assert(part2("input_test.txt") == 145);
    assert(part1("input.txt") == 514394);
    assert(part2("input.txt") == 236358);
    return 0;
}
#endif
