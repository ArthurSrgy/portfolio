package daytime

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// update le temps de la journée
func UpdateDaytime() {
	// on ajoute une unité de temps à la journée
	// une journée vaut 7200 maximum (2 minutes si la speed est à 1)
	configuration.Global.Daytime = (configuration.Global.Daytime + configuration.Global.DaytimeSpeed) % 7200
	//fmt.Println(configuration.Global.Daytime)
}
