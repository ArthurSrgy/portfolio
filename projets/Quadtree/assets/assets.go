package assets

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed newfloor.png
var floorBytes []byte

// FloorImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents éléments qui peuvent s'afficher au sol
// (herbe, sable, etc).
// Dans la version du projet qui vous est fournie, ces éléments sont des
// carrés de 16 pixels de côté. Vous pourrez changer cela si vous le voulez.
var FloorImage *ebiten.Image

//go:embed tinySLATESbyIvanVoirol.png
var FusedTileBytes []byte

// FusedTilesImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents tuiles fusionné au sol fournis.
var FusedTilesImage *ebiten.Image

//go:embed character.png
var characterBytes []byte

//go:embed ghost.png
var ghostBytes []byte

var GhostImage *ebiten.Image

// CharacterImage contient une version compatible avec Ebitengine de
// l'image qui contient les différentes étapes de l'animation du
// personnage.
// Dans la version du projet qui vous est fournie, ce personnage tient
// dans un carré de 16 pixels de côté. Vous pourrez changer cela si vous
// le voulez.
var CharacterImage *ebiten.Image

//go:embed particule.png
var particuleBytes []byte

// ParticuleImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents particule pouvant s'afficgher
// (empreinte, pluie, neige...)
var ParticuleImage *ebiten.Image

// Load est la fonction en charge de transformer, à l'exécution du programme,
// les images du jeu en structures de données compatibles avec Ebitengine.
// Ces structures de données sont stockées dans les variables définies ci-dessus.
func Load() {
	decoded, _, err := image.Decode(bytes.NewReader(floorBytes))
	if err != nil {
		log.Fatal(err)
	}
	FloorImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(FusedTileBytes))
	if err != nil {
		log.Fatal(err)
	}
	FusedTilesImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(characterBytes))
	if err != nil {
		log.Fatal(err)
	}
	CharacterImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(ghostBytes))
	if err != nil {
		log.Fatal(err)
	}
	GhostImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(particuleBytes))
	if err != nil {
		log.Fatal(err)
	}
	ParticuleImage = ebiten.NewImageFromImage(decoded)
}
