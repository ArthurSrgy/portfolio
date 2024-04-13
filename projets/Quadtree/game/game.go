package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particule"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/teleporteur"
)

// Game est le type permettant de représenter les données du jeu.
// Aucun champs n'est exporté pour le moment.
//
// Les champs non exportés sont :
//   - camera : la représentation de la caméra
//   - floor : la représentation du terrain
//   - character : la représentation du personnage
type Game struct {
	camera            camera.Camera
	floor             floor.Floor
	character         character.Character
	Liste_Empreinte   particule.FootprintList
	Pluie             particule.Pluie
	Liste_teleporteur teleporteur.Liste_teleporteur
}
