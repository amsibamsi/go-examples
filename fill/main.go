package main

/*
#cgo pkg-config: glfw3 glesv2

#include <GLFW/glfw3.h>
#include <GLES2/gl2.h>
#include "main.h"
*/
import "C"

import (
	"fmt"
	"os"
)

//export glfwError
func glfwError(error C.int, description *C.char) {
	fmt.Fprintf(os.Stderr, "GLFW Error %d: %s\n", int(error), C.GoString(description))
}

func main() {
	C.setGlfwErrorCallback()
	if C.glfwInit() != C.GL_TRUE {
		fmt.Fprintln(os.Stderr, "Failed to initialize GLFW")
		os.Exit(-1)
	}
	C.glfwWindowHint(C.GLFW_CLIENT_API, C.GLFW_OPENGL_ES_API)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 2)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 0)
	window := C.glfwCreateWindow(1024, 768, C.CString("OpenGL ES Window"), nil, nil)
	if window == nil {
		fmt.Fprintln(os.Stderr, "Failed to open GLFW window")
		C.glfwTerminate()
		os.Exit(-1)
	}
	C.glfwMakeContextCurrent(window)
	C.glfwSetInputMode(window, C.GLFW_STICKY_KEYS, C.GL_TRUE)
	C.glClearColor(0.4, 0.5, 0.6, 0.0)
	for C.glfwGetKey(window, C.GLFW_KEY_Q) != C.GLFW_PRESS && C.glfwWindowShouldClose(window) == 0 {
		C.glClear(C.GL_COLOR_BUFFER_BIT)
		C.glfwSwapBuffers(window)
		C.glfwPollEvents()
	}
}
