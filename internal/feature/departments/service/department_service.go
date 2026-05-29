package departments_service

import (
	"errors"
	"fmt"
	"strings"

	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrConflict   = errors.New("conflict")
	ErrBadRequest = errors.New("bad request")
)

type DepartmentService struct {
	repo *departments_repository_postgres.DepartmentRepository
}

func New(repo *departments_repository_postgres.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) Create(name string, parentID *uint) (*model.Department, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, fmt.Errorf("%w: name is required", ErrBadRequest)
	}

	if parentID != nil {
		parent, err := s.repo.GetByID(*parentID)
		if err != nil {
			return nil, err
		}
		if parent == nil {
			return nil, fmt.Errorf("%w: parent department", ErrNotFound)
		}
	}

	dept := &model.Department{Name: name, ParentID: parentID}
	return dept, s.repo.Create(dept)
}
func (s *DepartmentService) GetWithTree(id uint, depth int, includeEmployees bool) (*model.Department, error) {
	if depth < 1 {
		depth = 1
	}
	if depth > 5 {
		depth = 5
	}

	dept, err := s.repo.GetWithTree(id, depth, includeEmployees)
	if err != nil {
		return nil, err
	}
	if dept == nil {
		return nil, ErrNotFound
	}
	return dept, nil
}

func (s *DepartmentService) GetByID(id uint) (*model.Department, error) {
	dept, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dept == nil {
		return nil, ErrNotFound
	}
	return dept, nil
}

func (s *DepartmentService) Update(id uint, name *string, parentID *uint) (*model.Department, error) {
	dept, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dept == nil {
		return nil, ErrNotFound
	}

	if name != nil {
		trimmed := strings.TrimSpace(*name)
		if trimmed == "" {
			return nil, fmt.Errorf("%w: name cannot be empty", ErrBadRequest)
		}
		dept.Name = trimmed
	}

	if parentID != nil {
		if *parentID == id {
			return nil, fmt.Errorf("%w: department cannot be its own parent", ErrConflict)
		}
		dept.ParentID = parentID
	}

	return dept, s.repo.Update(dept)
}

func (s *DepartmentService) Delete(id uint) error {
	dept, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if dept == nil {
		return ErrNotFound
	}
	return s.repo.Delete(id)
}
