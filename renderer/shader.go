package renderer

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/all-core/gl"
)

type Shader struct {
	program    uint32
	uniforms   map[string]int32
	attributes map[string]uint32
}

func createProgram(vertexShaderFile, fragmentShaderFile string, uniforms, attributes []string) (*Shader, error) {

	// Configure the vertex and fragment shaders
	vertexShaderSource, err := ioutil.ReadFile(vertexShaderFile)
	vertexShaderSource = append(vertexShaderSource, 0x00)
	fragmentShaderSource, err := ioutil.ReadFile(fragmentShaderFile)
	fragmentShaderSource = append(fragmentShaderSource, 0x00)

	vertexShader, err := compileShader(string(vertexShaderSource), gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(string(fragmentShaderSource), gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	gl.UseProgram(program)

	programUniforms := make(map[string]int32)
	for _, uniformName := range uniforms {
		if uniformName != "" {
			glstr := gl.Str(uniformName + "\x00")
			programUniforms[uniformName] = gl.GetUniformLocation(program, glstr)
		}
	}

	programAttributes := make(map[string]uint32)
	for _, attribName := range attributes {
		if attribName != "" {
			glstr := gl.Str(attribName + "\x00")
			programAttributes[attribName] = uint32(gl.GetAttribLocation(program, glstr))
		}
	}
	return &Shader{
		program:    program,
		uniforms:   programUniforms,
		attributes: programAttributes,
	}, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
