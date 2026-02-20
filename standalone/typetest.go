package main

import ("fmt"; "strconv")

// typecast between 'int' and 'float': explicit

// strconv: ParseFloat, FormatFloat
//          cf: https://pkg.go.dev/strconv#ParseFloat
//          cf: https://pkg.go.dev/strconv#FormatFloat

// division: integer/integer is integer
//           float/integer is float

// use global variable before declaration: possible

func main() {
   fmt.Printf("testing conversions\n")
   var i1 int32 = 1234
// var f1 float64 = i1 -- cannot use i1 (variable of type int32) as float64 value in variable declaration
   var f1 float64 = float64(i1)
   var i2 int32 = 13/5     // result is 2
   var f2 float64 = 13/5   // result is 2.0
   var f3 float64 = 13.0/5 // result is 2.6
   fmt.Printf("i1=%d f1=%f i2=%d f2=%f f3=%f pie=%f\n", i1, f1, i2, f2, f3, pie)

   var s string = "12.45E4"
   var fs float64
   fs, _ = strconv.ParseFloat(s, 64) // second parameter: bitsize
   fmt.Printf("ParseFloat(\"%s\") is %f\n", s, fs);

   fmt.Printf("FormatFloat(%f,'G',3,64) is \"%s\"\n", fs, strconv.FormatFloat(fs, 'G', 3, 64))
   fmt.Printf("FormatFloat(%f,'G',1,64) is \"%s\"\n", fs, strconv.FormatFloat(fs, 'G', 1, 64))
}

var pie float64 = 3.14159 // declared _after_ usage: allowed
