package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestHandle(t *testing.T) {
	pl, err := ioutil.ReadFile("testdata/payload.json")
	if err != nil {
		t.Fatalf("could not read payload.json: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(pl))
	if err != nil {
		t.Fatalf("could not create test request: %v", err)
	}
	rec := httptest.NewRecorder()
	handle(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code %s", res.Status)
	}
	defer res.Body.Close()

	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read result payload: %v", err)
	}

	if exp := "pull request id: 191568743"; string(msg) != exp {
		t.Fatalf("expected message %q; got %q", exp, msg)
	}
}

func BenchmarkHandle(b *testing.B) {
	b.StopTimer()

	pl, err := ioutil.ReadFile("testdata/payload.json")
	if err != nil {
		b.Fatalf("could not read payload.json: %v", err)
	}

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(pl))
		if err != nil {
			b.Fatalf("could not create test request: %v", err)
		}
		rec := httptest.NewRecorder()

		b.StartTimer()
		handle(rec, req)
		b.StopTimer()
	}
}
