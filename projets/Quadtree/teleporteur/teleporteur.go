package teleporteur

//Cette structure va stocker les données des téléporteurs
type Liste_teleporteur struct {
	TPA      *Teleporteur
	TPB      *Teleporteur
}

//Cette structure garde en mémoire les coordonnées pour chaque téléporteur de Liste_teleporteur
type Teleporteur struct {
	X      int
	Y      int
}