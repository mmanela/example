  //go:build go1.21
  
  package indenthandler
//        ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/
//        documentation
//        > package indenthandler
  
  import (
   "context"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/
   "fmt"
//  ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
   "io"
//  ^^ reference scip-go gomod github.com/golang/go/src go1.18 io/
   "log/slog"
//  ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
   "runtime"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/
   "slices"
//  ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 slices/
   "strconv"
//  ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strconv/
   "sync"
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
   "time"
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/
  )
  
  // !+IndentHandler
  type IndentHandler struct {
//     ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//     documentation
//     > ```go
//     > type IndentHandler struct
//     > ```
//     documentation
//     > !+IndentHandler
//     documentation
//     > ```go
//     > struct {
//     >     opts Options
//     >     preformatted []byte
//     >     unopenedGroups []string
//     >     indentLevel int
//     >     mu *Mutex
//     >     out Writer
//     > }
//     > ```
//     relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler# implementation
   opts           Options
// ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#opts.
// documentation
// > ```go
// > struct field opts golang.org/x/example/slog-handler-guide/indenthandler4.Options
// > ```
//                ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#
   preformatted   []byte   // data from WithGroup and WithAttrs
// ^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#preformatted.
// documentation
// > ```go
// > struct field preformatted []byte
// > ```
   unopenedGroups []string // groups from WithGroup that haven't been opened
// ^^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
// documentation
// > ```go
// > struct field unopenedGroups []string
// > ```
   indentLevel    int      // same as number of opened groups so far
// ^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#indentLevel.
// documentation
// > ```go
// > struct field indentLevel int
// > ```
   mu             *sync.Mutex
// ^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#mu.
// documentation
// > ```go
// > struct field mu *sync.Mutex
// > ```
//                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
//                      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#
   out            io.Writer
// ^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#out.
// documentation
// > ```go
// > struct field out io.Writer
// > ```
//                ^^ reference scip-go gomod github.com/golang/go/src go1.18 io/
//                   ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 io/Writer#
  }
  
  //!-IndentHandler
  
  type Options struct {
//     ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#
//     documentation
//     > ```go
//     > type Options struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Level Leveler
//     > }
//     > ```
   // Level reports the minimum level to log.
   // Levels with lower levels are discarded.
   // If nil, the Handler uses [slog.LevelInfo].
   Level slog.Leveler
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#Level.
// documentation
// > ```go
// > struct field Level log/slog.Leveler
// > ```
//       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Leveler#
  }
  
  func New(out io.Writer, opts *Options) *IndentHandler {
//     ^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/New().
//     documentation
//     > ```go
//     > func New(out Writer, opts *Options) *IndentHandler
//     > ```
//         ^^^ definition local 0
//             ^^ reference scip-go gomod github.com/golang/go/src go1.18 io/
//                ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 io/Writer#
//                        ^^^^ definition local 1
//                              ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#
//                                        ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
   h := &IndentHandler{out: out, mu: &sync.Mutex{}}
// ^ definition local 2
//       ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                     ^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#out.
//                          ^^^ reference local 0
//                               ^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#mu.
//                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
//                                         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#
   if opts != nil {
//    ^^^^ reference local 1
    h.opts = *opts
//  ^ reference local 2
//    ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#opts.
//            ^^^^ reference local 1
   }
   if h.opts.Level == nil {
//    ^ reference local 2
//      ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#opts.
//           ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#Level.
    h.opts.Level = slog.LevelInfo
//  ^ reference local 2
//    ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#opts.
//         ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#Level.
//                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                      ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/LevelInfo.
   }
   return h
//        ^ reference local 2
  }
  
  func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
//      ^ definition local 3
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                        ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#Enabled().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).Enabled(ctx Context, level Level) bool
//                        > ```
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#Enabled. implementation
//                                ^^^ definition local 4
//                                    ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/
//                                            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/Context#
//                                                     ^^^^^ definition local 5
//                                                           ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                                ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Level#
   return level >= h.opts.Level.Level()
//        ^^^^^ reference local 5
//                 ^ reference local 3
//                   ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#opts.
//                        ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/Options#Level.
//                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Leveler#Level.
  }
  
  // !+WithGroup
  func (h *IndentHandler) WithGroup(name string) slog.Handler {
//      ^ definition local 6
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                        ^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#WithGroup().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).WithGroup(name string) Handler
//                        > ```
//                        documentation
//                        > !+WithGroup
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#WithGroup. implementation
//                                  ^^^^ definition local 7
//                                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                    ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#
   if name == "" {
//    ^^^^ reference local 7
    return h
//         ^ reference local 6
   }
   h2 := *h
// ^^ definition local 8
//        ^ reference local 6
   // Add an unopened group to h2 without modifying h.
   h2.unopenedGroups = make([]string, len(h.unopenedGroups)+1)
// ^^ reference local 8
//    ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
//                                        ^ reference local 6
//                                          ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
   copy(h2.unopenedGroups, h.unopenedGroups)
//      ^^ reference local 8
//         ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
//                         ^ reference local 6
//                           ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
   h2.unopenedGroups[len(h2.unopenedGroups)-1] = name
// ^^ reference local 8
//    ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
//                       ^^ reference local 8
//                          ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
//                                               ^^^^ reference local 7
   return &h2
//         ^^ reference local 8
  }
  
  //!-WithGroup
  
  // !+WithAttrs
  func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
