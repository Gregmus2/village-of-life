package area

import (
	"github.com/gregmus2/village-of-life/character"
	"github.com/gregmus2/village-of-life/util"
)

type Service struct {
	c *character.Service
	u *util.Service
}

func NewService(c *character.Service, u *util.Service) *Service {
	return &Service{c, u}
}
