package floor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image"
	"math"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, PosCamX, PosCamY int, Character character.Character) {
	// xInc et yInc du personnage
	// (impossible d'utiliser c.xInc ou c.yInc donc une fonction à été crée)
	CharaxInc, CharayInc := Character.GetCoordInc()
	// biais permettant de décaler les tuiles lors de la marche
	biais := 0
	if configuration.Global.CameraMode == 3 {
		// si la camera est fluide alors on augmente le biais
		biais = Character.GetShift()
	}
	// on récupère le BetterContent du sol
	f.GetBetterContent(PosCamX, PosCamY)

	if configuration.Global.FuseFloor {
		for y := 1; y < len(f.content)+1; y++ {
			for x := 1; x < len(f.content[0])+1; x++ {
				switch {
				case f.betterContent[y][x] == 1:
					f.FuseSand(screen, f.betterContent, x, y, biais, CharaxInc, CharayInc, false)
				case f.betterContent[y][x] == 4:
					f.FuseWater(screen, f.betterContent, x, y, biais, CharaxInc, CharayInc, false, PosCamX, PosCamY)
				case f.betterContent[y][x] == 2:
					f.FuseWall(screen, f.betterContent, x, y, biais, CharaxInc, CharayInc, false)
				case f.betterContent[y][x] != -1:
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64((x-1)*configuration.Global.TileSize-biais*CharaxInc), float64((y-1)*configuration.Global.TileSize-biais*CharayInc))

					shiftX := f.betterContent[y][x] * configuration.Global.TileSize
					//fmt.Println(" f.content[", y, "][", x, "]")
					shiftY := ((configuration.Global.FrameCount / 4) % 16) * configuration.Global.TileSize

					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}
			}
			//###### HORIZONTAL ######
			// s'il on bouge horizontalement et que la caméra doit être fluide
			// on doit rajouter des tuiles fusionnés à gauche ou à droite
			if CharaxInc != 0 && configuration.Global.CameraMode == 3 {
				// on récupère le BestContent du sol
				BestContent := f.GetBestContent(PosCamX, PosCamY)
				//le cas où l'on bouge à droite
				départ := 0
				x := len(BestContent[0]) - 4
				//le cas où l'on bouge à gauche
				if CharaxInc == -1 {
					départ = -1
					x = -1
				}
				for y := départ; y < len(f.content)+1; y++ {
					switch {
					case BestContent[y+2][x+2] == 1:
						f.FuseSand(screen, BestContent, x+2, y+2, biais, CharaxInc, CharayInc, true)
					case BestContent[y+2][x+2] == 4:
						f.FuseWater(screen, BestContent, x+2, y+2, biais, CharaxInc, CharayInc, true, PosCamX, PosCamY)
					case BestContent[y+2][x+2] == 2:
						f.FuseWall(screen, BestContent, x+2, y+2, biais, CharaxInc, CharayInc, true)
					default:
						op := &ebiten.DrawImageOptions{}
						op.GeoM.Translate(float64(x*configuration.Global.TileSize-biais*CharaxInc), float64(y*configuration.Global.TileSize))
						shiftX := BestContent[y+2][x+2] * configuration.Global.TileSize
						// shiftY pour l'eau et ses animations
						shiftY := ((configuration.Global.FrameCount / 4) % configuration.Global.TileSize) * configuration.Global.TileSize

						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
						).(*ebiten.Image), op)
					}
				}
			}
			//###### VERTICAL ######
			// s'il on bouge verticalement et que la caméra doit être fluide
			// on doit rajouter des tuiles fusionné en bas ou en haut
			if CharayInc != 0 && configuration.Global.CameraMode == 3 {
				// on récupère le BestContent du sol
				BestContent := f.GetBestContent(PosCamX, PosCamY)
				//le cas où l'on bouge en bas
				départ := 0
				y := len(BestContent) - 4
				//le cas où l'on bouge en haut
				if CharayInc == -1 {
					départ = -1
					y = -1
				}
				for x := départ; x < len(f.content[0])+1; x++ {
					switch {
					case BestContent[y+2][x+2] == 1:
						f.FuseSand(screen, BestContent, x+2, y+2, biais, CharaxInc, CharayInc, true)
					case BestContent[y+2][x+2] == 4:
						f.FuseWater(screen, BestContent, x+2, y+2, biais, CharaxInc, CharayInc, true, PosCamX, PosCamY)
					case BestContent[y+2][x+2] == 2:
						f.FuseWall(screen, BestContent, x+2, y+2, biais, CharaxInc, CharayInc, true)
					default:
						op := &ebiten.DrawImageOptions{}
						op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize-biais*CharayInc))
						shiftX := BestContent[y+2][x+2] * configuration.Global.TileSize
						// shiftY pour l'eau et ses animations
						shiftY := 0

						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
						).(*ebiten.Image), op)
					}
				}
			}
		}
	} else {
		//####### s'il on ne fusionne pas les tuiles ######
		for y := range f.content {
			for x := range f.content[y] {
				if f.content[y][x] != -1 {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*configuration.Global.TileSize-biais*CharaxInc), float64(y*configuration.Global.TileSize-biais*CharayInc))

					shiftX := f.content[y][x] * configuration.Global.TileSize
					// shiftY pour l'eau et ses animations
					shiftY := 0
					if configuration.Global.AnimeFloor {
						shiftY = ((configuration.Global.FrameCount / 4) % configuration.Global.TileSize) * configuration.Global.TileSize
					}

					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}
			}
		}
		//###### HORIZONTAL ######
		// s'il on bouge horizontalement et que la caméra doit être fluide
		// on doit rajouter des tuiles à gauche ou à droite
		if CharaxInc != 0 && configuration.Global.CameraMode == 3 {
			//le cas où l'on bouge à droite
			départ := 0
			x := len(f.betterContent[0]) - 2
			//le cas où l'on bouge à gauche
			if CharaxInc == -1 {
				départ = -1
				x = -1
			}
			for y := départ; y < len(f.content)+1; y++ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize-biais*CharaxInc), float64(y*configuration.Global.TileSize))
				shiftX := f.betterContent[y+1][x+1] * configuration.Global.TileSize
				// shiftY pour l'eau et ses animations
				shiftY := 0
				if configuration.Global.AnimeFloor {
					shiftY = ((configuration.Global.FrameCount / 4) % configuration.Global.TileSize) * configuration.Global.TileSize
				}

				screen.DrawImage(assets.FloorImage.SubImage(
					image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
				).(*ebiten.Image), op)
			}
		}
		//###### VERTICAL ######
		// s'il on bouge verticalement et que la caméra doit être fluide
		// on doit rajouter des tuiles en bas ou en haut
		if CharayInc != 0 && configuration.Global.CameraMode == 3 {
			//le cas où l'on bouge en bas
			départ := 0
			y := len(f.betterContent) - 2
			//le cas où l'on bouge en haut
			if CharayInc == -1 {
				départ = -1
				y = -1
			}
			for x := départ; x < len(f.content[0])+1; x++ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize-biais*CharayInc))
				shiftX := f.betterContent[y+1][x+1] * configuration.Global.TileSize
				// shiftY pour l'eau et ses animations
				shiftY := 0
				if configuration.Global.AnimeFloor {
					shiftY = ((configuration.Global.FrameCount / 4) % configuration.Global.TileSize) * configuration.Global.TileSize
				}

				screen.DrawImage(assets.FloorImage.SubImage(
					image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
				).(*ebiten.Image), op)
			}
		}
	}
}

