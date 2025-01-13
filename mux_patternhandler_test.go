package pathmux_test

import (
	"github.com/reiver/go-pathmux"

	"net/http"

	"testing"
)

func TestMuxPatternHandler(t *testing.T) {

	tests := []struct{
		Patterns []string
		PathMatches   []string
		PathNoMatches []string
	}{
		{
			Patterns: []string{
				"/v1/users/{user_id}",
				"/v1/users/{user_id}/wallets/{wallet_id}",
				"/v1/merchants/{merchant_id}",
			},
			PathMatches: []string{
				"/v1/users/123",
				"/v1/users/456789",
				"/v1/users/joeblow",
				"/v1/users/123/wallets/45678",
				"/v1/users/123/wallets/bitcoin",
				"/v1/merchants/789123",
				"/v1/merchants/acme",
			},
			PathNoMatches: []string{
				"/v1/users/",
				"/v1/users/123/wallets/",
				"/v1/merchants/",

				"/v1/users.html",
				"/v1/users/123/wallets.html",
				"/v1/merchants.html",

				"/v1/users/123/",
				"/v1/users/456789/",
				"/v1/users/joeblow/",
				"/v1/users/123/wallets/45678/",
				"/v1/users/123/wallets/bitcoin/",
				"/v1/merchants/789123/",
				"/v1/merchants/acme/",

				"/v2/users/123",
				"/v2/users/456789",
				"/v2/users/joeblow",
				"/v2/users/123/wallets/45678",
				"/v2/users/123/wallets/bitcoin",
				"/v2/merchants/789123",
				"/v2/merchants/acme",

				"/apple",
				"/apple/",
				"/apple.html",
				"/apple/banana",
				"/apple/banana/",
				"/apple/banana.html",
				"/apple/banana/cherry",
				"/apple/banana/cherry/",
				"/apple/banana/cherry.html",

				"//v1/users/123",
				"//v1/users/456789",
				"//v1/users/joeblow",
				"//v1/users/123/wallets/45678",
				"//v1/users/123/wallets/bitcoin",
				"//v1/merchants/789123",
				"//v1/merchants/acme",

				"/v1//users/123",
				"/v1//users/456789",
				"/v1//users/joeblow",
				"/v1//users/123/wallets/45678",
				"/v1//users/123/wallets/bitcoin",
				"/v1//merchants/789123",
				"/v1//merchants/acme",

				"/v1/users//123",
				"/v1/users//456789",
				"/v1/users//joeblow",
				"/v1/users//123/wallets/45678",
				"/v1/users//123/wallets/bitcoin",
				"/v1/merchants//789123",
				"/v1/merchants//acme",

				"/v1/users/123/wallets//45678",
				"/v1/users/123/wallets//bitcoin",
			},
		},



		{
			Patterns: []string{
				"/account={account_id}",
				"/account={account_id}/user={user_id}",
				"/company={company_id}",
				"/files={name}/",
			},
			PathMatches: []string{
				"/account=123",
				"/account=comp",
				"/account=123/user=45678",
				"/account=comp/user=joeblow",
				"/company=248",
				"/company=acme",
				"/files=stuff/",
			},
			PathNoMatches: []string{
				"/account=123/",
				"/account=comp/",
				"/account=123/user=45678/",
				"/account=comp/user=joeblow/",
				"/company=248/",
				"/company=acme/",
				"/files=stuff",

				"/apple",
				"/apple/",
				"/apple.html",
				"/apple/banana",
				"/apple/banana/",
				"/apple/banana.html",
				"/apple/banana/cherry",
				"/apple/banana/cherry/",
				"/apple/banana/cherry.html",
			},
		},
	}

	for testNumber, test := range tests {

		var mux pathmux.Mux

		for patternNumber, pattern := range test.Patterns {
			var producer pathmux.Producer = nopProducer{
				Name: pattern,
			}

			if err := mux.HandlePatternUsingProducer(producer, pattern); nil != err {
				t.Errorf("For test #%d and pattern #%d, did not expect an error, but actually got one: (%T) %q", testNumber, patternNumber, err, err)
				continue
			}
		}

		for pathNumber, path := range test.PathMatches {
			var handler http.Handler = mux.Handler(path)
			if nil == handler {
				t.Errorf("For test #%d and matched path #%d, did not expect nil handler, but actually got it: %#v", testNumber, pathNumber, handler)
				t.Errorf("\tPATH: %q", path)
				continue
			}

//			nop, casted := handler.(nopHandler)
//			if !casted {
//				t.Errorf("For test #%d and matched path #%d, expected cast, but did not actually: %t", testNumber, matchedPathNumber, casted)
//				t.Errorf("\tPATH: %q", path)
//				continue
//			}
		}

		for pathNumber, path := range test.PathNoMatches {
			var handler http.Handler = mux.Handler(path)
			if nil != handler {
				t.Errorf("For test #%d and no-match path #%d, expected nil handler, but did not actually get one: %#v", testNumber, pathNumber, handler)
				t.Errorf("\tPATH: %q", path)
				t.Errorf("\tPATTERN: ...")
				for _, pattern := range test.Patterns {
					t.Errorf("\t\t%q", pattern)
				}
				continue
			}

//			nop, casted := handler.(nopHandler)
//			if !casted {
//				t.Errorf("For test #%d and matched path #%d, expected cast, but did not actually: %t", testNumber, matchedPathNumber, casted)
//				t.Errorf("\tPATH: %q", path)
//				continue
//			}
		}






	}
}
