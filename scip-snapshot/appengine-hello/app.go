  // Copyright 2023 The Go Authors. All rights reserved.
  // Use of this source code is governed by a BSD-style
  // license that can be found in the LICENSE file.
  
  // Package hello is a simple App Engine application that replies to requests
  // on /hello with a welcoming message.
  package hello
//        ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/appengine-hello`/
//        documentation
//        > Package hello is a simple App Engine application that replies to requests
//        > on /hello with a welcoming message.
  
  import (
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "net/http"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
  )
  
  // init is run before the application starts serving.
  func init() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/appengine-hello`/init().
//     documentation
//     > ```go
//     > func init()
//     > ```
//     documentation
//     > init is run before the application starts serving.
   // Handle all requests with path /hello with the helloHandler function.
   http.HandleFunc("/hello", helloHandler)
// ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//      ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/HandleFunc().
//                           ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/appengine-hello`/helloHandler().
  }
  
  func helloHandler(w http.ResponseWriter, r *http.Request) {
//     ^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/appengine-hello`/helloHandler().
//     documentation
//     > ```go
//     > func helloHandler(w ResponseWriter, r *Request)
//     > ```
//                  ^ definition local 0
//                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                         ^^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/ResponseWriter#
//                                         ^ definition local 1
//                                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                                                 ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/Request#
   fmt.Fprintln(w, "Hello from the Go app")
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Fprintln().
//              ^ reference local 0
  }
  
