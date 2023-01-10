# R1.01.SAE.eq09_JESUS-Zacchary_ROBINEAU-Anton

## Name
Project_Particule

## Description
Création d'un générateur de particules dans le language go à l'aide d'un moteur 2D nommée Ebiten

## Pré-requis
- Golang - https://go.dev

## Installation
Installer le code source sur gitlab directement ou en passant par par le terminal avec la commande suivante :<br>
```
git clone https://gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae/projets/groupe2/r1.01.sae.eq09_jesus-zacchary_robineau-anton.git
```
Une fois installer, il faudra ce rendre dans le terminal et aller dans le dossier installer auparavant et executer les commandes suivantes :
```
go install
go build
```

## Execution

 Si vous désirez lancer une génération de particule et que vous avez déjà accomplie les étapes précédentes, il ne vous reste plus qu'a lancer le fichier 'project-particles.exe'

## Test 

 Si vous désirez lancer le jeu de tests depuis le dossier particles dans la console executer la commande go test
 Le jeu de test evoluera au cours du developpement de ce projet 


## Config.json
> ### <ins>Configuration de la fenêtre</ins>
> - WindowTitle | Choisir le nom de la fenêtre.
> - Debug | Activer l'affichage des informations en haut à gauche de l'écran (true/false)
> - Fullscreen | Choisir si l'application se lance en plein écran (true/false) à venir
> #### Taille de la fenêtre
> - WindowSizeX | Longueur de la fenêtre (en pixel)
> - WindowSizeY | Largeur de la fenêtre (en pixel)
> ### <ins>Configuration des particules</ins>
> - ParticleImage | L'image afficher pour les particules (chemin)
> - SpawnRate | Nombre de particule a chaque appel de la fonction update(60/s)
> - InitNumParticles | Le nombre de particule à l'écran au démarrage
> - RandomSpawn | Apparition des particules à un endroit aléatoire (true/false)
> - Gravity | Rajoute un effet de gravité au particules (true/false)
> ### <ins>Gamemod</ins>
> - 0 Pour des particules imobiles
> - 1 Pour des particules avec une vitesse aléatoire 
> - 2 Pour des particules qui forment un cercle mais necessite un RandomSpawn sur false et un InitNumParticle > 1000 pour un joli rendu 
> #### Seulement quand RandomSpawn est false
> - SpawnX | Gestion de la coordonnée X des particules au spawn (de 0 à longueur de la fenêtre en pixel)
> - SpawnY | Gestion de la coordonnée Y des particules au spawn (de 0 à largeur de la fenêtre en pixel)

## Visuels

![](/Projet_Particule/assets/screen1.png)

![](/Projet_Particule/assets/screen2.png)

Plus de visuels à venir

## Project status
Ce projet est toujours en cours de développement.