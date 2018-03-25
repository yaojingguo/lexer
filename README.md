A simple calculator syntax given on page 54  of `Programming Language Pragmatics, Fourth Edition`

```
assign -> :=
plus -> +
minus -> -
times -> *
div -> /
lparen -> (
rparen -> )
id -> letter ( letter | digit )* 
      except for read and write
number -> digit digit * | digit * ( . digit | digit . ) digit *
comment -> /* ( non-* | * non-/ )* *+ /
           | // ( non-newline )* newline
```
