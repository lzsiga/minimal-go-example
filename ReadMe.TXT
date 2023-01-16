Oh well, I know nothing about Golang,
but maybe an example created by an outsider
will be useful for other outsiders.

QuickStart:
    make clean
    make test01
    ./test01 *

    make test01
    ./test02

test01: minimal example, using
  function in the same file
  function in another file, same package
  function in another package
  function written in C

test01: go calls C: dlsym (dlfcn.h)
