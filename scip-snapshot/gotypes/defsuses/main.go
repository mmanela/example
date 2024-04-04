  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/defsuses`/
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
  
  const hello = `package main
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/defsuses`/hello.
//      documentation
//      > ```go
//      > const hello untyped string = "package main\n\nimport \"fmt\"\n\nfunc main() {\n        fmt.Println...
//      > ```
  
  import "fmt"
  
  func main() {
          fmt.Println("Hello, world")
  }
  `
  
  // !+
  func PrintDefsUses(fset *token.FileSet, files ...*ast.File) error {
//     ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/defsuses`/PrintDefsUses().
//     documentation
//     > ```go
//     > func PrintDefsUses(fset *FileSet, files ...*File) error
//     > ```
//     documentation
//     > !+
//                   ^^^^ definition local 0
//                         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                               ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#
//                                        ^^^^^ definition local 1
//                                                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#
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
    Defs: make(map[*ast.Ident]types.Object),
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Defs.
//                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#
//                            ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#
    Uses: make(map[*ast.Ident]types.Object),
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Uses.
//                  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#
//                            ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#
   }
   _, err := conf.Check("hello", fset, files, info)
//    ^^^ definition local 4
//           ^^^^ reference local 2
//                ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Check().
//                               ^^^^ reference local 0
//                                     ^^^^^ reference local 1
//                                            ^^^^ reference local 3
   if err != nil {
//    ^^^ reference local 4
    return err // type error
//         ^^^ reference local 4
   }
  
   for id, obj := range info.Defs {
//     ^^ definition local 5
//         ^^^ definition local 6
//                      ^^^^ reference local 3
//                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Defs.
    fmt.Printf("%s: %q defines %v\n",
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
     fset.Position(id.Pos()), id.Name, obj)
//   ^^^^ reference local 0
//        ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#Position().
//                 ^^ reference local 5
//                    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#Pos().
//                            ^^ reference local 5
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#Name.
//                                     ^^^ reference local 6
   }
   for id, obj := range info.Uses {
//     ^^ definition local 7
//         ^^^ definition local 8
//                      ^^^^ reference local 3
//                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Uses.
    fmt.Printf("%s: %q uses %v\n",
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
     fset.Position(id.Pos()), id.Name, obj)
//   ^^^^ reference local 0
//        ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#Position().
//                 ^^ reference local 7
//                    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#Pos().
//                            ^^ reference local 7
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Ident#Name.
//                                     ^^^ reference local 8
   }
   return nil
  }
  
  //!-
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/defsuses`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   // Parse one file.
   fset := token.NewFileSet()
// ^^^^ definition local 9
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//               ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NewFileSet().
   f, err := parser.ParseFile(fset, "hello.go", hello, 0)
// ^ definition local 10
//    ^^^ definition local 11
//           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseFile().
//                            ^^^^ reference local 9
//                                              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/defsuses`/hello.
   if err != nil {
//    ^^^ reference local 11
    log.Fatal(err) // parse error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 11
   }
   if err := PrintDefsUses(fset, f); err != nil {
//    ^^^ definition local 12
//           ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/defsuses`/PrintDefsUses().
//                         ^^^^ reference local 9
//                               ^ reference local 10
//                                   ^^^ reference local 12
    log.Fatal(err) // type error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 12
   }
  }
  
  /*
  //!+output
  $ go build golang.org/x/example/gotypes/defsuses
  $ ./defsuses
  hello.go:1:9: "main" defines <nil>
  hello.go:5:6: "main" defines func hello.main()
  hello.go:6:9: "fmt" uses package fmt
  hello.go:6:13: "Println" uses func fmt.Println(a ...interface{}) (n int, err error)
  //!-output
  */
  
