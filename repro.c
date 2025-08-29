#include "repro.h"
#include <stdlib.h>

char* allocate() {
  return (char*)malloc(100);
}

void free_allocated(char* c) {
  free(c);
}