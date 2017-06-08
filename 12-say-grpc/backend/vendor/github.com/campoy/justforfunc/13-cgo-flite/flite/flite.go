// Copyright 2107 Google Inc. All rights reserved.
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

// +build linux

// Package flite provides a cgo wrapping over the flite library available at
// http://www.festvox.org/flite/.
// It simply provides a wrapper over the text_to_speech function, and a helper
// that returns a slice of bytes instead.
package flite

// #cgo CFLAGS: -I /usr/include/flite
// #cgo LDFLAGS: -lflite -lflite_cmu_us_kal -lflite_usenglish
// #include <flite.h>
// #include <stdlib.h>
// cst_voice *register_cmu_us_kal(const char *voxdir);
import "C"

import (
	"fmt"
	"io/ioutil"
	"unsafe"
)

var voice *C.cst_voice

func init() {
	C.flite_init()
	voice = C.register_cmu_us_kal(nil)
}

// TextToSpeechFile synthesizes the text in the given string and writes it to
// a file in the given path.
func TextToSpeechFile(path, text string) error {
	if voice == nil {
		return fmt.Errorf("could not find default voice")
	}

	ctext := C.CString(text)
	cpath := C.CString(path)
	C.flite_text_to_speech(ctext, voice, cpath)
	C.free(unsafe.Pointer(ctext))
	C.free(unsafe.Pointer(cpath))
	return nil
}

// TextToSpeech synthesizes the text in the string and returns a slice of bytes
// containing the generated output.
func TextToSpeechBytes(text string) ([]byte, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	f.Close()

	if err := TextToSpeechFile(f.Name(), text); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}
	return data, nil
}