//      ^ definition local 9
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                        ^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#WithAttrs().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).WithAttrs(attrs []Attr) Handler
//                        > ```
//                        documentation
//                        > !+WithAttrs
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#WithAttrs. implementation
//                                  ^^^^^ definition local 10
//                                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
//                                                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                          ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#
   if len(attrs) == 0 {
//        ^^^^^ reference local 10
    return h
//         ^ reference local 9
   }
   h2 := *h
// ^^ definition local 11
//        ^ reference local 9
   // Force an append to copy the underlying array.
   pre := slices.Clip(h.preformatted)
// ^^^ definition local 12
//        ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 slices/
//               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 slices/Clip().
//                    ^ reference local 9
//                      ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#preformatted.
   // Add all groups from WithGroup that haven't already been added.
   h2.preformatted = h2.appendUnopenedGroups(pre, h2.indentLevel)
// ^^ reference local 11
//    ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#preformatted.
//                   ^^ reference local 11
//                      ^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendUnopenedGroups().
//                                           ^^^ reference local 12
//                                                ^^ reference local 11
//                                                   ^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#indentLevel.
   // Each of those groups increased the indent level by 1.
   h2.indentLevel += len(h2.unopenedGroups)
// ^^ reference local 11
//    ^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#indentLevel.
//                       ^^ reference local 11
//                          ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
   // Now all groups have been opened.
   h2.unopenedGroups = nil
// ^^ reference local 11
//    ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
   // Pre-format the attributes.
   for _, a := range attrs {
//        ^ definition local 13
//                   ^^^^^ reference local 10
    h2.preformatted = h2.appendAttr(h2.preformatted, a, h2.indentLevel)
//  ^^ reference local 11
//     ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#preformatted.
//                    ^^ reference local 11
//                       ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                                  ^^ reference local 11
//                                     ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#preformatted.
//                                                   ^ reference local 13
//                                                      ^^ reference local 11
//                                                         ^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#indentLevel.
   }
   return &h2
//         ^^ reference local 11
  }
  
  func (h *IndentHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
//      ^ definition local 14
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                        ^^^^^^^^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendUnopenedGroups().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).appendUnopenedGroups(buf []byte, indentLevel int) []byte
//                        > ```
//                                             ^^^ definition local 15
//                                                         ^^^^^^^^^^^ definition local 16
   for _, g := range h.unopenedGroups {
//        ^ definition local 17
//                   ^ reference local 14
//                     ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
    buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
//  ^^^ reference local 15
//        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                    ^^^ reference local 15
//                                     ^^^^^^^^^^^ reference local 16
//                                                        ^ reference local 17
    indentLevel++
//  ^^^^^^^^^^^ reference local 16
   }
   return buf
//        ^^^ reference local 15
  }
  
  //!-WithAttrs
  
  // !+Handle
  func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
//      ^ definition local 18
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                        ^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#Handle().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).Handle(ctx Context, r Record) error
//                        > ```
//                        documentation
//                        > !+Handle
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#Handle. implementation
//                               ^^^ definition local 19
//                                   ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/
//                                           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/Context#
//                                                    ^ definition local 20
//                                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#
   bufp := allocBuf()
// ^^^^ definition local 21
//         ^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/allocBuf().
   buf := *bufp
// ^^^ definition local 22
//         ^^^^ reference local 21
   defer func() {
    *bufp = buf
//   ^^^^ reference local 21
//          ^^^ reference local 22
    freeBuf(bufp)
//  ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/freeBuf().
//          ^^^^ reference local 21
   }()
   if !r.Time.IsZero() {
//     ^ reference local 20
//       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Time.
//            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/Time#IsZero().
    buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
//  ^^^ reference local 22
//        ^ reference local 18
//          ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                     ^^^ reference local 22
//                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Time().
//                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/TimeKey.
//                                                  ^ reference local 20
//                                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Time.
   }
   buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
