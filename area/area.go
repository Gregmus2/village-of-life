package area

import (
	"errors"
	"github.com/gregmus2/village-of-life/character"
	"github.com/lafriks/go-tiled"
)

type Obstacle struct{}
type ObstaclesArea map[int]map[int]*Obstacle
type CharactersArea map[int]map[int]*character.Character
type SpritesArea [][]*Sprite
type Sprite struct {
	ID          uint32
	TileSetName string
}

type Area struct {
	W          int
	H          int
	tileW      int
	tileH      int
	Obstacles  ObstaclesArea
	Characters CharactersArea
	Sprites    SpritesArea
	TileSets   map[string]*tiled.Tileset
}

const SpritesLayer string = "sprites"
const ObstaclesLayer string = "obstacles"
const CharactersLayer string = "characters"

func ParseArea(file string) (*Area, error) {
	a := &Area{
		Obstacles:  make(ObstaclesArea),
		Characters: make(CharactersArea),
		Sprites:    nil,
		TileSets:   make(map[string]*tiled.Tileset),
	}

	gameMap, err := tiled.LoadFromFile(file)
	if err != nil {
		return nil, err
	}

	a.tileW = gameMap.TileWidth
	a.tileH = gameMap.TileHeight
	a.W = gameMap.Width
	a.H = gameMap.Height

	for _, tileSet := range gameMap.Tilesets {
		a.TileSets[tileSet.Name] = tileSet
	}

	for _, layer := range gameMap.Layers {
		switch layer.Name {
		case SpritesLayer:
			a.ParseSprites(layer.Tiles)
		case ObstaclesLayer:
		case CharactersLayer:
		default:
			return nil, errors.New("unknown layer: " + layer.Name)
		}
	}

	return a, nil
}

func (a *Area) ParseSprites(tiles []*tiled.LayerTile) {
	a.Sprites = make(SpritesArea, 0, tiles[0].Tileset.Columns)
	for x := 0; x < a.W; x++ {
		a.Sprites = append(a.Sprites, make([]*Sprite, 0, a.H))
		for y := 0; y < a.H; y++ {
			tile := tiles[y+a.W*x]
			a.Sprites[x] = append(a.Sprites[x], &Sprite{tile.ID, tile.Tileset.Name})
		}
	}
}
