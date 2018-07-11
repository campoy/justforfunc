package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

var prPool = sync.Pool{
	New: func() interface{} { return new(pullRequest) },
}

type pullRequest struct {
	PullRequest struct{ ID int } `json:"pull_request"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	data := prPool.Get().(*pullRequest)
	defer prPool.Put(data)
	data.PullRequest.ID = 0

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		logrus.Errorf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "pull request id: %d", data.PullRequest.ID)
}
