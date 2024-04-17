# Projet Quadtree
__Résumé:__ Jeu générant un village aléatoirement et nous laissant se balader librement, comprend diverses options suplémentaires comme un système de téléporteur, de cycle de la journée, ou de météo.

## Comment télécharger ?
Ouvrez votre __terminal linux__ ou __Powershell__ (si sur windows), placer vous dans le répertoire/dossier voulu pour le télécherger avec la commande *'cd'*.
Une fois fait effectuer la commande *'git clone https://github.com/ArthurSrgy/portfolio.git'*, celà vous clonera les fichier du portfolio (projets inclus).

__OU__

Télécharger le projet via ce lien: *https://gitlab.univ-nantes.fr/iut_1/s1.01/eq_03_01/-/releases/RenduFinal*

## Pour le lancer

Ouvrez votre terminal, aller dans *[nom_de_votre_répertoire]/portfolio/projets/Quadtree/cmd/* avec la commande *'cd [nom_de_votre_répertoire]/portfolio/projets/Quadtree/cmd/'* puis entrer la commande *'./cmd'* une fois dedans, cela ouvrira le jeu, pour le quitter, cliquez simplement sur la croix de la fenêtre ouverte.
(les information entre [ ] sont à remplacer sans les crochets)

__OU__

Ouvrez votre terminal, aller dans *[nom_de_votre_répertoire]/RenduFinal/cmd/* avec la commande *'cd [nom_de_votre_répertoire]/portfolio/projets/Quadtree/cmd/'* puis entrer la commande *'./cmd'* une fois dedans, cela ouvrira le jeu, pour le quitter, cliquez simplement sur la croix de la fenêtre ouverte.

## Changer les paramètres

Pour changer les différents paramètres, rien de plus simple, ouvrez le fichier *./cmd/config.json*, éditez le conformément pour activer/désactiver les extensions.

## Comment paramétrer ?

- __Fullscreen__: démarre le jeu en plein écran si mit à true, appuyez sur Échap pour l'enlever / activer.

**1) Génération aléatoire:**
- __SeedOn:__ Prend en compte la seed proposé si mit à true.
- __Seed:__ Seed du terrain généré aléatoirement, permet de regénérer le même terrain.
- __GenerateRandomFloor:__ Génère un terrain aléatoire si mit à true.
		       * Génère des quadtrees aléatoire si FloorKind = 2.
		       * Génère un terrain aléatoire avec un algorithme WaveFonctionCollapse si floorKind = 1.
- __GenerateVillage:__ si mit à true avec FloorKind = 2 et GenerateRandomFloor = true, génère un
		   village aléatoirement en utilisant les terrains dans /floor-to-generate.

**2) Changement Sol:**
- __FrameCount:__ compte le nombre de Frame chaque seconde.
- __AnimeFloor:__ affiche les animations de l'eau si mit à true.
- __Fuse Floor:__ affiche les tuiles reliés entres elles .

**3) Changement Personnage:**
- __SwitchPersoOn:__ permet de changer de personnage avec la touche P.
- __Personnage:__ * si 0, le personnage est un chevalier (aucunne spécificité).
	      * si 1, le personnage est un fantôme (peut traverser eau et murs).

**4) Option Caméra:**
- __CameraMode:__ change le mode de caméra.
	      * 0 => la caméra ne suit pas le joueur.
	      * 1 & 2 => mobile mais non fluide.
	      * 3 => mobile et fluide.

**5) Particule:**
- __ParticuleOn:__ affiche les traces de pas du joueur.
- __Météo:__ affiche différentes météo.
	 * 0 => normale
	 * 1 => pluie
	 * 2 => neige

**6) Téléporteur:**
- __TeleporteurOn:__ active les téléporteurs, appuyez sur T pour les placer et s'y téléporter (si l'on se trouve dessus).

**7) Journée:**
- __DaytimeCycle:__ active le cycle de la journée.
- __Daytime:__ moment de la journée (entre 0 et 7200).
- __DaytimeSpeed:__ vitesse de la journée (si égale à 1 la journée durera 2 minutes, si égale à 2 une minute etc..)
