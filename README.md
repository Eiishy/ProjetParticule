# R1.01.SAE.eq09_JESUS-Zacchary_ROBINEAU-Anton

## Name
Project_Particle

## Description
Creation of a particle generator in the go language using a 2D engine named Ebiten

## Pre-requisites
- Golang - https://go.dev

## Setup
Install the source code on gitlab directly or through the terminal with the following command:<br>
```
git clone https://gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae/projets/groupe2/r1.01.sae.eq09_jesus-zacchary_robineau-anton.git
```
Once installed, you will have to go to the terminal and go to the installer folder before and execute the following commands:
```
go install
go build
```

## Run

 If you want to launch a particle generation and you have already completed the previous steps, you just have to launch the file 'project-particles.exe'

## Test 

 If you want to run the test set from the particles folder in the console run the command go test
 The test set will evolve during the development of this project 

## Config.json
> ### <ins>Window configuration</ins>
> - WindowTitle | Choose the name of the window.
> - Debug | Activate the information display at the top left of the screen (true/false)
> - Fullscreen | Choose if the application launches in full screen (true/false) coming soon
> #### Window size
> - WindowSizeX | Window length (in pixels)
> - WindowSizeY | Window widht (in pixels)
> ### <ins>Particle configuration</ins>
> - ParticleImage | Image display for particles (path)
> - SpawnRate | Number of particles at each call of the update function(60/s)
> - InitNumParticles | The number of particles on the screen at startup
> - RandomSpawn | Appearance of particles at a random location (true/false)
> - Gravity | Adds a gravity effect to the particles (true/false)
> - LiveSpanActivator | Activates a lifetime for the particles (true/false)
> - LiveSpan | Defines the lifetime of the particles (60 pour 1s)
> ### <ins>Gamemod</ins>
> - 0 For immobile particles
> - 1 For particles with random velocity 
> - 2 For particles that form a circle but need a RandomSpawn on false and an InitNumParticle > 1000 for a nice result 
> #### Only if RandomSpawn is false
> - SpawnX | Management of the X coordinate of the particles in the spawn (from 0 to length of the window in pixel)
> - SpawnY | Management of the Y coordinate of the particles in the spawn (from 0 to width of the window in pixel)

## Visual

![](/Projet_Particule/assets/screen1.png)

![](/Projet_Particule/assets/screen2.png)

More visuals to come

## Project status
This project is still under development.