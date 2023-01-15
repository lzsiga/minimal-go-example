# Makefile

GOPATH := ${PWD}/modules
	export GOPATH

ALL     := test01

clean:
	rm -f ${ALL} 2>/dev/null || true
	rm -rf ~/.cache/go-build/ 2>/dev/null || true

test01: test01_pck/src/main.go \
        test01_pck/src/main2.go \
        modules/src/echo_args/echo_args.go \
        modules/src/minimath/minimath.go \

%:	
	rm -f ./$* 2>&1 || true
	cd $*_pck/src/; go build -o ../../$* *.go
	./$*
