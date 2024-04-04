  // The weave command is a simple preprocessor for markdown files.
  // It builds a table of contents and processes %include directives.
  //
  // Example usage:
  //
  // $ go run internal/cmd/weave go-types.md > README.md
  //
  // The weave command copies lines of the input file to standard output, with two
  // exceptions:
  //
  // If a line begins with "%toc", it is replaced with a table of contents
  // consisting of links to the top two levels of headers ("#" and "##").
  //
  // If a line begins with "%include FILENAME TAG", it is replaced with the lines
  // of the file between lines containing "!+TAG" and  "!-TAG". TAG can be omitted,
  // in which case the delimiters are simply "!+" and "!-".
  //
  // Before the included lines, a line of the form
  //
  // // go get PACKAGE
  //
  // is output, where PACKAGE is constructed from the module path, the
  // base name of the current directory, and the directory of FILENAME.
  // This caption can be supressed by putting "-" as the final word of the %include line.
  package main
//        ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/
//        documentation
//        > The weave command is a simple preprocessor for markdown files.
//        > It builds a table of contents and processes %include directives.
//        > 
//        > Example usage:
//        > 
//        > 	$ go run internal/cmd/weave go-types.md > README.md
//        > 
//        > The weave command copies lines of the input file to standard output, with two
//        > exceptions:
//        > 
//        > If a line begins with "%toc", it is replaced with a table of contents
//        > consisting of links to the top two levels of headers ("#" and "##").
//        > 
//        > If a line begins with "%include FILENAME TAG", it is replaced with the lines
//        > of the file between lines containing "!+TAG" and  "!-TAG". TAG can be omitted,
//        > in which case the delimiters are simply "!+" and "!-".
//        > 
//        > Before the included lines, a line of the form
//        > 
//        > 	// go get PACKAGE
//        > 
//        > is output, where PACKAGE is constructed from the module path, the
//        > base name of the current directory, and the directory of FILENAME.
//        > This caption can be supressed by putting "-" as the final word of the %include line.
  
  import (
   "bufio"
//  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/
   "bytes"
//  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "log"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
   "os"
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
   "path/filepath"
//  ^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `path/filepath`/
   "regexp"
//  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/
   "strings"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
  )
  
  func main() {
//     ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   log.SetFlags(0)
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//     ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/SetFlags().
   log.SetPrefix("weave: ")
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//     ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/SetPrefix().
   if len(os.Args) != 2 {
//        ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
    log.Fatal("usage: weave input.md\n")
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
   }
  
   f, err := os.Open(os.Args[1])
// ^ definition local 0
//    ^^^ definition local 1
//           ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//              ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Open().
//                   ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Args.
   if err != nil {
//    ^^^ reference local 1
    log.Fatal(err)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 1
   }
   defer f.Close()
//       ^ reference local 0
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/File#Close().
  
   wd, err := os.Getwd()
// ^^ definition local 2
//     ^^^ reference local 1
//            ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//               ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Getwd().
   if err != nil {
//    ^^^ reference local 1
    log.Fatal(err)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^^ reference local 1
   }
   curDir := filepath.Base(wd)
// ^^^^^^ definition local 3
//           ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `path/filepath`/
//                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `path/filepath`/Base().
//                         ^^ reference local 2
  
   fmt.Println("<!-- Autogenerated by weave; DO NOT EDIT -->")
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//     ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
  
   // Pass 1: extract table of contents.
   var toc []string
//     ^^^ definition local 4
   in := bufio.NewScanner(f)
// ^^ definition local 5
//       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/
//             ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/NewScanner().
//                        ^ reference local 0
   for in.Scan() {
//     ^^ reference local 5
//        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Scan().
    line := in.Text()
//  ^^^^ definition local 6
//          ^^ reference local 5
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Text().
    if line == "" || (line[0] != '#' && line[0] != '%') {
//     ^^^^ reference local 6
//                    ^^^^ reference local 6
//                                      ^^^^ reference local 6
     continue
    }
    line = strings.TrimSpace(line)
//  ^^^^ reference local 6
//         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                 ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/TrimSpace().
//                           ^^^^ reference local 6
    if line == "%toc" {
//     ^^^^ reference local 6
     toc = nil
//   ^^^ reference local 4
    } else if strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") {
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                    ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/HasPrefix().
//                              ^^^^ reference local 6
//                                             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                                                     ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/HasPrefix().
//                                                               ^^^^ reference local 6
     words := strings.Fields(line)
//   ^^^^^ definition local 7
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                    ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Fields().
//                           ^^^^ reference local 6
     depth := len(words[0])
//   ^^^^^ definition local 8
//                ^^^^^ reference local 7
     words = words[1:]
//   ^^^^^ reference local 7
//           ^^^^^ reference local 7
     text := strings.Join(words, " ")
//   ^^^^ definition local 9
//           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                   ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Join().
//                        ^^^^^ reference local 7
     for i := range words {
//       ^ definition local 10
//                  ^^^^^ reference local 7
      words[i] = strings.ToLower(words[i])
//    ^^^^^ reference local 7
//          ^ reference local 10
//               ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/ToLower().
//                               ^^^^^ reference local 7
//                                     ^ reference local 10
     }
     line = fmt.Sprintf("%s1. [%s](#%s)",
//   ^^^^ reference local 6
//          ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//              ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Sprintf().
      strings.Repeat("\t", depth-1), text, strings.Join(words, "-"))
//    ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Repeat().
//                         ^^^^^ reference local 8
//                                   ^^^^ reference local 9
//                                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Join().
//                                                      ^^^^^ reference local 7
     toc = append(toc, line)
//   ^^^ reference local 4
//                ^^^ reference local 4
//                     ^^^^ reference local 6
    }
   }
   if in.Err() != nil {
//    ^^ reference local 5
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Err().
    log.Fatal(in.Err())
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^ reference local 5
//               ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Err().
   }
  
   // Pass 2.
   if _, err := f.Seek(0, os.SEEK_SET); err != nil {
//       ^^^ definition local 11
//              ^ reference local 0
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/File#Seek().
//                        ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//                           ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/SEEK_SET.
//                                      ^^^ reference local 11
    log.Fatalf("can't rewind input: %v", err)
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatalf().
//                                       ^^^ reference local 11
   }
   in = bufio.NewScanner(f)
