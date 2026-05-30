package departments_service

import (
	"fmt"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
)

func (s *DepartmentService) Delete(id uint, mode string, reassignTo *uint) error {
	dept, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if dept == nil {
		return core_errors.ErrNotFound
	}

	switch mode {
	case "cascade":
		return s.repo.Delete(id)

	case "reassign":
		if reassignTo == nil {
			return fmt.Errorf("%w: reassign_to_department_id is required", core_errors.ErrBadRequest)
		}
		target, err := s.repo.GetByID(*reassignTo)
		if err != nil {
			return err
		}
		if target == nil {
			return fmt.Errorf("%w: target department", core_errors.ErrNotFound)
		}
		if *reassignTo == id {
			return fmt.Errorf("%w: cannot reassign to the same department", core_errors.ErrConflict)
		}
		return s.repo.DeleteAndReassign(id, *reassignTo)

	default:
		return fmt.Errorf("%w: mode must be cascade or reassign", core_errors.ErrBadRequest)
	}
}
