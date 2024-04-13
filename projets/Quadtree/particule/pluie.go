package particule

// Liste des différentes gouttes de pluie
type Pluie struct {
	List []Goutte
}

// goutte est la structure prennant en compte une simple goutte de pluie.
// Si Global.meteo == 1 on a alors une goutte d'eau,
// si Global.meteo == 2 on a alors un flocon de neige.
//   - CaseX, CaseY : la position de la case où est la goutte.
//   - PosX, PosY 	: la position de la goutte en pyxels sur la case.
//   - lifetime   	: le temps restant avant sa disparition.
//   - speed	  	: vitesse de la goutte
type Goutte struct {
	CaseX    int
	CaseY    int
	PosX     float32
	PosY     float32
	lifetime int
	speed    float32
}
