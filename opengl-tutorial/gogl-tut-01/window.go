package main

// #cgo linux,amd64 pkg-config: glew glfw3
// #include <GL/glew.h>
// #include <GLFW/glfw3.h>
import "C"

import (
	"fmt"
	"os"
)

func main() {
	if C.glfwInit() != C.GL_TRUE {
		fmt.Fprintln(os.Stderr, "Failed to initialize GLFW")
		os.Exit(-1)
	}
	C.glfwWindowHint(C.GLFW_SAMPLES, 4)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 3)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 3)
	C.glfwWindowHint(C.GLFW_OPENGL_FORWARD_COMPAT, C.GL_TRUE)
	C.glfwWindowHint(C.GLFW_OPENGL_PROFILE, C.GLFW_OPENGL_CORE_PROFILE)
	window := C.glfwCreateWindow(1024, 768, C.CString("Tutorial 01"), nil, nil)
	if window == nil {
		fmt.Fprintln(os.Stderr, "Failed to open GLFW window. If you have an Intel GPU, they are not 3.3 compatible. Try the 2.1 version of the tutorials.")
		C.glfwTerminate()
		os.Exit(-1)
	}
	C.glfwMakeContextCurrent(window)
	C.glewExperimental = C.GL_TRUE
	if C.glewInit() != C.GLEW_OK {
		fmt.Fprintln(os.Stderr, "Failed to initialize GLEW")
		os.Exit(-1)
	}
	C.glfwSetInputMode(window, C.GLFW_STICKY_KEYS, C.GL_TRUE)
	for C.glfwGetKey(window, C.GLFW_KEY_ESCAPE) != C.GLFW_PRESS && C.glfwWindowShouldClose(window) == 0 {
		C.glfwSwapBuffers(window)
		C.glfwPollEvents()
	}
}
