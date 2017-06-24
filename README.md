# c2go-rating
Rating for https://github.com/elliotchance/c2go

C code base:
https://cis.temple.edu/~giorgio/cis71/code/
http://www.iu.hio.no/~mark/CTutorial/CTutorial.html

Table with results

| Date | Result |
|---|---|
| 22.06.2017 | 45 is Ok at 64 source c files |
| 24.06.2017 | 62 is Ok at 87 source c files |


Present list of mistake files:
```
Amount mistake source by gcc:  1
	Mistake file :  ./SingleCcode/kilo.c
Amount mistake c2go results:  25
	Mistake file :  ./SingleCcode/17-2.c
	Error:  panic: format.Node internal error (106:10: expected '==', found '=')
	Mistake file :  ./SingleCcode/19.c
	Error:  panic: format.Node internal error (137:37: expected ')', found '+=' (and 1 more errors))
	Mistake file :  ./SingleCcode/20.c
	Error:  panic: format.Node internal error (111:35: missing ',' in argument list)
	Mistake file :  ./SingleCcode/27.c
	Error:  panic: (*ast.FunctionDecl) (0x582680,0xc42043f700)
	Mistake file :  ./SingleCcode/31.c
	Error:  panic: (*ast.FunctionDecl) (0x582680,0xc4200ca880)
	Mistake file :  ./SingleCcode/array2.c
	Error:  panic: format.Node internal error (120:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/bubble.c
	Error:  panic: format.Node internal error (121:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/clean.c
	Error:  panic: format.Node internal error (114:9: expected '==', found '=' (and 1 more errors))
	Mistake file :  ./SingleCcode/counts.c
	Error:  panic: format.Node internal error (118:9: expected '==', found '=' (and 3 more errors))
	Mistake file :  ./SingleCcode/cpfile.c
	Error:  panic: format.Node internal error (233:10: expected '==', found '=' (and 2 more errors))
	Mistake file :  ./SingleCcode/cpintarray.c
	Error:  panic: format.Node internal error (103:6: expected ')', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/factorial.c
	Error:  panic: format.Node internal error (115:19: expected '==', found '=' (and 3 more errors))
	Mistake file :  ./SingleCcode/factors1.c
	Error:  panic: format.Node internal error (108:22: expected '==', found '=')
	Mistake file :  ./SingleCcode/fibo.c
	Error:  panic: format.Node internal error (114:19: expected '==', found '=')
	Mistake file :  ./SingleCcode/kilo.c
	Error:  panic: invalid identity: 'Not[]erow'
	Mistake file :  ./SingleCcode/linear.c
	Error:  panic: format.Node internal error (131:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/makebinfile.c
	Error:  panic: format.Node internal error (125:11: expected '==', found '=' (and 1 more errors))
	Mistake file :  ./SingleCcode/merge.c
	Error:  panic: format.Node internal error (113:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/number.c
	Error:  panic: format.Node internal error (141:6: expected ']', found '+=')
	Mistake file :  ./SingleCcode/selection.c
	Error:  panic: format.Node internal error (120:42: expected ']', found '+=' (and 10 more errors))
	Mistake file :  ./SingleCcode/shift.c
	Error:  panic: format.Node internal error (148:8: expected ']', found '+=' (and 1 more errors))
	Mistake file :  ./SingleCcode/sieve.c
	Error:  panic: format.Node internal error (113:21: expected '==', found '=' (and 2 more errors))
	Mistake file :  ./SingleCcode/sortmerge.c
	Error:  panic: format.Node internal error (107:9: expected '==', found '=' (and 10 more errors))
	Mistake file :  ./SingleCcode/studentarray.c
	Error:  panic: format.Node internal error (107:9: expected '==', found '=' (and 2 more errors))
	Mistake file :  ./SingleCcode/studentlist.c
	Error:  panic: format.Node internal error (117:9: expected '==', found '=' (and 1 more errors))
Result: 62 is Ok at 87 source c files - 71.264 procent. 
```
