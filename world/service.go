package world

import "time"

type Service struct {
	StartTime time.Time
}

func NewService() *Service {
	return &Service{
		StartTime: time.Now(),
	}
}
