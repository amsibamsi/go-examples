#include <stdlib.h>

int* array(int length) {
  return (int*)malloc(sizeof(int)*length);
}

void copy(int* src, int* dst, int length) {
  for (int i = 0; i < length; i++) {
    dst[i] = src[i];
  }
}

int check(int* arr, int length) {
  for (int i = 0; i < length; i++) {
    if (arr[i] != i) {
      return 1;
    }
  }
  return 0;
}
