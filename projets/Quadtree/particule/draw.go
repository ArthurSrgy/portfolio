package particule

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image"
	"image/color"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (Flist *FootprintList) Draw(screen *ebiten.Image, PosCamX, PosCamY, CharaxInc, CharayInc, shift int) {

	// biais permettant de décaler les tuiles lors de la marche
	biais := 0
	if configuration.Global.CameraMode == 3 {
		// si la camera est fluide alors on augmente le biais
		biais = shift
	}
	// on parcours la liste d'empreinte
	for indice := 0; indice < len(Flist.List); indice++ {
		// raccourci vers l'empreinte
		Empreinte := Flist.List[indice]

		// position de l'empreinte par rapport à l'écran
		x := Empreinte.PosX - PosCamX + configuration.Global.ScreenCenterTileX
		y := Empreinte.PosY - PosCamY + configuration.Global.ScreenCenterTileY

		//on regarde si l'empreinte peut être dans l'écran
		if (x > -1 && x < configuration.Global.NumTileX+1) && (y > -1 && y < configuration.Global.NumTileY+1) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((x)*configuration.Global.TileSize-biais*CharaxInc), float64((y)*configuration.Global.TileSize-biais*CharayInc))
			//vector.StrokeLine(op,10,10,20,20,,false)

			// on cherche où se trouve l'empreinte dans l'image
			shiftX := (3 - (Empreinte.lifetime/30)%4) * configuration.Global.TileSize
			shiftY := Empreinte.direction * configuration.Global.TileSize

			//on affiche
			screen.DrawImage(assets.ParticuleImage.SubImage(
				image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
			).(*ebiten.Image), op)
		}
	}

}

func (Pluie *Pluie) Draw(screen *ebiten.Image, camX, camY, CharaxInc, CharayInc, shift int) {
	// biais permettant de décaler les tuiles lors de la marche
	biais := 0
	if configuration.Global.CameraMode == 3 {
		// si la camera est fluide alors on augmente le biais
		biais = shift
	}

	// ####### PARTIE PLUIE #######

	if configuration.Global.Météo == 1 {
		// lorsqu'il pleut on met un voile noire pour les nuages
		vector.DrawFilledRect(screen, 0, 0, float32(configuration.Global.NumTileX*configuration.Global.TileSize), float32(configuration.Global.NumTileY*configuration.Global.TileSize), color.NRGBA{0, 0, 0, 0x30}, false)
		// on parcours la liste de gouttes
		for indice := 0; indice < len(Pluie.List); indice++ {
			// raccourci vers l'empreinte
			goutte := Pluie.List[indice]

			// position de la case de la goutte par rapport à l'écran
			x := (goutte.CaseX-camX+configuration.Global.NumTileX/2)*configuration.Global.TileSize - biais*CharaxInc
			y := (goutte.CaseY-camY+configuration.Global.NumTileY/2)*configuration.Global.TileSize - biais*CharayInc

			//on regarde si la goutte peut être dans l'écran pour l'afficher
			if (x > -1 && int(x) < (configuration.Global.NumTileX+1)*configuration.Global.TileSize) && (y > -1 && int(y) < (configuration.Global.NumTileY+1)*configuration.Global.TileSize) {
				//on affiche la goutte
				vector.DrawFilledRect(screen, float32(x)+goutte.PosX, float32(y)+goutte.PosY, 1, float32(goutte.lifetime)/4, color.NRGBA{0, 0, 0xe0, 0xc0}, false)
			}
		}

		// ####### PARTIE NEIGE #######
	} else {
		// lorsqu'il neige on met un voile blanc
		vector.DrawFilledRect(screen, 0, 0, float32(configuration.Global.NumTileX*configuration.Global.TileSize), float32(configuration.Global.NumTileY*configuration.Global.TileSize), color.NRGBA{0xff, 0xff, 0xff, 0x30}, false)
		// on parcours la liste de gouttes
		for indice := 0; indice < len(Pluie.List); indice++ {
			// raccourci vers l'empreinte
			goutte := Pluie.List[indice]

			// position de la case de la goutte par rapport à l'écran
			x := (goutte.CaseX-camX+configuration.Global.NumTileX/2)*configuration.Global.TileSize - biais*CharaxInc
			y := (goutte.CaseY-camY+configuration.Global.NumTileY/2)*configuration.Global.TileSize - biais*CharayInc

			//on regarde si la goutte peut être dans l'écran pour l'afficher
			if (x > -1 && int(x) < (configuration.Global.NumTileX+1)*configuration.Global.TileSize) && (y > -1 && int(y) < (configuration.Global.NumTileY+1)*configuration.Global.TileSize) {
				//on affiche la goutte
				vector.DrawFilledCircle(screen, float32(x)+goutte.PosX, float32(y)+goutte.PosY, (float32(goutte.lifetime)/10)/4, color.NRGBA{0xff, 0xff, 0xff, 0xc0}, false)
			}
		}
	}
}
