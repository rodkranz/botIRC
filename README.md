# Bot IRC [![Build Status](https://travis-ci.org/rodkranz/botIRC.svg?branch=master)](https://travis-ci.org/rodkranz/botIRC) [![Go Report Card](https://goreportcard.com/badge/github.com/rodkranz/botIRC)](https://goreportcard.com/report/github.com/rodkranz/botIRC)

---
Simple AI for boot IRC

### Execute  

- **RUN**:
    ```bash
    $ bra run
    ```
 
---
#### Build & Development tools

- **[Go Dep](https://github.com/tools/godep) - Golang Dependency Manager**

    - Tool Installation
        ```bash    
        $ go get -u github.com/golang/dep/cmd/dep
        ```

    - Usage (install dependencies)            
        ```bash
        $ dep ensure
        ```


- **[Go Bra](https://github.com/Unknwon/bra) - Brilliant Assistant - Hot Reload** 
    
    - Tool Installation
        ```bash   
        $ go get -u github.com/Unknwon/bra
        ```
        
    - Usage (run hot reload)
        ```bash
        $ bra run
        ```

- **[Go Lint](https://github.com/golang/lint/golint) - Golang Linter**
    
    - Tool Installation
        ```bash
        $ go get -u github.com/golang/lint/golint
        ```
        
    - Usage
        ```bash
        $ golint ./...
        ```
           