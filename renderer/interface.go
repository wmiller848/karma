package renderer

type GenericRenderer interface {
	Kill()
	ShouldDie() bool
	Render()

	SetCamera(windowWidth, windowHeight int)
	AddMesh(vertices []float32, texturePaths []string)

	UpdateCameraPos(vec [3]float32)
	UpdateCameraRot(vec [3]float32)

	UpdateMeshPos(id uint32, vec [3]float32)
	UpdateMeshRot(id uint32, vec [3]float32)
}
