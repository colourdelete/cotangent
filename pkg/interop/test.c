#include "cotangent.h"
#include <stdio.h>

int main(int argc, char **argv) {
  struct RenderAndSanitize_return r = RenderAndSanitize("# Hello, world!\n", 15);
  printf("%d %s", r.r1, r.r0);
}