// Fonction servant à "fusionner" les tuiles de "sable"/"chemin" selon les 8 tuiles autours d'elle.
// #Entrées:
//
//	 bettercontent 	 [][]int: content plus grand que de une case par coté, permettant de fuisonner les cases au bord de content
//		Xcase, Ycase  	  	 int: coordonnée de la case relative au content
//		biais				 int: décalage des cases lorsque la caméra et fluide
//	 CharaxInc, CharayInc int: donnent la direction où le joueur avance
//		loading				bool: bolléen permettant de savoir si la case à fusionner est une case à charger (lorsque CameraMode = 3)
func (f Floor) FuseSand(screen *ebiten.Image, bettercontent [][]int, Xcase, Ycase, biais, CharaxInc, CharayInc int, loading bool) {
	//tableau contenant la case et les 8 cases proches pour ne pas gener les test.
	var ToFuse [3][3]int = [3][3]int{
		{bettercontent[Ycase-1][Xcase-1], bettercontent[Ycase-1][Xcase], bettercontent[Ycase-1][Xcase+1]},
		{bettercontent[Ycase][Xcase-1], bettercontent[Ycase][Xcase], bettercontent[Ycase][Xcase+1]},
		{bettercontent[Ycase+1][Xcase-1], bettercontent[Ycase+1][Xcase], bettercontent[Ycase+1][Xcase+1]},
	}
	//raccourci taille tiles
	LenTiles := configuration.Global.TileSize
	op := &ebiten.DrawImageOptions{}

	if loading {
		op.GeoM.Translate(float64((Xcase-2)*configuration.Global.TileSize-biais*CharaxInc), float64((Ycase-2)*configuration.Global.TileSize-biais*CharayInc))

	} else {
		op.GeoM.Translate(float64((Xcase-1)*configuration.Global.TileSize-biais*CharaxInc), float64((Ycase-1)*configuration.Global.TileSize-biais*CharayInc))
	}
	//coordonnées de la Tile de base
	shiftX := 9 * LenTiles
	shiftY := LenTiles

	//On test toutes les possibilités de chemins différents (42 environ...)
	switch {
	//4 coins (1 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 1
	//coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 5
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 5
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 4
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 4
		//2 coins (6 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 8
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 9
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 8
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 8
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 9
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 9
		//3 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 4
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 4
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 5
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 5
		//bordure (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 0}, {1, 0, 0}, {-1, 0, 0}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, -1}, {0, 0, 1}, {0, 0, -1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 7
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 7
		//bordure largeur avec 1 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 1}, {1, 0, 0}, {-1, 0, 0}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 11
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 0}, {1, 0, 0}, {-1, 0, 1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 10
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, -1}, {0, 0, 1}, {0, 0, -1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 11
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, -1}, {0, 0, 1}, {1, 0, -1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 10
		//bordure hauteur avec 1 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 10
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 10
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 11
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 11
		//bordure avec 2 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 2
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 1}, {1, 0, 0}, {-1, 0, 1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 1
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, -1}, {0, 0, 1}, {1, 0, -1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 1
		//impasses (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, -1}, {1, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 2
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, -1}, {1, 0, 1}, {-1, 0, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 1
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 1}, {-1, 0, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 0
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 5
		shiftY = LenTiles * 0
		//4 coins petites tailles
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 0}, {-1, 0, 1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 0
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 1}, {1, 0, -1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 0
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, -1}, {0, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 2
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 1}, {1, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 4
		shiftY = LenTiles * 2
		//4 coins remplies
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 0}, {-1, 0, 0}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 1}, {0, 0, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, -1}, {0, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 7
		shiftY = LenTiles * 7
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 0}, {1, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 6
		shiftY = LenTiles * 7
	}
	screen.DrawImage(assets.FusedTilesImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)
}

