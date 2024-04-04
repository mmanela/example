  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/
//        documentation
//        > package main
  
  import (
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "go/ast"
//  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
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
  const input = `package main
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/input.
//      documentation
//      > ```go
//      > const input untyped string = "package main\n\nimport \"bytes\"\n\nfunc main() {\n\tvar buf bytes.B...
//      > ```
//      documentation
//      > !+input
  
  import "bytes"
  
  func main() {
   var buf bytes.Buffer
   if buf.Bytes == nil && bytes.Repeat != nil && main == nil {
    // ...
   }
  }
  `
  
  //!-input
  
  var fset = token.NewFileSet()
//    ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/fset.
//    documentation
//    > ```go
//    > var fset *FileSet
//    > ```
//           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                 ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NewFileSet().
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   f, err := parser.ParseFile(fset, "input.go", input, 0)
// ^ definition local 0
//    ^^^ definition local 1
//           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseFile().
//                            ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/fset.
//                                              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/input.
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
   info := &types.Info{
// ^^^^ definition local 3
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#
    Defs:  make(map[*ast.Ident]types.Object),
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Defs.
//                   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#
//                             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                   ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#
    Uses:  make(map[*ast.Ident]types.Object),
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Uses.
//                   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#
//                             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                   ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#
    Types: make(map[ast.Expr]types.TypeAndValue),
//  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Types.
//                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Expr#
//                           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                 ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#
   }
   if _, err = conf.Check("cmd/hello", fset, []*ast.File{f}, info); err != nil {
//       ^^^ reference local 1
//             ^^^^ reference local 2
//                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Check().
//                                     ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/fset.
//                                              ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#
//                                                       ^ reference local 0
//                                                           ^^^^ reference local 3
//                                                                  ^^^ reference local 1
    log.Fatal(err) // type error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 1
   }
  
   ast.Inspect(f, func(n ast.Node) bool {
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//     ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Inspect().
//             ^ reference local 0
//                     ^ definition local 4
//                       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Node#
    if n != nil {
//     ^ reference local 4
     CheckNilFuncComparison(info, n)
//   ^^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/CheckNilFuncComparison().
//                          ^^^^ reference local 3
//                                ^ reference local 4
    }
    return true
   })
  }
  
  // !+
  // CheckNilFuncComparison reports unintended comparisons
  // of functions against nil, e.g., "if x.Method == nil {".
  func CheckNilFuncComparison(info *types.Info, n ast.Node) {
//     ^^^^^^^^^^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/CheckNilFuncComparison().
//     documentation
//     > ```go
//     > func CheckNilFuncComparison(info *Info, n Node)
//     > ```
//     documentation
//     > !+
//     > CheckNilFuncComparison reports unintended comparisons
//     > of functions against nil, e.g., "if x.Method == nil {".
//                            ^^^^ definition local 5
//                                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#
//                                              ^ definition local 6
//                                                ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Node#
   e, ok := n.(*ast.BinaryExpr)
// ^ definition local 7
//    ^^ definition local 8
//          ^ reference local 6
//              ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                  ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#
   if !ok {
//     ^^ reference local 8
    return // not a binary operation
   }
   if e.Op != token.EQL && e.Op != token.NEQ {
//    ^ reference local 7
//      ^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Op.
//            ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/EQL.
//                         ^ reference local 7
//                           ^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Op.
//                                 ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                                       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NEQ.
    return // not a comparison
   }
  
   // If this is a comparison against nil, find the other operand.
   var other ast.Expr
//     ^^^^^ definition local 9
//           ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Expr#
   if info.Types[e.X].IsNil() {
//    ^^^^ reference local 5
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Types.
//               ^ reference local 7
//                 ^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#X.
//                    ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsNil().
    other = e.Y
//  ^^^^^ reference local 9
//          ^ reference local 7
//            ^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Y.
   } else if info.Types[e.Y].IsNil() {
//           ^^^^ reference local 5
//                ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Types.
//                      ^ reference local 7
//                        ^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Y.
//                           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#IsNil().
    other = e.X
//  ^^^^^ reference local 9
//          ^ reference local 7
//            ^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#X.
   } else {
    return // not a comparison against nil
   }
  
   // Find the object.
   var obj types.Object
//     ^^^ definition local 10
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#
   switch v := other.(type) {
//        ^ definition local 11
//             ^^^^^ reference local 9
   case *ast.Ident:
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#
    obj = info.Uses[v]
//  ^^^ reference local 10
//        ^^^^ reference local 5
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Uses.
//                  ^ reference local 11
//                  override_documentation
//                  > ```go
//                  > *go/ast.Ident
//                  > ```
   case *ast.SelectorExpr:
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//           ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/SelectorExpr#
    obj = info.Uses[v.Sel]
//  ^^^ reference local 10
//        ^^^^ reference local 5
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Uses.
//                  ^ reference local 11
//                  override_documentation
//                  > ```go
//                  > *go/ast.SelectorExpr
//                  > ```
//                    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/SelectorExpr#Sel.
   default:
    return // not an identifier or selection
   }
  
   if _, ok := obj.(*types.Func); !ok {
//       ^^ definition local 12
//             ^^^ reference local 10
//                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Func#
//                                 ^^ reference local 12
    return // not a function or method
   }
  
   fmt.Printf("%s: comparison of function %v %v nil is always %v\n",
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
    fset.Position(e.Pos()), obj.Name(), e.Op, e.Op == token.NEQ)
//  ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/nilfunc`/fset.
//       ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#Position().
//                ^ reference local 7
//                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Pos().
//                          ^^^ reference local 10
//                              ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Name.
//                                      ^ reference local 7
//                                        ^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Op.
//                                            ^ reference local 7
//                                              ^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/BinaryExpr#Op.
//                                                    ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                                                          ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NEQ.
  }
  
  //!-
  
  /*
  //!+output
  $ go build golang.org/x/example/gotypes/nilfunc
  $ ./nilfunc
  input.go:7:5: comparison of function Bytes == nil is always false
  input.go:7:25: comparison of function Repeat != nil is always true
  input.go:7:48: comparison of function main == nil is always false
  //!-output
  */
  
