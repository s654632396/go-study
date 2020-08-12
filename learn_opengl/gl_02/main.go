package main

import (
	"fmt"
	_ "github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"strings"
	"unsafe"
)

func main() {
	mainthread.Run(run)
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
			0.5, 0.5, 0.0, // 右上角
			0.5, -0.5, 0.0, // 右下角
			-0.5, -0.5, 0.0, // 左下角
			-0.5, 0.5, 0.0, // 左上角
		}
		var indices = []uint32{
			0, 1, 3, // 第一个三角
			1, 2, 3, // 第二个三角
		}
		var vao = makeVao(vertex, indices)

		for !win.ShouldClose() {
			// Do OpenGL stuff.
			// 处理输入
			processInput(win)

			draw(vao, win, program)

			glfw.PollEvents()
		}
		gl.DeleteShader(vao)
		win.Destroy()
	})

}

func draw(vao uint32, win *glfw.Window, prog uint32) {
	gl.ClearColor(0.2, 0.3, 0.3, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.UseProgram(prog)
	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
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
	//gl.Enable(gl.DEPTH_TEST) // Q: 开启后看不见三角形了.why?

	// //面剔除, 暂时不太懂,关了吧
	//gl.Enable(gl.CULL_FACE)
	//gl.CullFace(gl.FRONT)

	// 线框模式(Wireframe Mode), 可以配置gl.LINE来观察仅绘制了线条的图元
	// 第一个参数表示我们打算将其应用到所有的三角形的正面和背面，第二个参数告诉我们用线来绘制
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	// 默认模式
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

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

	var ebo uint32
	gl.GenBuffers(1, &ebo)

	var vao uint32
	gl.GenVertexArrays(1, &vao)

	gl.BindVertexArray(vao)

	// bind vbo buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertex), gl.Ptr(vertex), gl.STATIC_DRAW)

	// bind ebo buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3 * int32(unsafe.Sizeof(vertex[0])), gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// color
	//gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	//gl.EnableVertexAttribArray(1)

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
layout (location = 0) in vec3 position;
void main() {
	gl_Position = vec4(position.x, position.y, position.z, 1);
}
` + "\x00"

var fragmentShaderSrc = `
#version 330 core
out vec4 color;
void main() {
	color = vec4(1, 0.5, 0.2, 1);
}
` + "\x00"
