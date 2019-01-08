package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	handler := &Handler{}
	svr := httptest.NewServer(handler)
	defer svr.Close()

	for _, i := range []int{1, 2} {
		resp, err := http.Get(svr.URL)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}

		want := fmt.Sprintf("hits %d", i)
		got, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if want != string(got) {
			t.Errorf("got %q; want %q\n", got, want)
		}
	}
}
