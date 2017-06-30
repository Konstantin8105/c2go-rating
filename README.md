# c2go-rating
Rating for https://github.com/elliotchance/c2go

C code base:
https://cis.temple.edu/~giorgio/cis71/code/
http://www.iu.hio.no/~mark/CTutorial/CTutorial.html

## Instruction
1.	go get -u github.com/Konstantin8105/c2go-rating
2.	copy folder from $GOPATH/src/github.com/Konstantin8105/c2go-rating/SingleCcode to $GOPATH/bin
3.	run `./c2go-rating`

## Table with results

| Date | Result |
|---|---|
| 22.06.2017 | 45 is Ok at 64 source c files |
| 24.06.2017 | 62 is Ok at 87 source c files |
| 29.06.2017 | Added "comma" problem for FOR. 63 is Ok at 87 soure c files | 
|   | File "factors1.c" is Ok |
|   | file "sieve.c" have 1 error insteand of 2 |


## Present list of mistake files

```
Amount mistake source by gcc:  0
Amount mistake c2go results:  24
	Mistake file :  ./SingleCcode/17-2.c
	Error:  panic: format.Node internal error (107:10: expected '==', found '=')
	Mistake file :  ./SingleCcode/19.c
	Error:  panic: format.Node internal error (138:37: expected ')', found '+=' (and 1 more errors))
	Mistake file :  ./SingleCcode/20.c
	Error:  panic: format.Node internal error (112:35: missing ',' in argument list)
	Mistake file :  ./SingleCcode/27.c
	Error:  panic: (*ast.FunctionDecl) (0x5835c0,0xc4202d6d80)
	Mistake file :  ./SingleCcode/31.c
	Error:  panic: (*ast.FunctionDecl) (0x5835c0,0xc4203a6300)
	Mistake file :  ./SingleCcode/array2.c
	Error:  panic: format.Node internal error (121:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/bubble.c
	Error:  panic: format.Node internal error (122:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/clean.c
	Error:  panic: format.Node internal error (115:9: expected '==', found '=' (and 1 more errors))
	Mistake file :  ./SingleCcode/counts.c
	Error:  panic: format.Node internal error (119:9: expected '==', found '=' (and 3 more errors))
	Mistake file :  ./SingleCcode/cpfile.c
	Error:  panic: format.Node internal error (234:10: expected '==', found '=' (and 2 more errors))
	Mistake file :  ./SingleCcode/cpintarray.c
	Error:  panic: format.Node internal error (104:6: expected ')', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/factorial.c
	Error:  panic: interface conversion: ast.Node is nil, not *ast.CompoundStmt
	Mistake file :  ./SingleCcode/fibo.c
	Error:  panic: format.Node internal error (115:19: expected '==', found '=')
	Mistake file :  ./SingleCcode/kilo.c
	Error:  panic: invalid identity: 'Not[]erow'
	Mistake file :  ./SingleCcode/linear.c
	Error:  panic: format.Node internal error (132:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/makebinfile.c
	Error:  panic: format.Node internal error (126:11: expected '==', found '=' (and 1 more errors))
	Mistake file :  ./SingleCcode/merge.c
	Error:  panic: format.Node internal error (114:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/number.c
	Error:  panic: format.Node internal error (142:6: expected ']', found '+=')
	Mistake file :  ./SingleCcode/selection.c
	Error:  panic: format.Node internal error (121:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/shift.c
	Error:  panic: format.Node internal error (149:8: expected ']', found '+=' (and 1 more errors))
	Mistake file :  ./SingleCcode/sieve.c
	Error:  panic: format.Node internal error (121:18: expected ']', found '+=' (and 1 more errors))
	Mistake file :  ./SingleCcode/sortmerge.c
	Error:  panic: format.Node internal error (108:9: expected '==', found '=' (and 10 more errors))
	Mistake file :  ./SingleCcode/studentarray.c
	Error:  panic: format.Node internal error (108:9: expected '==', found '=' (and 2 more errors))
	Mistake file :  ./SingleCcode/studentlist.c
	Error:  panic: format.Node internal error (118:9: expected '==', found '=' (and 1 more errors))
Result: 63 is Ok at 87 source c files - 72.414 procent. 
```
