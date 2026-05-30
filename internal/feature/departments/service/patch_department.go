package departments_service

import (
	"fmt"
	"strings"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
)

func (s *DepartmentService) Update(id uint, name *string, parentID *uint) (*model.Department, error) {

	dept, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dept == nil {
		return nil, core_errors.ErrNotFound
	}

	if name != nil {
		trimmed := strings.TrimSpace(*name)
		if trimmed == "" {
			return nil, fmt.Errorf("%w: name cannot be empty", core_errors.ErrBadRequest)
		}
		dept.Name = trimmed
	}

	if parentID != nil {
		if *parentID == id {
			return nil, fmt.Errorf("%w: department cannot be its own parent", core_errors.ErrConflict)
		}

		isDesc, err := s.repo.IsDescendant(id, *parentID)
		if err != nil {
			return nil, err
		}
		if isDesc {
			return nil, fmt.Errorf("%w: cannot move department into its own subtree", core_errors.ErrConflict)
		}

		dept.ParentID = parentID
	}

	exists, err := s.repo.ExistsByParentAndName(dept.ParentID, dept.Name, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w: department with this name already exists under the same parent", core_errors.ErrConflict)
	}

	return dept, s.repo.Update(dept)
}
