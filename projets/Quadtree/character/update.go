package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particule"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, ListFootprint particule.FootprintList) (particule.FootprintList ) {
	//bolléen permettant de savoir si le personnage est statique (peut avancer)
	//IsStatic := (configuration.Global.MooveRunTimeTop == 0 && configuration.Global.MooveRunTimeLeft == 0 && configuration.Global.MooveRunTimeRight == 0 && configuration.Global.MooveRunTimeBottom == 0)
	if !c.moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if !blocking[1] {
				c.xInc = 1
				c.moving = true
				if configuration.Global.ParticuleOn {
					ListFootprint.CreateFootPrint(c.X, c.Y, 3)
				}
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.moving = true
				if configuration.Global.ParticuleOn {
					ListFootprint.CreateFootPrint(c.X, c.Y, 1)
				}
			}

		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if !blocking[0] {
				c.yInc = -1
				c.moving = true
				if configuration.Global.ParticuleOn {
					ListFootprint.CreateFootPrint(c.X, c.Y, 0)
				}
			}

		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if !blocking[2] {
				c.yInc = 1
				c.moving = true
				if configuration.Global.ParticuleOn {
					ListFootprint.CreateFootPrint(c.X, c.Y, 2)
				}
			}
		}
	} else {
		//fmt.Println("animationFrameCount=", c.animationFrameCount)
		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
			}
		}
	}
	return ListFootprint

}
