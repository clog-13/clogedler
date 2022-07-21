package main

import (
	"clogedler/business"
	"clogedler/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	business.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
