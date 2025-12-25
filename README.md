# A simple bencode decoder in golang, supports all bencode representations (or the ones that I'm aware of)

## Steps to install

clone the repo and cd into it
```
make build
```
after build you'll see a `bengoder` executable in your repo


then install the package globally while you're in the repo

```
go install
```

this should add the executable with path `/home/username/go/bin/bengoder` you can append this to your global path

for bash/zsh
```
$ export PATH=$PATH:/path/to/your/install/directory
```

for fish
```
fish_add_path /path/to/your/install/directory
```

to verify installation
```
which bengoder
```
this will throw out the path


## Usage

example
```
bengoder -path=./fixture/.torrent
```
