package quadtree

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func CreateRandomQuadtree() (Q Quadtree) {
	dimension := 256 // dimesions du quadtree généré (puissance de 2 obligé)
	if configuration.Global.SeedOn {
		rand.Seed(configuration.Global.Seed)
	}
	Q.height = dimension
	Q.width = dimension
	//on appelle CreateRandomNode pour créer les noeuds
	if configuration.Global.GenerateVillage == false {
		Q.root = CreateRandomNode(0, 0, Q.width, Q.height)
	} else {
		Q.root = CreateNodeVillage(0, 0, Q.width, Q.height)
	}

	//fmt.Println("~ le quadtree: ",ShowQuad(Q))
	return Q
}

func CreateRandomNode(topLeftX, topLeftY, width, height int) *node {
	//la valeur aléatoire
	valeur := rand.Intn(5)
	//on ne veut pas de case d'eau de moins de 8 de dimension
	if valeur == 4 && width < 8 {
		valeur = rand.Intn(4)
	}
	for valeur == 3 && width > 8 {
		valeur = rand.Intn(5)
	}
	//1/width chances de remplir entièrement le noeud
	var rempli int = rand.Intn(width) //

	if rempli == 0 || (width == 1 && height == 1) {
		//fmt.Println("~ on retourne un noeud rempli:",rempli," taille:",width,height,"coord:",topLeftX,topLeftY)
		return &node{
			topLeftX:        topLeftX,
			topLeftY:        topLeftY,
			width:           width,
			height:          height,
			content:         valeur,
			topLeftNode:     nil,
			topRightNode:    nil,
			bottomLeftNode:  nil,
			bottomRightNode: nil,
		}
	} else {
		//fmt.Println("~ on créer d'autres noeud:",rempli)
		return &node{
			topLeftX:        topLeftX,
			topLeftY:        topLeftY,
			width:           width,
			height:          height,
			content:         -1,
			topLeftNode:     CreateRandomNode(topLeftX, topLeftY, width/2, height/2),
			topRightNode:    CreateRandomNode(topLeftX+width/2, topLeftY, width/2, height/2),
			bottomLeftNode:  CreateRandomNode(topLeftX, topLeftY+height/2, width/2, height/2),
			bottomRightNode: CreateRandomNode(topLeftX+width/2, topLeftY+height/2, width/2, height/2),
		}
	}
}

// ###############################################//
// ########### Generate Random Village ###########//
// ###############################################//

func CreateNodeVillage(topLeftX, topLeftY, width, height int) *node {
	// valeur aléatoire permettant de savoir si le noeud doit être remplie ou divisé
	// si rempli vaut 0, on rempli la noeud, sinon on le divise
	rempli := rand.Intn(width / 5)
	// si le noeud est plus grand que 32, on le divise
	if width > 32 {
		rempli = 1
	}
	// dans un besoin d'harmonie on préfère augmenter les chances des noeud de taille 16 d'être remplis
	if width == 16 && rempli != 0 {
		rempli = rand.Intn(4)
	}

	// si le noeud est rempli ou à une dimension de 8 (minimum):
	if rempli == 0 || (width == 8 && height == 8) {
		var FloorName string
		//on lui attribut un fichier par rapport à sa taille
		switch {
		case width == 8:
			FloorName = ChoseFile8()
		case width == 16:
			FloorName = ChoseFile16()
		case width == 32:
			FloorName = ChoseFile32()
		}
		//on récupère le fichier en forme de matrice
		matrice := ReadFloorFromFile(FloorName)

		//on le tourne aléatoirement (sauf si c'est une chapelle, ou un espace vide)
		if FloorName != "../floor-to-generate/16church" && FloorName != "../floor-to-generate/8empty" {
			matrice = RotateMatrixRight(matrice, rand.Intn(4))
		}
		//puis on retourne le noeud
		return MakeFromArray(matrice, topLeftX, topLeftY).root

		// sinon on divise en 4 autres noeuds
	} else {
		return &node{
			topLeftX:        topLeftX,
			topLeftY:        topLeftY,
			width:           width,
			height:          height,
			content:         -1,
			topLeftNode:     CreateNodeVillage(topLeftX, topLeftY, width/2, height/2),
			topRightNode:    CreateNodeVillage(topLeftX+width/2, topLeftY, width/2, height/2),
			bottomLeftNode:  CreateNodeVillage(topLeftX, topLeftY+height/2, width/2, height/2),
			bottomRightNode: CreateNodeVillage(topLeftX+width/2, topLeftY+height/2, width/2, height/2),
		}
	}
}

