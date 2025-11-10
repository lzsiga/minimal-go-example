# Makefile

GOPATH := ${PWD}/modules
GO111MODULE := off

export GOPATH GO111MODULE

ALL    := test01 test02

all: ${ALL}
	@printf "Your binary files:\n"
	@ls -l ${ALL}
	@printf "\nTo see the binary files created behind your back:\n"
	@printf "(cd ~/.cache/go-build/; find . -type f -exec ls -l {} +)\n"

# test01: minimal example, using
#  function in the same file
#  function in another file, same package
#  function in another package
#  function written in C

# test02: go calls C: dlsym (dlfcn.h)

clean:
	rm -f ${ALL} 2>/dev/null || true
	rm -rf ~/.cache/go-build/ 2>/dev/null || true

%:
	rm -f ./$@ 2>&1 || true
	cd $@_pck/src; GOPATH='${GOPATH}' go build -o ../../$@ .
	./$*
