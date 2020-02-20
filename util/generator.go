package util

import "math/rand"

func (s *Service) GenerateName() string {
	return s.names[rand.Intn(len(s.names))]
}
