# dinesty.ninja-backend

## Setup GO developer environment with Homebrew
https://golang.org/doc/install

```sh
$ brew install go
```

## Create Workspace Directories

```sh
$ mkdir -p ~/go/{pkg,src,bin}
$ mkdir -p ~/go/src/github.com/seintun
$ cd ~/go/src/github.com/seintun && git clone https://github.com/seintun/dinesty.ninja-backend.git
```

## Set Environment variables (Depends if you use Bash or Zsh):

Bash
```sh
$ cd ~ && echo "export GOPATH=$HOME/go" >> .bash_profile && source ~/.bash_profile
```
(or)
Zsh
```sh
$ cd ~ && echo "export GOPATH=$HOME/go" >> .zshrc && source ~/.zshrc
```

## Installing dependencies
```sh
$ go get github.com/gorilla/handlers github.com/gorilla/mux gopkg.in/mgo.v2 gopkg.in/mgo.v2/bson github.com/BurntSushi/toml
```
### Gorilla Mux (HTTP request multiplexer)
Used for routing and CRUD Restful API
### Gorilla Handlers (HTTP middleware)
Used for HTTP logging
### MGo (MongoDB)
Used for dialing to MongoDB server
### MGo BSON (BSON)
Used for creating BSON ObjectID
### Gorilla Handlers (HTTP middleware)
Used for reading config.toml file for secret or apikeys

## Setting up Config with mLab (inside ./config.toml)
1. Create/Login account at https://www.yelp.com
2. Visit https://www.yelp.com/developers/v3/manage_app to obtain your API key
3. Fill-out form under Create New App for testing
4. Copy API Key after completion
5. Replace <INSERT YOUR SECRET mLAB userInfo HERE INSIDE THE QUOTES>

## Setting up Config with Yelp API (inside ./config.toml)
1. Create/Login account at https://mlab.com
2. Create a database name: dinesty_ninja_db
3. Replace server="<INSERT YOUR YELP API KEY HERE>" in the format of 
```sh
mongodb://<dbuser>:<dbpassword>@ds0000.mlab.com:0000/dinesty_ninja_db
```

## Run the GO server at port 8080
```sh
$ go run main.go
```