// ^^^ reference local 22
//       ^ reference local 18
//         ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                    ^^^ reference local 22
//                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                              ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Any().
//                                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                       ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/LevelKey.
//                                                 ^ reference local 20
//                                                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Level.
   if r.PC != 0 {
//    ^ reference local 20
//      ^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#PC.
    fs := runtime.CallersFrames([]uintptr{r.PC})
//  ^^ definition local 23
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/
//                ^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/CallersFrames().
//                                        ^ reference local 20
//                                          ^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#PC.
    f, _ := fs.Next()
//  ^ definition local 24
//          ^^ reference local 23
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/Frames#Next().
    // Optimize to minimize allocation.
    srcbufp := allocBuf()
//  ^^^^^^^ definition local 25
//             ^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/allocBuf().
    defer freeBuf(srcbufp)
//        ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/freeBuf().
//                ^^^^^^^ reference local 25
    *srcbufp = append(*srcbufp, f.File...)
//   ^^^^^^^ reference local 25
//                     ^^^^^^^ reference local 25
//                              ^ reference local 24
//                                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/Frame#File.
    *srcbufp = append(*srcbufp, ':')
//   ^^^^^^^ reference local 25
//                     ^^^^^^^ reference local 25
    *srcbufp = strconv.AppendInt(*srcbufp, int64(f.Line), 10)
//   ^^^^^^^ reference local 25
//             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strconv/
//                     ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strconv/AppendInt().
//                                ^^^^^^^ reference local 25
//                                               ^ reference local 24
//                                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/Frame#Line.
    buf = h.appendAttr(buf, slog.String(slog.SourceKey, string(*srcbufp)), 0)
//  ^^^ reference local 22
//        ^ reference local 18
//          ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                     ^^^ reference local 22
//                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/String().
//                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                           ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/SourceKey.
//                                                              ^^^^^^^ reference local 25
   }
  
   buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
// ^^^ reference local 22
//       ^ reference local 18
//         ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                    ^^^ reference local 22
//                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                              ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/String().
//                                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                          ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/MessageKey.
//                                                      ^ reference local 20
//                                                        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Message.
   // Insert preformatted attributes just after built-in ones.
   buf = append(buf, h.preformatted...)
// ^^^ reference local 22
//              ^^^ reference local 22
//                   ^ reference local 18
//                     ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#preformatted.
   if r.NumAttrs() > 0 {
//    ^ reference local 20
//      ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#NumAttrs().
    buf = h.appendUnopenedGroups(buf, h.indentLevel)
//  ^^^ reference local 22
//        ^ reference local 18
//          ^^^^^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendUnopenedGroups().
//                               ^^^ reference local 22
//                                    ^ reference local 18
//                                      ^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#indentLevel.
    r.Attrs(func(a slog.Attr) bool {
//  ^ reference local 20
//    ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Attrs().
//               ^ definition local 26
//                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
     buf = h.appendAttr(buf, a, h.indentLevel+len(h.unopenedGroups))
//   ^^^ reference local 22
//         ^ reference local 18
//           ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                      ^^^ reference local 22
//                           ^ reference local 26
//                              ^ reference local 18
//                                ^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#indentLevel.
//                                                ^ reference local 18
//                                                  ^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#unopenedGroups.
     return true
    })
   }
   buf = append(buf, "---\n"...)
// ^^^ reference local 22
//              ^^^ reference local 22
   h.mu.Lock()
// ^ reference local 18
//   ^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#mu.
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#Lock().
   defer h.mu.Unlock()
//       ^ reference local 18
//         ^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#mu.
//            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#Unlock().
   _, err := h.out.Write(buf)
//    ^^^ definition local 27
//           ^ reference local 18
//             ^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#out.
//                 ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 io/Writer#Write.
//                       ^^^ reference local 22
   return err
//        ^^^ reference local 27
  }
  
  //!-Handle
  
  func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
//      ^ definition local 28
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#
//                        ^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).appendAttr(buf []byte, a Attr, indentLevel int) []byte
//                        > ```
//                                   ^^^ definition local 29
//                                               ^ definition local 30
//                                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
//                                                            ^^^^^^^^^^^ definition local 31
   // Resolve the Attr's value before doing anything else.
   a.Value = a.Value.Resolve()
// ^ reference local 30
//   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//           ^ reference local 30
//             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                   ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Resolve().
   // Ignore empty Attrs.
   if a.Equal(slog.Attr{}) {
//    ^ reference local 30
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Equal().
//            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
    return buf
//         ^^^ reference local 29
   }
   // Indent 4 spaces per level.
   buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
