package main

import (
	"github.com/faiface/glhf"
	_ "github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
)

/** learn from: https://learnopengl-cn.github.io/01%20Getting%20started/03%20Hello%20Window/ */

func main() {

	mainthread.Run(run)

}

func loadImage(path string) (*image.NRGBA, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	nrgba := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)
	return nrgba, nil
}

func run() {

	var win *glfw.Window

	mainthread.Call(func() {
		width, height := 640, 480
		win = initGL(width, height)
	})



	mainthread.Call(func() {

		var vshaderSrcStr = `
#version 330 core

layout (location = 0) in vec3 position;

void main()
{
    gl_Position = vec4(position.x, position.y, position.z, 1.0);
}
`

		var fshaderSrcStr = `
#version 330 core

out vec4 color;

void main()
{
    color = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
`


		var (
			shader  *glhf.Shader
			slice   *glhf.VertexSlice
			texture *glhf.Texture

			err error
		)

		gopherImage, err := loadImage("celebrate.png")
		if err != nil {
			panic(err)
		}
		var vertexFormat = glhf.AttrFormat{
			{Name: "position", Type: glhf.Vec4},
			{Name: "color", Type: glhf.Vec4},
		}
		shader, err = glhf.NewShader(vertexFormat, glhf.AttrFormat{}, vshaderSrcStr, fshaderSrcStr)
		if err != nil {
			log.Println(err)
		}

		texture = glhf.NewTexture(
			gopherImage.Bounds().Dx(),
			gopherImage.Bounds().Dy(),
			true,
			gopherImage.Pix,
		)

		slice = glhf.MakeVertexSlice(shader, 1, 1)

		slice.Begin()

		var bData = []float32{
			0.5, 0.5, 0.0, // Top Right
			0.5, -0.5, 0.0, // Bottom Right
			-0.5, -0.5, 0.0, // Bottom Left
			-0.5, 0.5, 0.0, // Top Left
		}
		slice.SetVertexData(bData)
		slice.End()

		var indices = []int32{
			0, 1, 3,
			1, 2, 3,
		}
		log.Println(vshaderSrcStr, fshaderSrcStr, bData, indices)

		for !win.ShouldClose() {
			// Do OpenGL stuff.
			// 处理输入
			processInput(win)
			// glhf.Clear(1, 1, 1, 1)
			gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			gl.Clear(gl.COLOR_BUFFER_BIT)

			shader.Begin()
			texture.Begin()
			slice.Begin()
			slice.Draw()
			slice.End()
			texture.End()
			shader.End()

			win.SwapBuffers()
			glfw.PollEvents()
		}

		win.Destroy()
	})

}

func initGL(w, h int) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	// 创建一个窗口对象，这个窗口对象存放了所有和窗口相关的数据，而且会被GLFW的其他函数频繁地用到
	window, err := glfw.CreateWindow(w, h, "My First Step", nil, nil)
	if err != nil {
		panic(err)
	}
	// 创建完窗口我们就可以通知GLFW将我们窗口的上下文设置为当前线程的主上下文了。
	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}

	glfw.SwapInterval(1) // enable vsync
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)

	return window
}

func processInput(w *glfw.Window) {

	// 如果按下esc
	if w.GetKey(glfw.KeyEscape) == glfw.Press {
		log.Println("get close signal!")
		w.SetShouldClose(true)
	}
}
