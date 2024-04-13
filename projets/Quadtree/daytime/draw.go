package daytime

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image/color"
)

func Draw(screen *ebiten.Image, PosCamX, PosCamY, CharcX, CharcY, IncX, IncY, shift int) {
	// la nuit ne commence qu'à partir de 3600
	if configuration.Global.Daytime > 3600 {
		//si la caméra est pas fluide
		if configuration.Global.CameraMode == 3 {
			shift = 0
		}
		//on récupère la position de la lumière (même que le joueur)
		x := CharcX - PosCamX + configuration.Global.ScreenCenterTileX
		y := CharcY - PosCamY + configuration.Global.ScreenCenterTileY

		//raccourci de la taille de la camera sur l'écran
		CamWidth := configuration.Global.NumTileX * configuration.Global.TileSize
		CamHeight := configuration.Global.NumTileY * configuration.Global.TileSize
		tailleDemiTuile := configuration.Global.TileSize / 2

		// on obtient la transparance de l'ombre via ce calcul
		Shadow := uint8((float64(configuration.Global.Daytime/10) - 360) * (720 - float64(configuration.Global.Daytime/10)) / 140)
		// le maximum de l'ombre est 0xc4 (196)
		if Shadow > 0xc4 {
			Shadow = 0xc4
		}
		//on affiche l'ombre
		vector.DrawFilledRect(screen, 0, 0, float32(CamWidth), float32(CamHeight), color.NRGBA{0, 0, 0, Shadow}, false)
		// notre lumière ne s'active que lorsque l'ombre est à 110 ou plus
		if Shadow >= 110 {
			// on obtient la transparance de notre lumière via ce calcul
			Light := uint8((float64(configuration.Global.Daytime/10) - 410) * (670 - float64(configuration.Global.Daytime/10)) / 200)
			// le maximum de l'ombre est 0x40 (64)
			if Light > 0x40 {
				Light = 0x40
			}
			//on afffiche la lumière
			vector.DrawFilledCircle(screen, float32(x*configuration.Global.TileSize+tailleDemiTuile+IncX*shift), float32(y*configuration.Global.TileSize+tailleDemiTuile+IncY*shift), 48, color.NRGBA{0xff, 0xff, 0xff, Light}, false)
			vector.DrawFilledCircle(screen, float32(x*configuration.Global.TileSize+tailleDemiTuile+IncX*shift), float32(y*configuration.Global.TileSize+tailleDemiTuile+IncY*shift), 16, color.NRGBA{0x40, 0x40, 0x10, Light}, false)
			vector.DrawFilledCircle(screen, float32(x*configuration.Global.TileSize+tailleDemiTuile+IncX*shift), float32(y*configuration.Global.TileSize+tailleDemiTuile+IncY*shift), 32, color.NRGBA{0x40, 0x40, 0x10, Light}, false)
			vector.DrawFilledCircle(screen, float32(x*configuration.Global.TileSize+tailleDemiTuile+IncX*shift), float32(y*configuration.Global.TileSize+tailleDemiTuile+IncY*shift), 40, color.NRGBA{0x40, 0x40, 0x10, Light}, false)
		}
	}
}
