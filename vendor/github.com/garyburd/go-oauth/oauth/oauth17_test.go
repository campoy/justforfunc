// +build go1.7

package oauth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetContext_Cancel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c := Client{}
	_, err := c.GetContext(ctx, &Credentials{}, ts.URL, nil)
	if err == nil {
		t.Error("error should not be nil")
	}
}
