package departments_service

import (
	"fmt"
	"strings"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
)

func (s *DepartmentService) Create(name string, parentID *uint) (*model.Department, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, fmt.Errorf("%w: name is required", core_errors.ErrBadRequest)
	}

	if parentID != nil {
		parent, err := s.repo.GetByID(*parentID)
		if err != nil {
			return nil, err
		}
		if parent == nil {
			return nil, fmt.Errorf("%w: parent department", core_errors.ErrNotFound)
		}
	}
	exists, err := s.repo.ExistsByParentAndName(parentID, name, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w: depratment with this name already exists under the same parent", core_errors.ErrConflict)
	}

	dept := &model.Department{Name: name, ParentID: parentID}
	return dept, s.repo.Create(dept)
}
