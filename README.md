# c4go-rating
Rating for https://github.com/Konstantin8105/c4go

C code base:
* https://cis.temple.edu/~giorgio/cis71/code/
* http://www.iu.hio.no/~mark/CTutorial/CTutorial.html
* https://www.cs.cmu.edu/~quake/triangle.html
* http://www.sqlite.org/cgi/src/doc/trunk/README.md
* https://github.com/duxing2007/books-examples
* https://github.com/roktas/apue2e
* https://github.com/caisah/K-and-R-exercises-and-examples
* https://github.com/eugenetriguba/programming-in-c
* https://github.com/olegbukatchuk/book-c-the-examples-and-tasks
* https://github.com/KushalP/k-and-r
* https://github.com/team6612/ac-book
* https://github.com/xavriley/c_programming_language_book
* https://github.com/Emmetttt/C-Deitel-Book

## For future

Projects (single file):

http://www.iro.umontreal.ca/~felipe/IFT2030-Automne2002/Complements/tinyc.c
https://github.com/shenfeng/tiny-web-server/blob/master/tiny.c
https://github.com/adamdunkels/ubasic
https://github.com/fabianishere/brainfuck/tree/master/src
https://github.com/zserge/partcl/blob/master/tcl.c
https://github.com/zserge/expr
https://github.com/phase/o/blob/master/o.c
https://github.com/kgabis/brainfuck-c
https://github.com/bbu/simple-interpreter/tree/master/src
https://github.com/nanoflite/basic
https://github.com/fragglet/yoctolisp
https://github.com/paladin-t/my_basic/tree/master/core
https://raw.githubusercontent.com/dtschump/CImg/master/CImg.h (cpp)
https://github.com/vurtun/nuklear
Collections:

http://www.ioccc.org/years.html
https://cis.temple.edu/~ingargio/cis71/code/
https://github.com/nothings/single_file_libs
Projects (multi file):

https://github.com/larmel/lacc
https://github.com/kagemusha666/interpreter/tree/master/src
https://github.com/steve-m/librtlsdr
https://github.com/skx/simple.vm
https://github.com/brackeen/ok-file-formats
https://github.com/aligrudi/neatcc
http://runtimeterror.com/tech/lil/
Weird edge cases

http://blog.robertelder.org/weird-c-syntax/
https://news.ycombinator.com/item?id=9991528


## Instruction

Now, that software checked only on linux system and may be not work on Win, Darwin.

1.	go get -u github.com/Konstantin8105/c4go-rating
2.	copy folders from $GOPATH/src/github.com/Konstantin8105/c4go-rating/SingleCcode to $GOPATH/bin
3.	copy folders from $GOPATH/src/github.com/Konstantin8105/c4go-rating/sqlite to $GOPATH/bin
4.	run `./c4go-rating`

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
| 08.07.2017 | After rebase to upstream master c4go |
|   | 65 is Ok at 90 source c files |
|   | File 'counts.c' have 1 error instand of 3 |
|   | File "fibo.c" is Ok |
| 26.08.2017 | After big Elliot PR |
|   | Before: 25 mistakes in short file tests |
|   | After : just 5 mistakes |
| 20.10.2017 | Many modifications |
|   | Success with sqlite source |
|   | 4 panic error |


## Present list of mistake files

```
Amount mistake source by gcc:  0
Amount mistake c4go results:  4
	Mistake file :  SingleCcode/27.c
		Error:  panic: (*ast.FunctionDecl) (0x73d7e0,0xc4204288f0)
	Mistake file :  SingleCcode/31.c
		Error:  panic: (*ast.FunctionDecl) (0x73d7e0,0xc4203b1340)
	Mistake file :  SingleCcode/kilo.c
		Error:  panic: invalid identity: 'NotNoarch.File' [recovered]
	Mistake file :  SingleCcode/triangle.c
		Error:  panic: invalid identity: 'insertvertexresult ()' [recovered]
Result: 86 is Ok at 90 source c files - 95.556 procent. 
```
