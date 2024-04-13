package character

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw permet d'afficher le personnage dans une *ebiten.Image
// (en pratique, celle qui représente la fenêtre de jeu) en
// fonction des charactéristiques du personnage (position, orientation,
// étape d'animation, etc) et de la position de la caméra (le personnage
// est affiché relativement à la caméra).
func (c Character) Draw(screen *ebiten.Image, camX, camY int) {

	xShift := 0
	yShift := 0
	if configuration.Global.Personnage%2 == 1 {
		yShift += (configuration.Global.FrameCount % 60) / 30
	}

	switch c.orientation {
	case orientedDown:
		if configuration.Global.CameraMode != 3 {
			yShift = c.shift
		}

		if configuration.Global.Personnage%2 == 1 {
			yShift += (configuration.Global.FrameCount % 60) / 30
		}
	case orientedUp:
		if configuration.Global.CameraMode != 3 {
			yShift = -c.shift
		}

		if configuration.Global.Personnage%2 == 1 {
			yShift += (configuration.Global.FrameCount % 60) / 30
		}
	case orientedLeft:
		if configuration.Global.CameraMode != 3 {
			xShift = -c.shift
		}

	case orientedRight:
		if configuration.Global.CameraMode != 3 {
			xShift = c.shift
		}
	}

	xTileForDisplay := c.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := c.Y - camY + configuration.Global.ScreenCenterTileY
	xPos := (xTileForDisplay)*configuration.Global.TileSize + xShift
	yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := configuration.Global.TileSize
	if c.moving {
		shiftX += c.animationStep * configuration.Global.TileSize
	}
	shiftY := c.orientation * configuration.Global.TileSize

	if configuration.Global.Personnage%2 == 0 {
		screen.DrawImage(assets.CharacterImage.SubImage(
			image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	} else if configuration.Global.Personnage%2 == 1 {
		screen.DrawImage(assets.GhostImage.SubImage(
			image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}

}
