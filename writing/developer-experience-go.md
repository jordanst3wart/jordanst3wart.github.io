---
title: Golang Global
date: '2024-2-21'
tags: ['go', 'developer experience', 'golang', 'Typescript']
summary: 'Go has an amazing developer experience when compared to Typescript'
---

I find the development process with TypeScript painful in every step of development due to divergent tooling. Conversely, I find using go borrowing as it's very simple in comparison.

You can experience this enjoy by simply quickly setting up a basic webserver, as you go through a lot of the steps in the development process. So here we go...

How fast can you deploy a server in TypeScript from scratch?

3... 2... 1... go!

### 1. Initial setup:

```
mkdir go-deployment-minimal
cd go-deployment-minimal
go mod init github.com/jordanst3wart/go-deployment-minimal
```

This creates the go.mod file, which stores the version of go, the package name, and dependencies.

(Now we go straight to coding with Typescript this is step 5!!!)

### 2. write a web server

Create a main file.

```shell
touch main.go
```

And let's hit chatgpt to quickly get some code:

```Go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	// Respond to the request
	_, err := fmt.Fprintf(w, "Hello, world!")
	if err != nil {
		return
	}
}

func main() {
	serverAddress := ":8080"
	http.HandleFunc("/", helloHandler)
	log.Printf("Starting server on %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		panic(err)
	}
}

```

We are using the default go logger, and webserver, which are very popular.

We are also more aware of what errors could occur with writing to the http response.

We could create a `src/` directory, but at the moment that feels a bit unnecessary.

### compile and run

compile and run the code with:

```shell
go run main.go
```

You can compile an arm binary with:

```shell
GOOS=linux GOARCH=arm64 go build -o main
```

The flags of `GOOS` and `GOARCH` look a tiny bit intimidating, and you should compile a binary in a pipeline not locally on a laptop. But this is fine.

### deploy it - it's not done until it is running in production

Let's make the binary executable:

```shell
chmod +x main
```

Let's create a tiny EC2 instance on amazon web services (AWS), like t4g.nano spot instance. We will use ssh, and scp to setup the server.

Let's make an app directory (I can't to bothered to ssh):

```shell
ssh -i ~/.ssh/ec2-key.pem ec2-user@52.64.95.107 'mkdir app'
```

Let's copy the binary to the server:

```shell
scp -i ~/.ssh/ec2-key.pem main ec2-user@52.64.95.107:/home/ec2-user/app
```

Let's run the binary on the server (once again, I can't be bothered to ssh):

```shell
ssh -i ~/.ssh/ec2-key.pem ec2-user@52.64.95.107 'nohup ./app/main > ./app/application.log &'
```

That shell commands isn't perfect as it doesn't exit... but the go binary keeps running.

### thoughts on developer experience

And that's it, we have a server listening on port 8080, serving content. As a bonus you get `go fmt` to format your code.

I can deploy a web server extremely quickly, I didn't even use an IDE, and I feel no confusion or frustration.

Let's compare this to the minimistic typescript server I made before:

- instead of a `package.json`, `package-lock.json`, `tsconfig.json` (and maybe .prettier, and .eslintrc) we have a `go.mod` file that was auto generated as is
- the `go.mod` is 3 lines long compared to the 100 lines we started with in the tsconfig.json file, and the 21 lines of the package.json
- we have really well-defined error conditions with go, rather than the catch-all of the TypeScript server
- the binary has to be compiled for ARM, but that is easy with two flags
- I don't need to install anything on the server, instead of having to install a node version manager, node, and the production packages required for Typescript
- I can copy a single binary to the server instead of two files, and a directory (`dist` directory, `package.json`, and `package-lock.json`)
- I can write a test as `main_test.go` and run that test with `go test` as well. I don't need a tool, and a config file like with `jest`, or `vitest`
- instead of interacting with `npm`, `node`, `tsc`, and potentially `vitest`, I call `go mod ...`, `go run ...`, `go build ...` and `go get ...`

It's a very stream lined experience, and it's so simple that it's quiet boring.

It's removed several steps in the process, and simplified a lot of the others.
