package g4

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type FrameBuffer struct {
	Texture *Texture

	FramebufferId uint32
}

func NewFrameBuffer(width, height int32) *FrameBuffer {
	f := &FrameBuffer{}

	f.Texture = NewTexture()

	f.Texture.Allocate(width, height)

	gl.GenFramebuffers(1, &f.FramebufferId)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FramebufferId)
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, f.Texture.TextureId, 0)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return f
}

func (f *FrameBuffer) Begin() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FramebufferId);
}

func (f *FrameBuffer) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (f *FrameBuffer) Free() {
	f.Texture.Free()

	gl.DeleteFramebuffers(1, &f.FramebufferId);
}
