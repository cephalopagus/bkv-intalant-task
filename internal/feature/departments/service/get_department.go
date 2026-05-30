package departments_service

import (
	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	"go.uber.org/zap"
)

func (s *DepartmentService) GetWithTree(id uint, depth int, includeEmployees bool) (*model.Department, error) {
	s.logger.Info("getting department with tree",
		zap.Uint("id", id),
		zap.Int("depth", depth),
		zap.Bool("include_employees", includeEmployees),
	)

	if depth < 1 {
		depth = 1
	}
	if depth > 5 {
		depth = 5
	}

	dept, err := s.repo.GetWithTree(id, depth, includeEmployees)
	if err != nil {
		s.logger.Error("failed to get department with tree", zap.Uint("id", id), zap.Error(err))
		return nil, err
	}
	if dept == nil {
		return nil, core_errors.ErrNotFound
	}
	return dept, nil
}

func (s *DepartmentService) GetByID(id uint) (*model.Department, error) {
	s.logger.Info("getting department", zap.Uint("id", id))

	dept, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("failed to get department", zap.Uint("id", id), zap.Error(err))
		return nil, err
	}
	if dept == nil {
		return nil, core_errors.ErrNotFound
	}
	return dept, nil
}
