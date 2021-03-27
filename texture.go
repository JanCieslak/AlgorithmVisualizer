package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"image"
	"image/draw"
	"log"
	"os"
)

type Texture struct {
	textureId uint32
}

func newTexture(textureName string) *Texture {
	var texture = Texture{}

	imgFile, err := os.Open("resources/images/" + textureName)
	if err != nil {
		log.Fatal("Couldn't load the texture: ", err)
		return nil
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal("Couldn't load the texture: ", err)
		return nil
	}

	bounds := img.Bounds()
	nrgba := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)

	gl.GenTextures(1, &texture.textureId)
	gl.BindTexture(gl.TEXTURE_2D, texture.textureId)
	defer gl.BindTexture(gl.TEXTURE_2D, 0)

	gl.TexParameteri(texture.textureId, gl.TEXTURE_WRAP_R, gl.REPEAT)
	gl.TexParameteri(texture.textureId, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(texture.textureId, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.textureId, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(nrgba.Rect.Size().X), int32(nrgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(nrgba.Pix))

	gl.GenerateMipmap(gl.TEXTURE_2D)

	return &texture
}

func (it *Texture) Bind()  {
	gl.BindTexture(gl.TEXTURE_2D, it.textureId)
}

func (it *Texture) Unbind()  {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

