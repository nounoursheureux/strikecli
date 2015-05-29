# strikecli

Un client [getstrike](https://getstrike.net/torrents) en ligne de commande.

### Installation

`go get github.com/nounoursheureux/strikecli`

### Utilisation

Il n'y a pour le moment que deux commandes: `search` et `info`.

#### `search`

Recherche les mots-clés.       
Syntaxe: `strikecli search [mots-clés]`         
Options:            
- `--limit [nombre]` ou `-l [nombre]`: limite le nombre maximum de réponses
- `--category [catégorie]` or `-c [catégorie]`: filtre par catégorie
- `--format [pattern]` or `-f [pattern]`: applique un pattern en remplaçant les [séquences suivantes](#Formatting)
- `--script-mode`: supprime la 1e ligne qui indique le nombre de résultas et les mots-clés utilisés

#### `info`

Retourne les informations à propos de torrents particuliers.        
Syntaxe: `strikecli info [hashes]`          
Options: voir ci-dessus       

#### Formatting

Le programme va chercher et remplacer les séquences suivantes: 
- `%t`: le titre du torrent
- `%m`: le lien magnet du torrent
- `%h`: le hash du torrent
- `%S`: le nombre de seeders
- `%l`: le nombre de leechers
- `%s`: la taille du torrent
- `%d`: la date de mise en ligne du torrent
- `%c`: la catégorie du torrent
- `%sc`: la sous-catégorie du torrent
- `%f`: le nombre de fichiers du torrent
- `%u`: le nom d'utilisateur de l'uploader du torrent
