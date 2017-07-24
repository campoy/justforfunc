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

package sum

import (
	"testing"
)

func TestInts(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if s != tc.sum {
				t.Fatalf("sum of %v should be %v; got %v", tc.name, tc.sum, s)
			}
		})
	}
}
