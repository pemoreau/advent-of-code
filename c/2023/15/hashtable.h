//
// Created by pem on 16/12/2023.
//

#ifndef ADVENT_OF_CODE_HASHTABLE_H
#define ADVENT_OF_CODE_HASHTABLE_H

typedef struct _Entry {
    char *key;
    int value;
    struct _Entry *next;
} Entry;

typedef struct _Hashtable {
    Entry **entries;
    int size;
} Hashtable;

Hashtable *newHashtable();

int hash(char *key);

void put(Hashtable *ht, char *key, int value);

void removeEntry(Hashtable *ht, char *key);

void print_table(Hashtable *ht);
void print_entries(Entry *entry);

#endif //ADVENT_OF_CODE_HASHTABLE_H
