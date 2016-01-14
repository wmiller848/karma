package renderer

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Bottom
	-1.0, -1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	// Top
	-1.0, 1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 1.0,

	// Left
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,

	// Right
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
}

type Shaders struct {
	textureFlat *Shader
}

type Renderer struct {
	Window     *glfw.Window
	Shaders    *Shaders
	Projection mgl32.Mat4
	Camera     mgl32.Mat4
	ModelView  mgl32.Mat4
	Fps        int
}

func CreateRenderer(windowWidth, windowHeight int) *Renderer {

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	// defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Karma", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	shaders := Shaders{}

	textureFlatUniforms := make([]string, 1)
	textureFlatUniforms = append(textureFlatUniforms, "projection", "camera", "modelView", "tex")
	textureFlatAttributes := make([]string, 1)
	textureFlatAttributes = append(textureFlatAttributes, "vert", "vertTexCoord")

	fmt.Println("%+v", textureFlatUniforms)
	fmt.Println("%+v", textureFlatAttributes)

	shader, err := createProgram("./assets/shaders/texture_flat.vs", "./assets/shaders/texture_flat.fs", textureFlatUniforms, textureFlatAttributes)
	if err != nil {
		panic(err)
	}
	shaders.textureFlat = shader

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/float32(windowHeight), 0.1, 10.0)
	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	modelView := mgl32.Ident4()

	program := shader.program
	gl.UseProgram(program)

	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/float32(windowHeight), 0.1, 10.0)
	// projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	// gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])
	//
	// camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	// cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	// gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])
	//
	// modelView := mgl32.Ident4()
	// modelUniform := gl.GetUniformLocation(program, gl.Str("modelView\x00"))
	// gl.UniformMatrix4fv(modelUniform, 1, false, &modelView[0])
	//
	// textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	// gl.Uniform1i(textureUniform, 0)
	//
	// gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	//
	// // Load the texture
	// texture, err := createTexture("square.png")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Configure the vertex data
	// var vao uint32
	// gl.GenVertexArrays(1, &vao)
	// gl.BindVertexArray(vao)
	//
	// var vbo uint32
	// gl.GenBuffers(1, &vbo)
	// gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
	//
	// vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	// gl.EnableVertexAttribArray(vertAttrib)
	// gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	//
	// texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	// gl.EnableVertexAttribArray(texCoordAttrib)
	// gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	//
	// // Configure global settings
	// gl.Enable(gl.DEPTH_TEST)
	// gl.DepthFunc(gl.LESS)
	// gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	//
	// angle := 0.0
	// previousTime := glfw.GetTime()
	//
	// for !window.ShouldClose() {
	// 	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	//
	// 	// Update
	// 	time := glfw.GetTime()
	// 	elapsed := time - previousTime
	// 	previousTime = time
	//
	// 	angle += elapsed
	// 	modelView = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})
	//
	// 	// Render
	// 	gl.UseProgram(program)
	// 	gl.UniformMatrix4fv(modelUniform, 1, false, &modelView[0])
	//
	// 	gl.BindVertexArray(vao)
	//
	// 	gl.ActiveTexture(gl.TEXTURE0)
	// 	gl.BindTexture(gl.TEXTURE_2D, texture)
	//
	// 	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
	//
	// 	// Maintenance
	// 	window.SwapBuffers()
	// 	glfw.PollEvents()
	// }

	return &Renderer{
		Window:     window,
		Shaders:    &shaders,
		Projection: projection,
		Camera:     camera,
		ModelView:  modelView,
	}
}

func (r *Renderer) Render() {
	defer glfw.Terminate()
	shader := r.Shaders.textureFlat
	program := shader.program
	//
	gl.UseProgram(program)
	//
	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/float32(windowHeight), 0.1, 10.0)
	// projectionUniform := gl.GetUniformLocation(program, )
	gl.UniformMatrix4fv(shader.uniforms["projection"], 1, false, &r.Projection[0])
	//
	// camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	// cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(shader.uniforms["camera"], 1, false, &r.Camera[0])
	//
	// model := mgl32.Ident4()
	// modelUniform := gl.GetUniformLocation(program, gl.Str("ModelView\x00"))
	gl.UniformMatrix4fv(shader.uniforms["modelView"], 1, false, &r.ModelView[0])
	//
	// textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(shader.uniforms["tex"], 0)
	//
	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	//
	// // Load the texture
	texture, err := createTexture("square.png")
	if err != nil {
		panic(err)
	}
	//
	// // Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	//
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
	//
	// vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(shader.attributes["vert"])
	gl.VertexAttribPointer(shader.attributes["vert"], 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	//
	// texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(shader.attributes["vertTexCoord"])
	gl.VertexAttribPointer(shader.attributes["vertTexCoord"], 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	//
	// // Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	//
	angle := 0.0
	previousTime := glfw.GetTime()
	//
	for !r.Window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		angle += elapsed
		r.ModelView = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

		// Render
		// gl.UseProgram(program)
		gl.UniformMatrix4fv(shader.uniforms["modelView"], 1, false, &r.ModelView[0])

		gl.BindVertexArray(vao)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture)

		gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

		// Maintenance
		r.Window.SwapBuffers()
		glfw.PollEvents()
	}
}
