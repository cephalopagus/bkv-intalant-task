package employee_service

import (
	"fmt"
	"strings"
	"time"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	employee_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/repository/postgres"
)

type EmployeeService struct {
	empRepo  *employee_repository_postgres.EmployeeRepository
	deptRepo *departments_repository_postgres.DepartmentRepository
}

func NewEmployeeService(empRepo *employee_repository_postgres.EmployeeRepository, deptRepo *departments_repository_postgres.DepartmentRepository) *EmployeeService {
	return &EmployeeService{empRepo: empRepo, deptRepo: deptRepo}
}

func (s *EmployeeService) Create(deptID uint, fullName, position string, hiredAt *time.Time) (*model.Employee, error) {
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
	return emp, s.empRepo.Create(emp)
}
