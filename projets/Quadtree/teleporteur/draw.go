package teleporteur

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image"
)

func (List_tp *Liste_teleporteur) Draw(screen *ebiten.Image, PosCamX, PosCamY, CharaxInc, CharayInc, shift int) {
	//Cette fonction affiche les sprites des téléporteurs sur la map

	// biais permettant de décaler les tuiles lors de la marche
	biais := 0
	if configuration.Global.CameraMode == 3 {
		// si la camera est fluide alors on augmente le biais
		biais = shift
	}

	//On affiche le TPA si il existe en récupérant ses coordonnées et en les adaptant à la map
	if List_tp.TPA != nil {
		op := &ebiten.DrawImageOptions{}
		x := List_tp.TPA.X - PosCamX + configuration.Global.ScreenCenterTileX
		y := List_tp.TPA.Y - PosCamY + configuration.Global.ScreenCenterTileY
		//on regarde si l'empreinte peut être dans l'écran
		if (x > -1 && x < configuration.Global.NumTileX+1) && (y > -1 && y < configuration.Global.NumTileY+1) {
			op.GeoM.Translate(float64((x)*configuration.Global.TileSize-biais*CharaxInc), float64((y)*configuration.Global.TileSize-biais*CharayInc))
			shiftX := 5 * configuration.Global.TileSize
			shiftY := ((configuration.Global.FrameCount / 12) % 16) * configuration.Global.TileSize
			screen.DrawImage(assets.FloorImage.SubImage(
				image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
			).(*ebiten.Image), op)
		}
	}

	//On affiche le TPB si il existe en récupérant ses coordonnées et en les adaptant à la map
	if List_tp.TPB != nil {
		op := &ebiten.DrawImageOptions{}
		x := List_tp.TPB.X - PosCamX + configuration.Global.ScreenCenterTileX
		y := List_tp.TPB.Y - PosCamY + configuration.Global.ScreenCenterTileY

		if (x > -1 && x < configuration.Global.NumTileX+1) && (y > -1 && y < configuration.Global.NumTileY+1) {
			op.GeoM.Translate(float64((x)*configuration.Global.TileSize-biais*CharaxInc), float64((y)*configuration.Global.TileSize-biais*CharayInc))
			shiftX := 5 * configuration.Global.TileSize
			shiftY := ((configuration.Global.FrameCount / 12) % 16) * configuration.Global.TileSize
			screen.DrawImage(assets.FloorImage.SubImage(
				image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
			).(*ebiten.Image), op)
		}
	}
}
