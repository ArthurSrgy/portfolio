package particule

// Liste des différentes trace de pas
type FootprintList struct {
	List []footprint
}

// Footprint est la structure prennant en compte une trace de pas.
//   - PosX, PosY : la position de l'empreinte.
//   - lifetime : le temps restant avant sa disparition.
//   - direction: la direction de l'empreinte (4 sens)
type footprint struct {
	PosX      int
	PosY      int
	lifetime  int
	direction int
}
