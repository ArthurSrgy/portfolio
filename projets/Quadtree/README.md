# Projet Quadtree
Résumé: Jeu générant un village aléatoirement et nous laissant se balader librement, comprend diverses options suplémentaires comme un système de téléporteur, de cycle de la journée, ou de météo.




## Pour le lancer

Ouvrez votre terminal, aller dans ./cmd/ avec la commande "cd cmd/" puis entrer la commande "./cmd"une fois dedans, cela ouvrira le jeu, pour le quitter, cliquez simplement sur la croix de la fenêtre ouverte.

## Changer les paramètres

Pour changer les différents paramètres, rien de plus simple, ouvrez le fichier ./cmd/config.json, éditez le conformément pour activer/désactiver les extensions.

## Comment paramétrer ?

- Fullscreen: démarre le jeu en plein écran si mit à true, appuyez sur Échap pour l'enlever / activer.

**1) Génération aléatoire:**
- SeedOn: Prend en compte la seed proposé si mit à true.
- Seed: Seed du terrain généré aléatoirement, permet de regénérer le même terrain.
- GenerateRandomFloor: Génère un terrain aléatoire si mit à true.
		       * Génère des quadtrees aléatoire si FloorKind = 2.
		       * Génère un terrain aléatoire avec un algorithme WaveFonctionCollapse si floorKind = 1.
- GenerateVillage: si mit à true avec FloorKind = 2 et GenerateRandomFloor = true, génère un
		   village aléatoirement en utilisant les terrains dans /floor-to-generate.

**2) Changement Sol:**
- FrameCount: compte le nombre de Frame chaque seconde.
- AnimeFloor: affiche les animations de l'eau si mit à true.
- Fuse Floor: affiche les tuiles reliés entres elles .

**3) Changement Personnage:**
- SwitchPersoOn: permet de changer de personnage avec la touche P.
- Personnage: * si 0, le personnage est un chevalier (aucunne spécificité).
	      * si 1, le personnage est un fantôme (peut traverser eau et murs).

**4) Option Caméra:**
- CameraMode: change le mode de caméra.
	      * 0 => la caméra ne suit pas le joueur.
	      * 1 & 2 => mobile mais non fluide.
	      * 3 => mobile et fluide.

**5) Particule:**
- ParticuleOn: affiche les traces de pas du joueur.
- Météo: affiche différentes météo.
	 * 0 => normale
	 * 1 => pluie
	 * 2 => neige

**6) Téléporteur:**
- TeleporteurOn: active les téléporteurs, appuyez sur T pour les placer et s'y téléporter (si l'on se trouve dessus).

**7) Journée:**
- DaytimeCycle: active le cycle de la journée.
- Daytime: moment de la journée (entre 0 et 7200).
- DaytimeSpeed: vitesse de la journée (si égale à 1 la journée durera 2 minutes, si égale à 2 une minute etc..)
