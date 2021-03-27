package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"log"
	"strings"
)

type Shader struct {
	programId uint32
}

func (it *Shader) Bind()  {
	gl.UseProgram(it.programId)
}

func (it *Shader) Unbind()  {
	gl.UseProgram(0)
}

func (it *Shader) Projection(matrix mgl32.Mat4)  {
	it.UniformMatrix("u_projection", matrix)
}

func (it *Shader) Uniform1i(name string, value int32) {
	gl.Uniform1i(it.GetLocation(name), value)
}

func (it *Shader) Uniform1fv(name string, value []float32) {
	gl.Uniform1fv(it.GetLocation(name), 12, &value[0])
}

func (it *Shader) UniformMatrix(name string, matrix mgl32.Mat4) {
	gl.UniformMatrix4fv(it.GetLocation(name), 1, false, &matrix[0])
}

func (it *Shader) GetLocation(name string) int32 {
	return gl.GetUniformLocation(it.programId, gl.Str(name + "\x00"))
}

func newShader(shaderName string) *Shader {
	vertexId := loadShader("resources/shaders/" + shaderName + ".vert", gl.VERTEX_SHADER)
	fragmentId := loadShader("resources/shaders/" + shaderName + ".frag", gl.FRAGMENT_SHADER)
	shader := Shader { gl.CreateProgram() }
	gl.AttachShader(shader.programId, vertexId)
	gl.AttachShader(shader.programId, fragmentId)
	gl.LinkProgram(shader.programId)
	gl.ValidateProgram(shader.programId)
	return &shader
}

func loadShader(filename string, shaderType uint32) uint32 {
	shaderId := gl.CreateShader(shaderType)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Couldn't load a shader {}", filename)
	}

	shaderSource, free := gl.Strs(string(content) + "\x00")
	gl.ShaderSource(shaderId, 1, shaderSource, nil)
	free()
	gl.CompileShader(shaderId)

	var result int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &result)
	if result == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)

		shaderLog := strings.Repeat("\x00", int(logLength + 1))
		gl.GetShaderInfoLog(shaderId, logLength, nil, gl.Str(shaderLog))
		log.Fatal("Shader error: ", shaderLog)
	}

	return shaderId
}
