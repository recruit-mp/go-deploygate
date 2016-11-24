package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

// testClient is the test client.
var testClient = DefaultClient()

func record(t *testing.T, fixture string, f func(*Client)) {
	r, err := recorder.New("fixtures/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := r.Stop(); err != nil {
			t.Fatal(err)
		}
	}()

	client := DefaultClient()
	client.HTTPClient.Transport = r

	f(client)
}
