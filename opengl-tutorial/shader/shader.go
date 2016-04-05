package shader

// #cgo freebsd,amd64 CFLAGS: -I/usr/local/include
// #cgo freebsd,amd64 LDFLAGS: -L/usr/local/lib -lGLEW
// #cgo linux,amd64 pkg-config: glew
//
// #include <GL/glew.h>
//
// GLuint _glCreateShader(GLenum shaderType) {
//   return glCreateShader(shaderType);
// }
//
// void _glShaderSource(GLuint shader, GLsizei count, const GLchar **string, const GLint *length) {
//   glShaderSource(shader, count, string, length);
// }
//
// void _glCompileShader(GLuint shader) {
//   glCompileShader(shader);
// }
//
// GLuint _glCreateProgram() {
//   return glCreateProgram();
// }
//
// void _glAttachShader(GLuint program, GLuint shader) {
//   glAttachShader(program, shader);
// }
//
// void _glLinkProgram(GLuint program) {
//   glLinkProgram(program);
// }
//
// void _glDeleteShader(GLuint shader) {
//   glDeleteShader(shader);
// }
//
// void _glGetShaderiv(GLuint shader, GLenum pname, GLint *params) {
//   glGetShaderiv(shader, pname, params);
// }
//
// void _glGetShaderInfoLog(GLuint shader, GLsizei maxLength, GLsizei *length, GLchar *infoLog) {
//   glGetShaderInfoLog(shader, maxLength, length, infoLog);
// }
//
// void _glGetProgramiv(GLuint program, GLenum pname, GLint *params) {
//   glGetProgramiv(program, pname, params);
// }
//
// void _glGetProgramInfoLog(GLuint program, GLsizei maxLength, GLsizei *length, GLchar *programLog) {
//   glGetProgramInfoLog(program, maxLength, length, programLog);
// }
//
import "C"

import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

func Load(vertex_file_path string, fragment_file_path string) (ProgramID C.GLuint, error error) {

	VertexShaderID := C._glCreateShader(C.GL_VERTEX_SHADER)
	FragmentShaderID := C._glCreateShader(C.GL_FRAGMENT_SHADER)

	VertexShaderString, error := ioutil.ReadFile(vertex_file_path)
	if error != nil {
		return
	}
	VertexShaderCode := (*C.GLchar)(unsafe.Pointer(&VertexShaderString[0]))

	FragmentShaderString, error := ioutil.ReadFile(fragment_file_path)
	if error != nil {
		return
	}
	FragmentShaderCode := (*C.GLchar)(unsafe.Pointer(&FragmentShaderString[0]))

	Result := (C.GLint)(C.GL_FALSE)
	var InfoLogLength int

	fmt.Println("Compiling shader : " + vertex_file_path)
	C._glShaderSource(VertexShaderID, 1, &VertexShaderCode, nil)
	C._glCompileShader(VertexShaderID)

	C._glGetShaderiv(VertexShaderID, C.GL_COMPILE_STATUS, &Result)
	C._glGetShaderiv(VertexShaderID, C.GL_INFO_LOG_LENGTH, (*C.GLint)(unsafe.Pointer(&InfoLogLength)))
	if InfoLogLength > 0 {
		VertexShaderErrorMessage := make([]byte, InfoLogLength+1)
		C._glGetShaderInfoLog(VertexShaderID, C.GLsizei(InfoLogLength), nil, (*C.GLchar)(unsafe.Pointer(&VertexShaderErrorMessage[0])))
		fmt.Fprintln(os.Stderr, C.GoString((*C.char)(unsafe.Pointer(&VertexShaderErrorMessage[0]))))
	}

	fmt.Println("Compiling shader : " + fragment_file_path)
	C._glShaderSource(FragmentShaderID, 1, &FragmentShaderCode, nil)
	C._glCompileShader(FragmentShaderID)

	C._glGetShaderiv(FragmentShaderID, C.GL_COMPILE_STATUS, &Result)
	C._glGetShaderiv(FragmentShaderID, C.GL_INFO_LOG_LENGTH, (*C.GLint)(unsafe.Pointer(&InfoLogLength)))
	if InfoLogLength > 0 {
		FragmentShaderErrorMessage := make([]byte, InfoLogLength+1)
		C._glGetShaderInfoLog(FragmentShaderID, C.GLsizei(InfoLogLength), nil, (*C.GLchar)(unsafe.Pointer(&FragmentShaderErrorMessage[0])))
		fmt.Fprintln(os.Stderr, C.GoString((*C.char)(unsafe.Pointer(&FragmentShaderErrorMessage[0]))))
	}

	fmt.Println("Linking program")
	ProgramID = C._glCreateProgram()
	C._glAttachShader(ProgramID, VertexShaderID)
	C._glAttachShader(ProgramID, FragmentShaderID)
	C._glLinkProgram(ProgramID)

	C._glGetProgramiv(ProgramID, C.GL_LINK_STATUS, &Result)
	C._glGetProgramiv(ProgramID, C.GL_INFO_LOG_LENGTH, (*C.GLint)(unsafe.Pointer(&InfoLogLength)))
	if InfoLogLength > 0 {
		var ProgramErrorMessage [1000]byte
		C._glGetProgramInfoLog(ProgramID, C.GLsizei(InfoLogLength), nil, (*C.GLchar)(unsafe.Pointer(&ProgramErrorMessage[0])))
		fmt.Fprintln(os.Stderr, C.GoString((*C.char)(unsafe.Pointer(&ProgramErrorMessage[0]))))
	}

	C._glDeleteShader(VertexShaderID)
	C._glDeleteShader(FragmentShaderID)

	return

}