// choisi un des fichiers de dimension 8.
func ChoseFile8() (fname string) {
	//liste de toutes les possibilités de fichiers de dimension 8
	listFile8 := []string{"../floor-to-generate/8cabane", "../floor-to-generate/8cabane", "path", "path", "path", "water_well", "../floor-to-generate/8empty"}
	//indice du fichier choisi
	contenu := rand.Intn(len(listFile8))

	// listes des différents fichiers de chemins / puits
	listeWaterWell := []string{"../floor-to-generate/8water_well1", "../floor-to-generate/8water_well2"}
	listePath := []string{"../floor-to-generate/8path1", "../floor-to-generate/8path2", "../floor-to-generate/8path3", "../floor-to-generate/8path4", "../floor-to-generate/8path5", "../floor-to-generate/8path6", "../floor-to-generate/8path7"}

	fname = listFile8[contenu]
	//s'il on a sélectionné un chemin on regarde dans la liste de chemin
	if fname == "path" {
		fname = listePath[rand.Intn(len(listePath))]
	}
	//s'il on a sélectionné un puit on regarde dans la liste de puits
	if fname == "water_well" {
		fname = listeWaterWell[rand.Intn(len(listeWaterWell))]
	}
	return fname
}

// choisi un des fichiers de dimension 16.
func ChoseFile16() (fname string) {
	//liste de toutes les possibilités de fichiers de dimension 16
	listFile16 := []string{"../floor-to-generate/16fortress", "bridge", "../floor-to-generate/16village1", "../floor-to-generate/16village2", "../floor-to-generate/16village3", "../floor-to-generate/16house", "../floor-to-generate/16church"}

	//indice du fichier choisi
	contenu := rand.Intn(len(listFile16))

	// listes des différents fichiers pont
	listeBridge := []string{"../floor-to-generate/16bridgeA", "../floor-to-generate/16bridgeB", "../floor-to-generate/16bridgeC", "../floor-to-generate/16bridgeD", "../floor-to-generate/16lake"}

	fname = listFile16[contenu]
	//s'il on a sélectionné un pont on regarde dans la liste de pont
	if fname == "bridge" {
		fname = listeBridge[rand.Intn(len(listeBridge))]
	}
	return fname
}

// choisi un des fichiers de dimension 32.
func ChoseFile32() (fname string) {
	//liste de toutes les possibilités de fichiers de dimension 32
	listFile32 := []string{"../floor-to-generate/32island1", "../floor-to-generate/32island2", "../floor-to-generate/32cemetery"}
	//indice du fichier choisi
	contenu := rand.Intn(len(listFile32))

	return listFile32[contenu]
}

// Fait tourner une matrice à droite n fois (n entre 0 et 3)
func RotateMatrixRight(matrix [][]int, n int) [][]int {
	for i := 0; i < n; i++ {
		matrix = RotateMatrixRightOne(matrix)
	}
	return matrix
}

func RotateMatrixRightOne(matrix [][]int) [][]int {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

// ReadFromFileFloor mais ici dans /quadtree car les imports cassent la tête
func ReadFloorFromFile(fileName string) (floorContent [][]int) {
	var myFile *os.File
	var err error
	myFile, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var scanner *bufio.Scanner = bufio.NewScanner(myFile)

	for scanner.Scan() {
		var tab_ligne []int = make([]int, len(string(scanner.Text())))
		for indice_caractère := 0; indice_caractère < len(string(scanner.Text())); indice_caractère++ {
			tab_ligne[indice_caractère], _ = strconv.Atoi(string(scanner.Text()[indice_caractère]))
		}
		floorContent = append(floorContent, tab_ligne)
	}
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	return floorContent
}
