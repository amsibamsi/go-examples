package main

//#cgo pkg-config: OpenCL
//#include "opencl.h"
import "C"

import (
	"flag"
	"log"
)

var (
	platformInd = flag.Int("platform", 0, "OpenCL platform to use")
	deviceInd   = flag.Int("device", 0, "OpenCL device to use")
)

func main() {
	flag.Parse()
	var device C.cl_device_id
	if err := C.get_device(C.int(*platformInd), C.int(*deviceInd), &device); err != nil {
		log.Fatal(C.GoString(err))
	}
	var context C.cl_context
	if err := C.create_context(&device, &context); err != nil {
		log.Fatal(C.GoString(err))
	}
	var program C.cl_program
	if err := C.create_program(context, &device, &program); err != nil {
		log.Fatal(C.GoString(err))
	}
}
