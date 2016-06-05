package g4

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"g4/ace"
)

type TextureRect struct {
	program *ace.Program

	vao uint32
	vbo uint32
}

func NewTextureRect(vertexShaderFilename, fragmentShaderFilename string) *TextureRect {
	r := &TextureRect{}

	r.program = ace.NewProgram(vertexShaderFilename, fragmentShaderFilename)

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)

	return r
}

func (r *TextureRect) Draw(	texture *Texture,
				left int32, top int32,
				width int32, height int32,
				leftTopRightBottomAlphas []float32,
				projection *float32 ) {

	r.program.Activate()

	texture.Activate(gl.TEXTURE0)

	gl.Uniform1i(r.program.GetUniformLocation("Sampler"), 0)
	gl.UniformMatrix4fv(r.program.GetUniformLocation("Projection"), 1, false, projection)
	gl.Uniform4f(r.program.GetUniformLocation("Alphas"), leftTopRightBottomAlphas[0], leftTopRightBottomAlphas[1], leftTopRightBottomAlphas[2], leftTopRightBottomAlphas[3]);

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), 0.0, 0.0,
		float32(right), float32(top), 1.0, 0.0 ,
		float32(right), float32(bottom), 1.0, 1.0,
		float32(left), float32(bottom), 0.0, 1.0 }

	setVertexData2(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)

	texture.Deactivate()
}

func (r *TextureRect) DrawString( texture *Texture,
				left int32, top int32,
				width int32, height int32,
				rgb []float32,
				bg []float32,
				alpha float32,
				projection *float32 ) {

	r.program.Activate()

	texture.Activate(gl.TEXTURE0)

	gl.Uniform1i(r.program.GetUniformLocation("Sampler"), 0)
	gl.UniformMatrix4fv(r.program.GetUniformLocation("Projection"), 1, false, projection)
	gl.Uniform3f(r.program.GetUniformLocation("RGB"), rgb[0],rgb[1],rgb[2]);
	gl.Uniform3f(r.program.GetUniformLocation("Bg"), bg[0],bg[1],bg[2]);
	gl.Uniform1f(r.program.GetUniformLocation("Alpha"), alpha);

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	right := left + width
	bottom := top + height

	vertices := []float32{
		float32(left), float32(top), 0.0, 1.0,
		float32(right), float32(top), 1.0, 1.0 ,
		float32(right), float32(bottom), 1.0, 0.0,
		float32(left), float32(bottom), 0.0, 0.0 }

	setVertexData2(vertices)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.BindVertexArray(0)

	texture.Deactivate()
}

func (r *TextureRect) Free() {
	gl.DeleteVertexArrays(1, &r.vao)
	gl.DeleteBuffers(1, &r.vbo)

	r.program.Free()
}

func setVertexData2(data []float32) {

	// copy vertices data into VBO (it needs to be bound first)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 2 /*posPartCount*/ *4 + 2 /*texPartCount*/ *4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 2 /*posPartCount*/, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 2 /*posPartCount*/ * 4

	// texture
	gl.VertexAttribPointer(1, 2 /*texPartCount*/, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
}
