package departments_service

import (
	"fmt"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
	"go.uber.org/zap"
)

func (s *DepartmentService) Delete(id uint, mode string, reassignTo *uint) error {
	s.logger.Info("deleting department", zap.Uint("id", id), zap.String("mode", mode))

	dept, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("failed to get department", zap.Uint("id", id), zap.Error(err))
		return err
	}
	if dept == nil {
		return core_errors.ErrNotFound
	}

	switch mode {
	case "cascade":
		if err := s.repo.Delete(id); err != nil {
			s.logger.Error("failed to delete department", zap.Uint("id", id), zap.Error(err))
			return err
		}
		s.logger.Info("department deleted", zap.Uint("id", id), zap.String("mode", "cascade"))
		return nil

	case "reassign":
		if reassignTo == nil {
			return fmt.Errorf("%w: reassign_to_department_id is required", core_errors.ErrBadRequest)
		}
		target, err := s.repo.GetByID(*reassignTo)
		if err != nil {
			s.logger.Error("failed to get target department", zap.Uint("id", *reassignTo), zap.Error(err))
			return err
		}
		if target == nil {
			return fmt.Errorf("%w: target department", core_errors.ErrNotFound)
		}
		if *reassignTo == id {
			return fmt.Errorf("%w: cannot reassign to the same department", core_errors.ErrConflict)
		}
		if err := s.repo.DeleteAndReassign(id, *reassignTo); err != nil {
			s.logger.Error("failed to delete and reassign department", zap.Uint("id", id), zap.Error(err))
			return err
		}
		s.logger.Info("department deleted", zap.Uint("id", id), zap.String("mode", "reassign"), zap.Uint("reassign_to", *reassignTo))
		return nil

	default:
		return fmt.Errorf("%w: mode must be cascade or reassign", core_errors.ErrBadRequest)
	}
}
