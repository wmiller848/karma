package renderer

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Shaders struct {
	textureFlat *Shader
}

type Renderer struct {
	Ready        bool
	PreviousTime float64
	Window       *glfw.Window
	Shaders      *Shaders
	Projection   mgl32.Mat4
	Camera       mgl32.Mat4
	Meshes       []*Mesh
	Fps          int
}

func testCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Println("Callback")
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
	if windowWidth == 0 && windowHeight == 0 {
		windowWidth, windowHeight = glfw.GetPrimaryMonitor().GetPhysicalSize()
	}
	fmt.Println("Window Width -", windowWidth, "Window Height -", windowHeight)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Karma", glfw.GetPrimaryMonitor(), nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(testCallback)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	shaders := Shaders{}

	textureFlatUniforms := []string{"projection", "camera", "modelView", "tex"}
	textureFlatAttributes := []string{"vert", "vertTexCoord"}

	fmt.Println(textureFlatUniforms)
	fmt.Println(textureFlatAttributes)

	shader, err := createProgram("./assets/shaders/texture_flat.vs", "./assets/shaders/texture_flat.fs", textureFlatUniforms, textureFlatAttributes)
	if err != nil {
		panic(err)
	}
	shaders.textureFlat = shader

	meshes := []*Mesh{}

	previousTime := glfw.GetTime()

	return &Renderer{
		PreviousTime: previousTime,
		Window:       window,
		Shaders:      &shaders,
		Meshes:       meshes,
	}
}

func (r *Renderer) Kill() {
	glfw.Terminate()
}

func (r *Renderer) ShouldDie() bool {
	if r.Ready == true {
		return r.Window.ShouldClose()
	} else {
		return false
	}
}

func (r *Renderer) Render() {
	// defer glfw.Terminate()
	shader := r.Shaders.textureFlat
	program := shader.program
	//
	gl.UseProgram(program)
	//
	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	// // Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	//
	// angle += elapsed
	// r.Mesh.modelView = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

	// Render
	// gl.UniformMatrix4fv(shader.uniforms["modelView"], 1, false, &r.Mesh.modelView[0])

	time := glfw.GetTime()
	_ = time - r.PreviousTime
	r.PreviousTime = time

	// fmt.Println(elapsed * 100)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UniformMatrix4fv(shader.uniforms["projection"], 1, false, &r.Projection[0])
	gl.UniformMatrix4fv(shader.uniforms["camera"], 1, false, &r.Camera[0])

	// TODO : batch triangles and use multiple textures
	for _, mesh := range r.Meshes {
		gl.UniformMatrix4fv(shader.uniforms["modelView"], 1, false, &mesh.modelView[0])
		gl.Uniform1i(shader.uniforms["tex"], 0)

		gl.BindVertexArray(mesh.vao)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, mesh.textures[0])

		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(mesh.verticies)/5))
	}

	// Maintenance
	r.Window.SwapBuffers()
	glfw.PollEvents()
	if r.Ready == false {
		r.Ready = true
	}
}

func (r *Renderer) SetCamera(windowWidth, windowHeight int) {
	r.Projection = mgl32.Perspective(mgl32.DegToRad(65.0), float32(windowWidth)/float32(windowHeight), 0.1, 1000.0)
	r.Camera = mgl32.Translate3D(0, 0, 0)
	// camera := camera.Mul4(mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}))
}

func (r *Renderer) AddMesh(vertices []float32, texturePaths []string) {
	mesh, err := CreateMesh(vertices, texturePaths, r.Shaders.textureFlat)
	if err != nil {
		panic(err)
	}
	r.Meshes = append(r.Meshes, mesh)
}

func (r *Renderer) UpdateCameraPos(vec [3]float32) {
	// gl.UniformMatrix4fv(shader.uniforms["camera"], 1, false, &r.Camera[0])
	r.Camera = mgl32.Translate3D(vec[0], vec[1], vec[2])
}

func (r *Renderer) UpdateCameraRot(vec [3]float32) {
	// gl.UniformMatrix4fv(shader.uniforms["camera"], 1, false, &r.Camera[0])
}

func (r *Renderer) UpdateMeshPos(id uint32, vec [3]float32) {
	// gl.UniformMatrix4fv(shader.uniforms["camera"], 1, false, &r.Camera[0])
	r.Meshes[id].modelView = mgl32.Translate3D(vec[0], vec[1], vec[2])
}

func (r *Renderer) UpdateMeshRot(id uint32, vec [3]float32) {
	// gl.UniformMatrix4fv(shader.uniforms["camera"], 1, false, &r.Camera[0])
}
