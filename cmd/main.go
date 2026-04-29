package main

import (
	"log"
	"net/http"

	"github.com/su1uv/boxrec/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	//timeout := time.Duration(env.ContextTimeout) * time.Second

	log.Fatal(http.ListenAndServe(env.ServerAddress, nil))
}
