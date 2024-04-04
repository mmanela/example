  // The doc command prints the doc comment of a package-level object.
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/doc`/
//        documentation
//        > The doc command prints the doc comment of a package-level object.
  
  import (
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "go/ast"
//  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
   "log"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
   "os"
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
  
   "golang.org/x/tools/go/ast/astutil"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/ast/astutil`/
   "golang.org/x/tools/go/packages"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
   "golang.org/x/tools/go/types/typeutil"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/types/typeutil`/
  )
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/doc`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   if len(os.Args) != 3 {
//        ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
    log.Fatal("Usage: doc <package> <object>")
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
   }
   //!+part1
   pkgpath, name := os.Args[1], os.Args[2]
// ^^^^^^^ definition local 0
//          ^^^^ definition local 1
//                  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
//                              ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
  
   // Load complete type information for the specified packages,
   // along with type-annotated syntax.
   // Types for dependencies are loaded from export data.
   conf := &packages.Config{Mode: packages.LoadSyntax}
// ^^^^ definition local 2
//          ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                   ^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Config#
//                          ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Config#Mode.
//                                ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                                         ^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/LoadSyntax.
   pkgs, err := packages.Load(conf, pkgpath)
// ^^^^ definition local 3
//       ^^^ definition local 4
//              ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                       ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Load().
//                            ^^^^ reference local 2
//                                  ^^^^^^^ reference local 0
   if err != nil {
//    ^^^ reference local 4
    log.Fatal(err) // failed to load anything
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 4
   }
   if packages.PrintErrors(pkgs) > 0 {
//    ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//             ^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/PrintErrors().
//                         ^^^^ reference local 3
    os.Exit(1) // some packages contained errors
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Exit().
   }
  
   // Find the package and package-level object.
   pkg := pkgs[0]
// ^^^ definition local 5
//        ^^^^ reference local 3
   obj := pkg.Types.Scope().Lookup(name)
// ^^^ definition local 6
//        ^^^ reference local 5
//            ^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Types.
//                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Scope().
//                          ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Scope#Lookup().
//                                 ^^^^ reference local 1
   if obj == nil {
//    ^^^ reference local 6
    log.Fatalf("%s.%s not found", pkg.Types.Path(), name)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatalf().
//                                ^^^ reference local 5
//                                    ^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Types.
//                                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Path().
//                                                  ^^^^ reference local 1
   }
   //!-part1
   //!+part2
  
   // Print the object and its methods (incl. location of definition).
   fmt.Println(obj)
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//             ^^^ reference local 6
   for _, sel := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
//        ^^^ definition local 7
//                     ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/types/typeutil`/
//                              ^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/types/typeutil`/IntuitiveMethodSet().
//                                                 ^^^ reference local 6
//                                                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Type.
    fmt.Printf("%s: %s\n", pkg.Fset.Position(sel.Obj().Pos()), sel)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                         ^^^ reference local 5
//                             ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Fset.
//                                  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#Position().
//                                           ^^^ reference local 7
//                                               ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Selection#Obj().
//                                                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Pos.
//                                                             ^^^ reference local 7
   }
  
   // Find the path from the root of the AST to the object's position.
   // Walk up to the enclosing ast.Decl for the doc comment.
   for _, file := range pkg.Syntax {
//        ^^^^ definition local 8
//                      ^^^ reference local 5
//                          ^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Syntax.
    pos := obj.Pos()
//  ^^^ definition local 9
//         ^^^ reference local 6
//             ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Pos.
    if !(file.FileStart <= pos && pos < file.FileEnd) {
//       ^^^^ reference local 8
//            ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#FileStart.
//                         ^^^ reference local 9
//                                ^^^ reference local 9
//                                      ^^^^ reference local 8
//                                           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#FileEnd.
     continue // not in this file
    }
    path, _ := astutil.PathEnclosingInterval(file, pos, pos)
//  ^^^^ definition local 10
//             ^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/ast/astutil`/
//                     ^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/ast/astutil`/PathEnclosingInterval().
//                                           ^^^^ reference local 8
//                                                 ^^^ reference local 9
//                                                      ^^^ reference local 9
    for _, n := range path {
//         ^ definition local 11
//                    ^^^^ reference local 10
     switch n := n.(type) {
//          ^ definition local 12
//               ^ reference local 11
     case *ast.GenDecl:
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/GenDecl#
      fmt.Println("\n", n.Doc.Text())
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//                      ^ reference local 12
//                      override_documentation
//                      > ```go
//                      > *go/ast.GenDecl
//                      > ```
//                        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/GenDecl#Doc.
//                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/CommentGroup#Text().
      return
     case *ast.FuncDecl:
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//             ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/FuncDecl#
      fmt.Println("\n", n.Doc.Text())
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//                      ^ reference local 12
//                      override_documentation
//                      > ```go
//                      > *go/ast.FuncDecl
//                      > ```
//                        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/FuncDecl#Doc.
//                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/CommentGroup#Text().
      return
     }
    }
   }
   //!-part2
  }
  
  // (The $GOROOT below is the actual string that appears in file names
  // loaded from export data for packages in the standard library.)
  
  /*
  //!+output
  $ ./doc net/http File
  type net/http.File interface{Readdir(count int) ([]os.FileInfo, error); Seek(offset int64, whence int) (int64, error); Stat() (os.FileInfo, error); io.Closer; io.Reader}
  $GOROOT/src/io/io.go:92:2: method (net/http.File) Close() error
  $GOROOT/src/io/io.go:71:2: method (net/http.File) Read(p []byte) (n int, err error)
  /go/src/net/http/fs.go:65:2: method (net/http.File) Readdir(count int) ([]os.FileInfo, error)
  $GOROOT/src/net/http/fs.go:66:2: method (net/http.File) Seek(offset int64, whence int) (int64, error)
  /go/src/net/http/fs.go:67:2: method (net/http.File) Stat() (os.FileInfo, error)
  
   A File is returned by a FileSystem's Open method and can be
  served by the FileServer implementation.
  
  The methods should behave the same as those on an *os.File.
  //!-output
  */
  
