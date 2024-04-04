  //go:build go1.21
  
  package indenthandler
//        ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/
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
   "sync"
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
   "time"
//  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/
  )
  
  // !+IndentHandler
  type IndentHandler struct {
//     ^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
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
//     >     goas []groupOrAttrs
//     >     mu *Mutex
//     >     out Writer
//     > }
//     > ```
//     relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler# implementation
   opts Options
// ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#opts.
// documentation
// > ```go
// > struct field opts golang.org/x/example/slog-handler-guide/indenthandler2.Options
// > ```
//      ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#
   goas []groupOrAttrs
// ^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
// documentation
// > ```go
// > struct field goas []golang.org/x/example/slog-handler-guide/indenthandler2.groupOrAttrs
// > ```
//        ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#
   mu   *sync.Mutex
// ^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#mu.
// documentation
// > ```go
// > struct field mu *sync.Mutex
// > ```
//       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
//            ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#
   out  io.Writer
// ^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#out.
// documentation
// > ```go
// > struct field out io.Writer
// > ```
//      ^^ reference scip-go gomod github.com/golang/go/src go1.18 io/
//         ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 io/Writer#
  }
  
  //!-IndentHandler
  
  type Options struct {
//     ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#
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
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#Level.
// documentation
// > ```go
// > struct field Level log/slog.Leveler
// > ```
//       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Leveler#
  }
  
  // !+gora
  // groupOrAttrs holds either a group name or a list of slog.Attrs.
  type groupOrAttrs struct {
//     ^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#
//     documentation
//     > ```go
//     > type groupOrAttrs struct
//     > ```
//     documentation
//     > !+gora
//     > groupOrAttrs holds either a group name or a list of slog.Attrs.
//     documentation
//     > ```go
//     > struct {
//     >     group string
//     >     attrs []Attr
//     > }
//     > ```
   group string      // group name if non-empty
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#group.
// documentation
// > ```go
// > struct field group string
// > ```
   attrs []slog.Attr // attrs if non-empty
// ^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#attrs.
// documentation
// > ```go
// > struct field attrs []log/slog.Attr
// > ```
//         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//              ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
  }
  
  //!-gora
  
  func New(out io.Writer, opts *Options) *IndentHandler {
//     ^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/New().
//     documentation
//     > ```go
//     > func New(out Writer, opts *Options) *IndentHandler
//     > ```
//         ^^^ definition local 0
//             ^^ reference scip-go gomod github.com/golang/go/src go1.18 io/
//                ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 io/Writer#
//                        ^^^^ definition local 1
//                              ^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#
//                                        ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
   h := &IndentHandler{out: out, mu: &sync.Mutex{}}
// ^ definition local 2
//       ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                     ^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#out.
//                          ^^^ reference local 0
//                               ^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#mu.
//                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/
//                                         ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#
   if opts != nil {
//    ^^^^ reference local 1
    h.opts = *opts
//  ^ reference local 2
//    ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#opts.
//            ^^^^ reference local 1
   }
   if h.opts.Level == nil {
//    ^ reference local 2
//      ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#opts.
//           ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#Level.
    h.opts.Level = slog.LevelInfo
//  ^ reference local 2
//    ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#opts.
//         ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#Level.
//                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                      ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/LevelInfo.
   }
   return h
//        ^ reference local 2
  }
  
  func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
//      ^ definition local 3
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                        ^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#Enabled().
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
//                   ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#opts.
//                        ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/Options#Level.
//                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Leveler#Level.
  }
  
  // !+withs
  func (h *IndentHandler) WithGroup(name string) slog.Handler {
//      ^ definition local 6
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                        ^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#WithGroup().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).WithGroup(name string) Handler
//                        > ```
//                        documentation
//                        > !+withs
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#WithGroup. implementation
//                                  ^^^^ definition local 7
//                                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                    ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#
   if name == "" {
//    ^^^^ reference local 7
    return h
//         ^ reference local 6
   }
   return h.withGroupOrAttrs(groupOrAttrs{group: name})
//        ^ reference local 6
//          ^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#withGroupOrAttrs().
//                           ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#
//                                        ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#group.
//                                               ^^^^ reference local 7
  }
  
  func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
//      ^ definition local 8
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                        ^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#WithAttrs().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).WithAttrs(attrs []Attr) Handler
//                        > ```
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#WithAttrs. implementation
//                                  ^^^^^ definition local 9
//                                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
//                                                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                          ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#
   if len(attrs) == 0 {
//        ^^^^^ reference local 9
    return h
//         ^ reference local 8
   }
   return h.withGroupOrAttrs(groupOrAttrs{attrs: attrs})
