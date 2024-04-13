package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.




func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {

	blocking[0] = false
	blocking[1] = false
	blocking[2] = false
	blocking[3] = false

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY
	if !configuration.Global.DebugMode {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}

	if !configuration.Global.DebugMode && configuration.Global.Personnage == 0 {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1 || f.content[relativeYPos-1][relativeXPos] == 4 || f.content[relativeYPos-1][relativeXPos] == 2
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1 || f.content[relativeYPos][relativeXPos+1] == 4 || f.content[relativeYPos][relativeXPos+1] == 2
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1 || f.content[relativeYPos+1][relativeXPos] == 4 || f.content[relativeYPos+1][relativeXPos] == 2
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1 || f.content[relativeYPos][relativeXPos-1] == 4 || f.content[relativeYPos][relativeXPos-1] == 2
	}

	return blocking
}
