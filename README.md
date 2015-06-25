# strikecli

Use [getstrike](https://getstrike.net/torrents) search from command line.

### Install

If you want an executable see the [releases](https://github.com/nounoursheureux/strikecli/releases) page. Or you can compile it yourself:        

`go get github.com/nounoursheureux/strikecli`

### Usage 

There are currently two commands: `search` and `info`             

#### `search`

Search for the given keywords.         
Basic syntax: `strikecli search [keywords]`         
Options:          
- `--limit [number]` or `-l [number]`: set the maximum number of responses
- `--category [category name]` or `-c [category name]`: filter by category
- `--format [format string]` or `-f [format string]`: format the output by matching the following [sequences](#Formatting)

#### `info`

Get informations on the given hashes.    
Basic syntax: `strikecli info [torrents hashes]`         
Options: see above

#### Formatting

The program will match and replace the following sequences:        
- `%t`: the torrent title
- `%m`: the torrent magnet link
- `%h`: the torrent hash
- `%S`: the number of seeders
- `%l`: the number of leechers
- `%s`: the torrent size
- `%d`: the torrent upload date
- `%c`: the torrent category
- `%sc`: the torrent sub category
- `%f`: the torrent file count
- `%u`: the uploader username
