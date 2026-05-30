package departments_service_test

import (
	"errors"
	"testing"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	departments_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/service"
)

func TestCreate_EmptyName(t *testing.T) {
	svc := departments_service.NewDepartmentService(nil, core_logger.NewNop())

	_, err := svc.Create("   ", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, core_errors.ErrBadRequest) {
		t.Fatalf("expected ErrBadRequest, got: %v", err)
	}
}
