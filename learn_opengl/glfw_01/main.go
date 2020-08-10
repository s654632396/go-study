package main

import (
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

/** learn from: https://learnopengl-cn.github.io/01%20Getting%20started/03%20Hello%20Window/ */

func main() {

	mainthread.Run(run)

}

func run() {

	var win *glfw.Window

	mainthread.Call(func() {
		width, height := 640, 480
		win = initGL(width, height)
	})

	mainthread.Call(func() {
		for {
			// Do OpenGL stuff.
			// 处理输入
			processInput(win)

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
			if  win.ShouldClose() {
				win.Destroy()
				break
			}
		}
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
