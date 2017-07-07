/*
	Copyright 2017, Google, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

// Command logpipe is a service that will let you pipe logs directly to Stackdriver Logging.
package main

import (
	"bufio"
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"

	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	var opts struct {
		ProjectID string `short:"p" long:"project" description:"Google Cloud Platform Project ID" required:"true"`
		LogName   string `short:"l" long:"logname" description:"The name of the log to write to" default:"default"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(2)
	}

	projectID := &opts.ProjectID
	logName := &opts.LogName

	// Check if Standard In is coming from a pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		errorf("Could not stat standard input: %v\n", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		errorf("Nothing is piped in so there is nothing to log!\n")
	}

	// Creates a client.
	client, err := logging.NewClient(ctx, *projectID)
	if err != nil {
		errorf("Failed to create client: %v", err)
	}

	// Selects the log to write to.
	logger := client.Logger(*logName)

	// Read from Stdin and log it to Stdout and Stackdriver
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		fmt.Println(text)
		logger.Log(logging.Entry{Payload: text})
	}

	// Closes the client and flushes the buffer to the Stackdriver Logging
	// service.
	if err := client.Close(); err != nil {
		errorf("Failed to close client: %v", err)
	}

	if err := s.Err(); err != nil {
		errorf("Failed to scan input: %v", err)
	}
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(2)
}
