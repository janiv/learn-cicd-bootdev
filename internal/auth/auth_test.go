package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	var headerWithMissingKey = http.Header{
		"Content-Type": {"application/json"},
	}
	var headerWithKey = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"ApiKey 007secretkey"},
	}
	cases := []struct {
		input http.Header
		key   string
		err   error
	}{
		{input: headerWithMissingKey, key: "", err: ErrNoAuthHeaderIncluded},
		{input: headerWithKey, key: "007secretkey", err: nil},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			output, outErr := GetAPIKey(c.input)
			if output != c.key || !errors.Is(c.err, outErr) {
				t.Errorf("test case %v failed, got %v:%v instead of %v:%v", i, output, outErr, c.key, c.err)
				os.Exit(1)
				return
			}
		})
	}

}
