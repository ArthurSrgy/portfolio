package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/daytime"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {
	//on augmente le compte des frames de 1 (en le limitant à 60)
	configuration.Global.FrameCount = (configuration.Global.FrameCount + 1) % 60

	//on fait avancer la journée si le cycle est activé
	if configuration.Global.DaytimeCycle {
		daytime.UpdateDaytime()
	}

	//##Update des touches##

	//touche Debug
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}
	//touche changement perso
	if inpututil.IsKeyJustPressed(ebiten.KeyP) && configuration.Global.SwitchPersoOn {
		configuration.Global.Personnage++
		configuration.Global.Personnage = configuration.Global.Personnage % 2
	}
	//touche pour changer le fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		configuration.Global.Fullscreen = !configuration.Global.Fullscreen
		ebiten.SetFullscreen(configuration.Global.Fullscreen)
	}
	// touche utiliser / placer téléporteur
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		if g.Liste_teleporteur.TPA != nil && g.Liste_teleporteur.TPB != nil {
			if (g.Liste_teleporteur.TPA.X == g.character.X && g.Liste_teleporteur.TPA.Y == g.character.Y) || (g.Liste_teleporteur.TPB.X == g.character.X && g.Liste_teleporteur.TPB.Y == g.character.Y) {
				g.character.X, g.character.Y, g.camera.X, g.camera.Y = g.Liste_teleporteur.Deplacement(g.character, g.camera)
			} else {
				g.Liste_teleporteur.CreateTeleporteur(g.character)
			}
		} else {
			g.Liste_teleporteur.CreateTeleporteur(g.character)
		}
	}

	//##Update des différents attribut de g##

	// mise à jour de la liste d'empreinte via l'update du personnage
	g.Liste_Empreinte = g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y), g.Liste_Empreinte)
	if configuration.Global.ParticuleOn {
		g.Liste_Empreinte.UpdateFootprint()
	}
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)

	//mise à jour de la pluie si  activé
	if configuration.Global.Météo == 1 || configuration.Global.Météo == 2 {

		//coordonnée de la case en haut à gauche
		topLeftX := g.camera.X - configuration.Global.NumTileX/2
		topLeftY := g.camera.Y - configuration.Global.NumTileY/2
		//fmt.Println("topLeft:", topLeftY, topLeftX)
		g.Pluie.UpdatePluie(topLeftX, topLeftY)
	}

	return nil
}
