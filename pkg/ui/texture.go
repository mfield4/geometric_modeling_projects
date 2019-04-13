package ui

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadTexture(renderer *sdl.Renderer, imagePath string) *sdl.Texture {
	image, err := img.Load(imagePath)
	if err != nil {
		panic("Image failed to load!")
	}
	// Create texture from image
	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic("Texture failed to load")
	}
	return texture
}
