package status

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
)

type StatusService struct {
	StatusRepository IStatusRepository `inject:"status_repository"`
}

type IStatusService = *StatusService

func NewStatusService() *StatusService {
	return proinject.Resolve(&StatusService{})
}

func (s *StatusService) GetStatus() Status {
	return Status{
		UpdatedAt: time.Now(),
		Dependencies: StatusDependencies{
			Database: s.StatusRepository.GetStatusDatabase(),
		},
	}
}
