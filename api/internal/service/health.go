package service

import (
	"time"
)

var healthService = &HealthServiceStruct{}

type HealthServiceStruct struct{}

func NewHealthService() *HealthServiceStruct {
	return healthService
}

func (h *HealthServiceStruct) OpenHealth() HealthDto {
	return HealthDto{IsUp: true, Time: time.Now()}
}

func (h *HealthServiceStruct) CloseHealth() HealthDto {
	return HealthDto{IsUp: true, Time: time.Now()}
}

type HealthDto struct {
	IsUp bool      `json:"isUp"`
	Time time.Time `json:"time"`
	User string    `json:"-"`
}
