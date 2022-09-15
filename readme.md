## A Simple CLI based TO-DO application using golang and cobra


-------------------------------------------
Tri is a Todo cli library built using Golang and Cobra,
that helps to CRUD, search todo's
using your terminal. Created using Go,
based on the workshop by [spf13 - Building an Awesome CLI App in Go – OSCON 2017](https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/)

-------------------------------------------

Usage:
    
```
tri [command]

Available Commands:
  add         Add a new TODO
  completion  Generate the autocompletion script for the specified shell
  done        Marks todo as done
  edit        Edit a given todo
  help        Help about any command
  list        List all Todos
  search      Searches all todos using a keyword

Flags:
  --config string   config file override (Default is $HOME/.tri.yml),
  -h, --help            help for tri
```

## Installation
- Build from source
```
git clone
cd tri
go build main.go
``` 

## Work Remaining
- [ ] Add tests
- [ ] Add more commands
- [ ] Fix Edit command
- [x] Add Database support 
