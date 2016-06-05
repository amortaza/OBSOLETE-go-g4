package g4

type Canvas struct {
	Framebuffer *FrameBuffer
	Width, Height int32
}

func NewCanvas(width, height int32) *Canvas {
	canvas := &Canvas{}

	canvas.Framebuffer = NewFrameBuffer(width, height)

	canvas.Width, canvas.Height = width, height

	return canvas
}

func (c *Canvas) Begin() {
	c.Framebuffer.Begin()

	texture := c.Framebuffer.Texture

	PushViewport(texture.Width, texture.Height)
	PushOrtho(texture.Width, texture.Height)
}

func (c *Canvas) Clear(red, green, blue float32) {
	ClearRect(c.Width, c.Height, red, green, blue)
}

var opaque = []float32{1,1,1,1}

func (c *Canvas) Paint(left, top int32, alphas []float32) {
	if alphas == nil {
		alphas = opaque
	}

	DrawTextureRectUpsideDown(c.Framebuffer.Texture, left, top, c.Width,c.Height,alphas)
}

func (c *Canvas) End() {
	PopOrtho()
	PopViewport()

	c.Framebuffer.End()
}

func (c *Canvas) Free() {
	c.Framebuffer.Free()
}