// même fonction que FuseSand mais pour l'eau
// Fonction servant à "fusionner" les tuiles de d'eau selon les 8 tuiles autours d'elle.
// Elle sert aussi à l'animation des vaguelettes lorsque configuration.Global.AnimeFloor = true.
// #Entrées:
//
//	bettercontent 	 [][]int: content plus grand que de une case par coté, permettant de fuisonner les cases au bord de content
//	Xcase, Ycase  	  	 int: coordonnée de la case relative au content
//	biais				 int: décalage des cases lorsque la caméra et fluide
//	CharaxInc, CharayInc int: donnent la direction où le joueur avance
//	loading				bool: bolléen permettant de savoir si la case à fusionner est une case à charger (lorsque CameraMode = 3)
//	PosCamX, PosCamY 	 int: position de la caméra, sert lors de l'animation des vaguelette pour faire un lien entre position rellative à content et position relative à la caméra.
func (f Floor) FuseWater(screen *ebiten.Image, bettercontent [][]int, Xcase, Ycase, biais, CharaxInc, CharayInc int, loading bool, PosCamX, PosCamY int) {
	//raccourci taille tiles
	LenTiles := configuration.Global.TileSize

	op := &ebiten.DrawImageOptions{}

	if loading {
		op.GeoM.Translate(float64((Xcase-2)*configuration.Global.TileSize-biais*CharaxInc), float64((Ycase-2)*configuration.Global.TileSize-biais*CharayInc))

	} else {
		op.GeoM.Translate(float64((Xcase-1)*configuration.Global.TileSize-biais*CharaxInc), float64((Ycase-1)*configuration.Global.TileSize-biais*CharayInc))
	}

	//coordonnées de la Tile de base
	shiftX := LenTiles * 22
	shiftY := LenTiles * 12
	//tableau contenant la case et les 8 cases proches pour ne pas gener les test.
	var ToFuse [3][3]int = [3][3]int{
		{bettercontent[Ycase-1][Xcase-1], bettercontent[Ycase-1][Xcase], bettercontent[Ycase-1][Xcase+1]},
		{bettercontent[Ycase][Xcase-1], bettercontent[Ycase][Xcase], bettercontent[Ycase][Xcase+1]},
		{bettercontent[Ycase+1][Xcase-1], bettercontent[Ycase+1][Xcase], bettercontent[Ycase+1][Xcase+1]},
	}
	//On test toutes les possibilités d'eau différents (42 environ...)
	switch {
	//4 coins (1 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 1
	//petits coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 5
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 5
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 4
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 4
		//2 petits coins (6 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 8
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 9
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 8
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 8
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 9
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 9
		//bordure (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 0}, {1, 0, 0}, {-1, 0, 0}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, -1}, {0, 0, 1}, {0, 0, -1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 7
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {0, 0, 0}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 0}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 7
		//bordure largeur avec 1 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 1}, {1, 0, 0}, {-1, 0, 0}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 11
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 0}, {1, 0, 0}, {-1, 0, 1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 10
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, -1}, {0, 0, 1}, {0, 0, -1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 11
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, -1}, {0, 0, 1}, {1, 0, -1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 10
		//bordure hauteur avec 1 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 10
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 10
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 11
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 11
		//bordure avec 2 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 2
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 1}, {1, 0, 0}, {-1, 0, 1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 1
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, -1}, {0, 0, 1}, {1, 0, -1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 1
		//impasses (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, -1}, {1, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 2
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, -1}, {1, 0, 1}, {-1, 0, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 1
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 1}, {-1, 0, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 0
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 3
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 0
		//3 coins (4 poss)
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 0}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 4
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 1}, {0, 0, 0}, {0, 0, 1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 4
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, 0}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 5
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, 1}, {0, 0, 0}, {1, 0, 1}}):
		shiftX = LenTiles * 21
		shiftY = LenTiles * 5
		//4 coins petites tailles
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 0}, {-1, 0, 1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 0
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 1}, {1, 0, -1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 0
	case EstDisposeComme(ToFuse, [3][3]int{{1, 0, -1}, {0, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 2
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 1}, {1, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 20
		shiftY = LenTiles * 2
		//4 coins remplies
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {1, 0, 0}, {-1, 0, 0}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 1, -1}, {0, 0, 1}, {0, 0, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 6
	case EstDisposeComme(ToFuse, [3][3]int{{0, 0, -1}, {0, 0, 1}, {-1, 1, -1}}):
		shiftX = LenTiles * 23
		shiftY = LenTiles * 7
	case EstDisposeComme(ToFuse, [3][3]int{{-1, 0, 0}, {1, 0, 0}, {-1, 1, -1}}):
		shiftX = LenTiles * 22
		shiftY = LenTiles * 7
	}
	screen.DrawImage(assets.FusedTilesImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)

	// animation des vaguelette de l'eau
	if configuration.Global.AnimeFloor {
		shiftY = LenTiles*(6+((configuration.Global.FrameCount/8+int(math.Abs(float64((Ycase+Xcase+PosCamY+PosCamX)%2))))%4)) - (Ycase+Xcase+PosCamY+PosCamX)%2

		shiftX = LenTiles * 0
		screen.DrawImage(assets.FusedTilesImage.SubImage(
			image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}
}

func (f Floor) FuseWall(screen *ebiten.Image, BtrCont [][]int, Xcase, Ycase, biais, CharaxInc, CharayInc int, loading bool) {
	//raccourci taille tiles
	LenTiles := configuration.Global.TileSize
	op := &ebiten.DrawImageOptions{}

	if loading {
		op.GeoM.Translate(float64((Xcase-2)*configuration.Global.TileSize-biais*CharaxInc), float64((Ycase-2)*configuration.Global.TileSize-biais*CharayInc))

	} else {
		op.GeoM.Translate(float64((Xcase-1)*configuration.Global.TileSize-biais*CharaxInc), float64((Ycase-1)*configuration.Global.TileSize-biais*CharayInc))
	}
	//on affiche la tile de base
	shiftX := LenTiles * 29
	shiftY := LenTiles * 1
	screen.DrawImage(assets.FusedTilesImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)

	//on test toutes les possibilités
	// y a t-il un mur à coté ? si non, on met des limites.
	if BtrCont[Ycase][Xcase+1] != 2 {
		screen.DrawImage(assets.FusedTilesImage.SubImage(
			image.Rect(0, LenTiles*11, 0+configuration.Global.TileSize, LenTiles*11+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}
	if BtrCont[Ycase+1][Xcase] != 2 {
		screen.DrawImage(assets.FusedTilesImage.SubImage(
			image.Rect(0, LenTiles*10, 0+configuration.Global.TileSize, LenTiles*10+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}
	if BtrCont[Ycase-1][Xcase] != 2 {
		screen.DrawImage(assets.FusedTilesImage.SubImage(
			image.Rect(LenTiles, LenTiles*11, LenTiles+configuration.Global.TileSize, LenTiles*11+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}
	if BtrCont[Ycase][Xcase-1] != 2 {
		screen.DrawImage(assets.FusedTilesImage.SubImage(
			image.Rect(LenTiles, LenTiles*10, LenTiles+configuration.Global.TileSize, LenTiles*10+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}

}

// Compare les tableau pour savoir s'il sont dans la même dispossition.
// La fonction étant utilisé pour tester tout types de tuile,
// les valeurs de tabtest sont utilisé ainsi:
// -1 pour des cases qui ne nous interresse pas
// 0 pour les cases devants être comme celle testé
// 1 pour dire qu'elle ne doit pas être comme les cases testé
func EstDisposeComme(tab, tabtest [3][3]int) bool {
	ok := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tabtest[i][j] == 0 {
				ok = ok && tab[1][1] == tab[i][j]
			} else if tabtest[i][j] == 1 {
				ok = ok && tab[1][1] != tab[i][j]
			}
			if !ok {
				return false
			}
		}
	}
	return ok
}
