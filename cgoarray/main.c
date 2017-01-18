#include <stdlib.h>

// Returned array must be freed manually.
int* get_array(int* len) {
  int* array;
  *len = 3;
  array = (int*)malloc(sizeof(int) * (*len));
  for (int i = 0; i < (*len); i++) {
    *(array+i) = i;
  }
  return array;
}
