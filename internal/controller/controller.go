package controller

import (
	"awesomeProject/internal/service"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type PortService interface {
	Read(id int) (int, error)
	Write(id, value int) error
}

type PortController struct {
	PortService PortService
}

func (p *PortController) Read(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusBadRequest)
		// Не использовал структуры для экономии времени
		render.JSON(w, r, map[string]string{"error": "ID is required"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "ID must be an integer"})
		return
	}
	value, err := p.PortService.Read(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Port not found"})
		return
	}
	render.JSON(w, r, map[string]int{"value": value})
}

func (p *PortController) Write(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id    int `json:"id"`
		Value int `json:"value"`
	}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request"})
		return
	}
	if err := p.PortService.Write(req.Id, req.Value); err != nil {
		if errors.Is(err, service.ErrBadValue) {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Port not found"})
		return
	}
	fmt.Printf("OUT порт %d получил значение %d\n", req.Id, req.Value)
	render.JSON(w, r, map[string]string{"status": "ok"})
}

func NewPortController(portService PortService) *PortController {
	return &PortController{
		PortService: portService,
	}
}
