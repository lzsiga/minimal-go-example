Note:
  file-names 'foobar_test.go' don't work
  (as they have speciel meaning),
  use file-name like 'foobartest.go' instead

go build ReadStdin.go  && ./ReadStdin
go build gcd.go        && ./gcd
go build randtest.go   && ./randtest
go build switchtest.go && ./switchtest
go build ptrtest.go    && ./ptrtest     # it has pointers (including pointer-pointers and array-pointers)
go build stringtest.go && ./stringtest
gi build q8.go         && ./q8          # it has structure and pointer as function-paramater
