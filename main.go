package main

/*
#cgo LDFLAGS: -L . -lfoo
#include <stdio.h>
extern char* handleRequest(char* path);

extern void registerHandler(char* h());
*/
import "C"
import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"unsafe"
)

//export StartServer
func StartServer(port C.int) {
	httpServer := http.Server{
		Addr: ":" + strconv.Itoa(int(port)),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(C.GoString(C.handleRequest(C.CString(r.URL.Path)))))
		},
		),
	}

	httpServer.ListenAndServe()
	print("Server started")
	// wait for interrupt signal to gracefully shutdown the server with
	// Create a channel to listen for OS signals

	ch := make(chan os.Signal, 1)

	// relay incoming signals to channel
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	// Block main routine until a signal is received
	<-ch
}

//export InitEngine
func InitEngine(api unsafe.Pointer) {
}

func main() {
}
