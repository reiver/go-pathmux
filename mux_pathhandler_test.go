package pathmux_test

import (
	"github.com/reiver/go-pathmux"

	"net/http"

	"testing"
)

func TestMuxPathHandler(t *testing.T) {

	tests := []struct{
		Paths []string
	}{
		{
			Paths: []string{
				"/v1/help",
				"/upgrade",
				"/v1/users",
				"/v1/users/123",
				"/v2/users",
			},
		},

		{
			Paths: []string{
				"/about",
				"/contacts",
				"/faq",
				"/location",
				"/help",
			},
		},

		{
			Paths: []string{
				"/apple",
				"/banana",
				"/cherry",
				"/date",
				"/apple/banana/cherry",
			},
		},

		{
			Paths: []string{
				"/",
				"/index.html",
				"/resume",
				"/resume.php",
				"/robots.txt",
				"/humans.txt",
			},
		},
	}

	for testNumber, test := range tests {

		var mux pathmux.Mux

		for pathNumber, path := range test.Paths {
			var handler http.Handler = nopHandler{
				Name: path,
			}

			if err := mux.HandlePath(handler, path); nil != err {
				t.Errorf("For test #%d and path #%d, did not expect an error, but actually got one: (%T) %q", testNumber, pathNumber, err, err)
				continue
			}
		}

		for pathNumber, path := range test.Paths {
			var handler http.Handler = mux.PathHandler(path)
			if nil == handler {
				t.Errorf("For test #%d and path #%d, did not expect nil handler, but actually got it: %#v", testNumber, pathNumber, handler)
				continue
			}

			nop, casted := handler.(nopHandler)
			if !casted {
				t.Errorf("For test #%d and path #%d, expected cast, but did not actually: %t", testNumber, pathNumber, casted)
				continue
			}

			if expected, actual := path, nop.Name; expected != actual {
				t.Errorf("For test #%d and path #%d ...", testNumber, pathNumber)
				t.Errorf("\tEXPECTED: %q", expected)
				t.Errorf("\tACTUAL:   %q", actual)
				continue
			}
		}

		{
			var handler http.Handler = mux.PathHandler("/does/not/exist")
			if nil != handler {
				t.Errorf("For test #%d, expected nil handler, but did not actually get that: %#v", testNumber, handler)
				continue
			}
		}

		{
			var handler http.Handler = mux.PathHandler("/v3/does-not-exist/a/b/c/d/e/f/g.html")
			if nil != handler {
				t.Errorf("For test #%d, expected nil handler, but did not actually get that: %#v", testNumber, handler)
				continue
			}
		}
	}
}
