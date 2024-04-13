# Info

### Machine

git ne supporte pas de dossier vide donc convention : faire un fichier .gitkeep

Ctrl J : ouvrir terminal dans VSCode

retenir mdp et user name:

```bash
git config --global credential.helper store
```

### Configuration mdp et email

```bash
git config --global user.name "E232826X"
git config --global user.mail "E232826X@etu.univ-nantes.fr"
```

Si conflit :

```bash
git config pull.rebase false
```

ordre :

```bash
 git add .
 git commit -m "message"
 git pull
 git push
```

Ctrl Shift V : Open Preview

poubelle commit :
creer .gitignore puis mettre tous les chemin qu'on ne veut pas commit


pour plus tard : git tag -> release
