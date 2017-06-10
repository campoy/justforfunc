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

package flite

// #cgo CFLAGS: -I /usr/include/flite/
// #cgo LDFLAGS: -lflite -lflite_cmu_us_kal
// #include "flite.h"
// cst_voice* register_cmu_us_kal(const char *voxdir);
import "C"

import (
	"fmt"
	"unsafe"
)

var voice *C.cst_voice

func init() {
	C.flite_init()
	voice = C.register_cmu_us_kal(nil)
}

func TextToSpeech(path, text string) error {
	if voice == nil {
		return fmt.Errorf("could not find default voice")
	}

	ctext := C.CString(text)
	cout := C.CString(path)
	C.flite_text_to_speech(ctext, voice, cout)
	C.free(unsafe.Pointer(ctext))
	C.free(unsafe.Pointer(cout))
	return nil
}
