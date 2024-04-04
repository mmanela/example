  // The skeleton command prints the boilerplate for a concrete type
  // that implements the specified interface type.
  //
  // Example:
  //
  // $ ./skeleton io ReadWriteCloser buffer
  // // *buffer implements io.ReadWriteCloser.
  // type buffer struct{ /* ... */ }
  // func (b *buffer) Close() error { panic("unimplemented") }
  // func (b *buffer) Read(p []byte) (n int, err error) { panic("unimplemented") }
  // func (b *buffer) Write(p []byte) (n int, err error) { panic("unimplemented") }
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/
//        documentation
//        > The skeleton command prints the boilerplate for a concrete type
//        > that implements the specified interface type.
//        > 
//        > Example:
//        > 
//        > 	$ ./skeleton io ReadWriteCloser buffer
//        > 	// *buffer implements io.ReadWriteCloser.
//        > 	type buffer struct{ /* ... */ }
//        > 	func (b *buffer) Close() error { panic("unimplemented") }
//        > 	func (b *buffer) Read(p []byte) (n int, err error) { panic("unimplemented") }
//        > 	func (b *buffer) Write(p []byte) (n int, err error) { panic("unimplemented") }
  
  import (
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "go/types"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
   "log"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
   "os"
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
   "strings"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
   "unicode"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 unicode/
   "unicode/utf8"
//  ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `unicode/utf8`/
  
   "golang.org/x/tools/go/packages"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
  )
  
  const usage = "Usage: skeleton <package> <interface> <concrete>"
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/usage.
//      documentation
//      > ```go
//      > const usage untyped string = "Usage: skeleton <package> <interface> <concrete>"
//      > ```
  
  // !+
  func PrintSkeleton(pkg *types.Package, ifacename, concname string) error {
//     ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/PrintSkeleton().
//     documentation
//     > ```go
//     > func PrintSkeleton(pkg *Package, ifacename string, concname string) error
//     > ```
//     documentation
//     > !+
//                   ^^^ definition local 0
//                        ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                              ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#
//                                       ^^^^^^^^^ definition local 1
//                                                  ^^^^^^^^ definition local 2
   obj := pkg.Scope().Lookup(ifacename)
// ^^^ definition local 3
//        ^^^ reference local 0
//            ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Scope().
//                    ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Scope#Lookup().
//                           ^^^^^^^^^ reference local 1
   if obj == nil {
//    ^^^ reference local 3
    return fmt.Errorf("%s.%s not found", pkg.Path(), ifacename)
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Errorf().
//                                       ^^^ reference local 0
//                                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Path().
//                                                   ^^^^^^^^^ reference local 1
   }
   if _, ok := obj.(*types.TypeName); !ok {
//       ^^ definition local 4
//             ^^^ reference local 3
//                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                         ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeName#
//                                     ^^ reference local 4
    return fmt.Errorf("%v is not a named type", obj)
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Errorf().
//                                              ^^^ reference local 3
   }
   iface, ok := obj.Type().Underlying().(*types.Interface)
// ^^^^^ definition local 5
//        ^^ definition local 6
//              ^^^ reference local 3
//                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Type.
//                         ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Type#Underlying.
//                                        ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                              ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Interface#
   if !ok {
//     ^^ reference local 6
    return fmt.Errorf("type %v is a %T, not an interface",
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Errorf().
     obj, obj.Type().Underlying())
//   ^^^ reference local 3
//        ^^^ reference local 3
//            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Type.
//                   ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Type#Underlying.
   }
   // Use first letter of type name as receiver parameter.
   if !isValidIdentifier(concname) {
//     ^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/isValidIdentifier().
//                       ^^^^^^^^ reference local 2
    return fmt.Errorf("invalid concrete type name: %q", concname)
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Errorf().
//                                                      ^^^^^^^^ reference local 2
   }
   r, _ := utf8.DecodeRuneInString(concname)
// ^ definition local 7
//         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `unicode/utf8`/
//              ^^^^^^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `unicode/utf8`/DecodeRuneInString().
//                                 ^^^^^^^^ reference local 2
  
   fmt.Printf("// *%s implements %s.%s.\n", concname, pkg.Path(), ifacename)
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                          ^^^^^^^^ reference local 2
//                                                    ^^^ reference local 0
//                                                        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Path().
//                                                                ^^^^^^^^^ reference local 1
   fmt.Printf("type %s struct{}\n", concname)
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
//                                  ^^^^^^^^ reference local 2
   mset := types.NewMethodSet(iface)
// ^^^^ definition local 8
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//               ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/NewMethodSet().
//                            ^^^^^ reference local 5
   for i := 0; i < mset.Len(); i++ {
//     ^ definition local 9
//             ^ reference local 9
//                 ^^^^ reference local 8
//                      ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/MethodSet#Len().
//                             ^ reference local 9
    meth := mset.At(i).Obj()
//  ^^^^ definition local 10
//          ^^^^ reference local 8
//               ^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/MethodSet#At().
//                  ^ reference local 9
//                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Selection#Obj().
    sig := types.TypeString(meth.Type(), (*types.Package).Name)
//  ^^^ definition local 11
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//               ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/TypeString().
//                          ^^^^ reference local 10
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Type.
//                                         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/
//                                               ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#
//                                                        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Package#Name().
    fmt.Printf("func (%c *%s) %s%s {\n\tpanic(\"unimplemented\")\n}\n",
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
     r, concname, meth.Name(),
//   ^ reference local 7
//      ^^^^^^^^ reference local 2
//                ^^^^ reference local 10
//                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `go/types`/Object#Name.
     strings.TrimPrefix(sig, "func"))
//   ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//           ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/TrimPrefix().
//                      ^^^ reference local 11
   }
   return nil
  }
  
  //!-
  
  func isValidIdentifier(id string) bool {
//     ^^^^^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/isValidIdentifier().
//     documentation
//     > ```go
//     > func isValidIdentifier(id string) bool
//     > ```
//                       ^^ definition local 12
   for i, r := range id {
//     ^ definition local 13
//        ^ definition local 14
//                   ^^ reference local 12
    if !unicode.IsLetter(r) &&
//      ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 unicode/
//              ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 unicode/IsLetter().
//                       ^ reference local 14
     !(i > 0 && unicode.IsDigit(r)) {
//     ^ reference local 13
//              ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 unicode/
//                      ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 unicode/IsDigit().
//                              ^ reference local 14
     return false
    }
   }
   return id != ""
//        ^^ reference local 12
  }
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   if len(os.Args) != 4 {
//        ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
    log.Fatal(usage)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/usage.
   }
   pkgpath, ifacename, concname := os.Args[1], os.Args[2], os.Args[3]
