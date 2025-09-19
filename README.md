# Projet Red - Jeu tour par tour en Go

## Description
*Projet Red* est un jeu tour par tour inspiré de *Matrix*, développé en **Go**.  
Le joueur incarne **Néo**, le personnage principal, et affronte différents adversaires dans des combats d’arts martiaux.

---

## Fonctionnalités

### Menu
- Titre du jeu  
- Sélection des arènes / niveaux  
- Boutons interactifs : **PLAY**, **FORGE**, **EXIT**, **NOTICE**  
- Arrière-plan : cascade de code vert rappelant l’univers de Matrix

![Menu Matrix](asset\images\presentation du readme\menu matrix.png)

### Marchand
- Achat de pilules **Rouge** et **Bleue**  
- Description et prix des pilules  
- Effets sur le combat :  
  - **Pilule Rouge** : +1 point dans la manche  
  - **Pilule Bleue** : enlève 1 point à l’adversaire  

### Inventaire
- Affichage et gestion des pilules achetées  
- Arts martiaux disponibles :  
  - Boxe  
  - Judo  
  - Brazilian Jiu-Jutsu  
  - Karate  
  - Wrestling  

### Équipement
- Équipement de Néo : tenue Wrestling, tenue Karate  

### Combat
- Tour par tour basé sur un système “pierre-feuille-ciseaux” adapté aux arts martiaux  
- Système de manches et points  
- Utilisation des pilules pour modifier les scores  
- Affichage des **rounds** et du résultat du combat  
- Affichage des règles via le bouton **NOTICE**  

### Forge
- Fabrication et déblocage d’arts martiaux  
- Gestion des améliorations et objets spéciaux  

![Forge Matrix](asset\images\presentation du readme\forge matrix.png)


### Interface
- Boutons interactifs avec détection des clics pour éviter les répétitions  
- Notice avec **contour vert fluo** pour afficher les règles de victoire  
- Transition fluide entre les arènes et le menu  

---

## Combat : règles de victoire
- **Boxe** → Victoire : Judo, Jujutsu | Défaite : Karate, Lutte  
- **Judo** → Victoire : Jujutsu, Karate | Défaite : Lutte, Boxe  
- **Jujutsu** → Victoire : Karate, Lutte | Défaite : Boxe, Judo  
- **Karate** → Victoire : Lutte, Boxe | Défaite : Judo, Jujutsu  
- **Lutte** → Victoire : Boxe, Judo | Défaite : Jujutsu, Karate  

---

## Instructions d’installation et lancement
1. Cloner le dépôt :  
   ```bash
   git clone https://github.com/totolasticot-ynov/projet-red_trinity.git
   cd projet-red_trinity
   go run main.go