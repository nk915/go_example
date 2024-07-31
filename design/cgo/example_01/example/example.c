#include "example.h"
#include <stdio.h>
#include <string.h>

void exampleFunction() {
    printf("Hello from C!\n");
}

void printExampleStruct(ExampleStruct es) {
    printf("ID: %d, Name: %s\n", es.id, es.name);
}
