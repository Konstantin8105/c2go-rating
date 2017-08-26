# c2go-rating
Rating for https://github.com/elliotchance/c2go

C code base:
* https://cis.temple.edu/~giorgio/cis71/code/
* http://www.iu.hio.no/~mark/CTutorial/CTutorial.html
* https://www.cs.cmu.edu/~quake/triangle.html
* http://www.sqlite.org/cgi/src/doc/trunk/README.md

## Instruction

Now, that software checked only on linux system and may be not work on Win, Darwin.

1.	go get -u github.com/Konstantin8105/c2go-rating
2.	copy folders from $GOPATH/src/github.com/Konstantin8105/c2go-rating/SingleCcode to $GOPATH/bin
3.	copy folders from $GOPATH/src/github.com/Konstantin8105/c2go-rating/sqlite to $GOPATH/bin
4.	run `./c2go-rating`

## Table with results

| Date | Result |
|---|---|
| 22.06.2017 | 45 is Ok at 64 source c files |
| 24.06.2017 | 62 is Ok at 87 source c files |
| 29.06.2017 | Added "comma" problem for FOR. 63 is Ok at 87 soure c files | 
|   | File "factors1.c" is Ok |
|   | File "sieve.c" have 1 error instand of 2 |
| 03.07.2017 | Added "comma" problem for VARIABLES. 64 is Ok at 87 soure c files | 
|   | File "fibo.c" is Ok |
|   | File "counts.c" have 1 error instand of 3 |
| 06.07.2017 | Add new file "triangle.c" |
|   | 64 is Ok at 88 source c files |
| 06.07.2017 | Add analising of 2 sqlite source code |
|   | 64 is Ok at 90 source c files |
| 08.07.2017 | No change |
|   | 64 is Ok at 90 source c files |
| 08.07.2017 | After rebase to upstream master c2go |
|   | 65 is Ok at 90 source c files |
|   | File 'counts.c' have 1 error instand of 3 |
|   | File "fibo.c" is Ok |
| 26.08.2017 | After big Elliot PR |
|   | Before: 25 mistakes in short file tests |
|   | After : just 5 mistakes |


## Present list of mistake files

```
Amount mistake source by gcc:  0
Amount mistake c2go results:  5
	Mistake file :  SingleCcode/27.c
		Error:  panic: (*ast.FunctionDecl) (0x589700,0xc420180400)
	Mistake file :  SingleCcode/31.c
		Error:  panic: (*ast.FunctionDecl) (0x589700,0xc4204b6000)
	Mistake file :  SingleCcode/kilo.c
		Error:  panic: invalid identity: 'NotNoarch.File' [recovered]
	Mistake file :  SingleCcode/triangle.c
		Error:  panic: invalid identity: 'insertvertexresult ()' [recovered]
	Mistake file :  sqlite/shell.c
		Error:  panic: unknown node type: 'Field 0x1e8f970 '' 'union sigcontext::(anonymous at /usr/include/x86_64-linux-gnu/bits/sigcontext.h:165:17)''
Result: 85 is Ok at 90 source c files - 94.444 procent.
```
