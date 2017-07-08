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


## Present list of mistake files

```
Amount mistake source by gcc:  0
Amount mistake c2go results:  25
	Mistake file :  SingleCcode/17-2.c
		Error:  panic: format.Node internal error (107:10: expected '==', found '=')
	Mistake file :  SingleCcode/19.c
		Error:  panic: format.Node internal error (138:37: expected ')', found '+=' (and 1 more errors))
	Mistake file :  SingleCcode/20.c
		Error:  panic: format.Node internal error (112:35: missing ',' in argument list)
	Mistake file :  SingleCcode/27.c
		Error:  panic: (*ast.FunctionDecl) (0x5865c0,0xc420350000)
	Mistake file :  SingleCcode/31.c
		Error:  panic: (*ast.FunctionDecl) (0x5865c0,0xc4203ca700)
	Mistake file :  SingleCcode/array2.c
		Error:  panic: format.Node internal error (121:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  SingleCcode/bubble.c
		Error:  panic: format.Node internal error (122:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  SingleCcode/clean.c
		Error:  panic: format.Node internal error (115:9: expected '==', found '=' (and 1 more errors))
	Mistake file :  SingleCcode/counts.c
		Error:  panic: format.Node internal error (119:9: expected '==', found '=' (and 1 more errors))
	Mistake file :  SingleCcode/cpfile.c
		Error:  panic: format.Node internal error (234:10: expected '==', found '=' (and 2 more errors))
	Mistake file :  SingleCcode/cpintarray.c
		Error:  panic: format.Node internal error (104:6: expected ')', found '+=' (and 10 more errors))
	Mistake file :  SingleCcode/kilo.c
		Error:  panic: invalid identity: 'Not[]erow'
	Mistake file :  SingleCcode/linear.c
		Error:  panic: format.Node internal error (132:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  SingleCcode/makebinfile.c
		Error:  panic: format.Node internal error (126:11: expected '==', found '=' (and 1 more errors))
	Mistake file :  SingleCcode/merge.c
		Error:  panic: format.Node internal error (114:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  SingleCcode/number.c
		Error:  panic: format.Node internal error (142:6: expected ']', found '+=')
	Mistake file :  SingleCcode/selection.c
		Error:  panic: format.Node internal error (121:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  SingleCcode/shift.c
		Error:  panic: format.Node internal error (149:8: expected ']', found '+=' (and 1 more errors))
	Mistake file :  SingleCcode/sieve.c
		Error:  panic: format.Node internal error (121:18: expected ']', found '+=' (and 1 more errors))
	Mistake file :  SingleCcode/sortmerge.c
		Error:  panic: format.Node internal error (108:9: expected '==', found '=' (and 10 more errors))
	Mistake file :  SingleCcode/studentarray.c
		Error:  panic: format.Node internal error (108:9: expected '==', found '=' (and 2 more errors))
	Mistake file :  SingleCcode/studentlist.c
		Error:  panic: format.Node internal error (118:9: expected '==', found '=' (and 1 more errors))
	Mistake file :  SingleCcode/triangle.c
		Error:  panic: interface conversion: ast.Node is *ast.MemberExpr, not *ast.DeclRefExpr
	Mistake file :  sqlite/shell.c
		Error:  panic: unknown node type: 'IndirectFieldDecl 0x1d599c8 <line:167:25> col:25 implicit fpstate 'struct _fpstate *''
	Mistake file :  sqlite/sqlite3.c
		Error:  panic: interface conversion: interface {} is nil, not string
Result: 65 is Ok at 90 source c files - 72.222 procent. 
```
