package g4

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/amortaza/go-g4/ace"
)

type ColorRect struct {
	program *ace.Program

	vao uint32
	vbo uint32
}

func NewColorRect() *ColorRect {
	r := &ColorRect{}

	r.program = ace.NewProgram("github.com/amortaza/go-g4/shader/rgb.vertex.txt", "github.com/amortaza/go-g4/shader/rgb.fragment.txt")

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)

	return r
}

func (r *ColorRect) Draw(	left int32, top int32,
				width int32, height int32,
				leftTopColor []float32,
				rightTopColor []float32,
				rightBottomColor []float32,
				leftBottomColor []float32,
				projection *float32 ) {

	r.program.Activate()

	gl.UniformMatrix4fv(r.program.GetUniformLocation("project"), 1, false, projection)

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), leftTopColor[0], leftTopColor[1], leftTopColor[2], leftTopColor[3],
			float32(right), float32(top), rightTopColor[0], rightTopColor[1], rightTopColor[2], rightTopColor[3],
				float32(right), float32(bottom), rightBottomColor[0], rightBottomColor[1], rightBottomColor[2], rightBottomColor[3],
					float32(left), float32(bottom), leftBottomColor[0], leftBottomColor[1], leftBottomColor[2], leftBottomColor[3] }

	colorRect_setVertexData(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)
}

func (r *ColorRect) DrawSolid(	left int32, top int32,
				width int32, height int32,
				red, green, blue float32,
				projection *float32 ) {

	r.program.Activate()

	gl.UniformMatrix4fv(r.program.GetUniformLocation("project"), 1, false, projection)

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), red, green, blue, 1,
			float32(right), float32(top), red, green, blue, 1,
				float32(right), float32(bottom), red, green, blue, 1,
					float32(left), float32(bottom), red, green, blue, 1 }

	colorRect_setVertexData(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)
}

func (r *ColorRect) Free() {
	gl.DeleteVertexArrays(1, &r.vao)
	gl.DeleteBuffers(1, &r.vbo)

	r.program.Free()
}

func colorRect_setVertexData(data []float32) {

	// copy vertices data into VBO (it needs to be bound first)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 2 /*posPartCount*/ *4 + 4 /*colorPartCount*/ *4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 2 /*posPartCount*/, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 2 /*posPartCount*/ * 4

	// color
	gl.VertexAttribPointer(1, 4 /*colorPartCount*/, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
}
