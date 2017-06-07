package main

// #cgo CFLAGS: -I /usr/include/flite
// #include <flite.h>
// #cgo LDFLAGS: -L /usr/lib/x86_64-linux-gnu -lflite -lflite_cmu_us_kal
// extern cst_voice *register_cmu_us_kal(const char *voxdir);
import "C"
import "log"

func main() {
	C.flite_init()
	voice := C.register_cmu_us_kal(nil)
	if voice == nil {
		log.Fatal("could not find voice awb")
	}

	C.flite_text_to_speech(C.CString("hello, Cameron!"), voice, C.CString("out.wav"))
}
