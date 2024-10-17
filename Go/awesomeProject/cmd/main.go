package main

import (
	"awesomeProject/internal/controller"
	"awesomeProject/internal/repository/port"
	"awesomeProject/internal/service"
	"awesomeProject/pkg"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	portSizes, err := pkg.GetPortSizes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("IN port size: %d\nOUT port size: %d\n", portSizes["IN"], portSizes["OUT"])
	portRepo := port.NewPortSystem(portSizes["IN"], portSizes["OUT"])
	portService := service.NewPortService(portRepo)
	portController := controller.NewPortController(portService)

	router := chi.NewRouter()
	router.Get("/read/{id}", portController.Read)
	router.Post("/write", portController.Write)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
