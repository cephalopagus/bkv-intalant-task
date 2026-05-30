package departments_service

import (
	"fmt"
	"strings"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	"go.uber.org/zap"
)

func (s *DepartmentService) Create(name string, parentID *uint) (*model.Department, error) {
	s.logger.Info("creating department", zap.String("name", name), zap.Any("parent_id", parentID))

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, fmt.Errorf("%w: name is required", core_errors.ErrBadRequest)
	}

	if parentID != nil {
		parent, err := s.repo.GetByID(*parentID)
		if err != nil {
			s.logger.Error("failed to get parent department", zap.Uint("parent_id", *parentID), zap.Error(err))
			return nil, err
		}
		if parent == nil {
			return nil, fmt.Errorf("%w: parent department", core_errors.ErrNotFound)
		}
	}

	exists, err := s.repo.ExistsByParentAndName(parentID, name, nil)
	if err != nil {
		s.logger.Error("failed to check department name", zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w: department with this name already exists under the same parent", core_errors.ErrConflict)
	}

	dept := &model.Department{Name: name, ParentID: parentID}
	if err := s.repo.Create(dept); err != nil {
		s.logger.Error("failed to create department", zap.Error(err))
		return nil, err
	}

	s.logger.Info("department created", zap.Uint("id", dept.ID), zap.String("name", dept.Name))
	return dept, nil
}
