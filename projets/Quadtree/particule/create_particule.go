package particule

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"math/rand"
)

// créer une empreinte
func (Flist *FootprintList) CreateFootPrint(PersoPosX, PersoPosY, direction int) {
	//on créer des empreinte seulement si le personnage n'es pas un fantome
	if configuration.Global.Personnage == 0 {
		Empreinte := footprint{PosX: PersoPosX, PosY: PersoPosY, lifetime: 120, direction: direction}
		Flist.List = append(Flist.List, Empreinte)
	}
}

// créer une goutte de PLuie
func (p *Pluie) CreateGoutte(TopLeftX, TopLeftY int) {
	//on créer une seule goutte à une position aléatoire
	goutte := Goutte{
		CaseX: TopLeftX + rand.Intn(configuration.Global.NumTileX+1),
		CaseY: TopLeftY + rand.Intn(configuration.Global.NumTileY+1) - 3,
		PosX:  float32(rand.Intn(16)),
		PosY:  float32(rand.Intn(16)),
		speed: 3 + float32(rand.Intn(20))/10,
	}
	goutte.lifetime = goutte.CaseY - TopLeftY + configuration.Global.NumTileY
	p.List = append(p.List, goutte)
}

// créer un flocon de neige
// (la vitesse, la position X de départ ainsi que le temps de vie diffère
func (p *Pluie) CreateFlocons(TopLeftX, TopLeftY int) {
	//on créer une seule goutte à une position aléatoire
	flocon := Goutte{
		CaseX: TopLeftX + rand.Intn(configuration.Global.NumTileX+1) + 3,
		CaseY: TopLeftY + rand.Intn(configuration.Global.NumTileY+1) - 3,
		PosX:  float32(rand.Intn(16)),
		PosY:  float32(rand.Intn(16)),
		speed: 1 + float32(rand.Intn(20))/10,
	}
	flocon.lifetime = 120 + rand.Intn(60)
	p.List = append(p.List, flocon)
}