// ^^ reference local 5
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/
//            ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/NewScanner().
//                       ^ reference local 0
   for in.Scan() {
//     ^^ reference local 5
//        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Scan().
    line := in.Text()
//  ^^^^ definition local 12
//          ^^ reference local 5
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Text().
    switch {
    case strings.HasPrefix(line, "%toc"): // ToC
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//               ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/HasPrefix().
//                         ^^^^ reference local 12
     for _, h := range toc {
//          ^ definition local 13
//                     ^^^ reference local 4
      fmt.Println(h)
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//                ^ reference local 13
     }
    case strings.HasPrefix(line, "%include"):
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//               ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/HasPrefix().
//                         ^^^^ reference local 12
     words := strings.Fields(line)
//   ^^^^^ definition local 14
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                    ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Fields().
//                           ^^^^ reference local 12
     if len(words) < 2 {
//          ^^^^^ reference local 14
      log.Fatal(line)
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//        ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//              ^^^^ reference local 12
     }
     filename := words[1]
//   ^^^^^^^^ definition local 15
//               ^^^^^ reference local 14
  
     // Show caption unless '-' follows.
     if len(words) < 4 || words[3] != "-" {
//          ^^^^^ reference local 14
//                        ^^^^^ reference local 14
      fmt.Printf(" // go get golang.org/x/example/%s/%s\n\n",
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Printf().
       curDir, filepath.Dir(filename))
//     ^^^^^^ reference local 3
//             ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `path/filepath`/
//                      ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `path/filepath`/Dir().
//                          ^^^^^^^^ reference local 15
     }
  
     section := ""
//   ^^^^^^^ definition local 16
     if len(words) > 2 {
//          ^^^^^ reference local 14
      section = words[2]
//    ^^^^^^^ reference local 16
//              ^^^^^ reference local 14
     }
     s, err := include(filename, section)
//   ^ definition local 17
//      ^^^ definition local 18
//             ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/include().
//                     ^^^^^^^^ reference local 15
//                               ^^^^^^^ reference local 16
     if err != nil {
//      ^^^ reference local 18
      log.Fatal(err)
//    ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//        ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//              ^^^ reference local 18
     }
     fmt.Println("```")
//   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
     fmt.Println(cleanListing(s)) // TODO(adonovan): escape /^```/ in s
//   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//               ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/cleanListing().
//                            ^ reference local 17
     fmt.Println("```")
//   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
    default:
     fmt.Println(line)
//   ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//       ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Println().
//               ^^^^ reference local 12
    }
   }
   if in.Err() != nil {
//    ^^ reference local 5
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Err().
    log.Fatal(in.Err())
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 log/Fatal().
//            ^^ reference local 5
//               ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Err().
   }
  }
  
  // include processes an included file, and returns the included text.
  // Only lines between those matching !+tag and !-tag will be returned.
  // This is true even if tag=="".
  func include(file, tag string) (string, error) {
//     ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/include().
//     documentation
//     > ```go
//     > func include(file string, tag string) (string, error)
//     > ```
//     documentation
//     > include processes an included file, and returns the included text.
//     > Only lines between those matching !+tag and !-tag will be returned.
//     > This is true even if tag=="".
//             ^^^^ definition local 19
//                   ^^^ definition local 20
   f, err := os.Open(file)
// ^ definition local 21
//    ^^^ definition local 22
//           ^^ reference scip-go gomod github.com/golang/go/src go1.18 os/
//              ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/Open().
//                   ^^^^ reference local 19
   if err != nil {
//    ^^^ reference local 22
    return "", err
//             ^^^ reference local 22
   }
   defer f.Close()
//       ^ reference local 21
//         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 os/File#Close().
  
   startre, err := regexp.Compile("!\\+" + tag + "$")
// ^^^^^^^ definition local 23
//          ^^^ reference local 22
//                 ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/
//                        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/Compile().
//                                         ^^^ reference local 20
   if err != nil {
//    ^^^ reference local 22
    return "", err
//             ^^^ reference local 22
   }
   endre, err := regexp.Compile("!\\-" + tag + "$")
// ^^^^^ definition local 24
//        ^^^ reference local 22
//               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/
//                      ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/Compile().
//                                       ^^^ reference local 20
   if err != nil {
//    ^^^ reference local 22
    return "", err
//             ^^^ reference local 22
   }
  
   var text bytes.Buffer
//     ^^^^ definition local 25
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/
//                ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#
   in := bufio.NewScanner(f)
// ^^ definition local 26
//       ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/
//             ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/NewScanner().
//                        ^ reference local 21
   var on bool
//     ^^ definition local 27
   for in.Scan() {
//     ^^ reference local 26
//        ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Scan().
    line := in.Text()
//  ^^^^ definition local 28
//          ^^ reference local 26
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Text().
    switch {
    case startre.MatchString(line):
//       ^^^^^^^ reference local 23
//               ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/Regexp#MatchString().
//                           ^^^^ reference local 28
     on = true
//   ^^ reference local 27
    case endre.MatchString(line):
//       ^^^^^ reference local 24
//             ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 regexp/Regexp#MatchString().
//                         ^^^^ reference local 28
     on = false
//   ^^ reference local 27
    case on:
//       ^^ reference local 27
     text.WriteByte('\t')
//   ^^^^ reference local 25
//        ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#WriteByte().
     text.WriteString(line)
//   ^^^^ reference local 25
//        ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#WriteString().
//                    ^^^^ reference local 28
     text.WriteByte('\n')
//   ^^^^ reference local 25
//        ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#WriteByte().
    }
   }
   if in.Err() != nil {
//    ^^ reference local 26
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Err().
    return "", in.Err()
//             ^^ reference local 26
//                ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bufio/Scanner#Err().
   }
   if text.Len() == 0 {
//    ^^^^ reference local 25
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#Len().
    return "", fmt.Errorf("no lines of %s matched tag %q", file, tag)
//             ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//                 ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Errorf().
//                                                         ^^^^ reference local 19
//                                                               ^^^ reference local 20
   }
   return text.String(), nil
//        ^^^^ reference local 25
//             ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 bytes/Buffer#String().
  }
  
  func isBlank(line string) bool { return strings.TrimSpace(line) == "" }
