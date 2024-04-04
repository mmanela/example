  // !+
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hello`/
//        documentation
//        > !+
  
  import "fmt"
//        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/gotypes/hello`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   fmt.Println("Hello, 世界")
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
  }
  
  //!-
  
