package service

import (
	"time"

	"github.com/ribeirosaimon/skadi/api/internal/skadiEngine"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) OpenHealth() HealthDto {
	return HealthDto{Environment: skadiEngine.GetEnvironment(), Time: time.Now()}
}

func (h *HealthService) CloseHealth() HealthDto {
	// sot
	return HealthDto{Environment: skadiEngine.GetEnvironment(), Time: time.Now()}
}

type HealthDto struct {
	Environment string    `json:"environment"`
	Time        time.Time `json:"time"`
	User        string    `json:"-"`
}
