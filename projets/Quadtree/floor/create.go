package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"math/rand"
)

// Initialisation de l'algorithme WaveFonctionCollapse
func (f *Floor) InitWaveFunctionCollapse() {
	//on créer la seed du terrain
	if configuration.Global.SeedOn {
		rand.Seed(configuration.Global.Seed)
	}
	dimension := 32
	if dimension < 8 {
		dimension = 8
	}
	// on créer un fullContent de 32x32.
	// Et on le rempli de -1.
	f.fullContent = make([][]int, dimension)
	for i := range f.fullContent {
		f.fullContent[i] = make([]int, dimension)
		for j := range f.fullContent[i] {
			f.fullContent[i][j] = -1
		}
	}

	//on créer une liste de possibilités de toutes les cases
	var ListePossibilites [][][]int
	ListePossibilites = make([][][]int, dimension)
	for i := range ListePossibilites {
		ListePossibilites[i] = make([][]int, dimension)
		for j := range ListePossibilites[i] {
			// lors de sa création, chaques cases à comme possibilité d'être soit:
			// de la terre (0), du sable (1) ou de l'eau(4).
			ListePossibilites[i][j] = []int{0, 1, 4}
		}
	}
	//##########################//
	//on pose la première tuile.
	//ils s'agit soit de sable, soit de terre (soit 0 soit 1)
	valeur := rand.Intn(2)
	f.fullContent[dimension/2+dimension%2][dimension/2+dimension%2] = valeur
	f.fullContent[0][0] = 0
	//on modifie la liste de possibilités
	ListePossibilites[dimension/2+dimension%2][dimension/2+dimension%2] = []int{valeur}
	//on propage la modification
	ListePossibilites = Propagation(ListePossibilites)
	//on fait des étapes jusqu'à sa remplition
	for f.IsNotFull() {
		ListePossibilites = f.StepWaveFunctionCollapse(ListePossibilites)
	}
}

// Execution d'une étape de l'algorithme WaveFonctionCollapse
// Une étape est effectué en 3 partie:
//  1. Cell Selection : selection d'une cases avec le minimum d'entropie
//  2. Collapse Cell  : assigner une seule de ses possibilités au hasard
//  3. Propagation	  : on propage les entropies aux cellules proches
func (f *Floor) StepWaveFunctionCollapse(ListePossibilites [][][]int) [][][]int {
	// 1) Cell Selection
	// on sélectionne une case d'entropie minimum au hasard
	ListeCaseMin := f.SelectMinimumEntropie(ListePossibilites)
	SelectedPosition := ListeCaseMin[rand.Intn(len(ListeCaseMin))]
	y, x := SelectedPosition[0], SelectedPosition[1] // les coordonnées la case sélectionnés
	SelectedPosibilities := ListePossibilites[y][x]  //les posibilités de la case selectionné
	// 2) Collapse Cell
	//on modifie la valeur de possibilité et la valeur f.fullcontent de cette case.
	valeur := SelectedPosibilities[rand.Intn(len(SelectedPosibilities))]
	f.fullContent[y][x] = valeur
	ListePossibilites[y][x] = []int{valeur}
	// 3) Propagation
	//on propage les possibilités aux autres.
	ListePossibilites = Propagation(ListePossibilites)

	//on vérifie si l'on peut définir une des cases en même temps.
	f.UpdateTerrain(ListePossibilites)
	ListePossibilites = Propagation(ListePossibilites)
	ListePossibilites = Propagation(ListePossibilites)
	ListePossibilites = Propagation(ListePossibilites)

	return ListePossibilites
}

func IsIn(valeur int, tableau []int) bool {
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == valeur {
			return true
		}
	}
	return false
}

// IsNotFull est une fonction qui vérifie s'il reste bien des -1 dans fullcontent.
// Elle permet de créer une boucle jusqu'à ce que le terrain soit rempli
func (f *Floor) IsNotFull() bool {
	for i := range f.fullContent {
		for j := range f.fullContent[i] {
			if f.fullContent[i][j] == -1 {
				return true
			}
		}
	}
	return false
}

