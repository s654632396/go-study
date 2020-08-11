package main

import (
	"fmt"
	_ "github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

/** learn from: https://learnopengl-cn.github.io/01%20Getting%20started/03%20Hello%20Window/ */

func main() {

	mainthread.Run(run)

}

var vShaderSrcStr = `
#version 330 core

layout (location = 0) in vec3 position;

void main()
{
    gl_Position = vec4(position.x, position.y, position.z, 1.0);
}
`

var fShaderSrcStr = `
#version 330 core

out vec4 color;

void main()
{
    color = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
`

func run() {

	var win *glfw.Window

	mainthread.Call(func() {
		width, height := 640, 480
		win = initGL(width, height)
	})

	mainthread.Call(func() {

		var vshader uint32
		{
			vshader = gl.CreateShader(gl.VERTEX_SHADER)

			src, free := gl.Strs(vShaderSrcStr)
			defer free()
			var lenSrc = int32(len(vShaderSrcStr))
			gl.ShaderSource(vshader, 1, src, &lenSrc)
			gl.CompileShader(vshader)
			var success int32
			gl.GetShaderiv(vshader, gl.COMPILE_STATUS, &success)
			if success == gl.FALSE {
				var logLen int32
				gl.GetShaderiv(vshader, gl.INFO_LOG_LENGTH, &logLen)
				infoLog := make([]byte, logLen)
				gl.GetShaderInfoLog(vshader, logLen, nil, &infoLog[0])
				log.Println(
					fmt.Errorf("error compiling vertex shader: %s", string(infoLog)),
				)
			}

		}

		// fragment shader
		fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
		{

			src, free := gl.Strs(fShaderSrcStr)
			defer free()
			var lenSrc = int32(len(fShaderSrcStr))
			gl.ShaderSource(fshader, 1, src, &lenSrc)
			gl.CompileShader(fshader)
			var success int32
			gl.GetShaderiv(fshader, gl.COMPILE_STATUS, &success)
			if success == gl.FALSE {
				var logLen int32
				gl.GetShaderiv(fshader, gl.INFO_LOG_LENGTH, &logLen)
				infoLog := make([]byte, logLen)
				gl.GetShaderInfoLog(fshader, logLen, nil, &infoLog[0])
				log.Println(
					fmt.Errorf("error compiling fragment shader: %s", string(infoLog)),
				)
			}
		}

		shaderProgram := gl.CreateProgram()
		{
			gl.AttachShader(shaderProgram, vshader)
			gl.AttachShader(shaderProgram, fshader)
			gl.LinkProgram(shaderProgram)
			var success int32
			gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
			if success == gl.FALSE {
				var logLen int32
				gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLen)

				infoLog := make([]byte, logLen)
				gl.GetProgramInfoLog(shaderProgram, logLen, nil, &infoLog[0])
				log.Println(
					fmt.Errorf("error linking shader program: %s", string(infoLog)),
				)
			}
			gl.DeleteShader(vshader)
			gl.DeleteShader(fshader)
		}

		var bData = []float32{
			-0.5, -0.5, 0.0, // Left
			0.5, -0.5, 0.0, // Right
			0.0, 0.5, 0.0, // Top
		}

		var vao uint32
		var vbo uint32
		gl.GenVertexArrays(1, &vao)
		gl.GenBuffers(1, &vbo)

		// 1. 绑定VAO
		gl.BindVertexArray(vao)

		// 2. 把顶点数组复制到缓冲中供OpenGL使用
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(bData), gl.Ptr(bData), gl.STATIC_DRAW)

		// 3. 设置顶点属性指针
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*int32(len(bData)), gl.PtrOffset(0))
		gl.EnableVertexArrayAttrib(vao, 0)

		gl.BindBuffer(gl.ARRAY_BUFFER, 0)

		// 4. 解绑VAO
		gl.BindVertexArray(0)

		for !win.ShouldClose() {
			// Do OpenGL stuff.
			// 处理输入
			processInput(win)

			gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			gl.Clear(gl.COLOR_BUFFER_BIT)

			gl.UseProgram(shaderProgram)
			gl.BindVertexArray(vao)
			gl.DrawArrays(gl.TRIANGLES, 0, 3)
			gl.BindVertexArray(0)

			// glfwSwapBuffers函数会交换颜色缓冲(它是一个储存着GLFW窗口每一个像素颜色值的大缓冲)
			// 它在这一迭代中被用来绘制，并且将会作为输出显示在屏幕上。
			// ### 知识点 ###
			// ** 双缓冲(Double Buffer) **
			// 应用程序使用单缓冲绘图时可能会存在图像闪烁的问题。
			//这是因为生成的图像不是一下子被绘制出来的，而是按照从左到右，由上而下逐像素地绘制而成的。
			//最终图像不是在瞬间显示给用户，而是通过一步一步生成的，这会导致渲染的结果很不真实。
			//为了规避这些问题，我们应用双缓冲渲染窗口应用程序。
			//前缓冲保存着最终输出的图像，它会在屏幕上显示；而所有的的渲染指令都会在后缓冲上绘制。
			//当所有的渲染指令执行完毕后，我们交换(Swap)前缓冲和后缓冲，这样图像就立即呈显出来，之前提到的不真实感就消除了。
			win.SwapBuffers()
			glfw.PollEvents()
		}

		gl.DeleteVertexArrays(1, &vao)
		gl.DeleteBuffers(1, &vbo)

		win.Destroy()
	})

}

func drawScene() {

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
