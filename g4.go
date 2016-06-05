package g4

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v3.3-core/gl"
	"adt"
)

var g_projection mgl32.Mat4

var g_colorRect *ColorRect
var g_textureRect *TextureRect
var g_stringRect *TextureRect

var g_viewportWidthStack  adt.Stack
var g_viewportHeightStack adt.Stack
var g_orthoStack adt.Stack

func Init() {
	gl.ClearColor(0.1, 0.4, 0.4, 1.0)

	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)

	// blending is required to be able to render text
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	g_colorRect = NewColorRect()
	g_textureRect = NewTextureRect("g4/shader/texture-rect.vertex.txt", "g4/shader/texture-rect.fragment.txt")
	g_stringRect = NewTextureRect("g4/shader/font.vertex.txt", "g4/shader/font.fragment.txt")
}

func Clear(red,green,blue,alpha float32) {
	gl.ClearColor(red,green,blue,alpha)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Uninit() {
	g_stringRect.Free()
	g_textureRect.Free()
	g_colorRect.Free()
}

func PushView(width, height int32) {
	PushViewport(width, height)
	PushOrtho(width,height)
}

func PopView() {
	PopViewport()
	PopOrtho()
}

func PushViewport(width, height int32) {
	g_viewportWidthStack.Push(width)
	g_viewportHeightStack.Push(height)

	gl.Viewport(0, 0, width, height);
}

func PopViewport() {
	g_viewportWidthStack.Pop()
	g_viewportHeightStack.Pop()

	if g_viewportWidthStack.Size != 0 {
		width, _ := g_viewportWidthStack.Top().(int32)
		height, _ := g_viewportHeightStack.Top().(int32)

		gl.Viewport(0, 0, width, height);
	}
}

func PushOrtho(width, height int32) {
	g_projection = mgl32.Ortho2D(0, float32(width), float32(height), 0)
	g_orthoStack.Push(g_projection)
}

func PopOrtho() {
	g_orthoStack.Pop()

	if g_orthoStack.Size != 0 {
		g_projection = g_orthoStack.Top().(mgl32.Mat4)
	}
}