// ^^^^^^^ definition local 15
//          ^^^^^^^^^ definition local 16
//                     ^^^^^^^^ definition local 17
//                                 ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
//                                             ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                                                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
//                                                         ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                                                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
  
   // Load only exported type information for the specified package.
   conf := &packages.Config{Mode: packages.NeedTypes}
// ^^^^ definition local 18
//          ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                   ^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Config#
//                          ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Config#Mode.
//                                ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                                         ^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/NeedTypes.
   pkgs, err := packages.Load(conf, pkgpath)
// ^^^^ definition local 19
//       ^^^ definition local 20
//              ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//                       ^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Load().
//                            ^^^^ reference local 18
//                                  ^^^^^^^ reference local 15
   if err != nil {
//    ^^^ reference local 20
    log.Fatal(err) // failed to load anything
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 20
   }
   if packages.PrintErrors(pkgs) > 0 {
//    ^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/
//             ^^^^^^^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/PrintErrors().
//                         ^^^^ reference local 19
    os.Exit(1) // some packages contained errors
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Exit().
   }
   if err := PrintSkeleton(pkgs[0].Types, ifacename, concname); err != nil {
//    ^^^ definition local 21
//           ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/skeleton`/PrintSkeleton().
//                         ^^^^ reference local 19
//                                 ^^^^^ reference scip-go gomod golang.org/x/tools v0.14.0 `golang.org/x/tools/go/packages`/Package#Types.
//                                        ^^^^^^^^^ reference local 16
//                                                   ^^^^^^^^ reference local 17
//                                                              ^^^ reference local 21
    log.Fatal(err)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 21
   }
  }
  
  /*
  //!+output1
  $ ./skeleton io ReadWriteCloser buffer
  // *buffer implements io.ReadWriteCloser.
  type buffer struct{}
  func (b *buffer) Close() error {
   panic("unimplemented")
  }
  func (b *buffer) Read(p []byte) (n int, err error) {
   panic("unimplemented")
  }
  func (b *buffer) Write(p []byte) (n int, err error) {
   panic("unimplemented")
  }
  //!-output1
  
  //!+output2
  $ ./skeleton net/http Handler myHandler
  // *myHandler implements net/http.Handler.
  type myHandler struct{}
  func (m *myHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
   panic("unimplemented")
  }
  //!-output2
  */
  