//        ^ reference local 8
//          ^^^^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#withGroupOrAttrs().
//                           ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#
//                                        ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#attrs.
//                                               ^^^^^ reference local 9
  }
  
  //!-withs
  
  // !+withgora
  func (h *IndentHandler) withGroupOrAttrs(goa groupOrAttrs) *IndentHandler {
//      ^ definition local 10
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                        ^^^^^^^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#withGroupOrAttrs().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).withGroupOrAttrs(goa groupOrAttrs) *IndentHandler
//                        > ```
//                        documentation
//                        > !+withgora
//                                         ^^^ definition local 11
//                                             ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#
//                                                            ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
   h2 := *h
// ^^ definition local 12
//        ^ reference local 10
   h2.goas = make([]groupOrAttrs, len(h.goas)+1)
// ^^ reference local 12
//    ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
//                  ^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#
//                                    ^ reference local 10
//                                      ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
   copy(h2.goas, h.goas)
//      ^^ reference local 12
//         ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
//               ^ reference local 10
//                 ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
   h2.goas[len(h2.goas)-1] = goa
// ^^ reference local 12
//    ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
//             ^^ reference local 12
//                ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
//                           ^^^ reference local 11
   return &h2
//         ^^ reference local 12
  }
  
  //!-withgora
  
  // !+handle
  func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
//      ^ definition local 13
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                        ^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#Handle().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).Handle(ctx Context, r Record) error
//                        > ```
//                        documentation
//                        > !+handle
//                        relationship scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Handler#Handle. implementation
//                               ^^^ definition local 14
//                                   ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/
//                                           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 context/Context#
//                                                    ^ definition local 15
//                                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#
   buf := make([]byte, 0, 1024)
// ^^^ definition local 16
   if !r.Time.IsZero() {
//     ^ reference local 15
//       ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Time.
//            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/Time#IsZero().
    buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
//  ^^^ reference local 16
//        ^ reference local 13
//          ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                     ^^^ reference local 16
//                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                               ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Time().
//                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                         ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/TimeKey.
//                                                  ^ reference local 15
//                                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Time.
   }
   buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
// ^^^ reference local 16
//       ^ reference local 13
//         ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                    ^^^ reference local 16
//                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                              ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Any().
//                                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                       ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/LevelKey.
//                                                 ^ reference local 15
//                                                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Level.
   if r.PC != 0 {
//    ^ reference local 15
//      ^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#PC.
    fs := runtime.CallersFrames([]uintptr{r.PC})
//  ^^ definition local 17
//        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/
//                ^^^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/CallersFrames().
//                                        ^ reference local 15
//                                          ^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#PC.
    f, _ := fs.Next()
//  ^ definition local 18
//          ^^ reference local 17
//             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/Frames#Next().
    buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
//  ^^^ reference local 16
//        ^ reference local 13
//          ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                     ^^^ reference local 16
//                          ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                               ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/String().
//                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                           ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/SourceKey.
//                                                      ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//                                                          ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Sprintf().
//                                                                           ^ reference local 18
//                                                                             ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/Frame#File.
//                                                                                   ^ reference local 18
//                                                                                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 runtime/Frame#Line.
   }
   buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
// ^^^ reference local 16
//       ^ reference local 13
//         ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                    ^^^ reference local 16
//                         ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                              ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/String().
//                                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                          ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/MessageKey.
//                                                      ^ reference local 15
//                                                        ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Message.
   indentLevel := 0
// ^^^^^^^^^^^ definition local 19
   // Handle state from WithGroup and WithAttrs.
   goas := h.goas
// ^^^^ definition local 20
//         ^ reference local 13
//           ^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#goas.
   if r.NumAttrs() == 0 {
//    ^ reference local 15
//      ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#NumAttrs().
    // If the record has no Attrs, remove groups at the end of the list; they are empty.
    for len(goas) > 0 && goas[len(goas)-1].group != "" {
//          ^^^^ reference local 20
//                       ^^^^ reference local 20
//                                ^^^^ reference local 20
//                                         ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#group.
     goas = goas[:len(goas)-1]
//   ^^^^ reference local 20
//          ^^^^ reference local 20
//                    ^^^^ reference local 20
    }
   }
   for _, goa := range goas {
//        ^^^ definition local 21
//                     ^^^^ reference local 20
    if goa.group != "" {
//     ^^^ reference local 21
//         ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#group.
     buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", goa.group)
//   ^^^ reference local 16
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                     ^^^ reference local 16
//                                      ^^^^^^^^^^^ reference local 19
//                                                         ^^^ reference local 21
//                                                             ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#group.
     indentLevel++
//   ^^^^^^^^^^^ reference local 19
    } else {
     for _, a := range goa.attrs {
//          ^ definition local 22
//                     ^^^ reference local 21
//                         ^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/groupOrAttrs#attrs.
      buf = h.appendAttr(buf, a, indentLevel)
//    ^^^ reference local 16
//          ^ reference local 13
//            ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                       ^^^ reference local 16
//                            ^ reference local 22
//                               ^^^^^^^^^^^ reference local 19
     }
    }
   }
   r.Attrs(func(a slog.Attr) bool {
// ^ reference local 15
//   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Record#Attrs().
//              ^ definition local 23
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                     ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
    buf = h.appendAttr(buf, a, indentLevel)
//  ^^^ reference local 16
//        ^ reference local 13
//          ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                     ^^^ reference local 16
//                          ^ reference local 23
//                             ^^^^^^^^^^^ reference local 19
    return true
   })
   buf = append(buf, "---\n"...)
