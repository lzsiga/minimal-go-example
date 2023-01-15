// gcd.go

package minimath

func Gcd(a int64, b int64) int64 {
    if a<0 { a = -a }
    if b<0 { b = -b }
    if b>a {
        var tmp int64 = b
        b = a
        a = tmp
    }
    for b!=0 {
        var tmp int64 = a%b;
        a = b
        b = tmp
    }
    return a;
}

func Lcm(a int64, b int64) int64 {
    if a<0 { a = -a; }
    if b<0 { b = -b; }
    if b>a {
        var tmp int64 = b;
        b = a;
        a = tmp;
    }
    if b==0 {
        return 0
    } else {
        return a/Gcd(a,b)*b
    }
}