// Selectionne dans un tableau les positions des cases avec le
// minimum d'entropie.
// (l'une d'entre elle sera enusite sélectionné au hasard)
func (f *Floor) SelectMinimumEntropie(ListPossibilites [][][]int) (TabCaseMin [][2]int) {
	// on parcours toutes les entropie.
	min := 3
	for y := range ListPossibilites {
		for x := range ListPossibilites[y] {
			// si l'entropie de la case (vide) est plus petites que le minimum trouvé jusqu'à présent.
			if f.fullContent[y][x] == -1 && len(ListPossibilites[y][x]) < min {
				TabCaseMin = [][2]int{{y, x}}
			} else if f.fullContent[y][x] == -1 && len(ListPossibilites[y][x]) == min {
				TabCaseMin = append(TabCaseMin, [2]int{y, x})
			}
		}
	}
	return TabCaseMin
}

func Propagation(ListPossibilites [][][]int) [][][]int {
	//on parcours toutes les cases
	for y := range ListPossibilites {
		for x := range ListPossibilites[y] {
			var Voisins [4][]int //contient les tableaux des propriétés voisines
			//on ajoute les possibilités du voisin gauche
			if x-1 >= 0 {
				Voisins[0] = ListPossibilites[y][x-1] //s'il existe
			} else {
				Voisins[0] = []int{} //s'il n'existe pas on crée une liste de d'aucunnes probabilités
			}
			//on ajoute les possibilités du voisin droit
			if x+1 > len(ListPossibilites[0]) {
				Voisins[1] = ListPossibilites[y][x+1] //s'il existe
			} else {
				Voisins[1] = []int{} //s'il n'existe pas on crée une liste de d'aucunnes probabilités
			}
			//on ajoute les possibilités du voisin du haut
			if y-1 >= 0 {
				Voisins[2] = ListPossibilites[y-1][x] //s'il existe
			} else {
				Voisins[2] = []int{} //s'il n'existe pas on crée une liste de d'aucunnes probabilités
			}
			//on ajoute les possibilités du voisin du bas
			if y+1 > len(ListPossibilites[0]) {
				Voisins[3] = ListPossibilites[y+1][x] //s'il existe
			} else {
				Voisins[3] = []int{} //s'il n'existe pas on crée une liste de d'aucunnes probabilités
			}
			//on parcours les listes de probabilités.
			for indice := range Voisins {
				if len(Voisins[indice]) < 3 { // si mon voisin a des possibilités modifiés
					if IsIn(1, Voisins[indice]) { //s'il a un voisin sable
						ListPossibilites[y][x] = enlever_possibilités(4, ListPossibilites[y][x]) //on enlève la possibilité d'être de l'eau
					}
					if IsIn(4, Voisins[indice]) { //s'il a un voisin eau
						ListPossibilites[y][x] = enlever_possibilités(1, ListPossibilites[y][x]) //on enlève la possibilité d'être du sable
					}
				}
			}
		}
	}
	return ListPossibilites

}

func enlever_possibilités(valeur int, Posibilites []int) []int {
	var TabRetour []int
	for indice := range Posibilites {
		if Posibilites[indice] != valeur {
			TabRetour = append(TabRetour, Posibilites[indice])
		}
	}
	return Posibilites
}

// Fonction qui change le terrain lorsque sa liste de probabilité ne contient qu'un seul élément.
func (f *Floor) UpdateTerrain(ListPossibilites [][][]int) {
	for y := range f.fullContent {
		for x := range f.fullContent[y] {
			//s'il on a pas modifié le fullcontent et qu'une des case n'à qu'une seule possibilité.
			if f.fullContent[y][x] == -1 && len(ListPossibilites[y][x]) == 1 {
				f.fullContent[y][x] = ListPossibilites[y][x][0] // on attribue la seule possibilité à fullContent.
			}
		}
	}
}
