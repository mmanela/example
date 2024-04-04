  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/implements`/
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
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/implements`/input.
//      documentation
//      > ```go
//      > const input untyped string = "package main\n\ntype A struct{}\nfunc (*A) f()\n\ntype B int\nfunc (...
//      > ```
//      documentation
//      > !+input
  
  type A struct{}
  func (*A) f()
  
  type B int
  func (B) f()
  func (*B) g()
  
  type I interface { f() }
  type J interface { g() }
  `
  
  //!-input
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/implements`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   // Parse one file.
   fset := token.NewFileSet()
// ^^^^ definition local 0
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/
//               ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/token`/NewFileSet().
   f, err := parser.ParseFile(fset, "input.go", input, 0)
// ^ definition local 1
//    ^^^ definition local 2
//           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/
//                  ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/parser`/ParseFile().
//                            ^^^^ reference local 0
//                                              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/implements`/input.
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
   pkg, err := conf.Check("hello", fset, []*ast.File{f}, nil)
// ^^^ definition local 4
//      ^^^ reference local 2
//             ^^^^ reference local 3
//                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Config#Check().
//                                 ^^^^ reference local 0
//                                          ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/
//                                              ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/ast`/File#
//                                                   ^ reference local 1
   if err != nil {
//    ^^^ reference local 2
    log.Fatal(err) // type error
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 2
   }
  
   //!+implements
   // Find all named types at package level.
   var allNamed []*types.Named
//     ^^^^^^^^ definition local 5
//                 ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Named#
   for _, name := range pkg.Scope().Names() {
//        ^^^^ definition local 6
//                      ^^^ reference local 4
//                          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Scope().
//                                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Scope#Names().
    if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
//     ^^^ definition local 7
//          ^^ definition local 8
//                ^^^ reference local 4
//                    ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Scope().
//                            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Scope#Lookup().
//                                   ^^^^ reference local 6
//                                           ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                                 ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeName#
//                                                            ^^ reference local 8
     allNamed = append(allNamed, obj.Type().(*types.Named))
//   ^^^^^^^^ reference local 5
//                     ^^^^^^^^ reference local 5
//                               ^^^ reference local 7
//                                   ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/object#Type().
//                                            ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Named#
    }
   }
  
   // Test assignability of all distinct pairs of
   // named types (T, U) where U is an interface.
   for _, T := range allNamed {
//        ^ definition local 9
//                   ^^^^^^^^ reference local 5
    for _, U := range allNamed {
//         ^ definition local 10
//                    ^^^^^^^^ reference local 5
     if T == U || !types.IsInterface(U) {
//      ^ reference local 9
//           ^ reference local 10
//                 ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                       ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/IsInterface().
//                                   ^ reference local 10
      continue
     }
     if types.AssignableTo(T, U) {
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//            ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/AssignableTo().
//                         ^ reference local 9
//                            ^ reference local 10
      fmt.Printf("%s satisfies %s\n", T, U)
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                    ^ reference local 9
//                                       ^ reference local 10
     } else if !types.IsInterface(T) &&
//              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                    ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/IsInterface().
//                                ^ reference local 9
      types.AssignableTo(types.NewPointer(T), U) {
//    ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//          ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/AssignableTo().
//                       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                             ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/NewPointer().
//                                        ^ reference local 9
//                                            ^ reference local 10
      fmt.Printf("%s satisfies %s\n", types.NewPointer(T), U)
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                    ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                          ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/NewPointer().
//                                                     ^ reference local 9
//                                                         ^ reference local 10
     }
    }
   }
   //!-implements
  }
  
  /*
  //!+output
  $ go build golang.org/x/example/gotypes/implements
  $ ./implements
  *hello.A satisfies hello.I
  hello.B satisfies hello.I
  *hello.B satisfies hello.J
  //!-output
  */
  
