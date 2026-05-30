package employee_service

import (
	"fmt"
	"strings"
	"time"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	employee_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/repository/postgres"
	"go.uber.org/zap"
)

type EmployeeService struct {
	empRepo  *employee_repository_postgres.EmployeeRepository
	deptRepo *departments_repository_postgres.DepartmentRepository
	logger   *core_logger.Logger
}

func NewEmployeeService(
	empRepo *employee_repository_postgres.EmployeeRepository,
	deptRepo *departments_repository_postgres.DepartmentRepository,
	logger *core_logger.Logger) *EmployeeService {
	return &EmployeeService{empRepo: empRepo, deptRepo: deptRepo, logger: logger}
}

func (s *EmployeeService) Create(deptID uint, fullName, position string, hiredAt *time.Time) (*model.Employee, error) {
	s.logger.Info("creating employee",
		zap.Uint("department_id", deptID),
		zap.String("full_name", fullName),
	)

	fullName = strings.TrimSpace(fullName)
	position = strings.TrimSpace(position)

	if fullName == "" {
		return nil, fmt.Errorf("%w: full_name is required", core_errors.ErrBadRequest)
	}
	if position == "" {
		return nil, fmt.Errorf("%w: position is required", core_errors.ErrBadRequest)
	}

	dept, err := s.deptRepo.GetByID(deptID)
	if err != nil {
		s.logger.Error("failed to get department", zap.Uint("id", deptID), zap.Error(err))
		return nil, err
	}
	if dept == nil {
		return nil, fmt.Errorf("%w: department", core_errors.ErrNotFound)
	}

	emp := &model.Employee{
		DepartmentID: deptID,
		FullName:     fullName,
		Position:     position,
		HiredAt:      hiredAt,
	}
	if err := s.empRepo.Create(emp); err != nil {
		s.logger.Error("failed to create employee", zap.Error(err))
		return nil, err
	}

	s.logger.Info("employee created", zap.Uint("id", emp.ID), zap.Uint("department_id", deptID))
	return emp, nil
}
