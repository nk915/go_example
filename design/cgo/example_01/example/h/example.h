#ifndef EXAMPLE_H
#define EXAMPLE_H

#include <string.h>
#include <stdlib.h>

#define MAX 10

typedef struct {
    int id;
    char name[50];
} ExampleStruct;

void exampleFunction();
void printExampleStruct(ExampleStruct es);

//typedef boolean (*fnVCScan_ReloadPattern)(void *lpReloadPattern, unsigned long ulSize);
typedef void 	(*fnVCScan_Clean)();

#endif // EXAMPLE_H
