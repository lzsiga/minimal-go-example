# Makefile

GOPATH := ${PWD}/modules

ALL     := test01 test02

clean:
	rm -f ${ALL} 2>/dev/null || true
	rm -rf ~/.cache/go-build/ 2>/dev/null || true

%:
	rm -f ./$@ 2>&1 || true
	cd $@_pck/src; GOPATH='${GOPATH}' go build -o ../../$@ .
	./$*