// ^^^ reference local 29
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                   ^^^ reference local 29
//                               ^^^^^^^^^^^ reference local 31
   switch a.Value.Kind() {
//        ^ reference local 30
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Kind().
   case slog.KindString:
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//           ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/KindString.
    // Quote string values, to make them easy to parse.
    buf = append(buf, a.Key...)
//  ^^^ reference local 29
//               ^^^ reference local 29
//                    ^ reference local 30
//                      ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
    buf = append(buf, ": "...)
//  ^^^ reference local 29
//               ^^^ reference local 29
    buf = strconv.AppendQuote(buf, a.Value.String())
//  ^^^ reference local 29
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strconv/
//                ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 strconv/AppendQuote().
//                            ^^^ reference local 29
//                                 ^ reference local 30
//                                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                                         ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#String().
    buf = append(buf, '\n')
//  ^^^ reference local 29
//               ^^^ reference local 29
   case slog.KindTime:
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//           ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/KindTime.
    // Write times in a standard way, without the monotonic time.
    buf = append(buf, a.Key...)
//  ^^^ reference local 29
//               ^^^ reference local 29
//                    ^ reference local 30
//                      ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
    buf = append(buf, ": "...)
//  ^^^ reference local 29
//               ^^^ reference local 29
    buf = a.Value.Time().AppendFormat(buf, time.RFC3339Nano)
//  ^^^ reference local 29
//        ^ reference local 30
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Time().
//                       ^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/Time#AppendFormat().
//                                    ^^^ reference local 29
//                                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/
//                                              ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/RFC3339Nano.
    buf = append(buf, '\n')
//  ^^^ reference local 29
//               ^^^ reference local 29
   case slog.KindGroup:
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//           ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/KindGroup.
    attrs := a.Value.Group()
//  ^^^^^ definition local 32
//           ^ reference local 30
//             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Group().
    // Ignore empty groups.
    if len(attrs) == 0 {
//         ^^^^^ reference local 32
     return buf
//          ^^^ reference local 29
    }
    // If the key is non-empty, write it out and indent the rest of the attrs.
    // Otherwise, inline the attrs.
    if a.Key != "" {
//     ^ reference local 30
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
     buf = fmt.Appendf(buf, "%s:\n", a.Key)
//   ^^^ reference local 29
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                     ^^^ reference local 29
//                                   ^ reference local 30
//                                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
     indentLevel++
//   ^^^^^^^^^^^ reference local 31
    }
    for _, ga := range attrs {
//         ^^ definition local 33
//                     ^^^^^ reference local 32
     buf = h.appendAttr(buf, ga, indentLevel)
//   ^^^ reference local 29
//         ^ reference local 28
//           ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/IndentHandler#appendAttr().
//                      ^^^ reference local 29
//                           ^^ reference local 33
//                               ^^^^^^^^^^^ reference local 31
    }
  
   default:
    buf = append(buf, a.Key...)
//  ^^^ reference local 29
//               ^^^ reference local 29
//                    ^ reference local 30
//                      ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
    buf = append(buf, ": "...)
//  ^^^ reference local 29
//               ^^^ reference local 29
    buf = append(buf, a.Value.String()...)
//  ^^^ reference local 29
//               ^^^ reference local 29
//                    ^ reference local 30
//                      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#String().
    buf = append(buf, '\n')
//  ^^^ reference local 29
//               ^^^ reference local 29
   }
   return buf
//        ^^^ reference local 29
  }
  
  // !+pool
  var bufPool = sync.Pool{
//    ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/bufPool.
//    documentation
//    > ```go
//    > var bufPool Pool
//    > ```
//    documentation
//    > !+pool
//              ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
//                   ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Pool#
   New: func() any {
// ^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Pool#New.
    b := make([]byte, 0, 1024)
//  ^ definition local 34
    return &b
//          ^ reference local 34
   },
  }
  
  func allocBuf() *[]byte {
//     ^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/allocBuf().
//     documentation
//     > ```go
//     > func allocBuf() *[]byte
//     > ```
   return bufPool.Get().(*[]byte)
//        ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/bufPool.
//                ^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Pool#Get().
  }
  
  func freeBuf(b *[]byte) {
//     ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/freeBuf().
//     documentation
//     > ```go
//     > func freeBuf(b *[]byte)
//     > ```
//             ^ definition local 35
   // To reduce peak allocation, return only smaller buffers to the pool.
   const maxBufferSize = 16 << 10
//       ^^^^^^^^^^^^^ definition local 36
   if cap(*b) <= maxBufferSize {
//         ^ reference local 35
//               ^^^^^^^^^^^^^ reference local 36
    *b = (*b)[:0]
//   ^ reference local 35
//         ^ reference local 35
    bufPool.Put(b)
//  ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler4`/bufPool.
//          ^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Pool#Put().
//              ^ reference local 35
   }
  }
  
  //!-pool
  
