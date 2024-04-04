  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/
//        documentation
//        > package main
  
  import (
   "bytes"
//  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "go/ast"
//  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
   "go/format"
//  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/format`/
   "go/importer"
//  ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/
   "go/parser"
//  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
   "go/token"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
   "go/types"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
   "log"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
  )
  
  // !+input
  const input = `
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/input.
//      documentation
//      > ```go
//      > const input untyped string = "\npackage main\n\nvar m = make(map[string]int)\n\nfunc main() {\n\tv...
//      > ```
//      documentation
//      > !+input
  package main
  
  var m = make(map[string]int)
  
  func main() {
   v, ok := m["hello, " + "world"]
   print(rune(v), ok)
  }
  `
  
  //!-input
  
  var fset = token.NewFileSet()
//    ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/fset.
//    documentation
//    > ```go
//    > var fset *FileSet
//    > ```
//           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                 ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NewFileSet().
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   f, err := parser.ParseFile(fset, "hello.go", input, 0)
// ^ definition local 0
//    ^^^ definition local 1
//           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseFile().
//                            ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/fset.
//                                              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/input.
   if err != nil {
//    ^^^ reference local 1
    log.Fatal(err) // parse error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 1
   }
  
   conf := types.Config{Importer: importer.Default()}
// ^^^^ definition local 2
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#
//                      ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Importer.
//                                ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/
//                                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/Default().
   info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
// ^^^^ definition local 3
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#
//                     ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Types.
//                                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Expr#
//                                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                                    ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#
   if _, err := conf.Check("cmd/hello", fset, []*ast.File{f}, info); err != nil {
//       ^^^ definition local 4
//              ^^^^ reference local 2
//                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Check().
//                                      ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/fset.
//                                               ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                                   ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#
//                                                        ^ reference local 0
//                                                            ^^^^ reference local 3
//                                                                   ^^^ reference local 4
    log.Fatal(err) // type error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 4
   }
  
   //!+inspect
   // f is a parsed, type-checked *ast.File.
   ast.Inspect(f, func(n ast.Node) bool {
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//     ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Inspect().
//             ^ reference local 0
//                     ^ definition local 5
//                       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Node#
    if expr, ok := n.(ast.Expr); ok {
//     ^^^^ definition local 6
//           ^^ definition local 7
//                 ^ reference local 5
//                    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Expr#
//                               ^^ reference local 7
     if tv, ok := info.Types[expr]; ok {
//      ^^ definition local 8
//          ^^ definition local 9
//                ^^^^ reference local 3
//                     ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Types.
//                           ^^^^ reference local 6
//                                  ^^ reference local 9
      fmt.Printf("%-24s\tmode:  %s\n", nodeString(expr), mode(tv))
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                     ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/nodeString().
//                                                ^^^^ reference local 6
//                                                       ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/mode().
//                                                            ^^ reference local 8
      fmt.Printf("\t\t\t\ttype:  %v\n", tv.Type)
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                      ^^ reference local 8
//                                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#Type.
      if tv.Value != nil {
//       ^^ reference local 8
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#Value.
       fmt.Printf("\t\t\t\tvalue: %v\n", tv.Value)
//     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//         ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                       ^^ reference local 8
//                                          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#Value.
      }
     }
    }
    return true
   })
   //!-inspect
  }
  
  // nodeString formats a syntax tree in the style of gofmt.
  func nodeString(n ast.Node) string {
//     ^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/nodeString().
//     documentation
//     > ```go
//     > func nodeString(n Node) string
//     > ```
//     documentation
//     > nodeString formats a syntax tree in the style of gofmt.
//                ^ definition local 10
//                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Node#
   var buf bytes.Buffer
//     ^^^ definition local 11
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/
//               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#
   format.Node(&buf, fset, n)
// ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/format`/
//        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/format`/Node().
//              ^^^ reference local 11
//                   ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/fset.
//                         ^ reference local 10
   return buf.String()
//        ^^^ reference local 11
//            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#String().
  }
  
  // mode returns a string describing the mode of an expression.
  func mode(tv types.TypeAndValue) string {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/typeandvalue`/mode().
//     documentation
//     > ```go
//     > func mode(tv TypeAndValue) string
//     > ```
//     documentation
//     > mode returns a string describing the mode of an expression.
//          ^^ definition local 12
//             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                   ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#
   s := ""
// ^ definition local 13
   if tv.IsVoid() {
//    ^^ reference local 12
//       ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsVoid().
    s += ",void"
//  ^ reference local 13
   }
   if tv.IsType() {
//    ^^ reference local 12
//       ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsType().
    s += ",type"
//  ^ reference local 13
   }
   if tv.IsBuiltin() {
//    ^^ reference local 12
//       ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsBuiltin().
    s += ",builtin"
//  ^ reference local 13
   }
   if tv.IsValue() {
//    ^^ reference local 12
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsValue().
    s += ",value"
//  ^ reference local 13
   }
   if tv.IsNil() {
//    ^^ reference local 12
//       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsNil().
    s += ",nil"
//  ^ reference local 13
   }
   if tv.Addressable() {
//    ^^ reference local 12
//       ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#Addressable().
    s += ",addressable"
//  ^ reference local 13
   }
   if tv.Assignable() {
//    ^^ reference local 12
//       ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#Assignable().
    s += ",assignable"
//  ^ reference local 13
   }
   if tv.HasOk() {
//    ^^ reference local 12
//       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#HasOk().
    s += ",ok"
//  ^ reference local 13
   }
   return s[1:]
//        ^ reference local 13
  }
  
  /*
  //!+output
  $ go build golang.org/x/example/gotypes/typeandvalue
  $ ./typeandvalue
  make(map[string]int)            mode:  value
                                  type:  map[string]int
  make                            mode:  builtin
                                  type:  func(map[string]int) map[string]int
  map[string]int                  mode:  type
                                  type:  map[string]int
  string                          mode:  type
                                  type:  string
  int                             mode:  type
                                  type:  int
  m["hello, "+"world"]            mode:  value,assignable,ok
                                  type:  (int, bool)
  m                               mode:  value,addressable,assignable
                                  type:  map[string]int
  "hello, " + "world"             mode:  value
                                  type:  string
                                  value: "hello, world"
  "hello, "                       mode:  value
                                  type:  untyped string
                                  value: "hello, "
  "world"                         mode:  value
                                  type:  untyped string
                                  value: "world"
  print(rune(v), ok)              mode:  void
                                  type:  ()
  print                           mode:  builtin
                                  type:  func(rune, bool)
  rune(v)                         mode:  value
                                  type:  rune
  rune                            mode:  type
                                  type:  rune
  ...more not shown...
  //!-output
  */
  
