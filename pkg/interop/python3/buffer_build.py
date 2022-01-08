#!/bin/env python3

from cffi import FFI

ffibuilder = FFI()

ffibuilder.cdef("char* render_and_sanitize(char* src);")

ffibuilder.set_source(
    "_buffer",  # name of the output C extension
    """
    #include "buffer.h"
""",
    sources=["buffer.c", "../cotangent.a"],
)

if __name__ == "__main__":
    ffibuilder.compile(verbose=True)
