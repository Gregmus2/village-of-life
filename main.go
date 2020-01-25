package main

import (
	//"github.com/gregmus2/village-of-life/storage"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//s := storage.NewBolt("village")

	server := NewJSONServer()

	server.Start(12050)
}
