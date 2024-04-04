# SCIP Indexing Example

This is an example of the output for [SCIP](https://sourcegraph.com/blog/announcing-scip) indexing using [SCIP-GO](https://github.com/sourcegraph/scip-go). I forked the [go-examples](https://github.com/golang/example) repo and ran scip-go over it.

## Step by step


- Ran `scip-go`:  This traversed the whole project and built the binary index output file `index.scip`


- Ran `scip snapshot`:  Using the [SCIP CLI](https://github.com/sourcegraph/scip) ran the snapshot function which generates a new version of the code (see [scip-snapshot](scip-snapshot) folder at the root) that contains all the annotations that the SCIP indexer produced.
![image](https://github.com/mmanela/example/assets/304410/7a39e7bf-1457-4bb1-acfb-e21905bcd49d)

  

- Ran `scip print --json index.scip | jq  > index.json`: Leveraging the SCIP CLI print command to make a json version of the index. You can view this in [index.json](index.json) as the root

![image](https://github.com/mmanela/example/assets/304410/443b1d5c-f593-4802-be7e-9e4f95560724)



------


# Go example projects

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
$ git clone https://go.googlesource.com/example
$ cd example
```
https://go.googlesource.com/example is the canonical Git repository.
It is mirrored at https://github.com/golang/example.

## [hello](hello/) and [hello/reverse](hello/reverse/)

```
$ cd hello
$ go build
$ ./hello -help
```
A trivial "Hello, world" program that uses a library package.

The [hello](hello/) command covers:

* The basic form of an executable command
* Importing packages (from the standard library and the local repository)
* Printing strings ([fmt](//golang.org/pkg/fmt/))
* Command-line flags ([flag](//golang.org/pkg/flag/))
* Logging ([log](//golang.org/pkg/log/))

The [reverse](hello/reverse/) reverse covers:

* The basic form of a library
* Conversion between string and []rune
* Table-driven unit tests ([testing](//golang.org/pkg/testing/))

## [helloserver](helloserver/)

```
$ cd helloserver
$ go run .
```

A trivial "Hello, world" web server.

Topics covered:

* Command-line flags ([flag](//golang.org/pkg/flag/))
* Logging ([log](//golang.org/pkg/log/))
* Web servers ([net/http](//golang.org/pkg/net/http/))

## [outyet](outyet/)

```
$ cd outyet
$ go run .
```
A web server that answers the question: "Is Go 1.x out yet?"

Topics covered:

* Command-line flags ([flag](//golang.org/pkg/flag/))
* Web servers ([net/http](//golang.org/pkg/net/http/))
* HTML Templates ([html/template](//golang.org/pkg/html/template/))
* Logging ([log](//golang.org/pkg/log/))
* Long-running background processes
* Synchronizing data access between goroutines ([sync](//golang.org/pkg/sync/))
* Exporting server state for monitoring ([expvar](//golang.org/pkg/expvar/))
* Unit and integration tests ([testing](//golang.org/pkg/testing/))
* Dependency injection
* Time ([time](//golang.org/pkg/time/))

## [appengine-hello](appengine-hello/)

A trivial "Hello, world" App Engine application intended to be used as the
starting point for your own code. Please see
[Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go)
and [Quickstart for Go in the App Engine Standard Environment](https://cloud.google.com/appengine/docs/standard/go/quickstart).

## [gotypes](gotypes/)

The `go/types` package is a type-checker for Go programs. It is one of the most
complex packages in Go's standard library, so we have provided this tutorial to
help you find your bearings. It comes with several example programs that you
can obtain using `go get` and play with as you learn to build tools that analyze
or manipulate Go programs.

## [template](template/)

A trivial web server that demonstrates the use of the
[`template` package](https://golang.org/pkg/text/template/)'s `block` feature.

## [slog-handler-guide](slog-handler-guide/)

The `log/slog` package supports structured logging.
It features a flexible backend in the form of a `Handler` interface.
This guide can help you write your own handler.
