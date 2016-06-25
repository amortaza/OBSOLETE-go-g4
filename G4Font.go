package g4

import (
	"io/ioutil"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"fmt"
)

type G4Font struct {
	Face *font.Face
	Height int32
	Ascent int32
	Descent int32
}

func LoadTrueTypeFromFile(fontFilename string) *truetype.Font {
	fmt.Println("Reading font file ", fontFilename)
	fontBytes, _ := ioutil.ReadFile(fontFilename)

	f, _ := truetype.Parse(fontBytes)

	return f
}

func NewG4Font(f *truetype.Font, fontSize int32) *G4Font {
	g4font := &G4Font{}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     196.0,
		Hinting: font.HintingNone })

	metrics := face.Metrics()

	g4font.Height = int32(metrics.Height.Ceil())
	g4font.Descent = int32(metrics.Descent.Ceil())
	g4font.Ascent = int32(metrics.Ascent.Ceil())
	g4font.Face = &face

	return g4font
}

func (f *G4Font) Width(str string) int32 {
	return int32(font.MeasureString(*f.Face, str)>>6)
}

func (f *G4Font) Free() {
	fmt.Println("Free has not been implemented for font.AceFont")
}