package user

import (
	"github.com/gregmus2/village-of-life/area"
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
	Target     area.Pos
}

type User struct {
	Skills struct {
		Trade      float64
		Parenthood float64
	}
	Relations       map[*User]float64
	Propensity      map[int]float64
	Characteristics struct {
		Social  float64
		Savvy   float64
		PhysDev float64
		Energy  float64
		Saving  float64
	}
	Age       float64
	BirthDate int
	Sex       Sex
	Goals     map[int]*Goal
}
