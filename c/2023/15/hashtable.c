//
// Created by pem on 16/12/2023.
//

#include "hashtable.h"

#include <stdlib.h>
#include <string.h>
#include <stdio.h>

#define INITIAL_SIZE 256

Hashtable *newHashtable() {
    Hashtable *ht = malloc(sizeof(Hashtable));
    ht->size = INITIAL_SIZE;
    ht->entries = malloc(sizeof(Entry *) * ht->size);
    memset(ht->entries, 0, sizeof(Entry *) * ht->size);
    return ht;
}

void print_entries(Entry *entry) {
    while (entry) {
        printf("key = %s, value = %d\n", entry->key, entry->value);
        entry = entry->next;
    }
}

void print_table(Hashtable *ht) {
    for (int i = 0; i < ht->size; i++) {
        if (!ht->entries[i]) {
            continue;
        }
        printf("i = %d --> ", i);
        print_entries(ht->entries[i]);
    }
}

int hash(char *key) {
    int h = 0;
    while (*key) {
        h += *key;
        h *= 17;
        h &= 0xff;
        key++;
    }
    return h;
}

void put(Hashtable *ht, char *key, int value) {
    int h = hash(key) % ht->size;

    Entry *e = ht->entries[h];
    Entry *prev = NULL;
    while (e) {
        if (strcmp(e->key, key) == 0) {
            e->value = value;
            return;
        }
        prev = e;
        e = e->next;
    }

    Entry *newEntry = malloc(sizeof(Entry));
    newEntry->key = malloc(strlen(key) + 1);
    strcpy(newEntry->key, key);
    newEntry->value = value;
    newEntry->next = NULL;
    if (prev) {
        prev->next = newEntry;
    } else {
        ht->entries[h] = newEntry;
    }
}

void removeEntry(Hashtable *ht, char *key) {
    int h = hash(key) % ht->size;

    Entry *e = ht->entries[h];
    Entry *prev = NULL;
    if (!e) {
        return;
    }
    if (strcmp(e->key, key) == 0) {
        ht->entries[h] = e->next;
        free(e->key);
        free(e);
        return;
    }
    prev = e;
    e = e->next;
    while (e) {
        if (strcmp(e->key, key) == 0) {
            prev->next = e->next;
            free(e);
            return;
        }
        prev = e;
        e = e->next;
    }
}
