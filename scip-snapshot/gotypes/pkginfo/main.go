  // !+
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/pkginfo`/
//        documentation
//        > !+
  
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
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/pkginfo`/hello.
//      documentation
//      > ```go
//      > const hello untyped string = "package main\n\nimport \"fmt\"\n\nfunc main() {\n        fmt.Println...
//      > ```
  
  import "fmt"
  
  func main() {
          fmt.Println("Hello, world")
  }`
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/pkginfo`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   fset := token.NewFileSet()
// ^^^^ definition local 0
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//               ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NewFileSet().
  
   // Parse the input string, []byte, or io.Reader,
   // recording position information in fset.
   // ParseFile returns an *ast.File, a syntax tree.
   f, err := parser.ParseFile(fset, "hello.go", hello, 0)
// ^ definition local 1
//    ^^^ definition local 2
//           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseFile().
//                            ^^^^ reference local 0
//                                              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/pkginfo`/hello.
   if err != nil {
//    ^^^ reference local 2
    log.Fatal(err) // parse error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 2
   }
  
   // A Config controls various options of the type checker.
   // The defaults work fine except for one setting:
   // we must specify how to deal with imports.
   conf := types.Config{Importer: importer.Default()}
// ^^^^ definition local 3
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#
//                      ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Importer.
//                                ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/
//                                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/importer`/Default().
  
   // Type-check the package containing only file f.
   // Check returns a *types.Package.
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
  
   fmt.Printf("Package  %q\n", pkg.Path())
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                             ^^^ reference local 4
//                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Path().
   fmt.Printf("Name:    %s\n", pkg.Name())
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                             ^^^ reference local 4
//                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Name().
   fmt.Printf("Imports: %s\n", pkg.Imports())
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                             ^^^ reference local 4
//                                 ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Imports().
   fmt.Printf("Scope:   %s\n", pkg.Scope())
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                             ^^^ reference local 4
//                                 ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Scope().
  }
  
  //!-
  
  /*
  //!+output
  $ go build golang.org/x/example/gotypes/pkginfo
  $ ./pkginfo
  Package  "cmd/hello"
  Name:    main
  Imports: [package fmt ("fmt")]
  Scope:   package "cmd/hello" scope 0x820533590 {
  .  func cmd/hello.main()
  }
  //!-output
  */
  
