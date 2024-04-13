package teleporteur

import 	(
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
)

func (L_tp *Liste_teleporteur) CreateTeleporteur (Personnage character.Character) {
	//Cette fonction remplis les champs de le structure Liste_teleporteur donc TPA et TPB

	//1ère initialisation de TPA
	if L_tp.TPA == nil {
		var TP Teleporteur
		TP.X = Personnage.X
		TP.Y = Personnage.Y
		L_tp.TPA = &TP

	//1ère initialisation de TPB
	} else if L_tp.TPB == nil {
		if L_tp.TPA.X != Personnage.X ||  L_tp.TPA.Y != Personnage.Y{
			var TP Teleporteur
			TP.X = Personnage.X
			TP.Y = Personnage.Y
			L_tp.TPB = &TP
		}

	//on place TPB dans TPA quand le personnage n'est pas sur un téléporteur et que TPA et TPB sont déjà renseigné puis on redéfinit TPB
	}else if L_tp.TPA != nil && L_tp.TPB != nil {
		if (Personnage.X != L_tp.TPA.X || Personnage.Y != L_tp.TPA.Y) && (Personnage.X != L_tp.TPB.X || Personnage.Y != L_tp.TPB.Y ){
			L_tp.TPA = L_tp.TPB
			var TP Teleporteur
			TP.X = Personnage.X
			TP.Y = Personnage.Y
			L_tp.TPB = &TP
		}
	}
}
