package dlapi

// #include "stdint.h"
// #include "include/dart_api_dl.h"
import "C"
import "unsafe"

func InitEngine(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		panic("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}
