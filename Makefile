# Makefile

GOPATH := ${PWD}/modules

ALL    := test01 test02

# test01: minimal example, using
#  function in the same file
#  function in another file, same package
#  function in another package
#  function written in C

# test01: go calls C: dlsym (dlfcn.h)

clean:
	rm -f ${ALL} 2>/dev/null || true
	rm -rf ~/.cache/go-build/ 2>/dev/null || true

%:
	rm -f ./$@ 2>&1 || true
	cd $@_pck/src; GOPATH='${GOPATH}' go build -o ../../$@ .
	./$*
