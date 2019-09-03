package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

func init() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	err := gl.Init()
	if err != nil {
		panic(err)
	}

	runtime.LockOSThread()
}

func main() {
	win, err := glfw.CreateWindow(800, 600, "hello world", nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()

	for !win.ShouldClose() {
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.DrawArrays(gl.POINTS, 0, 4)

		win.SwapBuffers()
		glfw.PollEvents()
		if win.GetKey(glfw.KeyEscape) == glfw.Press {
			win.SetShouldClose(true)
		}
	}
}
