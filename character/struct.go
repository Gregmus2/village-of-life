package character

import (
	"github.com/gregmus2/village-of-life/util"
	"math/rand"
)

type Skills struct {
	Social     float64
	Trade      float64
	Parenthood float64
}

type Sex byte

const (
	Male Sex = iota
	Female
)

type GoalType int8

const (
	Main GoalType = iota
	Secondary
	Tmp
)

type Goal struct {
	ID         int
	Importance float32
	Type       GoalType
	Target     util.Pos
}

type Character struct {
	Name   string
	Skills struct {
		Trade      float64
		Parenthood float64
	}
	Relations       map[*Character]float64
	Propensity      map[int]float64
	Characteristics struct {
		Social  float64
		Savvy   float64
		PhysDev float64
		Energy  float64
		Saving  float64
	}
	Age       float64
	BirthTime int
	Sex       Sex
	Goals     map[int]*Goal
}

func (s *Service) NewCharacter(name string, birthTime int) *Character {
	return &Character{
		Name: name,
		Skills: struct {
			Trade      float64
			Parenthood float64
		}{Trade: 0, Parenthood: 0},
		Relations:  make(map[*Character]float64),
		Propensity: make(map[int]float64),
		Characteristics: struct {
			Social  float64
			Savvy   float64
			PhysDev float64
			Energy  float64
			Saving  float64
		}{0, 0, 0, 0, 0},
		Age:       0,
		BirthTime: birthTime,
		Sex:       Sex(rand.Intn(2)),
		Goals:     make(map[int]*Goal),
	}
}
