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

err := mux.HandlePattern(productProducer, "/products/{key}")
err := mux.HandlePattern(articleCategoryProducer, "/articles/{category}/")
err := mux.HandlePattern(articleProducer, "/articles/{category}/{article_id}")

err := mux.HandlePath(productsHandler, "/products")
err := mux.HandlePath(articlesHandler, "/articles")
err := mux.HandlePath(homeHandler, "/")

// ...

err := http.ListenAndServe(":8080", mux)
```
