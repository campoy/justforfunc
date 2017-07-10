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
	"log"
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

	flags.Parse(&opts)

	projectID := &opts.ProjectID
	logName := &opts.LogName

	if *projectID == "" {
		fmt.Printf("Please specify a project ID\n")
		return
	}

	// Check if Standard In is coming from a pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Printf("Nothing is piped in so there is nothing to log!\n")
		return
	}

	// Creates a client.
	client, err := logging.NewClient(ctx, *projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Selects the log to write to.
	logger := client.Logger(*logName)

	// Read from Stdin and log it to Stdout and Stackdriver
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		logger.Log(logging.Entry{Payload: text})
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan input: %v", err)
	}

	// Closes the client and flushes the buffer to the Stackdriver Logging
	// service.
	if err := client.Close(); err != nil {
		log.Fatalf("Failed to close client: %v", err)
	}

	fmt.Printf("Finished logging\n")
}
