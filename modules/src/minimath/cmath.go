// cmath.go

package minimath

/*
double mypow (double base, long p, long q);
#cgo LDFLAGS: -lm
*/
import "C"

func Mypow(base float64, p int64, q int64) float64 {
    var retval C.double
    retval = C.mypow(C.double(base), C.long(p), C.long(q))
    return float64(retval)
}
