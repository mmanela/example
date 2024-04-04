  // The hugeparam command identifies by-value parameters that are larger than n bytes.
  //
  // Example:
  //
  // $ ./hugeparams encoding/xml
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hugeparam`/
//        documentation
//        > The hugeparam command identifies by-value parameters that are larger than n bytes.
//        > 
//        > Example:
//        > 
//        > 	$ ./hugeparams encoding/xml
  
  import (
   "flag"
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "go/ast"
//  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
   "go/token"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
   "go/types"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
   "log"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
   "os"
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
  
   "golang.org/x/tools/go/packages"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
  )
  
  // !+
  var bytesFlag = flag.Int("bytes", 48, "maximum parameter size in bytes")
//    ^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hugeparam`/bytesFlag.
//    documentation
//    > ```go
//    > var bytesFlag *int
//    > ```
//    documentation
//    > !+
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/
//                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/Int().
  
  func PrintHugeParams(fset *token.FileSet, info *types.Info, sizes types.Sizes, files []*ast.File) {
//     ^^^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hugeparam`/PrintHugeParams().
//     documentation
//     > ```go
//     > func PrintHugeParams(fset *FileSet, info *Info, sizes Sizes, files []*File)
//     > ```
//                     ^^^^ definition local 0
//                           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//                                 ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#
//                                          ^^^^ definition local 1
//                                                ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#
//                                                            ^^^^^ definition local 2
//                                                                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                                                        ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Sizes#
//                                                                               ^^^^^ definition local 3
//                                                                                        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                                                                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#
   checkTuple := func(descr string, tuple *types.Tuple) {
// ^^^^^^^^^^ definition local 4
//                    ^^^^^ definition local 5
//                                  ^^^^^ definition local 6
//                                         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                               ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Tuple#
    for i := 0; i < tuple.Len(); i++ {
//      ^ definition local 7
//              ^ reference local 7
//                  ^^^^^ reference local 6
//                        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Tuple#Len().
//                               ^ reference local 7
     v := tuple.At(i)
//   ^ definition local 8
//        ^^^^^ reference local 6
//              ^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Tuple#At().
//                 ^ reference local 7
     if sz := sizes.Sizeof(v.Type()); sz > int64(*bytesFlag) {
//      ^^ definition local 9
//            ^^^^^ reference local 2
//                  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Sizes#Sizeof.
//                         ^ reference local 8
//                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/object#Type().
//                                    ^^ reference local 9
//                                                ^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hugeparam`/bytesFlag.
      fmt.Printf("%s: %q %s: %s = %d bytes\n",
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
       fset.Position(v.Pos()),
//     ^^^^ reference local 0
//          ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/FileSet#Position().
//                   ^ reference local 8
//                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/object#Pos().
       v.Name(), descr, v.Type(), sz)
//     ^ reference local 8
//       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/object#Name().
//               ^^^^^ reference local 5
//                      ^ reference local 8
//                        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/object#Type().
//                                ^^ reference local 9
     }
    }
   }
   checkSig := func(sig *types.Signature) {
// ^^^^^^^^ definition local 10
//                  ^^^ definition local 11
//                       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                             ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Signature#
    checkTuple("parameter", sig.Params())
//  ^^^^^^^^^^ reference local 4
//                          ^^^ reference local 11
//                              ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Signature#Params().
    checkTuple("result", sig.Results())
//  ^^^^^^^^^^ reference local 4
//                       ^^^ reference local 11
//                           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Signature#Results().
   }
   for _, file := range files {
//        ^^^^ definition local 12
//                      ^^^^^ reference local 3
    ast.Inspect(file, func(n ast.Node) bool {
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//      ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Inspect().
//              ^^^^ reference local 12
//                         ^ definition local 13
//                           ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/Node#
     switch n := n.(type) {
//          ^ definition local 14
//               ^ reference local 13
     case *ast.FuncDecl:
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//             ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/FuncDecl#
      checkSig(info.Defs[n.Name].Type().(*types.Signature))
//    ^^^^^^^^ reference local 10
//             ^^^^ reference local 1
//                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Defs.
//                       ^ reference local 14
//                       override_documentation
//                       > ```go
//                       > *go/ast.FuncDecl
//                       > ```
//                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/FuncDecl#Name.
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Type.
//                                        ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                              ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Signature#
     case *ast.FuncLit:
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/FuncLit#
      checkSig(info.Types[n.Type].Type.(*types.Signature))
//    ^^^^^^^^ reference local 10
//             ^^^^ reference local 1
//                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Info#Types.
//                        ^ reference local 14
//                        override_documentation
//                        > ```go
//                        > *go/ast.FuncLit
//                        > ```
//                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/FuncLit#Type.
//                                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeAndValue#Type.
//                                       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                             ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Signature#
     }
     return true
    })
   }
  }
  
  //!-
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hugeparam`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   flag.Parse()
// ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/Parse().
  
   // Load complete type information for the specified packages,
   // along with type-annotated syntax and the "sizeof" function.
   // Types for dependencies are loaded from export data.
   conf := &packages.Config{Mode: packages.LoadSyntax}
// ^^^^ definition local 15
//          ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                   ^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Config#
//                          ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Config#Mode.
//                                ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                                         ^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/LoadSyntax.
   pkgs, err := packages.Load(conf, flag.Args()...)
// ^^^^ definition local 16
//       ^^^ definition local 17
//              ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                       ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Load().
//                            ^^^^ reference local 15
//                                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/
//                                       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 flag/Args().
   if err != nil {
//    ^^^ reference local 17
    log.Fatal(err) // failed to load anything
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 17
   }
   if packages.PrintErrors(pkgs) > 0 {
//    ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//             ^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/PrintErrors().
//                         ^^^^ reference local 16
    os.Exit(1) // some packages contained errors
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Exit().
   }
  
   for _, pkg := range pkgs {
//        ^^^ definition local 18
//                     ^^^^ reference local 16
    PrintHugeParams(pkg.Fset, pkg.TypesInfo, pkg.TypesSizes, pkg.Syntax)
//  ^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hugeparam`/PrintHugeParams().
//                  ^^^ reference local 18
//                      ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Fset.
//                            ^^^ reference local 18
//                                ^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#TypesInfo.
//                                           ^^^ reference local 18
//                                               ^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#TypesSizes.
//                                                           ^^^ reference local 18
//                                                               ^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Syntax.
   }
  }
  
  /*
  //!+output
  % ./hugeparam encoding/xml
  /go/src/encoding/xml/marshal.go:167:50: "start" parameter: encoding/xml.StartElement = 56 bytes
  /go/src/encoding/xml/marshal.go:734:97: "" result: encoding/xml.StartElement = 56 bytes
  /go/src/encoding/xml/marshal.go:761:51: "start" parameter: encoding/xml.StartElement = 56 bytes
  /go/src/encoding/xml/marshal.go:781:68: "start" parameter: encoding/xml.StartElement = 56 bytes
  /go/src/encoding/xml/xml.go:72:30: "" result: encoding/xml.StartElement = 56 bytes
  //!-output
  */
  