// ^^^ reference local 16
//              ^^^ reference local 16
   h.mu.Lock()
// ^ reference local 13
//   ^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#mu.
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#Lock().
   defer h.mu.Unlock()
//       ^ reference local 13
//         ^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#mu.
//            ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 sync/Mutex#Unlock().
   _, err := h.out.Write(buf)
//    ^^^ definition local 24
//           ^ reference local 13
//             ^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#out.
//                 ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 io/Writer#Write.
//                       ^^^ reference local 16
   return err
//        ^^^ reference local 24
  }
  
  //!-handle
  
  func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
//      ^ definition local 25
//         ^^^^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#
//                        ^^^^^^^^^^ definition scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                        documentation
//                        > ```go
//                        > func (*IndentHandler).appendAttr(buf []byte, a Attr, indentLevel int) []byte
//                        > ```
//                                   ^^^ definition local 26
//                                               ^ definition local 27
//                                                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                                                      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
//                                                            ^^^^^^^^^^^ definition local 28
   // Resolve the Attr's value before doing anything else.
   a.Value = a.Value.Resolve()
// ^ reference local 27
//   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//           ^ reference local 27
//             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                   ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Resolve().
   // Ignore empty Attrs.
   if a.Equal(slog.Attr{}) {
//    ^ reference local 27
//      ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Equal().
//            ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//                 ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#
    return buf
//         ^^^ reference local 26
   }
   // Indent 4 spaces per level.
   buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
// ^^^ reference local 26
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//           ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                   ^^^ reference local 26
//                               ^^^^^^^^^^^ reference local 28
   switch a.Value.Kind() {
//        ^ reference local 27
//          ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Kind().
   case slog.KindString:
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//           ^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/KindString.
    // Quote string values, to make them easy to parse.
    buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
//  ^^^ reference local 26
//        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                    ^^^ reference local 26
//                                     ^ reference local 27
//                                       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
//                                            ^ reference local 27
//                                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                                                    ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#String().
   case slog.KindTime:
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//           ^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/KindTime.
    // Write times in a standard way, without the monotonic time.
    buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
//  ^^^ reference local 26
//        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                    ^^^ reference local 26
//                                     ^ reference local 27
//                                       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
//                                            ^ reference local 27
//                                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                                                    ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Time().
//                                                           ^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/Time#Format().
//                                                                  ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/
//                                                                       ^^^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 time/RFC3339Nano.
   case slog.KindGroup:
//      ^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/
//           ^^^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/KindGroup.
    attrs := a.Value.Group()
//  ^^^^^ definition local 29
//           ^ reference local 27
//             ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
//                   ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Value#Group().
    // Ignore empty groups.
    if len(attrs) == 0 {
//         ^^^^^ reference local 29
     return buf
//          ^^^ reference local 26
    }
    // If the key is non-empty, write it out and indent the rest of the attrs.
    // Otherwise, inline the attrs.
    if a.Key != "" {
//     ^ reference local 27
//       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
     buf = fmt.Appendf(buf, "%s:\n", a.Key)
//   ^^^ reference local 26
//         ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//             ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                     ^^^ reference local 26
//                                   ^ reference local 27
//                                     ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
     indentLevel++
//   ^^^^^^^^^^^ reference local 28
    }
    for _, ga := range attrs {
//         ^^ definition local 30
//                     ^^^^^ reference local 29
     buf = h.appendAttr(buf, ga, indentLevel)
//   ^^^ reference local 26
//         ^ reference local 25
//           ^^^^^^^^^^ reference scip-go gomod golang.org/x/example 32022caedd6a `golang.org/x/example/slog-handler-guide/indenthandler2`/IndentHandler#appendAttr().
//                      ^^^ reference local 26
//                           ^^ reference local 30
//                               ^^^^^^^^^^^ reference local 28
    }
   default:
    buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
//  ^^^ reference local 26
//        ^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/
//            ^^^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 fmt/Appendf().
//                    ^^^ reference local 26
//                                     ^ reference local 27
//                                       ^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Key.
//                                            ^ reference local 27
//                                              ^^^^^ reference scip-go gomod github.com/golang/go/src go1.18 `log/slog`/Attr#Value.
   }
   return buf
//        ^^^ reference local 26
  }
  
