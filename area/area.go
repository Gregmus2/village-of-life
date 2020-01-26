package area

import (
	"github.com/gregmus2/village-of-life/character"
	"github.com/lafriks/go-tiled"
)

type Obstacle struct{}
type ObstaclesArea map[int]map[int]*Obstacle

type CharactersArea map[int]map[int]*character.Character

type SpritesArea [][]int

type Area struct {
	Obstacles  ObstaclesArea
	Characters CharactersArea
	Sprites    SpritesArea
}

func ParseArea(file string) (*Area, error) {
	gameMap, err := tiled.LoadFromFile(file)
	if err != nil {
		return nil, err
	}

	for _, layer := range gameMap.Layers {
		layer.
	}
}