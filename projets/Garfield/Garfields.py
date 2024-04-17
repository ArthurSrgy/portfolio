import pyxel
from math import*
from random import randint


# taille de l'écran
global cote ; cote = 300



class Balle :
    """
    objet Balle avec ses coordonnées, son différentiel en x et en yet sa direction d en degrés.
    Les balles sont représenté par des Garfields.
    (Garfield est un chat fictif issu du comic strip homonyme Garfield créé par Jim Davis et publié pour la première fois le 19 juin 1978 par United Feature Syndicate en syndication dans 41 journaux)
    """
    # le constructeur
    def __init__(self,classe):
        self.x = cote // 2
        self.y = cote //2
        self.d = randint(-179,180)


        #attribution des différences
        self.attribut = ["Lent","Normal","Rapide"][randint(0,2)]
        if classe == "Lent" :
            self.attribut = "Lent"
        elif classe == "Normal" :
            self.attribut = "Normal"
        elif classe == "Rapide" :
            self.attribut = "Rapide"

        if self.attribut == "Lent" : #un Garfield lent aura une moins bonne vitesse mais une zone de colision plus grosse
            self.vitesse = 1.25
            self.hitbox = 8.5 #zone de colision
            self.couleur = 2 #couleur du nez
        if self.attribut == "Normal" :#un Garfield normal aura une vitesse moyenne et une zone de colision moyenne
            self.vitesse = 1.5
            self.hitbox = 7 #zone de colision
            self.couleur = 14 #couleur du nez
        if self.attribut == "Rapide" : #un Garfield rapide aura une meilleure vitesse mais une zone de colision plus petite
            self.vitesse = 1.75
            self.hitbox = 6 #zone de colision
            self.couleur = 8 #couleur du nez

        self.nblasagne = 0 #initialisation du compteur du nombre de lasagne avalé

        temps_secondes = 20 # temps de la génération
        self.temps_vie = 30*temps_secondes # il y a 30 images par seconde




class Lasagne :

    def __init__(self):
        self.lx = randint(5,cote-5)
        self.ly = randint(5,cote-5)


    def distance(self,garf):
        """Fonction calculant la distance entre tout les garfields et une lasagne,
        si la distance est inférieur ou égal à la zone de colision d'un des garfield
        alors la lasagne changera aléatoirement de position et le compteur personel
        de lasagnes mangé par le garfield augmentera de 1"""
        for b in garf :
            if hypot(self.lx-b.x, self.ly-b.y) <= b.hitbox:
                b.nblasagne += 1
                #b.vitesse += 0.25
                self.lx = randint(5,cote-5)
                self.ly = randint(5,cote-5)
                if b.nblasagne%5 == 0 :
                    b.vitesse += 0.5
                #print("distance: ",hypot(self.lx-b.x, self.ly-b.y))







