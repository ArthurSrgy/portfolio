package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.character.Init()
	g.camera.Init()
	g.floor.Init()
	g.Liste_Empreinte.Init()
	// on initialise le fullscreen (ou non)
	ebiten.SetFullscreen(configuration.Global.Fullscreen)
}
