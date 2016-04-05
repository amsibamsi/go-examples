// Create a random texture and draw it with OpenGL as background. Use an orthographic projection and draw a 2D texture onto a quad.
package main

/*
#cgo pkg-config: glew glfw3

#include <GL/glew.h>
#include <GLFW/glfw3.h>
#include "main.h"
*/
import "C"

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"unsafe"
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
	C.glfwWindowHint(C.GLFW_CLIENT_API, C.GLFW_OPENGL_API)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 2)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 1)
	window := C.glfwCreateWindow(1000, 1000, C.CString("Bitmap"), nil, nil)
	if window == nil {
		fmt.Fprintln(os.Stderr, "Failed to open GLFW window")
		C.glfwTerminate()
		os.Exit(-1)
	}
	C.glfwMakeContextCurrent(window)
	C.glewExperimental = C.GL_TRUE
	if C.glewInit() != C.GLEW_OK {
		fmt.Fprintln(os.Stderr, "Failed to initialize GLEW")
		os.Exit(-1)
	}
	C.glClearColor(0.0, 0.0, 0.0, 0.0)
	var texture C.GLuint
	rand.Seed(time.Now().UTC().UnixNano())
	textureData := make([]byte, 3*1e6)
	for i := 0; i < len(textureData); i++ {
		textureData[i] = byte(rand.Intn(255))
	}
	C.glPixelStorei(C.GL_UNPACK_ALIGNMENT, 1)
	C.glPixelStorei(C.GL_PACK_ALIGNMENT, 1)
	C.glGenTextures(1, &texture)
	C.glBindTexture(C.GL_TEXTURE_2D, texture)
	C.glTexImage2D(C.GL_TEXTURE_2D, 0, C.GL_RGB8, 1000, 1000, 0, C.GL_RGB, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&textureData[0]))
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MIN_FILTER, C.GL_NEAREST)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MAG_FILTER, C.GL_NEAREST)
	C.glfwSetInputMode(window, C.GLFW_STICKY_KEYS, C.GL_TRUE)
	for C.glfwGetKey(window, C.GLFW_KEY_Q) != C.GLFW_PRESS && C.glfwWindowShouldClose(window) == 0 {
		C.glClear(C.GL_COLOR_BUFFER_BIT)
		C.glMatrixMode(C.GL_PROJECTION)
		C.glLoadIdentity()
		C.glOrtho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)
		C.glMatrixMode(C.GL_MODELVIEW)
		C.glLoadIdentity()
		C.glEnable(C.GL_TEXTURE_2D)
		C.glBindTexture(C.GL_TEXTURE_2D, texture)
		C.glBegin(C.GL_QUADS)
		C.glTexCoord2f(0.0, 0.0)
		C.glVertex2f(-1.0, -1.0)
		C.glTexCoord2f(1.0, 0.0)
		C.glVertex2f(1.0, -1.0)
		C.glTexCoord2f(0.0, 1.0)
		C.glVertex2f(1.0, 1.0)
		C.glTexCoord2f(1.0, 1.0)
		C.glVertex2f(-1.0, 1.0)
		C.glEnd()
		C.glfwSwapBuffers(window)
		C.glfwPollEvents()
	}
}
