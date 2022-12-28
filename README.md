# IA04 - Poker

## Explication du projet 

### 1) Une partie expliquant le fonctionnement du projet, i.e. «  Qu'est-il censé faire ? », « Que sommes-nous censés observer ? », « Quelle est la problématique à laquelle nous essayons de répondre ? »

#### « Qu'est-il censé faire ? »
 
Notre projet a pour but de représenter différentes parties de poker. Le type de jeu de poker utilisé est le Texas Hold’em. Ici, on représente une partie avec 5 joueurs autour de la table pour chacune d’entre elles. Chaque tour consiste en une mise successive de chacun des joueurs en fonction des cartes qu’il a en sa possession. D’abord avec uniquement les 2 cartes en main, ensuite avec les 3,4 et 5 cartes présentes au centre de la table représentant un tour de jeu à chaque fois. A la fin de ces 4 tours de jeu, soit il y a eu un vainqueur auparavant ayant effectué une mise supérieure aux autres, ceux-ci  n’ayant pas suivi. Soit le joueur ayant la meilleure combinaison de 5 cartes sur les 7 disponibles entre les joueurs restants gagne la partie. 
 
Concernant les caractéristiques du projet, celles-ci sont liées en grande partie aux spécificités des joueurs.  En effet, dans le but de faire jouer les différents joueurs de manière assez réaliste, nous avons décidé de mettre en place des attributs de jeu à chacun d’eux. On retrouve donc 4 attributs qui définissent chaque joueur, la timidité, qui décrit la tendance d'un joueur à juste suivre la mise actuelle ou à augmenter la mise. Ensuite, l’agressivité, décrivant à quel point le joueur fait monter la mise quand il joue, le risque, qui décrit la tendance d'un joueur à jouer (continuer de miser) selon la puissance de sa main (plus le risque sera élevé, plus il jouera même avec une main faible). Enfin la dernière caractéristique repose sur la capacité de bluff, c’est-à-dire la tendance d'un joueur à jouer (continuer de miser) alors qu'il ne devrait peut être pas.
 
#### « Que sommes-nous censés observer ? »
 
Au niveau de l’aperçu sur le front, nous sommes censés observer le déroulement, tour par tour, d’une ou plusieurs parties de poker sur différentes tables. Ainsi on peut choisir le nombre de tables et de parties que l’on souhaite lancer et observer chacune des tables à tout moment. Nous pouvons aussi choisir d’observer un joueur en particulier et découvrir ses statistiques qui s’affichent en dessous du choix du joueur. Au niveau de la table, on observe les cartes des joueurs de la table en cours et les cartes de la table à chaque tour. Afin de mieux comprendre le déroulement de chaque tour, l’interface affiche les mises par tour et et par partie des joueurs ainsi que son action en cours. Le gagnant d’une table voit son fond passer au vert tandis qu’un joueur n’ayant plus de jeton pour jouer passe au rouge. Au centre de la table, en plus de l’affichage des cartes, un point sur l’état actuel de la table, notamment le tour et la partie en cours, est affiché. 
 
 
#### « Quelle est la problématique à laquelle nous essayons de répondre ? »
 
Finalement la problématique à laquelle nous essayons de répondre est la suivante : comment mettre en place une ou plusieurs parties de poker informatiquement tout en créant une interface visuelle et compréhensible pour l’utilisateur ?
 
### 2) Une partie détaillant comment installer et lancer votre projet, ainsi qu’une brève description de votre architecture
 
### 3) Une partie contenant les points positifs et négatifs de votre projet, servant aussi de discussions sur l’analyse de variations de méta-paramètres de votre projet.




## Comment installer et lancer le projet

### Description brève de l'architecture



## Discussions

### Points positifs et négatifs

### Analyse des méta-paramètres