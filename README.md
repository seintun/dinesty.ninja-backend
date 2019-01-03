# dinesty.ninja-backend

## Setup GO developer environment with Homebrew
https://golang.org/doc/install

```sh
$ brew update && brew upgrade 
$ brew install go
```

## Create Workspace Directories

```sh
$ mkdir -p ~/go/{pkg,src,bin}
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

## Installing dependiencies
### Gorilla Mux (HTTP request multiplexer)
primarily used for routing and CRUD methods
### Gorilla Handlers (HTTP middleware)
primarily used for HTTP logging in this project
```sh
$ go get -u github.com/gorilla/mux github.com/gorilla/handlers
```

## Setting up Config and Yelp API
1. Visit https://www.yelp.com/developers/v3/manage_app to obtain your API key
2. Fill-out form under Create New App for testing
3. Copy API Key after completion
4. Change the name of "config/example.config.json" to "config/config.json"
5. Replace "INSERT YOUR YELP API KEY HERE" with your API Key

## Test endpoints with Postman
```sh
$ go run *.go
```
Send POST request to http://localhost:8080/biz/validate with following json data in the request body as raw with JSON(application/json)

```sh
{
	"businessID": "hog-island-oyster-co-san-francisco"
}
```