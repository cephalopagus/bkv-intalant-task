package departments_transport_http

import (
	"net/http"

	departments_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/service"
)

type Handler struct {
	svc *departments_service.DepartmentService
}

func NewDepartmentHandler(svc *departments_service.DepartmentService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /departments/", h.Create)
	mux.HandleFunc("GET /departments/{id}", h.GetByID)
	mux.HandleFunc("PATCH /departments/{id}", h.Update)
	mux.HandleFunc("DELETE /departments/{id}", h.Delete)
}
