/* stringtest.go */

package main

import ("fmt"; "os"; "strings"; "unicode/utf8")

func Test1 (s string) {
    fmt.Printf ("'%s' len=%d ToLopwer='%s' ToUpper='%s'\n", s,
        len(s), strings.ToLower(s), strings.ToUpper(s))
}

func Test2 (s string) {
    var len= len(s)
    fmt.Printf("\nTest2: process characters of '%s' bytewise (bytelen=%d)\n", s, len);
    var i int
    for i=0; i<len; i++ {
        var c= s[i:i+1]
        fmt.Printf ("[%d]='%s' code=%d 0x%x\n", i, c, c[0], c[0]);
    }
}

func Test3 (s string) {
    var len= utf8.RuneCountInString(s)
    fmt.Printf("\nTest3: process characters of '%s' characterwise (length=%d)\n", s, len);
    var pos int
    var code rune
    for pos, code = range s {
        fmt.Printf("pos %d: %s (code %d 0x%x)\n", pos, string(code), code, code)
    }
}

func reverseBytes (s string) string {
    var result string= ""
    var len= len(s)
    for i:=0; i<len; i++ {
        var c= s[i:i+1]
        result= c+result
    }
    return result
}

func reverseString (s string) string {
    var result string= ""
    var code rune
    for _,code = range s {
        result= string(code)+result
    }
    return result
}

func main () {
    var progname string = os.Args[0]
    var nargs int = len(os.Args)

    fmt.Printf("%s is running\n", progname)
    if nargs==1 {
        Test1 ("Árvíztűrő tükörfúrógép")
        Test1 ("ÁRVÍZTŰRŐ TÜKÖRFÚRÓGÉP")
        var s1= "tüzér\xc5\xb1l\u0151";
        Test2 (s1);
        Test3 (s1);
        var s2= "\xf0\x9f\x8d\x94";
        Test2 (s2);
        Test3 (s2);
    } else {
        var i int
        for i=1; i<nargs; i++ {
            var s string = os.Args[i]
            Test1 (s)
            Test2 (s)
            fmt.Printf ("Bytes reverse of '%s' is '%s'\n", s, reverseBytes(s))
            fmt.Printf ("String reverse of '%s' is '%s'\n", s, reverseString(s))
        }
    }
}
