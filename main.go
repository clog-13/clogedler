package main

import (
	"clogedler/business"
	"clogedler/framework"
	"clogedler/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()
	// core.Use(
	// 	middleware.Test1(),
	// 	middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	// core.Use(middleware.Timeout(1 * time.Second))

	business.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
