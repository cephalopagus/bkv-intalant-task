package departments_transport_http

import (
	"net/http"

	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	departments_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/service"
)

type Handler struct {
	svc    *departments_service.DepartmentService
	logger *core_logger.Logger
}

func NewDepartmentHandler(
	svc *departments_service.DepartmentService,
	logger *core_logger.Logger) *Handler {
	return &Handler{svc: svc, logger: logger}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /departments/", h.Create)
	mux.HandleFunc("GET /departments/{id}", h.GetByID)
	mux.HandleFunc("PATCH /departments/{id}", h.Update)
	mux.HandleFunc("DELETE /departments/{id}", h.Delete)
}
