package ace

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"io/ioutil"
)

type Shader struct {
	shaderId uint32
}

func NewVertexShader(filename string) *Shader {
	return newShader(filename, gl.VERTEX_SHADER)
}

func NewFragmentShader(filename string) *Shader {
	return newShader(filename, gl.FRAGMENT_SHADER)
}

func newShader(filename string, shaderType uint32) *Shader {
	src, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err.Error())
	}

	shader := &Shader{}

	shader.shaderId = gl.CreateShader(shaderType)

	glSrc, free := gl.Strs(string(src) + "\x00")
	gl.ShaderSource(shader.shaderId, 1, glSrc, nil)
	free()

	gl.CompileShader(shader.shaderId)

	var status int32
	gl.GetShaderiv(shader.shaderId, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		panic("There was a problem with compiling the shader")
	}

	return shader
}

func (s *Shader) Free() {
	gl.DeleteShader(s.shaderId);
}