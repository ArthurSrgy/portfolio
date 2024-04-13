package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int, TLX, TLY int) (q Quadtree) {
	return Quadtree{
		width:  len(floorContent[0]),
		height: len(floorContent),
		root:   makeNodeFromArray(TLX, TLY, len(floorContent[0]), len(floorContent), floorContent),
	}
}

func makeNodeFromArray(topLeftX, topLeftY, width, height int, floorContent [][]int) *node {
	if len(floorContent) == 0 || len(floorContent[0]) == 0 || floorContent[0][0] == -1 {
		return nil
	}

	firstElement := floorContent[0][0]
	IsAllSame := true
	// Regarde si tout les éléments sont les mêmes (cas de base)
	for i := 0; i < len(floorContent); i++ {
		for j := 0; j < len(floorContent[0]); j++ {
			if floorContent[i][j] != firstElement {
				IsAllSame = false // s'il on a un cas contraire on break
				break
			}
		}
	}
	//si tout est pareil on renvoie le node
	if IsAllSame {
		return &node{
			topLeftX:        topLeftX,
			topLeftY:        topLeftY,
			width:           width,
			height:          height,
			content:         firstElement,
			topLeftNode:     nil,
			topRightNode:    nil,
			bottomLeftNode:  nil,
			bottomRightNode: nil,
		}
	}
	//sinon on cherche à diviser le tableau en 4 (de parts égales)
	Xmiddle := topLeftX + width/2 + width%2
	Ymiddle := topLeftY + height/2 + height%2
	tabDivisons := divideFloorContent(floorContent)
	//et on retourne un quadtree le contenant
	return &node{
		topLeftX:        topLeftX,
		topLeftY:        topLeftY,
		width:           width,
		height:          height,
		content:         -1, // valeur par defaut (/!\ pas une case vide /!\)
		topLeftNode:     makeNodeFromArray(topLeftX, topLeftY, width/2+width%2, height/2+height%2, tabDivisons[0]),
		topRightNode:    makeNodeFromArray(Xmiddle, topLeftY, width/2, height/2+height%2, tabDivisons[1]),
		bottomLeftNode:  makeNodeFromArray(topLeftX, Ymiddle, width/2+width%2, height/2, tabDivisons[2]),
		bottomRightNode: makeNodeFromArray(Xmiddle, Ymiddle, width/2, height/2, tabDivisons[3]),
	}
}

func divideFloorContent(floorContent [][]int) [][][]int {
	HauteurMiddle := len(floorContent)/2 + len(floorContent)%2
	LargeurMiddle := len(floorContent[0])/2 + len(floorContent[0])%2
	//on créer les tableaux de tableaux que l'on renvera
	var tabHautGauche [][]int
	var tabHautDroit [][]int
	var tabBasGauche [][]int
	var tabBasDroit [][]int
	//on ajoute les éléments des 4 parties aux 4 tableaux de tableaux respectifs
	for y := 0; y < len(floorContent); y++ {
		if y < HauteurMiddle {
			tabHautGauche = append(tabHautGauche, floorContent[y][:LargeurMiddle])
			tabHautDroit = append(tabHautDroit, floorContent[y][LargeurMiddle:])
		} else {
			tabBasGauche = append(tabBasGauche, floorContent[y][:LargeurMiddle])
			tabBasDroit = append(tabBasDroit, floorContent[y][LargeurMiddle:])
		}
	}
	//on ajoute des -1 si les tableaux de tableaux ne sont pas carrée...
	tabHautGauche = ajout1neg(tabHautGauche, HauteurMiddle, LargeurMiddle)
	tabHautDroit = ajout1neg(tabHautDroit, HauteurMiddle, LargeurMiddle-len(floorContent[0])%2)
	tabBasGauche = ajout1neg(tabBasGauche, HauteurMiddle-len(floorContent)%2, LargeurMiddle)
	tabBasDroit = ajout1neg(tabBasDroit, HauteurMiddle-len(floorContent)%2, LargeurMiddle-len(floorContent[0])%2)
	//on renvoye les 4 dans un seul tableau
	return [][][]int{tabHautGauche, tabHautDroit, tabBasGauche, tabBasDroit}
}

func ajout1neg(tabEntree [][]int, midy, midx int) (tabSortie [][]int) {
	tabSortie = tabEntree
	//ajout de tableaux de -1 lorsqu'il en manque
	if len(tabSortie) < midy {
		slice := make([]int, midx)
		for indice := range slice {
			slice[indice] = -1
		}
		tabSortie = append(tabSortie, slice)
	}
	//on rajoute les tableaux possiblement manquant
	//fmt.Println("sortie",tabSortie,midy,midx,len(tabSortie),len(tabSortie[0]))
	for len(tabSortie) < midy {
		tabSortie = append(tabSortie, []int{})
	}

	// ajout des -1 dans les tablaux
	for i := 0; i < midy; i++ {
		//fmt.Println("un ptit tour",tabSortie,i)
		if len(tabSortie[i]) < midx {
			tabSortie[i] = append(tabSortie[i], -1)
		}
	}
	return tabSortie
}

func ShowQuad(Q Quadtree) (tab [][]int) {
	//fonction d'affichage des quadtree en un tableau de tableaux
	for i := 0; i < Q.height; i++ {
		tab = append(tab, make([]int, Q.width))
	}

	tab = ShowNode(Q.root, tab)
	return tab
}

func ShowNode(noeu *node, tab [][]int) [][]int {
	//on parcours tt les noeuds
	if noeu == nil {
		//s'il est nul on ne touche pas au tableau
		return tab
	}
	//s'il contient d'autres tableaux (node.content = -1) on parcours ses 4 noeuds
	if noeu.content < 0 {
		ShowNode(noeu.topLeftNode, tab)
		ShowNode(noeu.topRightNode, tab)
		ShowNode(noeu.bottomLeftNode, tab)
		ShowNode(noeu.bottomRightNode, tab)
	} else {
		//sinon on ajoute son content dans le tableau de tableaux.
		for i := 0; i < noeu.height; i++ {

			for j := 0; j < noeu.width; j++ {
				tab[noeu.topLeftY+i][noeu.topLeftX+j] = noeu.content
			}
		}
	}
	return tab
}
