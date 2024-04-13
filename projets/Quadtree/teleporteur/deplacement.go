package teleporteur

import 	(
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
)

func (L_tp *Liste_teleporteur) Deplacement (Personnage character.Character, Cam camera.Camera) (int, int, int, int){
	//Cette fonction change les coordonnées de la caméra et du personnage 
	//par celle du téléporteur sur lequel le personnage n'est pas

	//Personnage sur TPA ? => récupération des coordonnées de TPB
	if Personnage.X == L_tp.TPA.X && Personnage.Y == L_tp.TPA.Y	{
		Personnage.X = L_tp.TPB.X
		Personnage.Y = L_tp.TPB.Y
		Cam.X = L_tp.TPB.X
		Cam.Y = L_tp.TPB.Y
	
	//Personnage sur TPB ? => récupération des coordonnées de TPA
	} else if Personnage.X == L_tp.TPB.X && Personnage.Y == L_tp.TPB.Y{	
		Personnage.X = L_tp.TPA.X
		Personnage.Y = L_tp.TPA.Y
		Cam.X = L_tp.TPA.X
		Cam.Y = L_tp.TPA.Y
	}
	return Personnage.X, Personnage.Y, Cam.X, Cam.Y
}