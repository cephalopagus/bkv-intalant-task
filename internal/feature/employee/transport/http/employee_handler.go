package employee_transport_http

import (
	"encoding/json"
	"net/http"
	"time"

	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	core_response "github.com/cephalopagus/bkv-intalant-task/internal/core/response"
	employee_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/service"
	"go.uber.org/zap"
)

type EmployeeHandler struct {
	svc    *employee_service.EmployeeService
	logger *core_logger.Logger
}

func NewEmployeeHandler(svc *employee_service.EmployeeService, logger *core_logger.Logger) *EmployeeHandler {
	return &EmployeeHandler{svc: svc, logger: logger}
}

func (h *EmployeeHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /departments/{id}/employees/", h.Create)
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	deptID, ok := core_response.PathUint(r, "id")
	if !ok {
		core_response.WriteError(w, http.StatusBadRequest, "invalid department id")
		return
	}

	var body struct {
		FullName string  `json:"full_name"`
		Position string  `json:"position"`
		HiredAt  *string `json:"hired_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		core_response.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	var hiredAt *time.Time
	if body.HiredAt != nil {
		t, err := time.Parse("2006-01-02", *body.HiredAt)
		if err != nil {
			h.logger.Error("failed to parse hired_at", zap.String("hired_at", *body.HiredAt))
			core_response.WriteError(w, http.StatusBadRequest, "hired_at must be YYYY-MM-DD")
			return
		}
		hiredAt = &t
	}

	emp, err := h.svc.Create(deptID, body.FullName, body.Position, hiredAt)
	if err != nil {
		core_response.WriteServiceError(w, err)
		return
	}
	core_response.WriteJSON(w, http.StatusCreated, emp)
}
