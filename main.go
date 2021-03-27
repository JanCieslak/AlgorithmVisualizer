package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	mgl "github.com/go-gl/mathgl/mgl32"
	_ "image/png"
	"log"
	"runtime"
	"strconv"
)

const WindowWidth = 1280
const WindowHeight = 640

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	err := gl.Init()
	if err != nil {
		log.Fatal("Failed to initialize GL")
	}

	gl.ClearColor(0.2, 0.3, 0.8, 1.0)

	shader := newShader("basic")
	shader.Bind()

	generateTextures(shader)

	ortho := mgl.Ortho(0, WindowWidth, WindowHeight,0, -1.0, 1.0)
	shader.Projection(ortho)

	grid := newGrid(10, 20, 64, 64)

	renderer := newRenderer()
	isDone := false

	for !window.ShouldClose() {
		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		renderer.Begin()
		renderer.AddGrid(grid)
		renderer.End()

		if !isDone {
			go generateMaze(grid)
			isDone = true
		}

		window.SwapBuffers()
	}
}

func generateTextures(shader *Shader) {
	textures := make([]*Texture, 16)

	for i := 0; i < 16; i++ {
		textures[i] = newTexture(strconv.Itoa(i) + ".png")
		shader.Uniform1i("tex[" + strconv.Itoa(i) + "]", int32(i))
	}

	for i := 0; i < 16; i++ {
		gl.ActiveTexture(uint32(gl.TEXTURE0 + i))
		textures[i].Bind()
	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "Maze generator", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Center the window
	monitor := glfw.GetPrimaryMonitor()
	vidmode := monitor.GetVideoMode()
	monitorX, monitorY := monitor.GetPos()
	window.SetPos(monitorX + (vidmode.Width - WindowWidth) / 2,
		monitorY + (vidmode.Height - WindowHeight) / 2)


	return window
}