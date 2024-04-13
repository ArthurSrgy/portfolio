package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:

		f.UpdateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:

		f.updateQuadtreeFloor(camXPos, camYPos)

	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) UpdateFromFileFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.NumTileX/2 + len(f.fullContent[0])/2
	topLeftY := camYPos - configuration.Global.NumTileY/2 + len(f.fullContent)/2
	if configuration.Global.CameraMode == 0 {
		topLeftX = camXPos - configuration.Global.NumTileX + configuration.Global.NumTileX%2
		topLeftY = camYPos - configuration.Global.NumTileY + configuration.Global.NumTileY%2
	}
	for abscisse := 0; abscisse < configuration.Global.NumTileX; abscisse++ {
		for ordonnee := 0; ordonnee < configuration.Global.NumTileY; ordonnee++ {
			//Booléen vérifiant si la case sélectionnée est pas hors de l'écran
			IsTooFar := topLeftY+ordonnee < 0 || topLeftY+ordonnee > len(f.fullContent)-1
			IsTooFar = IsTooFar || (topLeftX+abscisse < 0 || topLeftX+abscisse > len(f.fullContent[0])-1)
			if IsTooFar == true {
				f.content[ordonnee][abscisse] = -1
			} else {
				f.content[ordonnee][abscisse] = f.fullContent[topLeftY+ordonnee][topLeftX+abscisse]
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
	f.fullContent = quadtree.ShowQuad(f.quadtreeContent)
}

// Pour obtenir un content étendu de une case aux 4 directions.
// Sert pour la fusion des tuiles ou pour la caméra fluide
func (f *Floor) GetBetterContent(camXPos, camYPos int) {
	BetterNumTileX := configuration.Global.NumTileX + 2
	BetterNumTileY := configuration.Global.NumTileY + 2
	//on créer betterContent
	f.betterContent = make([][]int, BetterNumTileY)
	for i := range f.betterContent {
		f.betterContent[i] = make([]int, BetterNumTileX)
	}

	topLeftX := camXPos - BetterNumTileX/2 + len(f.fullContent[0])/2
	topLeftY := camYPos - BetterNumTileY/2 + len(f.fullContent)/2
	if configuration.Global.CameraMode == 0 {
		topLeftX = camXPos - BetterNumTileX + (BetterNumTileX)%2 + 1
		topLeftY = camYPos - BetterNumTileY + (BetterNumTileY)%2 + 1
	}
	for abscisse := 0; abscisse < BetterNumTileX; abscisse++ {
		for ordonnee := 0; ordonnee < BetterNumTileY; ordonnee++ {
			//Booléen vérifiant si la case sélectionnée est pas hors de l'écran
			IsTooFar := topLeftY+ordonnee < 0 || topLeftY+ordonnee > len(f.fullContent)-1
			IsTooFar = IsTooFar || (topLeftX+abscisse < 0 || topLeftX+abscisse > len(f.fullContent[0])-1)
			if IsTooFar == true {
				f.betterContent[ordonnee][abscisse] = -1
			} else {
				f.betterContent[ordonnee][abscisse] = f.fullContent[topLeftY+ordonnee][topLeftX+abscisse]
			}
		}
	}
}

// GetBestContent sert à obtenir un content étendu de deux cases
// aux 4 directions. Elle est utilisé lorsque la caméra est fluide
// et que l'on doit fusionner les tuiles. (permettant de voir les
// tuiles suivantes bien fusionné lors de notre marche.
// Oui le nom de la fonction est un peu ridicule.
func (f *Floor) GetBestContent(camXPos, camYPos int) [][]int {
	BetterNumTileX := configuration.Global.NumTileX + 4 //+4 car on a étendu la largeur de 4
	BetterNumTileY := configuration.Global.NumTileY + 4 //+4 car on a étendu la hauteur de 4
	//on créer betterContent
	bestContent := make([][]int, BetterNumTileY)
	for i := range bestContent {
		bestContent[i] = make([]int, BetterNumTileX)
	}

	topLeftX := camXPos - BetterNumTileX/2 + len(f.fullContent[0])/2
	topLeftY := camYPos - BetterNumTileY/2 + len(f.fullContent)/2
	if configuration.Global.CameraMode == 0 {
		topLeftX = camXPos - BetterNumTileX + (BetterNumTileX)%2 + 1
		topLeftY = camYPos - BetterNumTileY + (BetterNumTileY)%2 + 1
	}
	for abscisse := 0; abscisse < BetterNumTileX; abscisse++ {
		for ordonnee := 0; ordonnee < BetterNumTileY; ordonnee++ {
			//Booléen vérifiant si la case sélectionnée est pas hors de l'écran
			IsTooFar := topLeftY+ordonnee < 0 || topLeftY+ordonnee > len(f.fullContent)-1
			IsTooFar = IsTooFar || (topLeftX+abscisse < 0 || topLeftX+abscisse > len(f.fullContent[0])-1)
			if IsTooFar == true {
				bestContent[ordonnee][abscisse] = -1
			} else {
				bestContent[ordonnee][abscisse] = f.fullContent[topLeftY+ordonnee][topLeftX+abscisse]
			}
		}
	}
	return bestContent
}
