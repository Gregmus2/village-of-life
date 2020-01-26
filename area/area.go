package area

import "github.com/gregmus2/village-of-life/character"

type Obstacle struct{}
type ObstaclesArea map[int]map[int]*Obstacle

type CharactersArea map[int]map[int]*character.Character

type SpritesArea [][]int

type Area struct {
	Obstacles  ObstaclesArea
	Characters CharactersArea
	Sprites    SpritesArea
}
