package main

import (
	"fmt"
	_ "github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"image"
	"image/draw"
	_ "image/jpeg"
	"log"
	"os"
	"strings"
	"unsafe"
)

func main() {
	mainthread.Run(run)
}

func GetTriangleTexture() image.Image {
	var picFile = "container.jpg"
	f, err := os.Open(picFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}

func run() {

	var win *glfw.Window
	var program uint32

	mainthread.Call(func() {
		width, height := 640, 480
		win = initGL(width, height)

	})

	mainthread.Call(func() {

		program = initOpenGL()
		var vertex = []float32{
			-0.5, 0.5, 0, // top-left vertex
			1, 1, 0, // yellow
			0, 1, // texture(0,1)
			0.5, 0.5, 0, // top-right vertex
			1, 0, 0, // red
			1, 1, // texture(1,1)
			-0.5, -0.5, 0, // bottom-left vertex
			0, 0, 1, // blue
			0, 0, // texture(0,0)
			0.5, -0.5, 0, // bottom-right vertex
			0, 1, 0, // green
			1, 0, // texture(1,0)
		}

		var indices = []uint32{
			0, 1, 3,
			0, 2, 3,
		}
		var vao = makeVao(vertex, indices)

		// make texture

		var myTexture uint32
		gl.GenTextures(1, &myTexture)

		gl.BindTexture(gl.TEXTURE_2D, myTexture)
		defer gl.BindTexture(gl.TEXTURE_2D, 0)

		// get our texture img
		img := GetTriangleTexture()
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)

		if rgba.Stride != rgba.Rect.Size().X*4 {
			panic("unsupported stride")
		}

		// Set the texture wrapping parameters
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
		// Set texture filtering parameters
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR) // minification filter
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR) // magnification filter

		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y),
			0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
		gl.GenerateMipmap(gl.TEXTURE_2D)

		for !win.ShouldClose() {
			// Do OpenGL stuff.
			// 处理输入
			processInput(win)

			runDraw(vao, myTexture, win, program)

			glfw.PollEvents()
		}
		gl.DeleteShader(vao)
		win.Destroy()
	})

}

func runDraw(vao uint32, texture uint32, win *glfw.Window, prog uint32) {
	gl.ClearColor(0.2, 0.3, 0.3, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.UseProgram(prog)

	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)

	// swap in rendered buffer
	win.SwapBuffers()
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

	glfw.SwapInterval(1) // enable vsync

	return window
}

func processInput(w *glfw.Window) {

	// 如果按下esc
	if w.GetKey(glfw.KeyEscape) == glfw.Press {
		log.Println("get close signal!")
		w.SetShouldClose(true)
	}
}

// opengl
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL Version:", version)

	vertexShader, err := compileShader(vertexShaderSrc, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	return program
}

func makeVao(vertex []float32, indices []uint32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)

	var vao uint32
	gl.GenVertexArrays(1, &vao)

	var ebo uint32
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)
	vsize := int32(unsafe.Sizeof(vertex[0]))
	isize := int32(unsafe.Sizeof(indices[0]))

	// bind vbo buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, int(vsize)*len(vertex), gl.Ptr(vertex), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int(isize)*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*vsize, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	//// color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*vsize, gl.PtrOffset(int(3*vsize)))
	gl.EnableVertexAttribArray(1)

	// texture
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*vsize, gl.PtrOffset(int(6*vsize)))
	gl.EnableVertexAttribArray(2)

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return vao
}

func compileShader(src string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	shaderSrc, free := gl.Strs(src)
	defer free()
	gl.ShaderSource(shader, 1, shaderSrc, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLen)
		log := strings.Repeat("\x00", int(logLen+1))
		gl.GetShaderInfoLog(shader, logLen, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile %v: %v", src, log)
	}

	return shader, nil
}

var vertexShaderSrc = `
#version 330 core
layout (location = 0) in vec3 position; // 位置变量的属性位置值为 0 
layout (location = 1) in vec3 tColor; 
layout (location = 2) in vec2 texCoord; 

out vec2 TexCoord;
out vec3 myColor;

void main() {
	gl_Position = vec4(position, 1.0);
 	TexCoord = texCoord;
	myColor	= tColor;
}
` + "\x00"

var fragmentShaderSrc = `
#version 330 core
in vec2 TexCoord;
in vec3 myColor;

out vec4 color;

uniform sampler2D ourTexture;

void main() {
	color = texture(ourTexture, TexCoord) * vec4(myColor, 1.0);
}
` + "\x00"
