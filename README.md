# go-pathmux

Package **pathmux** provides a path oriented ("middleware") HTTP handler, which can hand-off to other HTTP handler for each path,
for the Go programming language's built-in `"net/http"` library.

Package pathmux is a HTTP request router, and dispatcher.

The name mux stands for "HTTP request multiplexer".

Similar to the built-in standard http.ServeMux in the "net/http" package,
pathmux.Mux matches incoming requests against a list of registered routes
and calls a handler for the route that matches the URL or other conditions.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-pathmux

[![GoDoc](https://godoc.org/github.com/reiver/go-pathmux?status.svg)](https://godoc.org/github.com/reiver/go-pathmux)

## Example
```go
import "github.com/reiver/go-pathmux"

// ...

var mux pathmux.Mux

err := mux.HandlePath(homeHandler, "/")
err := mux.HandlePath(productsHandler, "/products")
err := mux.HandlePath(articlesHandler, "/articles")

err := mux.HandlePattern(productProducer, "/products/{key}")
err := mux.HandlePattern(articleCategoryProducer, "/articles/{category}/")
err := mux.HandlePattern(articleProducer, "/articles/{category}/{article_id}")

// handles: "/app", "/app/", "/app/apple", "/app/banana", "/app/banana/", "/app/banana/peel.jpeg",  "/app/cherry", etc.
err := mux.HandleDirectory(homeHandler, "/app")

// ...

err := http.ListenAndServe(":8080", mux)
```

## Import

To import package **pathmux** use `import` code like the follownig:

```
import "github.com/reiver/go-pathmux"
```

## Installation

To install package **pathmux** do the following:
```
GOPROXY=direct go get github.com/reiver/go-pathmux
```

## Author

Package **pathmux** was written by [Charles Iliya Krempeaux](http://reiver.link)
