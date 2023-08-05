#include <stdio.h>

static char* (*handler)(char* path) = 0;

void registerHandler(char* (*h)(char* path)) {
    handler = h;
    printf("Handler registered\n");
}

char* handleRequest(char* path) {
    if (handler == 0) {
        // Handle error: no handler registered
        return 0;
    }
    return handler(path);
}
