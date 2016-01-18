package renderer

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	textures  []uint32
	verticies []float32
	vao       uint32
	vbo       uint32
	modelView mgl32.Mat4
}

func CreateMesh(verticies []float32, texturePaths []string, shader *Shader) (*Mesh, error) {

	modelView := mgl32.Ident4()

	// // Load the texture
	textures := make([]uint32, 0)
	// fmt.Println(texturePaths)
	for _, texturePath := range texturePaths {
		// fmt.Println("Paths", texturePath, texturePaths[i])
		texture, err := createTexture(texturePath)
		if err != nil {
			return nil, err
		}
		textures = append(textures, texture)
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
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, gl.Ptr(verticies), gl.STATIC_DRAW)
	//
	gl.EnableVertexAttribArray(shader.attributes["vert"])
	gl.VertexAttribPointer(shader.attributes["vert"], 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	//
	gl.EnableVertexAttribArray(shader.attributes["vertTexCoord"])
	gl.VertexAttribPointer(shader.attributes["vertTexCoord"], 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return &Mesh{
		modelView: modelView,
		textures:  textures,
		verticies: verticies,
		vao:       vao,
		vbo:       vbo,
	}, nil
}
