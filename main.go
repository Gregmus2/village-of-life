package main

import (
	"github.com/gregmus2/village-of-life/area"
	"github.com/gregmus2/village-of-life/character"
	"github.com/gregmus2/village-of-life/config"
	"github.com/gregmus2/village-of-life/storage"
	"github.com/gregmus2/village-of-life/util"
	"github.com/gregmus2/village-of-life/world"
	"go.uber.org/dig"
	"log"

	//"github.com/gregmus2/village-of-life/storage"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := dig.New()
	FatalOnError(c.Provide(config.NewConfig))
	FatalOnError(c.Provide(area.NewService))
	FatalOnError(c.Provide(character.NewService))
	FatalOnError(c.Provide(storage.NewService))
	FatalOnError(c.Provide(world.NewService))
	FatalOnError(c.Provide(util.NewService))
	FatalOnError(c.Invoke(func(s *area.Service) error {
		_, err := s.ParseArea("untitled.tmx")
		if err != nil {
			return err
		}

		return nil
	}))
}

func FatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
