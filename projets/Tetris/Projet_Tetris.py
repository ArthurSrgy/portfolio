import pyxel
import os
from random import *

global largeur ; largeur = 14
global hauteur ; hauteur = 14
global Garfield ; Garfield = False
global Mute ; Mute = True

if not( os.path.exists("data.txt") ) : #si il n'existe pas de fichier de sauvegarde ("data.txt")
    sauvegarde = open("data.txt", "w") #on en créer un
    sauvegarde.write('0') #on met le score max à 0
    sauvegarde.close() #on le ferme (enregistre)

global fichier ; fichier = open("data.txt", "r") #initialisation de la variable globale fichier, permettant de lire et modifier la sauvegarde
fichier.close() #on le ferme (enregistre)

class Jeu :
    def __init__(self):
        global tableau
        tableau = [[ 0 for i in range(largeur)] for i in range(hauteur)]
        self.score = 0
        self.vitesse = 10
        self.temps = 0
        self.tetro_actif = Tetromino()
        self.tetro_suivant = Tetromino()
        self.state = "menu"
        self.cooldown = 0

        pyxel.run(self.update,self.draw)


    def creer_tetromino(self):
        self.tetro_actif = self.tetro_suivant
        self.tetro_suivant = Tetromino()
        self.tetro_suivant.forme = self.tetro_suivant.attribution_forme(0,0,0)


    def supprimer_ligne(self):
        global tableau
        h = 0
        multiplicateur = 1
        while h < hauteur:
            if tableau[h] == [1]*largeur :
                tableau.pop(h)
                tableau.insert(0,[0]*largeur)
                self.score += 50*multiplicateur
                self.changement_bestscore()
                self.changement_vitesse()
                multiplicateur += 1
                if Mute == False :
                    pyxel.play(2,7)

            else:
                h += 1



    def ajouter_tetro(self): #cette fonction ajoute le tetromino une fois tombé mais regarde avant s'il ne dépasse pas la ligne, sinon Game Over
        global tableau
        if self.tetro_actif.state == 'lock' :
            self.score += 10
            self.changement_vitesse()
            if Mute == False :
                pyxel.play(3,6)

            for coord in self.tetro_actif.forme :
                if coord[1] < 0 :
                    pyxel.stop()
                    self.state = "game over"
                    self.cooldown = 100
                    if Mute == False :
                        pyxel.play(2,8)

                else:
                    tableau[coord[1]][coord[0]] = 1
                    self.changement_bestscore()

            self.creer_tetromino()


    def changement_vitesse(self):
        #plus la vitesse augmente, plus le jeu est rapide
        self.vitesse = 10+((self.score*1.5)//500)
        if self.vitesse >= 25:
            self.vitesse = 25

    def changement_bestscore(self):
        global fichier
        fichier = open("data.txt", "r")
        if int(fichier.read()) < self.score :
            fichier = open("data.txt", "w")
            fichier.write(str(self.score))
            fichier = open("data.txt", "r")
            fichier.close()


    def update(self):

        if self.state == "play" :
            self.temps +=1
            if pyxel.btn(pyxel.KEY_ESCAPE):
                pyxel.quit()

            self.supprimer_ligne()

            if pyxel.btnp(pyxel.KEY_G) :
                global Garfield
                Garfield = True

            if pyxel.btnp(pyxel.KEY_M) :
                global Mute
                if Mute == False :
                    Mute = True
                    pyxel.stop()
                else :
                    Mute = False
                    if self.state in ["play", "pause"] :
                        pyxel.playm(0,0,True)


            if pyxel.btnp(pyxel.KEY_DOWN) or pyxel.btnp(pyxel.KEY_KP_5) or pyxel.btnp(pyxel.KEY_S):
                self.tetro_actif.tomber()
            if pyxel.btnp(pyxel.KEY_UP) or pyxel.btnp(pyxel.KEY_KP_8) or pyxel.btnp(pyxel.KEY_Z):
                self.tetro_actif.rotation()
            if pyxel.btnp(pyxel.KEY_LEFT) or pyxel.btnp(pyxel.KEY_KP_4) or pyxel.btnp(pyxel.KEY_Q):
                self.tetro_actif.deplacement_g()
            if pyxel.btnp(pyxel.KEY_RIGHT) or pyxel.btnp(pyxel.KEY_KP_6) or pyxel.btnp(pyxel.KEY_D):
                self.tetro_actif.deplacement_d()
            if pyxel.btnp(pyxel.KEY_SPACE) or pyxel.btnp(pyxel.KEY_RETURN):
                self.tetro_actif.deplacement_bas()
            if pyxel.btnp(pyxel.KEY_P):
                self.state = "pause"
            if pyxel.btnp(pyxel.KEY_CTRL) :
                pyxel.stop()
                self.__init__()

            self.tetro_actif.forme = self.tetro_actif.attribution_forme(self.tetro_actif.x, self.tetro_actif.y, self.tetro_actif.angle)

            if pyxel.frame_count%(30-self.vitesse) == 0 :
                self.tetro_actif.tomber()

            self.ajouter_tetro()
            self.tetro_actif.forme = self.tetro_actif.attribution_forme(self.tetro_actif.x, self.tetro_actif.y, self.tetro_actif.angle)

        elif self.state == "menu" :
            if pyxel.btnp(pyxel.KEY_SPACE) :
                self.temps = 0
                self.state = "play"
                if Mute == False :
                    pyxel.playm(0,0,True)


        elif self.state == "game over" :
            self.cooldown -= 1
            if self.cooldown == 0 :
                self.__init__()

        elif self.state == "pause" :
            if pyxel.btnp(pyxel.KEY_P):
                self.state = "play"



    def draw(self):
        #### création du fond #####
        pyxel.cls(0)
                   #pos. x , pos. y , tm? = 0 , pos. u , pos. v, taille du bloc, couleur transparante

        #afficher tetromino actif
        if not self.state in ["menu","game over"] :
            if Garfield == True and self.tetro_actif.type == 2 :
                    pyxel.bltm( self.tetro_actif.forme[0][0]*8+8 , self.tetro_actif.forme[0][1]*8+8, 0, 0, 32, 16, 16,1)
            else:
                for coord in self.tetro_actif.forme :
                    if Garfield == False:
                        pyxel.bltm(coord[0]*8+8, coord[1]*8+8, 0, self.tetro_actif.type*8, 0, 8, 8)
                    else :
                        pyxel.bltm(coord[0]*8+8, coord[1]*8+8, 0, 16, 32, 8, 8,0)

        #afficher tetro posés
        if self.state == "play" or self.state == "pause" :
            for i in range(len(tableau[0])) :
                for j in range(len(tableau)) :
                    if tableau[j][i] == 1 :
                        if Garfield == False:
                            pyxel.bltm(i*8+8, j*8+8, 0, 56, 0, 8, 8)
                        else:
                            pyxel.bltm(i*8+8, j*8+8, 0, 16, 32, 8, 8,0)

        if self.state == "game over" :
            if self.cooldown%20 > 10:
                for i in range(len(tableau[0])) :
                    for j in range(len(tableau)) :
                        if tableau[j][i] == 1 :
                            pyxel.bltm(i*8+8, j*8+8, 0, 56, 0, 8, 8) #on met en noir...


        #coins de la fenêtre entière
        pyxel.bltm(0, 0, 0, 0, 8, 8, 8)
        pyxel.bltm((largeur+8)*8, (hauteur+1)*8, 0, 16, 24, 8, 8)
        pyxel.bltm(0, (hauteur+1)*8, 0, 0, 24, 8, 8)
        pyxel.bltm((largeur+8)*8, 0, 0, 16, 8, 8, 8)


        for i in range(1,largeur+8):
            pyxel.bltm(i*8, 0, 0, 8, 8, 8, 8)
            pyxel.bltm(i*8, (hauteur+1)*8, 0, 8, 24, 8, 8)

        for i in range(1,hauteur+1):
            pyxel.bltm(0, i*8, 0, 0, 16, 8, 8)
            pyxel.bltm((largeur+8)*8, i*8, 0, 16, 16, 8, 8)

        for i in range(1,hauteur+1):
            for j in range(largeur+1,largeur+8):
                pyxel.bltm(j*8, i*8, 0, 8, 16, 8, 8)

        #coin de la fenêtre des stats
        pyxel.bltm((largeur+2)*8, 16, 0, 24, 8, 8, 8)
        pyxel.bltm((largeur+2)*8, (hauteur-1)*8 , 0, 24, 24, 8, 8)
        pyxel.bltm((largeur+7)*8, 16, 0, 40, 8, 8, 8)
        pyxel.bltm((largeur+7)*8, (hauteur-1)*8 , 0, 40, 24, 8, 8)

        for i in range(2,6):
            pyxel.bltm((largeur+1+i)*8, 16, 0, 32, 8, 8, 8)
            pyxel.bltm((largeur+1+i)*8, (hauteur-1)*8, 0, 32, 24, 8, 8)

        for i in range(hauteur-4):
            pyxel.bltm((largeur+7)*8, (i+3)*8, 0, 40, 16, 8, 8)
            pyxel.bltm((largeur+2)*8, (i+3)*8, 0, 24, 16, 8, 8)

        for i in range(4):
            for j in range(hauteur-4):
                pyxel.bltm((largeur+i+3)*8, (j+3)*8, 0, 32, 16, 8, 8)

        #afficher tetro suivant
        a = 4
        if self.tetro_suivant.type in [3,6]:
            a = 5
        if not self.state in ["menu", "game over"] :
            if Garfield == True and self.tetro_suivant.type == 2 :
                    pyxel.bltm( (largeur+a)*8 , 4*8, 0, 0, 32, 16, 16,1)

            else :
                for coord in self.tetro_suivant.attribution_forme(0,0,0) :
                    pyxel.bltm( ( coord[0] + largeur + a )*8 , (coord[1]+5)*8, 0, self.tetro_suivant.type*8, 0, 8, 8)

        ###AFFICHAGE STATS###
        #afficher score
        pyxel.text((largeur+3)*8,8*8,"Score:",10 - (pyxel.frame_count%30//15))
        pyxel.text((largeur+3)*8,9*8,str(self.score),9+ (pyxel.frame_count%30//15))

        #afficher meilleur score
        if hauteur >= 13 :
            global fichier
            fichier = open("data.txt", "r")
            pyxel.text((largeur+3)*8,10*8,"Best:",10 - (pyxel.frame_count%30//15))
            pyxel.text((largeur+3)*8,11*8,str(fichier.read()),9+ (pyxel.frame_count%30//15))
            fichier.close()

        #afficher vitesse
        if hauteur >= 14 :
            pyxel.text((largeur+3)*8,12*8,"Speed:",10 - (pyxel.frame_count%30//15))
            pyxel.text((largeur+3)*8,13*8,str(self.vitesse),9+ (pyxel.frame_count%30//15))

        #afficher temps
        if hauteur >= 16 :
            pyxel.text((largeur+3)*8,14*8,"Time:",10 - (pyxel.frame_count%30//15))
            pyxel.text((largeur+3)*8,15*8,str(self.temps//30),9+ (pyxel.frame_count%30//15))


        if self.state == "menu" :
            pyxel.text((largeur//2)*8-14,hauteur*8//2,"Press Space",int(not(pyxel.frame_count%30//15))*7)
            pyxel.text((largeur//2)*8-6,(hauteur+2)*8//2,"To Play",int(not(pyxel.frame_count%30//15))*7)





class Tetromino :
    def __init__(self) :
        self.type = randint(0,6)
        self.x = largeur // 2
        self.y = -1
        self.angle = 0
        self.state = 'fall'
        self.forme = self.attribution_forme(self.x, self.y, self.angle)

    def attribution_forme(self, x, y, angle): #on y entre le type, l'abcisse, l'ordonnée et l'angle du tetromino
        formes = [
            [ [(x,y-2),(x,y-1),(x,y),(x,y+1)]       , [(x-1,y),(x,y),(x+1,y),(x+2,y)]       , [(x,y-2),(x,y-1),(x,y),(x,y+1)]       , [(x-1,y),(x,y),(x+1,y),(x+2,y)] ]      , #Tetromino I
            [ [(x,y-1),(x,y),(x+1,y),(x,y+1)]       , [(x+1,y),(x,y),(x,y+1),(x-1,y)]       , [(x,y+1),(x,y),(x-1,y),(x,y-1)]       , [(x-1,y),(x,y),(x,y-1),(x+1,y)] ]      , #tetromino T
            [ [(x,y-1),(x+1,y-1),(x,y),(x+1,y)]     , [(x,y-1),(x+1,y-1),(x,y),(x+1,y)]     , [(x,y-1),(x+1,y-1),(x,y),(x+1,y)]     , [(x,y-1),(x+1,y-1),(x,y),(x+1,y)] ]    , #tetromino O
            [ [(x-1,y-1),(x-1,y),(x-1,y+1),(x,y+1)] , [(x+1,y-1),(x,y-1),(x-1,y-1),(x-1,y)] , [(x+1,y+1),(x+1,y),(x+1,y-1),(x,y-1)] , [(x-1,y+1),(x,y+1),(x+1,y+1),(x+1,y)] ], #tetromino L
            [ [(x+1,y-1),(x+1,y),(x+1,y+1),(x,y+1)] , [(x+1,y+1),(x,y+1),(x-1,y+1),(x-1,y)] , [(x-1,y+1),(x-1,y),(x-1,y-1),(x,y-1)] , [(x-1,y-1),(x,y-1),(x+1,y-1),(x+1,y)] ], #tetromino J
            [ [(x+1,y-1),(x+1,y),(x,y),(x,y+1)]     , [(x+1,y+1),(x,y+1),(x,y),(x-1,y)]     , [(x+1,y-1),(x+1,y),(x,y),(x,y+1)]     , [(x+1,y+1),(x,y+1),(x,y),(x-1,y)] ]    , #tetromino Z
            [ [(x-1,y-1),(x-1,y),(x,y),(x,y+1)]     , [(x-1,y+1),(x,y+1),(x,y),(x+1,y)]     , [(x-1,y-1),(x-1,y),(x,y),(x,y+1)]     , [(x-1,y+1),(x,y+1),(x,y),(x+1,y)] ]    , #tetromino S
            ]

        return formes[self.type][angle]


    def collision(self):
        retour = False #le renvoie est positif, le tetromino peut être placé/bougé
        for coord in self.forme :
            if not( 0 <= coord[0] < largeur ) : #on regarde si l'abcisse du tetromino est bien entre 0 et la largeur
                retour = True #si non, il y a collision

            elif not ( -3 <= coord[1] < hauteur ) : #on regarde si l'ordonné du tetromino est bien entre 0 et la hauteur
                retour = True #si non, il y a collision

            elif coord[0] >= 0 and coord[1] >= 0 :
                if tableau[coord[1]][coord[0]] == 1 :
                    retour = True

        return retour


    def rotation(self):
        if not(self.state == 'lock'):
            if self.type == 2: #si le tetromino est un carrée il ne tourne pas
                None
            else:
                self.angle = (self.angle + 1 )%4
                self.forme = self.attribution_forme(self.x, self.y, self.angle)
                while self.collision() == True : #on vérifie que le tetromino n'entre pas en collision
                    self.angle = (self.angle + 1 )%4
                    self.forme = self.attribution_forme(self.x, self.y, self.angle)

    def tomber(self):
        if not(self.state == 'lock'):
            self.y += 1
            self.forme = self.attribution_forme(self.x, self.y, self.angle)
            if self.collision() == True :
                self.y -= 1
                self.forme = self.attribution_forme(self.x, self.y, self.angle)
                self.lock()

    def deplacement_d(self):
        if not(self.state == 'lock'):
            self.x += 1
            self.forme = self.attribution_forme(self.x, self.y, self.angle)
            if self.collision() == True :
                self.x -= 1
                self.forme = self.attribution_forme(self.x, self.y, self.angle)


    def deplacement_g(self):
        if not(self.state == 'lock'):
            self.x -= 1
            self.forme = self.attribution_forme(self.x, self.y, self.angle)
            if self.collision() == True :
                self.x += 1
                self.forme = self.attribution_forme(self.x, self.y, self.angle)

    def deplacement_bas(self):
        if not(self.state == 'lock'):
            while self.state != 'lock' :
                self.tomber()

    def lock(self):
        self.state = 'lock'


if largeur < 8:
            largeur = 8
elif largeur > 22 :
            largeur = 22
if hauteur < 10:
            hauteur = 10
elif hauteur > 24 :
            hauteur = 24

global tableau ;  tableau = [[ 0 for i in range(largeur)] for i in range(hauteur)]

pyxel.init((largeur+9)*8, (hauteur+2)*8, title="Tetris",fps=30) #initialisation de l'écran
pyxel.load('Ressource_TetrisFinalMix.pyxres') #charge les ressources

Jeu()

