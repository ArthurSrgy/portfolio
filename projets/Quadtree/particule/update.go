package particule

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// update Liste Empreinte
func (Flist *FootprintList) UpdateFootprint() {
	//on parcours toutes les empreintes
	for indice := 0; indice < len(Flist.List); indice++ {
		//on leurs baisse leurs temps de vie
		Flist.List[indice].lifetime--
		//si l'une d'entre elle est à 0
		if Flist.List[indice].lifetime <= 0 {
			//on la supprime
			Flist.List = append(Flist.List[:indice], Flist.List[indice+1:]...)
		}
	}
}

// update Gouttes de Pluies
func (Pluie *Pluie) UpdatePluie(TopLeftX, TopLeftY int) {
	// minimum de gouttes
	min := (configuration.Global.NumTileX * configuration.Global.NumTileY) / 8
	//s'il on a moins de 20 gouttes à l'écran on en rajoute
	for len(Pluie.List) < min {
		if configuration.Global.Météo == 1 {
			Pluie.CreateGoutte(TopLeftX, TopLeftY)
		} else {
			Pluie.CreateFlocons(TopLeftX, TopLeftY)
		}

	}

	//si la météo est à 1, on fait l'update des gouttes de pluies.
	if configuration.Global.Météo == 1 {
		//on parcours toutes les gouttes
		for indice := 0; indice < len(Pluie.List); indice++ {
			//on les déplacent
			Pluie.List[indice].PosY += Pluie.List[indice].speed

			// si PosY (position en pyxel relative à la case) est plus grand que 16
			if Pluie.List[indice].PosY/16 > 1 {
				//on ajoute 1 à la case de la goutte
				Pluie.List[indice].CaseY++
				Pluie.List[indice].PosY = float32(int(Pluie.List[indice].PosY) % 16)
			}

			//on leurs baisse leurs temps de vie
			Pluie.List[indice].lifetime--

			//si l'une d'entre elle est à 0
			if Pluie.List[indice].lifetime <= 0 {
				//on la supprime
				Pluie.List = append(Pluie.List[:indice], Pluie.List[indice+1:]...)
			}

		}

		//si la météo est à 2, on fait l'update des flocons de neige.
	} else if configuration.Global.Météo == 2 {
		//on parcours tout les flocons
		for indice := 0; indice < len(Pluie.List); indice++ {
			// On change leurs positions
			Pluie.List[indice].PosY += Pluie.List[indice].speed
			Pluie.List[indice].PosX -= Pluie.List[indice].speed * 0.8

			if Pluie.List[indice].PosY/16 > 1 {
				Pluie.List[indice].CaseY++
				Pluie.List[indice].PosY = float32(int(Pluie.List[indice].PosY) % 16)
			}
			if Pluie.List[indice].PosX/16 > 1 {
				Pluie.List[indice].CaseX++
				Pluie.List[indice].PosX = float32(int(Pluie.List[indice].PosX) % 16)
			}

			//on leurs baisse leurs temps de vie
			Pluie.List[indice].lifetime--

			//si l'un d'entre eux est à 0
			if Pluie.List[indice].lifetime <= 0 {
				//on le supprime
				Pluie.List = append(Pluie.List[:indice], Pluie.List[indice+1:]...)
			}

		}
	}
}
