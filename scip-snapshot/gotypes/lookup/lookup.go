  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/lookup`/
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
   "strings"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
  )
  
  // !+input
  const hello = `
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/lookup`/hello.
//      documentation
//      > ```go
//      > const hello untyped string = "\npackage main\n\nimport \"fmt\"\n\n// append\nfunc main() {\n      ...
//      > ```
//      documentation
//      > !+input
  package main
  
  import "fmt"
  
  // append
  func main() {
          // fmt
          fmt.Println("Hello, world")
          // main
          main, x := 1, 2
          // main
          print(main, x)
          // x
  }
  // x
  `
  
  //!-input
  
  // !+main
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/lookup`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
//     documentation
//     > !+main
   fset := token.NewFileSet()
// ^^^^ definition local 0
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//               ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NewFileSet().
   f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
// ^ definition local 1
//    ^^^ definition local 2
//           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseFile().
//                            ^^^^ reference local 0
//                                              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/lookup`/hello.
//                                                     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                                                            ^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseComments.
   if err != nil {
//    ^^^ reference local 2
    log.Fatal(err) // parse error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 2
   }
  
   conf := types.Config{Importer: importer.Default()}
// ^^^^ definition local 3
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#
//                      ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Importer.
//                                ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/
//                                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/Default().
   pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
// ^^^ definition local 4
//      ^^^ reference local 2
//             ^^^^ reference local 3
//                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Check().
//                                     ^^^^ reference local 0
//                                              ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#
//                                                       ^ reference local 1
   if err != nil {
//    ^^^ reference local 2
    log.Fatal(err) // type error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 2
   }
  
   // Each comment contains a name.
   // Look up that name in the innermost scope enclosing the comment.
   for _, comment := range f.Comments {
//        ^^^^^^^ definition local 5
//                         ^ reference local 1
//                           ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#Comments.
    pos := comment.Pos()
//  ^^^ definition local 6
//         ^^^^^^^ reference local 5
//                 ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/CommentGroup#Pos().
    name := strings.TrimSpace(comment.Text())
//  ^^^^ definition local 7
//          ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/TrimSpace().
//                            ^^^^^^^ reference local 5
//                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/CommentGroup#Text().
    fmt.Printf("At %s,\t%q = ", fset.Position(pos), name)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                              ^^^^ reference local 0
//                                   ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#Position().
//                                            ^^^ reference local 6
//                                                  ^^^^ reference local 7
    inner := pkg.Scope().Innermost(pos)
//  ^^^^^ definition local 8
//           ^^^ reference local 4
//               ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Scope().
//                       ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Scope#Innermost().
//                                 ^^^ reference local 6
    if _, obj := inner.LookupParent(name, pos); obj != nil {
//        ^^^ definition local 9
//               ^^^^^ reference local 8
//                     ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Scope#LookupParent().
//                                  ^^^^ reference local 7
//                                        ^^^ reference local 6
//                                              ^^^ reference local 9
     fmt.Println(obj)
//   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//               ^^^ reference local 9
    } else {
     fmt.Println("not found")
//   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
    }
   }
  }
  
  //!-main
  
  /*
  //!+output
  $ go build golang.org/x/example/gotypes/lookup
  $ ./lookup
  At hello.go:6:1,        "append" = builtin append
  At hello.go:8:9,        "fmt" = package fmt
  At hello.go:10:9,       "main" = func cmd/hello.main()
  At hello.go:12:9,       "main" = var main int
  At hello.go:14:9,       "x" = var x int
  At hello.go:16:1,       "x" = not found
  //!-output
  */
  