//     ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/isBlank().
//     documentation
//     > ```go
//     > func isBlank(line string) bool
//     > ```
//             ^^^^ definition local 29
//                                        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                                                ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/TrimSpace().
//                                                          ^^^^ reference local 29
  
  func indented(line string) bool {
//     ^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/indented().
//     documentation
//     > ```go
//     > func indented(line string) bool
//     > ```
//              ^^^^ definition local 30
   return strings.HasPrefix(line, "    ") || strings.HasPrefix(line, "\t")
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/HasPrefix().
//                          ^^^^ reference local 30
//                                           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                                                   ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/HasPrefix().
//                                                             ^^^^ reference local 30
  }
  
  // cleanListing removes entirely blank leading and trailing lines from
  // text, and removes n leading tabs.
  func cleanListing(text string) string {
//     ^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/cleanListing().
//     documentation
//     > ```go
//     > func cleanListing(text string) string
//     > ```
//     documentation
//     > cleanListing removes entirely blank leading and trailing lines from
//     > text, and removes n leading tabs.
//                  ^^^^ definition local 31
   lines := strings.Split(text, "\n")
// ^^^^^ definition local 32
//          ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                  ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Split().
//                        ^^^^ reference local 31
  
   // remove minimum number of leading tabs from all non-blank lines
   tabs := 999
// ^^^^ definition local 33
   for i, line := range lines {
//     ^ definition local 34
//        ^^^^ definition local 35
//                      ^^^^^ reference local 32
    if strings.TrimSpace(line) == "" {
//     ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//             ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/TrimSpace().
//                       ^^^^ reference local 35
     lines[i] = ""
//   ^^^^^ reference local 32
//         ^ reference local 34
    } else {
     if n := leadingTabs(line); n < tabs {
//      ^ definition local 36
//           ^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/leadingTabs().
//                       ^^^^ reference local 35
//                              ^ reference local 36
//                                  ^^^^ reference local 33
      tabs = n
//    ^^^^ reference local 33
//           ^ reference local 36
     }
    }
   }
   for i, line := range lines {
//     ^ definition local 37
//        ^^^^ definition local 38
//                      ^^^^^ reference local 32
    if line != "" {
//     ^^^^ reference local 38
     line := line[tabs:]
//   ^^^^ definition local 39
//           ^^^^ reference local 38
//                ^^^^ reference local 33
     lines[i] = line // remove leading tabs
//   ^^^^^ reference local 32
//         ^ reference local 37
//              ^^^^ reference local 39
    }
   }
  
   // remove leading blank lines
   for len(lines) > 0 && lines[0] == "" {
//         ^^^^^ reference local 32
//                       ^^^^^ reference local 32
    lines = lines[1:]
//  ^^^^^ reference local 32
//          ^^^^^ reference local 32
   }
   // remove trailing blank lines
   for len(lines) > 0 && lines[len(lines)-1] == "" {
//         ^^^^^ reference local 32
//                       ^^^^^ reference local 32
//                                 ^^^^^ reference local 32
    lines = lines[:len(lines)-1]
//  ^^^^^ reference local 32
//          ^^^^^ reference local 32
//                     ^^^^^ reference local 32
   }
   return strings.Join(lines, "\n")
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strings/Join().
//                     ^^^^^ reference local 32
  }
  
  func leadingTabs(s string) int {
//     ^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/internal/cmd/weave`/leadingTabs().
//     documentation
//     > ```go
//     > func leadingTabs(s string) int
//     > ```
//                 ^ definition local 40
   var i int
//     ^ definition local 41
   for i = 0; i < len(s); i++ {
//     ^ reference local 41
//            ^ reference local 41
//                    ^ reference local 40
//                        ^ reference local 41
    if s[i] != '\t' {
//     ^ reference local 40
//       ^ reference local 41
     break
    }
   }
   return i
//        ^ reference local 41
  }
  
