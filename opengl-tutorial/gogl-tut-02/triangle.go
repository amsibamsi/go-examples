package main

// #cgo freebsd,amd64 CFLAGS: -I/usr/local/include
// #cgo freebsd,amd64 LDFLAGS: -L/usr/local/lib -lX11 -lXrandr -lXxf86vm -lXi -lGL -lm -lglfw3 -lGLEW
// #cgo linux,amd64 pkg-config: glew glfw3
// #include <GL/glew.h>
// #include <GLFW/glfw3.h>
//
// void _glGenVertexArrays(GLsizei n, GLuint* arrays) {
//   glGenVertexArrays(n, arrays);
// }
//
// void _glBindVertexArray(GLuint array) {
//   glBindVertexArray(array);
// }
//
// void _glGenBuffers(GLsizei n, GLuint *buffers) {
//   glGenBuffers(n, buffers);
// }
//
// void _glBindBuffer(GLenum target, GLuint buffer) {
//   glBindBuffer(target, buffer);
// }
//
// void _glBufferData(GLenum target, GLsizeiptr size, const GLvoid * data, GLenum usage) {
//   glBufferData(target, size, data, usage);
// }
//
// void _glUseProgram(GLuint program) {
//   glUseProgram(program);
// }
//
// void _glEnableVertexAttribArray(GLuint index) {
//   glEnableVertexAttribArray(index);
// }
//
// void _glVertexAttribPointer(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const GLvoid * pointer) {
//   glVertexAttribPointer(index, size, type, normalized, stride, pointer);
// }
//
// void _glDisableVertexAttribArray(GLuint index) {
//   glDisableVertexAttribArray(index);
// }
//
// void _glDeleteBuffers(GLsizei n, const GLuint * buffers) {
//   glDeleteBuffers(n, buffers);
// }
//
// void _glDeleteVertexArrays(GLsizei n, const GLuint *arrays) {
//   glDeleteVertexArrays(n, arrays);
// }
//
// void _glDeleteProgram(GLuint program) {
//   glDeleteProgram(program);
// }
//
import "C"

import (
	"fmt"
	"github.com/amsibamsi/go-opengl-tutorial/shader"
	"os"
	"unsafe"
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
	window := C.glfwCreateWindow(1024, 768, C.CString("Tutorial 02 - Red Triangle"), nil, nil)
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

	C.glClearColor(0.0, 0.0, 0.4, 0.0)

	var VertexArrayID C.GLuint
	C._glGenVertexArrays(1, &VertexArrayID)
	C._glBindVertexArray(VertexArrayID)

	programID, err := shader.Load("SimpleVertexShader.vertexshader", "SimpleFragmentShader.fragmentshader")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load shaders: error: "+err.Error())
		os.Exit(-1)
	}

	g_vertex_buffer_data := []float32{-1.0, -1.0, 0.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0}
	var vertexbuffer C.GLuint
	C._glGenBuffers(1, &vertexbuffer)
	C._glBindBuffer(C.GL_ARRAY_BUFFER, vertexbuffer)
	C._glBufferData(C.GL_ARRAY_BUFFER, C.GLsizeiptr(int(unsafe.Sizeof(g_vertex_buffer_data))*len(g_vertex_buffer_data)), unsafe.Pointer(&g_vertex_buffer_data[0]), C.GL_STATIC_DRAW)

	for C.glfwGetKey(window, C.GLFW_KEY_ESCAPE) != C.GLFW_PRESS && C.glfwWindowShouldClose(window) == 0 {

		C.glClear(C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)

		C._glUseProgram(C.GLuint(programID))

		C._glEnableVertexAttribArray(0)
		C._glBindBuffer(C.GL_ARRAY_BUFFER, vertexbuffer)
		C._glVertexAttribPointer(0, 3, C.GL_FLOAT, C.GL_FALSE, 0, nil)

		C.glDrawArrays(C.GL_TRIANGLES, 0, 3)

		C._glDisableVertexAttribArray(0)

		C.glfwSwapBuffers(window)
		C.glfwPollEvents()
	}

	C._glDeleteBuffers(1, &vertexbuffer)
	C._glDeleteVertexArrays(1, &VertexArrayID)
	C._glDeleteProgram(C.GLuint(programID))

	C.glfwTerminate()

}
