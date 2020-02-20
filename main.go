package main

import (
	"github.com/gregmus2/village-of-life/area"
	//"github.com/gregmus2/village-of-life/storage"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//s := storage.NewBolt("village")

	area.ParseArea("untitled.tmx")
	server := NewJSONServer()

	server.Start(12050)
}
