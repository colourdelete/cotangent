#!/bin/env python3
from cffi import FFI

ffibuilder = FFI()

ffibuilder.set_source(
    "_cotangent",
    """
#include "../cotangent.h"
""",
    libraries=["../cotangent.a"],
)
ffibuilder.cdef(
    """
struct RenderAndSanitize_return {
	char* r0; /* res */
	int r1; /* resLen */
};

struct RenderAndSanitize_return RenderAndSanitize(char* src, int srcLen);
"""
)


if __name__ == "__main__":
    ffibuilder.compile(verbose=True)
