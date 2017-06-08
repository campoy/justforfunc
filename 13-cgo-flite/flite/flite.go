package flite

// #cgo CFLAGS: -I /usr/include/flite
// #include <flite.h>
// #cgo LDFLAGS: -L /usr/lib/x86_64-linux-gnu -lflite -lflite_cmu_us_kal -lflite_usenglish
// extern cst_voice *register_cmu_us_kal(const char *voxdir);
import "C"
import (
	"fmt"
	"io/ioutil"
)

var voice *C.cst_voice

func init() {
	C.flite_init()
	voice := C.register_cmu_us_kal(nil)
}

func TextToSpeechFile(path, text string) error {
	if voice == nil {
		return fmt.Errorf("could not find default voice")
	}

	ctext := C.CString(text)
	cpath := C.CString(path)
	C.flite_text_to_speech(ctext, voice, cpath)
	C.free(ctext)
	C.free(c.path)
	return nil
}

// TextToSpeech synthesizes the text in a string and returns a slice of bytes
// containing the generated output.
func TextToSpeechBytes(text string) ([]byte, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	if err := TextToSpeechFile(f.Name(), text); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}
	return data, nil	
}
