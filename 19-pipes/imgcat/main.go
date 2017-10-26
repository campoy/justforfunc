// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/campoy/justforfunc/19-pipes/imgcat/imgcat"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}
}

type badWriter struct{}

func (badWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("bad writer")
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "could not open image")
	}
	defer f.Close()

	wc := imgcat.NewWriter(os.Stdout)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	return wc.Close()
}
