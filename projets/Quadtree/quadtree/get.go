package quadtree

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	//si CameraMode est à 0 on change les topLeft
	if configuration.Global.CameraMode == 0 {
		topLeftX = topLeftX - configuration.Global.NumTileX/2
		topLeftY = topLeftY - configuration.Global.NumTileY/2
	} else {
		topLeftX = topLeftX + q.width/2
		topLeftY = topLeftY + q.height/2
	}
	//on remplie contentHolder de -1
	for indiceAbs := range contentHolder {
		for indiceOrd := range contentHolder[0] {
			contentHolder[indiceAbs][indiceOrd] = -1
		}
	}
	//on avance en défillant les noeuds avec la variable recursive continue
	Continue(q.root, contentHolder, topLeftX, topLeftY)
}

func Continue(node *node, contentHolder [][]int, topLeftX int, topLeftY int) {
	//création de variable servant à une meilleur compréhension des coordonnées de
	var NodeDebX int = node.topLeftX
	var NodeDebY int = node.topLeftY
	var NodeFinX int = node.topLeftX + node.width
	var NodeFinY int = node.topLeftY + node.height

	var PrendreCeNoeud bool
	//booléen déterminant s'il on prend (continue le script avec) le noeud
	PrendreCeNoeud = ((topLeftX >= NodeDebX && topLeftY >= NodeDebY) || (topLeftX <= NodeFinX && topLeftY <= NodeFinY))
	/*
		if PrendreCeNoeud == true {
			fmt.Println("lesssgooo")
		}else{
			fmt.Println("topCamX",topLeftX,"topCamY",topLeftY)
		}
	*/
	//si on le prend et qu'il contient -1
	if PrendreCeNoeud && node.content == -1 {
		//on refait le programme sur ses noeuds (s'ils existent)
		if node.topLeftNode != nil {
			Continue(node.topLeftNode, contentHolder, topLeftX, topLeftY)
		}
		if node.topRightNode != nil {
			Continue(node.topRightNode, contentHolder, topLeftX, topLeftY)
		}
		if node.bottomLeftNode != nil {
			Continue(node.bottomLeftNode, contentHolder, topLeftX, topLeftY)
		}
		if node.bottomRightNode != nil {
			Continue(node.bottomRightNode, contentHolder, topLeftX, topLeftY)
		}
	}
	//si on le prend et qu'il ne contient pas -1
	if PrendreCeNoeud && node.content != -1 {
		//on le parcours et regarde s'il on peut ajouter une de ses cases aux coord corespondantes de contentHolder.
		for i := 0; i < node.height; i++ {
			for j := 0; j < node.width; j++ {
				PosCaseX := (j + node.topLeftX) - topLeftX
				PosCaseY := (i + node.topLeftY) - topLeftY
				if PosCaseY >= 0 && PosCaseX >= 0 && PosCaseY < len(contentHolder) && PosCaseX < len(contentHolder[0]) {
					//fmt.Println("####################\ncase sel:",PosCaseY,PosCaseX,"max",len(contentHolder),len(contentHolder[0]),"\ncontient:",node.content)
					//fmt.Println("node coord (x;y) :",node.topLeftX,node.topLeftY,"\n####################")
					contentHolder[PosCaseY][PosCaseX] = node.content
				}
			}
		}
	}
}
