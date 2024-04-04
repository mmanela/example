  // Copyright 2023 The Go Authors. All rights reserved.
  // Use of this source code is governed by a BSD-style
  // license that can be found in the LICENSE file.
  
  // Template is a trivial web server that uses the text/template (and
  // html/template) package's "block" feature to implement a kind of template
  // inheritance.
  //
  // It should be executed from the directory in which the source resides,
  // as it will look for its template files in the current directory.
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/
//        documentation
//        > Template is a trivial web server that uses the text/template (and
//        > html/template) package's "block" feature to implement a kind of template
//        > inheritance.
//        > 
//        > It should be executed from the directory in which the source resides,
//        > as it will look for its template files in the current directory.
  
  import (
   "html/template"
//  ^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/
   "log"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
   "net/http"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
   "strings"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
  )
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   http.HandleFunc("/", indexHandler)
// ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//      ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/HandleFunc().
//                      ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/indexHandler().
   http.HandleFunc("/image/", imageHandler)
// ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//      ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/HandleFunc().
//                            ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/imageHandler().
   log.Fatal(http.ListenAndServe("localhost:8080", nil))
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//     ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                ^^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/ListenAndServe().
  }
  
  // indexTemplate is the main site template.
  // The default template includes two template blocks ("sidebar" and "content")
  // that may be replaced in templates derived from this one.
  var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))
//    ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/indexTemplate.
//    documentation
//    > ```go
//    > var indexTemplate *Template
//    > ```
//    documentation
//    > indexTemplate is the main site template.
//    > The default template includes two template blocks ("sidebar" and "content")
//    > that may be replaced in templates derived from this one.
//                    ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/
//                             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Must().
//                                  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/
//                                           ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/ParseFiles().
  
  // Index is a data structure used to populate an indexTemplate.
  type Index struct {
//     ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#
//     documentation
//     > ```go
//     > type Index struct
//     > ```
//     documentation
//     > Index is a data structure used to populate an indexTemplate.
//     documentation
//     > ```go
//     > struct {
//     >     Title string
//     >     Body string
//     >     Links []Link
//     > }
//     > ```
   Title string
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Title.
// documentation
// > ```go
// > struct field Title string
// > ```
   Body  string
// ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Body.
// documentation
// > ```go
// > struct field Body string
// > ```
   Links []Link
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Links.
// documentation
// > ```go
// > struct field Links []golang.org/x/example/template.Link
// > ```
//         ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#
  }
  
  type Link struct {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#
//     documentation
//     > ```go
//     > type Link struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     URL string
//     >     Title string
//     > }
//     > ```
   URL, Title string
// ^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#URL.
// documentation
// > ```go
// > struct field URL string
// > ```
//      ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#Title.
//      documentation
//      > ```go
//      > struct field Title string
//      > ```
  }
  
  // indexHandler is an HTTP handler that serves the index page.
  func indexHandler(w http.ResponseWriter, r *http.Request) {
//     ^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/indexHandler().
//     documentation
//     > ```go
//     > func indexHandler(w ResponseWriter, r *Request)
//     > ```
//     documentation
//     > indexHandler is an HTTP handler that serves the index page.
//                  ^ definition local 0
//                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                         ^^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/ResponseWriter#
//                                         ^ definition local 1
//                                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                                                 ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/Request#
   data := &Index{
// ^^^^ definition local 2
//          ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#
    Title: "Image gallery",
//  ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Title.
    Body:  "Welcome to the image gallery.",
//  ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Body.
   }
   for name, img := range images {
//     ^^^^ definition local 3
//           ^^^ definition local 4
//                        ^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/images.
    data.Links = append(data.Links, Link{
//  ^^^^ reference local 2
//       ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Links.
//                      ^^^^ reference local 2
//                           ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Index#Links.
//                                  ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#
     URL:   "/image/" + name,
//   ^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#URL.
//                      ^^^^ reference local 3
     Title: img.Title,
//   ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Link#Title.
//          ^^^ reference local 4
//              ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Image#Title.
    })
   }
   if err := indexTemplate.Execute(w, data); err != nil {
//    ^^^ definition local 5
//           ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/indexTemplate.
//                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Template#Execute().
//                                 ^ reference local 0
//                                    ^^^^ reference local 2
//                                           ^^^ reference local 5
    log.Println(err)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Println().
//              ^^^ reference local 5
   }
  }
  
  // imageTemplate is a clone of indexTemplate that provides
  // alternate "sidebar" and "content" templates.
  var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("image.tmpl"))
//    ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/imageTemplate.
//    documentation
//    > ```go
//    > var imageTemplate *Template
//    > ```
//    documentation
//    > imageTemplate is a clone of indexTemplate that provides
//    > alternate "sidebar" and "content" templates.
//                    ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/
//                             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Must().
//                                  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/
//                                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Must().
//                                                ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/indexTemplate.
//                                                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Template#Clone().
//                                                                       ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Template#ParseFiles().
  
  // Image is a data structure used to populate an imageTemplate.
  type Image struct {
//     ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Image#
//     documentation
//     > ```go
//     > type Image struct
//     > ```
//     documentation
//     > Image is a data structure used to populate an imageTemplate.
//     documentation
//     > ```go
//     > struct {
//     >     Title string
//     >     URL string
//     > }
//     > ```
   Title string
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Image#Title.
// documentation
// > ```go
// > struct field Title string
// > ```
   URL   string
// ^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Image#URL.
// documentation
// > ```go
// > struct field URL string
// > ```
  }
  
  // imageHandler is an HTTP handler that serves the image pages.
  func imageHandler(w http.ResponseWriter, r *http.Request) {
//     ^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/imageHandler().
//     documentation
//     > ```go
//     > func imageHandler(w ResponseWriter, r *Request)
//     > ```
//     documentation
//     > imageHandler is an HTTP handler that serves the image pages.
//                  ^ definition local 6
//                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                         ^^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/ResponseWriter#
//                                         ^ definition local 7
//                                            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//                                                 ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/Request#
   data, ok := images[strings.TrimPrefix(r.URL.Path, "/image/")]
// ^^^^ definition local 8
//       ^^ definition local 9
//             ^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/images.
//                    ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                            ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/TrimPrefix().
//                                       ^ reference local 7
//                                         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/Request#URL.
//                                             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/url`/URL#Path.
   if !ok {
//     ^^ reference local 9
    http.NotFound(w, r)
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/
//       ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `net/http`/NotFound().
//                ^ reference local 6
//                   ^ reference local 7
    return
   }
   if err := imageTemplate.Execute(w, data); err != nil {
//    ^^^ definition local 10
//           ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/imageTemplate.
//                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `html/template`/Template#Execute().
//                                 ^ reference local 6
//                                    ^^^^ reference local 8
//                                           ^^^ reference local 10
    log.Println(err)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Println().
//              ^^^ reference local 10
   }
  }
  
  // images specifies the site content: a collection of images.
  var images = map[string]*Image{
//    ^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/images.
//    documentation
//    > ```go
//    > var images map[string]*Image
//    > ```
//    documentation
//    > images specifies the site content: a collection of images.
//                         ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/template`/Image#
   "go":     {"The Go Gopher", "https://golang.org/doc/gopher/frontpage.png"},
   "google": {"The Google Logo", "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"},
  }
  