# =====================================================
# == La class qui définit l'animation
# =====================================================
class Anim:
    # le constructeur
    def __init__(self):
        # taille de la fenetre cote * cote pixels
        pyxel.init(cote+95, cote+20, title="Garfields & lasagnes")

        # nombre de balles et de lasagnes créé
        nbGarfld = 80
        nbLasagna = 30

        #initialisation des variables comptant les totaux de garfields selon leurs attributs
        TotalGlent = 0
        TotalGnrml = 0
        TotalGrapd = 0

        # notre liste de balles qui n'en contient qu'une pour commencer
        self.balles = []
        self.lasagnes = []

        # numéro de la génération
        self.generation = 1
        self.finalText = ""
        self.finalList =[]

        for i in range(nbGarfld):
            self.balles.append(Balle(""))


        for i in range(nbLasagna):
            self.lasagnes.append(Lasagne())

        #on compte combien il y a de garfields lent, normaux, et rapide
        for b in self.balles :
            if b.attribut == "Lent":
                TotalGlent += 1
            if b.attribut == "Normal":
                TotalGnrml += 1
            if b.attribut == "Rapide":
                TotalGrapd += 1

        self.prcLent = TotalGlent * 100 // nbGarfld
        self.prcNrml = TotalGnrml * 100 // nbGarfld
        self.prcRapd = TotalGrapd * 100 // nbGarfld
        print("Nombre de Garfields lents:", TotalGlent,"(", self.prcLent ,"%)")
        print("Nombre de Garfields normaux:", TotalGnrml,"(", self.prcNrml ,"%)")
        print("Nombre de Garfields rapides:", TotalGrapd,"(", self.prcRapd ,"%)")
        print("---------")

        #scrolling
        self.scrollingX = 0

        # pour exécuter le jeu
        pyxel.run(self.update, self.draw)





    # =====================================================
    # == Des méthodes
    # =====================================================

    def deplacement(self):
        """déplacement des balles avec les rebonds sur les côtés ou au sol"""
        for b in self.balles :
            if b.x > cote or b.x  < 0 or b.y  > cote or b.y < 0:
                # on touche l'une des bordures
                # on modifie l'orientation de la balle
                b.d = b.d + randint(87,93)

            #la balle se déplace

            b.x += round( sin( radians(b.d) ) , 3) * b.vitesse
            b.y += round( cos( radians (b.d) ) , 3) * b.vitesse



    def vieillir(self) :
        """ vieillissement et mort des balles"""
        # on crée la liste des balles qui sont périmées
        liste_dead = []

        for i in range(len(self.balles)) :
            self.balles[i].temps_vie -= 1
            if self.balles[i].temps_vie == 0 :
                liste_dead.append(i)
        # on supprime les balles périmées de la liste en commençant par la fin

        if len(liste_dead) > 0:
            self.finalText = ""
            self.finalList =[]
            self.scrollingX = 0

        for i in range(len(liste_dead)-1,-1,-1) :
            for b in self.balles :
                if self.balles[0] == b :

                    #print(b.attribut, b.nblasagne)
                    self.finalText += str(b.attribut)+": "+str(b.nblasagne) +'\n'
                    self.finalList.append( (b.attribut,b.nblasagne) )
            #print("--------")

            self.balles.pop(0)




    def newGeneration(self):
        #on trie la liste des garfields de la génération précédente
        for i in range(1, len(self.finalList)):
            k = self.finalList[i]
            j = i-1
            while j >= 0 and k[1] > self.finalList[j][1] :
                    self.finalList[j + 1] = self.finalList[j]
                    j -= 1
            self.finalList[j + 1] = k

        #initialisation des variables comptant les totaux de garfields selon leurs attributs
        nbGarfld = len(self.finalList)
        TotalGlent = 0
        TotalGnrml = 0
        TotalGrapd = 0

        # notre liste de balles qui n'en contient qu'une pour commencer
        # numéro de la génération
        self.generation += 1



        # on recrée une génération
        for i in range(len(self.finalList)//2) :
            newBall = Balle(self.finalList[i][0])
            self.balles.append(newBall)
            newBall = Balle(self.finalList[i][0])
            self.balles.append(newBall)


        #on compte combien il y a de garfields lent, normaux, et rapide
        for b in self.balles :
            if b.attribut == "Lent":
                TotalGlent += 1
            if b.attribut == "Normal":
                TotalGnrml += 1
            if b.attribut == "Rapide":
                TotalGrapd += 1

        self.prcLent = TotalGlent * 100 // nbGarfld
        self.prcNrml = TotalGnrml * 100 // nbGarfld
        self.prcRapd = TotalGrapd * 100 // nbGarfld



    # =====================================================
    # == UPDATE
    # =====================================================
    def update(self):
        """mise à jour des variables (30 fois par seconde)"""
        # on effectue les mises à jour éventuelles
        self.deplacement()
        self.vieillir()
        for l in self.lasagnes :
            l.distance(self.balles)

        if pyxel.btn(pyxel.KEY_DOWN):
            self.scrollingX += 1.5
        if pyxel.btn(pyxel.KEY_UP):
            self.scrollingX -= 1.5

        #on quitte le jeu
        if pyxel.btn(pyxel.KEY_SPACE):
            pyxel.quit()





    # =====================================================
    # == DRAW
    # =====================================================
    def draw(self):
        """création et positionnement des différents objets (30 fois par seconde)"""
        # vide la fenêtre
        pyxel.cls(10)
        #on dessine les lasagnes
        for l in self.lasagnes :
            Dlasagna(l.lx , l.ly)

        # on dessine les balles
        for b in self.balles :
            Dgarfield(b.x , b.y, b.couleur)

        pyxel.rect(cote+1, 0, 95, cote+30, 0)
        pyxel.rect(0, cote+1, cote+1, 30, 0)
        pyxel.text(cote+5,50+int(self.scrollingX),"Generation Precedente:\n" +self.finalText,7)

        pyxel.rect(cote+1, 0, 95, 50, 0)
        pyxel.text(cote+5,10,"Generation: "+str(self.generation),7)
        pyxel.text(cote+5,20,"% Lent   : "+str(self.prcLent),2)
        pyxel.text(cote+5,28,"% Normal : "+str(self.prcNrml),14)
        pyxel.text(cote+5,36,"% Rapide : "+str(self.prcRapd),8)

        if len(self.balles) == 0 :
            self.newGeneration()




def Dgarfield(x, y, clr):
    """Fonction dessinant Garfield"""
    #1ère ligne
    l=-7
    pyxel.rect(x-4, y+l, 4, 1, 0)
    pyxel.rect(x+1, y+l, 3, 1, 0)
    #2nd ligne
    l+=1
    pyxel.rect(x-5, y+l, 1, 1, 0)
    pyxel.rect(x-4, y+l, 4, 1, 9)
    pyxel.rect(x  , y+l, 1, 1, 0)
    pyxel.rect(x+1, y+l, 3, 1, 9)
    pyxel.rect(x+4, y+l, 1, 1, 0)
    #3ème ligne
    l+=1
    pyxel.rect(x-6, y+l, 1, 1, 0)
    pyxel.rect(x-5, y+l, 10, 1, 9)
    pyxel.rect(x+5, y+l, 1, 1, 0)
    #4ème ligne
    l+=1
    pyxel.rect(x-7, y+l, 1, 1, 0)
    pyxel.rect(x-6, y+l, 3, 1, 9)
    pyxel.rect(x-3, y+l, 3, 1, 0)
    pyxel.rect(x  , y+l, 1, 1, 9)
    pyxel.rect(x+1, y+l, 3, 1, 0)
    pyxel.rect(x+4, y+l, 2, 1, 9)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #5ème ligne
    l+=1
    pyxel.rect(x-8, y+l, 1, 1, 0)
    pyxel.rect(x-7, y+l, 3, 1, 9)
    pyxel.rect(x-4, y+l, 1, 1, 0)
    pyxel.rect(x-3, y+l, 3, 1, 10)
    pyxel.rect(x, y+l, 1, 1, 0)
    pyxel.rect(x+1, y+l, 3, 1, 10)
    pyxel.rect(x+4, y+l, 2, 1, 9)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #6ème ligne
    l+=1
    pyxel.rect(x-8, y+l, 1, 1, 0)
    pyxel.rect(x-7, y+l, 3, 1, 9)
    pyxel.rect(x-4, y+l, 1, 1, 0)
    pyxel.rect(x-3, y+l, 3, 1, 10)
    pyxel.rect(x, y+l, 1, 1, 0)
    pyxel.rect(x+1, y+l, 3, 1, 10)
    pyxel.rect(x+4, y+l, 2, 1, 9)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #7ème ligne
    l+=1
    pyxel.rect(x-8, y+l, 1, 1, 0)
    pyxel.rect(x-7, y+l, 1, 1, 9)
    pyxel.rect(x-6, y+l, 3, 1, 10)
    pyxel.rect(x-3, y+l, 1, 1, 7)
    pyxel.rect(x-2, y+l, 3, 1, 0)
    pyxel.rect(x+1, y+l, 1, 1, 7)
    pyxel.rect(x+2, y+l, 2, 1, 0)
    pyxel.rect(x+4, y+l, 2, 1, 10)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #8ème ligne
    l+=1
    pyxel.rect(x-8, y+l, 1, 1, 0)
    pyxel.rect(x-7, y+l, 1, 1, 9)
    pyxel.rect(x-6, y+l, 3, 1, 10)
    pyxel.rect(x-3, y+l, 3, 1, 7)
    pyxel.rect(x, y+l, 1, 1, 0)
    pyxel.rect(x+1, y+l, 3, 1, 7)
    pyxel.rect(x+4, y+l, 2, 1, 10)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #9ème ligne
    l+=1
    pyxel.rect(x-8, y+l, 1, 1, 0)
    pyxel.rect(x-7, y+l, 2, 1, 9)
    pyxel.rect(x-5, y+l, 5, 1, 10)
    pyxel.rect(x, y+l, 2, 1, clr)
    pyxel.rect(x+2, y+l, 4, 1, 10)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #10ème ligne
    l+=1
    pyxel.rect(x-8, y+l, 1, 1, 0)
    pyxel.rect(x-7, y+l, 3, 1, 9)
    pyxel.rect(x-4, y+l, 4, 1, 10)
    pyxel.rect(x, y+l, 2, 1, 9)
    pyxel.rect(x+2, y+l, 3, 1, 10)
    pyxel.rect(x+5, y+l, 1, 1, 9)
    pyxel.rect(x+6, y+l, 1, 1, 0)
    #10ème ligne
    l+=1
    pyxel.rect(x-7, y+l, 2, 1, 0)
    pyxel.rect(x-5, y+l, 10, 1, 9)
    pyxel.rect(x+5, y+l, 1, 1, 0)
    #11ème ligne
    l+=1
    pyxel.rect(x-5, y+l, 10, 1, 0)


def Dlasagna(x , y):
    """Fonction dessinant les délicieuses lasagnes"""
    l=-3
    pyxel.rect(x-4, y+l, 7, 1, 9)
    l+=1
    pyxel.rect(x-4, y+l, 1, 1, 4)
    pyxel.rect(x-3, y+l, 8, 1, 9)
    l+=1
    pyxel.rect(x-4, y+l, 10, 1, 9)
    pyxel.rect(x-3, y+l, 3, 1, 4)
    l+=1
    pyxel.rect(x-4, y+l, 10, 1, 4)
    pyxel.rect(x-3, y+l, 2, 1, 9)
    l+=1
    pyxel.rect(x-4, y+l, 10, 1, 9)
    pyxel.rect(x-3, y+l, 1, 1, 4)
    l+=1
    pyxel.rect(x-3, y+l, 1, 1, 9)
    pyxel.rect(x-2, y+l, 8, 1, 4)
    l+=1
    pyxel.rect(x-2, y+l, 8, 1, 9)

Anim()
