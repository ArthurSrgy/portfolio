package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case fromFileFloor:
		if configuration.Global.GenerateRandomFloor == true {
			f.InitWaveFunctionCollapse()
		} else {
			f.fullContent = ReadFloorFromFile(configuration.Global.FloorFile)
		}

	case quadTreeFloor:
		if configuration.Global.GenerateRandomFloor {
			if configuration.Global.GenerateVillage == false {
				f.quadtreeContent = quadtree.CreateRandomQuadtree()
			} else {
				f.quadtreeContent = quadtree.CreateRandomQuadtree()
			}
		} else {
			f.quadtreeContent = quadtree.MakeFromArray(ReadFloorFromFile(configuration.Global.FloorFile), 0, 0)

		}
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func ReadFloorFromFile(fileName string) (floorContent [][]int) {
	// Ouverture du fichier
	var myFile *os.File
	var err error

	myFile, err = os.Open(fileName)
	if err != nil { //Vérification de l'existence du fichier
		log.Fatal(err)
	}

	//Préparation à la lecture du fichier
	var scanner *bufio.Scanner = bufio.NewScanner(myFile)

	//Récup donnée
	for scanner.Scan() { //Parcours des lignes du document ouvert
		var tab_ligne []int = make([]int, len(string(scanner.Text())))
		for indice_caractère := 0; indice_caractère < len(string(scanner.Text())); indice_caractère++ {
			tab_ligne[indice_caractère], _ = strconv.Atoi(string(scanner.Text()[indice_caractère]))
		}
		floorContent = append(floorContent, tab_ligne) //faire sans un append
	}

	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	return floorContent
}

// ###############################################//
// ########### Generate Random Village ###########//
// ###############################################//
