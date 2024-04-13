package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
//   - betterContent: partie du terrain qui doit être affichée à l'écran mais avec 1 case en plus dans les 4 directions.
type Floor struct {
	content         [][]int
	fullContent     [][]int
	betterContent   [][]int
	quadtreeContent quadtree.Quadtree
}

// types d'affichage du terrain disponibles
const (
	gridFloor int = iota
	fromFileFloor
	quadTreeFloor
)

func (f *Floor) GetFullContentLen() (width, height int) {
	return len(f.fullContent[0]), len(f.fullContent)
}
