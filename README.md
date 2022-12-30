# IA04 - Poker

| Information | Valeur |
| :------------ | :------------- |
| **Auteurs** | Théo Claussmann |
| | Thomas Bossuat |
| | Louis Wahart |
| | Paul Nivoix |
| **UV** | IA04 |

# Explication du projet 

## 1) Explication du fonctionnement du projet

### « Qu'est-il censé faire ? »
 
Notre projet a pour but de représenter différentes parties de poker. Le type de jeu de poker utilisé est le Texas Hold’em. Ici, on représente une partie avec 5 joueurs autour de la table pour chacune d’entre elles. Chaque tour consiste en une mise successive de chacun des joueurs en fonction des cartes qu’il a en sa possession. D’abord avec uniquement les 2 cartes en main, ensuite avec les 3,4 et 5 cartes présentes au centre de la table représentant un tour de jeu à chaque fois. A la fin de ces 4 tours de jeu, soit il y a eu un vainqueur auparavant ayant effectué une mise supérieure aux autres, ceux-ci n’ayant pas suivi. Soit le joueur ayant la meilleure combinaison de 5 cartes sur les 7 disponibles entre les joueurs restants gagne la partie. Ainsi notre projet est une simulation de parties de poker, un certain nombres de tables sont crées (choix fait par l'utilisateur) et chaque table gère les parties avec ses 5 joueurs.
 
Concernant les caractéristiques du projet, celles-ci sont liées en grande partie aux spécificités des joueurs.  En effet, dans le but de faire jouer les différents joueurs de manière assez réaliste, nous avons décidé de mettre en place des attributs de jeu à chacun d’eux. On retrouve donc 4 attributs qui définissent chaque joueur, la timidité, qui décrit la tendance d'un joueur à juste suivre la mise actuelle ou à augmenter la mise. Ensuite, l’agressivité, décrivant à quel point le joueur fait monter la mise quand il joue, le risque, qui décrit la tendance d'un joueur à jouer (continuer de miser) selon la puissance de sa main (plus le risque sera élevé, plus il jouera même avec une main faible). Enfin la dernière caractéristique repose sur la capacité de bluff, c’est-à-dire la tendance d'un joueur à jouer (continuer de miser) alors qu'il ne devrait peut être pas.
 
### « Que sommes-nous censés observer ? »
 
Au niveau de l’aperçu sur le front, nous sommes censés observer le déroulement, tour par tour, d’une ou plusieurs parties de poker sur différentes tables. Ainsi on peut choisir le nombre de tables et de parties que l’on souhaite lancer et observer chacune des tables à tout moment en sélectionnant celle de notre choix. Nous pouvons aussi choisir d’observer un joueur en particulier et découvrir ses statistiques qui s’affichent en dessous du choix du joueur. Nous pouvons modifier les statistiques du joueur afin de les régler à notre guise.

Au niveau de la table, on observe les cartes des joueurs de la table en cours et les cartes de la table à chaque tour. Un nouveau tour de jeu est réalisé toutes les 5 secondes (possiblement plus si il y a beaucoup de tables et donc beaucoup de calculs à réaliser). Nous avons choisi de n'afficher les résultats que tour par tour (et non action par action) afin que cela soit visuellement plus compréhensible et compact. Toutes les tables sont synchronisées, ainsi même si un vainqueur est désigné avant la fin d'une partie, la table attendra la fin de tous les tours de la partie pour passer à la suivante (et affichera donc les cartes qui auraient été affichées).

Afin de mieux comprendre le déroulement de chaque tour, l’interface affiche les mises par tour et par partie des joueurs ainsi que son action en cours, à la fin d'une partie est également affiché les bénéfices réalisés sur celle-ci. Les joueurs réalisant une petite ou grosse blind ont également une indication visuelle. Le gagnant d’une table voit son fond passer au vert tandis qu’un joueur n’ayant plus de jeton pour jouer passe au rouge. Au centre de la table, en plus de l’affichage des cartes, un point sur l’état actuel de la table, notamment le tour, la partie en cours et le pot total, est affiché. Il est possible à tout moment d'une partie de mettre celle-ci en pause, ou alors de reset et recommencer une configuration de simulation.
 
### « Quelle est la problématique à laquelle nous essayons de répondre ? »
 
Finalement la problématique à laquelle nous essayons de répondre est la suivante : Quel est l'impact des caractéristiques d'un joueur sur ses résultats dans des parties de poker multi-agents ? Comment le joueur et ses caractéristiques peuvent influencer les résultats des autres joueurs ? 

Ainsi nous pourrions dans le futur faire en sorte de réaliser des statistiques grâce à notre simulation de poker pour trouver quelles caractéristiques permettent de maximiser ses chances de victoire, les corrélations avec les caractéristiques des autres joueurs, voir si certaines sont très bonne face à certain type de joueur mais moins bonne face à d'autres, etc.
 
## 2) Installation et architecture du projet

### Installation et lancement du projet
*Remarque :* Toutes les commandes données seront à réaliser depuis le dossier racine du projet (ia04-poker). Nous donnerons les chemins pour Windows, à modifier pour Linux si nécessaire.

#### Installation
Le backend de notre projet est réalisé en GO et le frontend en React. 

Ainsi pour installer le projet GO et node.js sont requis. 
(à télécharger sur le site officiel : 
GO : https://go.dev/dl/  
Node.js : https://nodejs.org/en/download/)

Une fois cela fait, pour installer le frontend il faut installer les modules nécessaires via la commande suivante :

**cd .\front ; npm install**

Pour installer le backend il suffit de lancer la commande suivante (Rappel : les commandes doivent être lancées depuis le dossier racine du projet (ia04-poker)) :

**go install .\cmd\launch-server.go**

#### Lancement

Pour lancer le backend il suffira de faire :

**launch-server**
ou
**go run .\cmd\launch-server**

Pour lancer le frontend il suffira de faire :

**cd .\front ; npm start .\front**

La backend et le frontend fonctionneront en local. L'adresse du frontend doit obligatoirement être http://localhost:3000 pour que tout fonctionne correctement (valeur par défaut via npm start). Vous pourrez ainsi vous rendre via le navigateur à l'adresse http://localhost:3000 (s'ouvre normalement automatiquement avec npm start) pour avoir accès au frontend et agir sur celui ci.

### Architecture
Notre architecture peut être décomposée en 4 grandes parties, le backend composé des agents joueurs, des agents tables et de l'agent serveur et du frontend. Voici une brève description de chacune de ses parties. 

#### Agents Joueurs
Explication brève du joueur

#### Agents Tables
Explication brève de la table

#### Agent Serveur
Explication brève du serveur

#### Front
Explication brève du Front

 
## 3) Discussions

### Points positifs et négatifs / Améliorations possibles

Positifs :
Réalisation totalement fonctionnelle
Parties respectant les vrais règles
Joueur avec une certaine intelligence et réalisant de réelles actions
Communication importante entre table et joueur
Front fonctionnel, interaction avec l'utilisateur (choix du nb de tables, de parties, changement de table, de joueur, modification des stats, pause, reset, etc)
Indications visuelles

Négatifs / Améliorations :
Travail plus complexe sur l'intelligence de jeu d'un joueur, moins utiliser le hasard
Création de joueur ayant la capacité de tricher (soit connaissant toutes les cartes ou certains peuvent communiquer entre eux)
Amélioration visuelle du front
Encore plus de possibilités à l'utilisateur (changer les joueurs de tables, etc)

### Analyse des méta-paramètres
Nombre de tables
Nombre de parties
Timidité
Agressivité
Risque
Bluff
Nb Jetons
Blind
